package vertical

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/garyburd/redigo/redis"
)

type RedisConf struct {
	Addrs []string

	TestInterval time.Duration

	MaxActive   int
	MaxIdle     int
	IdleTimeout time.Duration

	ConnectTimeout time.Duration
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
}

var (
	redisConnMapping       map[string]*redis.Pool = make(map[string]*redis.Pool)
	Err_invalid_connection                        = errors.New("Invalid connection")
)

func InitRedis(configs map[string]RedisConf) error {
	for sn, config := range configs {
		rp := &redis.Pool{
			MaxActive:   config.MaxActive,
			MaxIdle:     config.MaxIdle,
			IdleTimeout: config.IdleTimeout,
			Wait:        true,
			Dial: func() (redis.Conn, error) {
				var (
					addr string
					conn redis.Conn
					err  error
				)
				addr = config.Addrs[rand.Intn(len(config.Addrs))]
				conn, err = redis.DialTimeout("tcp", addr, config.ConnectTimeout, config.ReadTimeout, config.WriteTimeout)
				if err != nil {
					GetLogger("error").Printf("[REDIS] connect to redis[%s] failed: %s", addr, err)
					return nil, err
				}

				_, err = conn.Do("PING")
				if err != nil {
					GetLogger("error").Printf("[REDIS] ping redis[%s] failed: %s", addr, err)
					return nil, err
				}
				return conn, nil
			},
			TestOnBorrow: func(conn redis.Conn, t time.Time) error {
				if time.Since(t) < config.TestInterval {
					return nil
				}
				_, err := conn.Do("PING")
				return err
			},
		}

		redisConnMapping[sn] = rp
	}
	return nil
}

func GetRedis(sn string) (*RedisConnWrapper, error) {
	if conn, exists := redisConnMapping[sn]; exists {
		return &RedisConnWrapper{Conn: conn.Get()}, nil
	}
	return nil, fmt.Errorf("have no mysql cluster: %s", sn)
}
