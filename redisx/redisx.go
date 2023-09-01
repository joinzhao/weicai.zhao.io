package redisx

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"io"
	"net"
	"sync"
	"time"
	"weicai.zhao.io/resource"
)

const (
	errorFormat = "[go-redis] err: %s"
)

type (
	redisX struct {
		client *redis.Client
		sync.Once
	}

	Config struct {
		Usage       string `json:"usage"`
		DSN         string `json:"dsn"`
		Password    string `json:"password"`
		Username    string `json:"username"`
		DB          int    `json:"db"`
		MaxIdleConn int    `json:"maxIdleConn"`
	}

	Manager struct {
		resources     *resource.Manager
		configs       map[string]*Config
		defaultConfig string
	}
)

func New(configs ...*Config) (*Manager, error) {
	if len(configs) == 0 {
		return nil, fmt.Errorf(errorFormat, "empty config")
	}

	var cfgs = make(map[string]*Config)
	for i := 0; i < len(configs); i++ {
		cfgs[configs[i].Usage] = configs[i]
	}

	return &Manager{
		resources:     resource.NewManager(),
		configs:       cfgs,
		defaultConfig: configs[0].Usage,
	}, nil
}

func (x *redisX) Close() error {
	return nil
}

func (m *Manager) DefaultUse() *redis.Client {
	return m.MustUse(m.defaultConfig)
}

func (m *Manager) MustUse(usage string) *redis.Client {
	cli, err := m.Use(usage)
	if err != nil {
		panic(err)
	}

	return cli
}

func (m *Manager) Use(usage string) (*redis.Client, error) {
	cfg, ok := m.configs[usage]

	if !ok {
		return nil, fmt.Errorf(errorFormat, "key not exists")
	}

	cli, err := m.getConn(cfg)

	if err != nil {
		return nil, err
	}

	return cli, nil
}

func (m *Manager) getConn(cfg *Config) (*redis.Client, error) {
	var err error
	x, _ := m.connRedis(cfg)

	x.Do(func() {
		err = x.client.Ping(context.Background()).Err()
	})

	if err != nil {
		return nil, err
	}

	return x.client, nil
}

func (m *Manager) connRedis(cfg *Config) (*redisX, error) {
	get, _ := m.resources.Get(cfg.Usage, func() (io.Closer, error) {
		client := m.create(cfg)

		return &redisX{client, sync.Once{}}, nil
	})

	return get.(*redisX), nil
}

func (m *Manager) create(cfg *Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		//连接信息
		Network: "tcp",   //网络类型，tcp or unix，默认tcp
		Addr:    cfg.DSN, //主机名+冒号+端口，默认localhost:6379
		//可自定义连接函数
		Dialer: func(ctx context.Context, network, addr string) (net.Conn, error) {
			return net.Dial(network, addr)
		},
		//钩子函数
		OnConnect: func(ctx context.Context, cn *redis.Conn) error {
			//仅当客户端执行命令时需要从连接池获取连接时，如果连接池需要新建连接时则会调用此钩子函数
			return nil
		},
		Username: "",
		Password: cfg.Password, //密码
		DB:       cfg.DB,       // redis数据库index
		//命令执行失败时的重试策略
		MaxRetries:      0,                      // 命令执行失败时，最多重试多少次，默认为0即不重试
		MinRetryBackoff: 8 * time.Millisecond,   //每次计算重试间隔时间的下限，默认8毫秒，-1表示取消间隔
		MaxRetryBackoff: 512 * time.Millisecond, //每次计算重试间隔时间的上限，默认512毫秒，-1表示取消间隔
		//超时
		DialTimeout:  5 * time.Second, //连接建立超时时间，默认5秒。
		ReadTimeout:  3 * time.Second, //读超时，默认3秒， -1表示取消读超时
		WriteTimeout: 3 * time.Second, //写超时，默认等于读超时
		PoolFIFO:     false,
		//连接池容量及闲置连接数量
		PoolSize:     15,              // 连接池最大socket连接数，默认为4倍CPU数， 4 * runtime.NumCPU
		MinIdleConns: cfg.MaxIdleConn, //在启动阶段创建指定数量的Idle连接，并长期维持idle状态的连接数不少于指定数量；。
		MaxConnAge:   0 * time.Second, //连接存活时长，从创建开始计时，超过指定时长则关闭连接，默认为0，即不关闭存活时长较长的连接
		PoolTimeout:  4 * time.Second, //当所有连接都处在繁忙状态时，客户端等待可用连接的最大等待时长，默认为读超时+1秒。
		IdleTimeout:  5 * time.Minute, //闲置超时，默认5分钟，-1表示取消闲置超时检查
		//闲置连接检查包括IdleTimeout，MaxConnAge
		IdleCheckFrequency: 60 * time.Second, //闲置连接检查的周期，默认为1分钟，-1表示不做周期性检查，只在客户端获取连接时对闲置连接进行处理。
	})
}
