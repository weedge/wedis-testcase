package injectors

import (
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisClientOptions struct {
	Addr string `mapstructure:"addr"`
	Db   int    `mapstructure:"db"`

	limiter redis.Limiter

	RedisClientCommonOptions

	// optional, more details see redis.Options
	// if use redis.Options, config options can't use
	redisOpts *redis.Options
}

func (m *RedisClientOptions) String() string {
	res := fmt.Sprintf("addr:%s", m.Addr)
	res += fmt.Sprintf("db:%d", m.Db)
	res += fmt.Sprintf("RedisClientCommonOptions:%+v", m.RedisClientCommonOptions)
	if m.redisOpts != nil {

		res += fmt.Sprintf("redisOpts:%+v", m.redisOpts)
	}
	return res
}

func WithRedisAddr(addr string) Option {
	return NewOpt(func(op OptPrinter) {
		o, ok := op.(*RedisClientOptions)
		if !ok {
			return
		}
		o.Addr = addr
	})
}

func WithRedisDB(db int) Option {
	return NewOpt(func(op OptPrinter) {
		o, ok := op.(*RedisClientOptions)
		if !ok {
			return
		}
		o.Db = db
	})
}

func WithLimiter(limiter redis.Limiter) Option {
	return NewOpt(func(op OptPrinter) {
		o, ok := op.(*RedisClientOptions)
		if !ok {
			return
		}
		o.limiter = limiter
	})
}

func WithGoRedisOpts(opts *redis.Options) Option {
	return NewOpt(func(op OptPrinter) {
		o, ok := op.(*RedisClientOptions)
		if !ok {
			return
		}
		o.redisOpts = opts
	})
}

func WithRedisClientCommonOptions(opts RedisClientCommonOptions) Option {
	return NewOpt(func(op OptPrinter) {
		o, ok := op.(*RedisClusterClientOptions)
		if ok {
			o.RedisClientCommonOptions = opts
			return
		}

		co, ok := op.(*RedisClientOptions)
		if ok {
			co.RedisClientCommonOptions = opts
			return
		}
	})
}

func DefaultRedisClientOptions() *RedisClientOptions {
	return &RedisClientOptions{
		Addr: "localhost:6379",
		RedisClientCommonOptions: RedisClientCommonOptions{
			Password: "",
			Username: "",

			MaxRetries:      3,
			MinRetryBackoff: 3 * time.Second,
			MaxRetryBackoff: 5 * time.Second,
			DialTimeout:     5 * time.Second,
			ReadTimeout:     3 * time.Second,
			WriteTimeout:    3 * time.Second,

			// connect pool
			PoolSize:     100,
			MinIdleConns: 10,
			MaxConnAge:   60 * time.Second,
			PoolTimeout:  5 * time.Second,
			IdleTimeout:  30 * time.Second,
		},
	}
}

// InitRedisClient init redis instance
func InitRedisClient(options ...Option) redis.UniversalClient {
	opts := getClientOptions(options...)
	redisOptions := &redis.Options{
		Addr:            opts.Addr,
		Username:        opts.Username,
		Password:        opts.Password,
		DB:              opts.Db,
		MaxRetries:      opts.MaxRetries,
		MinRetryBackoff: opts.MinRetryBackoff,
		MaxRetryBackoff: opts.MaxRetryBackoff,
		DialTimeout:     opts.DialTimeout,
		ReadTimeout:     opts.ReadTimeout,
		WriteTimeout:    opts.WriteTimeout,
		PoolSize:        opts.PoolSize,
		MinIdleConns:    opts.MinIdleConns,
		MaxIdleConns:    opts.MaxIdleConns,
		ConnMaxLifetime: opts.MaxConnAge,
		PoolTimeout:     opts.PoolTimeout,
		ConnMaxIdleTime: opts.IdleTimeout,
		//TLSConfig:          &tls.Config{},
		Limiter: opts.limiter,
	}

	if opts.redisOpts != nil {
		redisOptions = opts.redisOpts
	}

	return redis.NewClient(redisOptions)
}

func getClientOptions(opts ...Option) *RedisClientOptions {
	options := DefaultRedisClientOptions()
	for _, o := range opts {
		o.apply(options)
	}

	return options
}
