package redisx

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"io"
	"sync"
	"time"
	"weicai.zhao.io/resource"
)

type (
	Config struct {
		Usage    string // use to manage connection
		Addr     string // 127.0.0.1:6379
		DB       int    // select 0
		Password string // password
	}
	// sql is a db shadow
	sql struct {
		*redis.Client
		once sync.Once
	}
	// Manager store all redis db engine according to the conf usage key
	Manager struct {
		resources     *resource.Manager
		configs       map[string]*Config
		defaultConfig *Config
		writers       []io.Writer
	}
)

// New return a redis db Manager
func New(configs []*Config) *Manager {
	if len(configs) == 0 {
		panic("please complete your mysql conf")
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

// Use return engine from incoming key
func (m *Manager) Use(ctx context.Context, key string) (*redis.Client, error) {

	config, ok := m.configs[key]
	if !ok {
		return nil, fmt.Errorf("gormx: miss [%s] conf", key)
	}

	conn, err := m.getRedisConn(config)
	if err != nil {
		return nil, err
	}

	return conn.WithContext(ctx), nil
}

// MustUseUsage return engine from usage key
func (m *Manager) MustUseUsage(ctx context.Context, key string) *redis.Client {

	config, ok := m.configs[key]
	if !ok {
		panic(fmt.Errorf("gormx: miss [%s] conf", key))
	}

	conn, err := m.getRedisConn(config)
	if err != nil {
		panic(err)
	}

	return conn.WithContext(ctx)
}

// MustUse return default engine, default is the first conf in conf,
// if err, got panic
func (m *Manager) MustUse(ctx context.Context) *redis.Client {
	conn, err := m.getRedisConn(m.defaultConfig)
	if err != nil {
		panic(err)
	}

	return conn.WithContext(ctx)
}

// Default return default engine, default is the first conf in conf
func (m *Manager) Default() *redis.Client {
	conn, err := m.getRedisConn(m.defaultConfig)
	if err != nil {
		panic(err)
	}

	return conn.WithContext(context.Background())
}

// getRedisConn transfer sql to redis client
func (m *Manager) getRedisConn(config *Config) (*redis.Client, error) {
	conn, err := m.getSqlConn(config)
	if err != nil {
		return nil, err
	}

	conn.once.Do(func() {
		err = conn.Ping().Err()
	})

	if err != nil {
		return nil, err
	}
	return conn.Client, nil
}

// getSqlConn get sql conn from sharedCall conn
func (m *Manager) getSqlConn(cfg *Config) (*sql, error) {

	get, err := m.resources.Get(cfg.Usage, func() (io.Closer, error) {

		create, err := m.create(cfg)
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

// create connect redis from incoming sql conf
func (m *Manager) create(cfg *Config) (*sql, error) {
	db := redis.NewClient(&redis.Options{
		Addr:         cfg.Addr,
		Password:     cfg.Password,
		DB:           cfg.DB,
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolTimeout:  30 * time.Second,
	})

	return &sql{db, sync.Once{}}, nil
}

// Close implements io closer
func (s *sql) Close() error {
	return s.Client.Close()
}
