package redis

import (
	"errors"

	"github.com/garyburd/redigo/redis"
	"github.com/joycastle/cop/connector"
)

var (
	ErrNotFoundDesc = "redigo: nil returned"
	ErrNotFound     = errors.New("redigo: nil returned")
)

func Rds_GetString(sn string, key string) (string, error) {
	conn := connector.GetRedisConn(sn)
	defer conn.Close()

	r, err := redis.String(conn.Do("GET", key))

	if err != nil && err.Error() == ErrNotFoundDesc {
		return r, ErrNotFound
	}

	return r, err
}
func Rds_GetInt(sn string, key string) (int, error) {
	conn := connector.GetRedisConn(sn)
	defer conn.Close()

	r, err := redis.Int(conn.Do("GET", key))

	if err != nil && err.Error() == ErrNotFoundDesc {
		return r, ErrNotFound
	}

	return r, err
}

func Rds_Set(sn string, key string, v interface{}) (string, error) {
	conn := connector.GetRedisConn(sn)
	defer conn.Close()

	return redis.String(conn.Do("SET", key, v))
}
func Rds_SetEx(sn string, key string, v interface{}, ex int) (string, error) {
	conn := connector.GetRedisConn(sn)
	defer conn.Close()

	return redis.String(conn.Do("SETEX", key, ex, v))
}
func Rds_Del(sn string, key string) (int, error) {
	conn := connector.GetRedisConn(sn)
	defer conn.Close()

	return redis.Int(conn.Do("DEL", key))
}
func Rds_Expire(sn string, key string, ex int) (int, error) {
	conn := connector.GetRedisConn(sn)
	defer conn.Close()

	return redis.Int(conn.Do("EXPIRE", key, ex))
}
func Rds_TTL(sn string, key string) (int, error) {
	conn := connector.GetRedisConn(sn)
	defer conn.Close()

	return redis.Int(conn.Do("TTL", key))
}
