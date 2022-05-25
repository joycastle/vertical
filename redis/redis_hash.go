package redis

import (
	"github.com/garyburd/redigo/redis"
	"github.com/joycastle/cop/connector"
)

func Rds_HGetString(sn string, key string, subKey string) (string, error) {
	conn := connector.GetRedisConn(sn)
	defer conn.Close()

	r, err := redis.String(conn.Do("HGET", key, subKey))

	if err != nil && err.Error() == ErrNotFoundDesc {
		return r, ErrNotFound
	}

	return r, err
}
func Rds_HSet(sn string, key string, subKey string, v interface{}) (int, error) {
	conn := connector.GetRedisConn(sn)
	defer conn.Close()

	return redis.Int(conn.Do("HSET", key, subKey, v))
}
func Rds_HDel(sn string, key string, subKey string) (int, error) {
	conn := connector.GetRedisConn(sn)
	defer conn.Close()

	return redis.Int(conn.Do("HDel", key, subKey))
}
func Rds_HExists(sn string, key string, subKey string) (int, error) {
	conn := connector.GetRedisConn(sn)
	defer conn.Close()

	return redis.Int(conn.Do("HEXISTS", key, subKey))
}
func Rds_HGetAllString(sn string, key string) (map[string]string, error) {
	conn := connector.GetRedisConn(sn)
	defer conn.Close()

	return redis.StringMap(conn.Do("HGETALL", key))
}
func Rds_HGetAllInt(sn string, key string) (map[string]int, error) {
	conn := connector.GetRedisConn(sn)
	defer conn.Close()

	return redis.IntMap(conn.Do("HGETALL", key))
}
func Rds_HGetAllInt64(sn string, key string) (map[string]int64, error) {
	conn := connector.GetRedisConn(sn)
	defer conn.Close()

	return redis.Int64Map(conn.Do("HGETALL", key))
}
