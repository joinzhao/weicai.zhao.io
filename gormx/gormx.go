package gormx

import (
	"context"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"io"
	"log"
	"os"
	"sync"
	"time"
	"weicai.zhao.io/resource"
)

const (
	SQLModeDebug     = "debug"
	SQLModeRelease   = "release"
	slowSQLThreshold = time.Second * 2
	logPrefix        = "[gorm-log]"
	maxIdleConns     = 64
	maxOpenConns     = 64
	maxLifetime      = time.Minute
)

type (
	// sql is a db shadow
	sql struct {
		*gorm.DB
		once sync.Once
	}

	// Config init config
	Config struct {
		Usage       string // use to manage connection
		RunMode     string // debug, release
		DSN         string // root:root@tcp(127.0.0.1:3306)
		Database    string // default database
		MaxIdleConn int
		MaxOpenConn int
		MaxLifeTime int
	}

	// Manager store all gorm db engine according to the config usage key
	Manager struct {
		resources     *resource.Manager
		configs       map[string]*Config
		defaultConfig *Config
		writers       []io.Writer
	}
)

// Close implements io closer
func (s *sql) Close() error {
	return nil
}

// New return a gorm db Manager
func New(configs []*Config) *Manager {
	if len(configs) == 0 {
		panic("please complete your mysql config")
	}

	cfgMap := make(map[string]*Config)
	for _, config := range configs {
		cfgMap[config.Usage] = config
	}

	return &Manager{
		resources:     resource.NewManager(),
		configs:       cfgMap,
		defaultConfig: configs[0],
	}
}

// SetWriters set gorm log writers, default is stdout
func (m *Manager) SetWriters(writers ...io.Writer) {
	m.writers = writers
}

// Use return engine from incoming key
func (m *Manager) Use(ctx context.Context, key string) (*gorm.DB, error) {

	config, ok := m.configs[key]
	if !ok {
		return nil, fmt.Errorf("gormx: miss [%s] config", key)
	}

	conn, err := m.getGormConn(config)
	if err != nil {
		return nil, err
	}

	return conn.Session(&gorm.Session{NewDB: true, Context: ctx}), nil
}

// MustUseUsage return engine from usage key
func (m *Manager) MustUseUsage(ctx context.Context, key string) *gorm.DB {

	config, ok := m.configs[key]
	if !ok {
		panic(fmt.Errorf("gormx: miss [%s] config", key))
	}

	conn, err := m.getGormConn(config)
	if err != nil {
		panic(err)
	}

	return conn.Session(&gorm.Session{NewDB: true, Context: ctx})
}

// MustUse return default engine, default is the first config in config,
// if err, got panic
func (m *Manager) MustUse(ctx context.Context) *gorm.DB {
	conn, err := m.getGormConn(m.defaultConfig)
	if err != nil {
		panic(err)
	}

	return conn.Session(&gorm.Session{NewDB: true, Context: ctx})
}

// Default return default engine, default is the first config in config
func (m *Manager) Default() *gorm.DB {

	conn, err := m.getGormConn(m.defaultConfig)
	if err != nil {
		panic(err)
	}

	return conn.Session(&gorm.Session{NewDB: true})
}

// getGormConn transfer sql to gorm db
func (m *Manager) getGormConn(config *Config) (*gorm.DB, error) {

	conn, err := m.getSqlConn(config)
	if err != nil {
		return nil, err
	}

	db, err := conn.DB.DB()
	if err != nil {
		return nil, err
	}

	conn.once.Do(func() {
		err = db.Ping()
	})

	if err != nil {
		return nil, err
	}

	return conn.DB, nil
}

// getSqlConn get sql conn from sharedCall conn
func (m *Manager) getSqlConn(config *Config) (*sql, error) {

	get, err := m.resources.Get(config.Usage, func() (io.Closer, error) {

		create, err := m.create(config)
		if err != nil {
			return nil, err
		}

		return create, nil
	})
	if err != nil {
		return nil, err
	}

	return get.(*sql), nil

}

// create connect mysql from incoming sql config
func (m *Manager) create(singleCfg *Config) (*sql, error) {

	var (
		_db                       *gorm.DB
		err                       error
		logLevel                  logger.LogLevel
		writer                    io.Writer
		ignoreRecordNotFoundError bool
	)
	writer = os.Stdout
	if singleCfg.RunMode == SQLModeRelease {
		logLevel = logger.Error
		if m.writers != nil {
			writer = io.MultiWriter(m.writers...)
		}
		ignoreRecordNotFoundError = true
	} else {
		logLevel = logger.Info
	}

	newLogger := logger.New(
		log.New(writer, logPrefix, log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             slowSQLThreshold, // Slow SQL threshold
			LogLevel:                  logLevel,         // Log level,Silent, Error, Warn, Info
			Colorful:                  true,             // Disable color
			IgnoreRecordNotFoundError: ignoreRecordNotFoundError,
		},
	)

	_db, err = gorm.Open(mysql.New(mysql.Config{
		DSN: fmt.Sprintf("%s/%s?charset=utf8mb4&parseTime=True&loc=Local",
			singleCfg.DSN, singleCfg.Database), // DSN data source name, parse time is important !!!
		DefaultStringSize:         256,  // string default length
		SkipInitializeWithVersion: true, // auto config according to version
	}), &gorm.Config{
		SkipDefaultTransaction:                   false,
		NamingStrategy:                           nil,
		FullSaveAssociations:                     false,
		Logger:                                   newLogger,
		NowFunc:                                  nil,
		DryRun:                                   false,
		PrepareStmt:                              false,
		DisableAutomaticPing:                     false,
		DisableForeignKeyConstraintWhenMigrating: true,
		AllowGlobalUpdate:                        false,
		QueryFields:                              false,
		CreateBatchSize:                          1000,
		ClauseBuilders:                           nil,
		ConnPool:                                 nil,
		Dialector:                                nil,
		Plugins:                                  nil,
	})

	if err != nil {
		return nil, fmt.Errorf("[Error] init ( %s ),"+"error [ %s ] \n", singleCfg.Usage, err.Error())
	}

	_ = _db.Use(&OpenTracingPlugin{})

	sqlDB, err := _db.DB()
	if err != nil {
		return nil, fmt.Errorf("[Error] get mysql db ( %s ),"+"error [ %s ] \n", singleCfg.Usage, err.Error())
	}

	// conn setting
	sqlDB.SetConnMaxLifetime(time.Duration(singleCfg.MaxLifeTime) * time.Second)
	sqlDB.SetMaxIdleConns(singleCfg.MaxIdleConn)
	sqlDB.SetMaxOpenConns(singleCfg.MaxOpenConn)

	_ = _db.Callback().Create().Remove("gorm:save_before_associations")
	_ = _db.Callback().Update().Remove("gorm:save_before_associations")

	_ = _db.Callback().Create().Remove("gorm:save_after_associations")
	_ = _db.Callback().Update().Remove("gorm:save_after_associations")

	return &sql{_db, sync.Once{}}, nil
}
