package vertical

import (
	"github.com/garyburd/redigo/redis"
)

type RedisConnWrapper struct {
	Conn redis.Conn
}

func (c *RedisConnWrapper) Close() error {
	if c.Conn != nil {
		return c.Conn.Close()
	}
	return nil
}
func (c *RedisConnWrapper) Do(command string, argv ...interface{}) (interface{}, error) {
	if c.Conn == nil {
		GetLogger("error").Printf("[REDIS] invlaid connection. call [%s %v]", command, argv)
		return nil, Err_invalid_connection
	}
	return c.Conn.Do(command, argv...)
}
func (c *RedisConnWrapper) DoBool(command string, argv ...interface{}) (bool, error) {
	return redis.Bool(c.Do(command, argv...))
}
func (c *RedisConnWrapper) DoByteSlices(command string, argv ...interface{}) ([][]byte, error) {
	return redis.ByteSlices(c.Do(command, argv...))
}
func (c *RedisConnWrapper) DoBytes(command string, argv ...interface{}) ([]byte, error) {
	return redis.Bytes(c.Do(command, argv...))
}
func (c *RedisConnWrapper) DoFloat64(command string, argv ...interface{}) (float64, error) {
	return redis.Float64(c.Do(command, argv...))
}
func (c *RedisConnWrapper) DoInt(command string, argv ...interface{}) (int, error) {
	return redis.Int(c.Do(command, argv...))
}
func (c *RedisConnWrapper) DoInt64(command string, argv ...interface{}) (int64, error) {
	return redis.Int64(c.Do(command, argv...))
}
func (c *RedisConnWrapper) DoInt64Map(command string, argv ...interface{}) (map[string]int64, error) {
	return redis.Int64Map(c.Do(command, argv...))
}
func (c *RedisConnWrapper) DoIntMap(command string, argv ...interface{}) (map[string]int, error) {
	return redis.IntMap(c.Do(command, argv...))
}
func (c *RedisConnWrapper) DoInts(command string, argv ...interface{}) ([]int, error) {
	return redis.Ints(c.Do(command, argv...))
}
func (c *RedisConnWrapper) DoMultiBulk(command string, argv ...interface{}) ([]interface{}, error) {
	return redis.MultiBulk(c.Do(command, argv...))
}
func (c *RedisConnWrapper) DoPositions(command string, argv ...interface{}) ([]*[2]float64, error) {
	return redis.Positions(c.Do(command, argv...))
}
func (c *RedisConnWrapper) DoString(command string, argv ...interface{}) (string, error) {
	return redis.String(c.Do(command, argv...))
}
func (c *RedisConnWrapper) DoStringMap(command string, argv ...interface{}) (map[string]string, error) {
	return redis.StringMap(c.Do(command, argv...))
}
func (c *RedisConnWrapper) DoStrings(command string, argv ...interface{}) ([]string, error) {
	return redis.Strings(c.Do(command, argv...))
}
func (c *RedisConnWrapper) DoUint64(command string, argv ...interface{}) (uint64, error) {
	return redis.Uint64(c.Do(command, argv...))
}
func (c *RedisConnWrapper) DoValues(command string, argv ...interface{}) ([]interface{}, error) {
	return redis.Values(c.Do(command, argv...))
}

func (c *RedisConnWrapper) Del(k string) (bool, error) {
	return redis.Bool(c.Do("DEL", k))
}
func (c *RedisConnWrapper) Set(k string, v interface{}) (string, error) {
	return redis.String(c.Do("SET", k, v))
}
func (c *RedisConnWrapper) TTL(k string) (int, error) {
	return redis.Int(c.Do("TTL", k))
}
func (c *RedisConnWrapper) Expire(k string, expire int) (bool, error) {
	return redis.Bool(c.Do("EXPIRE", k, expire))
}
func (c *RedisConnWrapper) GetString(k string) (string, error) {
	return redis.String(c.Do("GET", k))
}
func (c *RedisConnWrapper) GetInt(k string) (int, error) {
	return redis.Int(c.Do("GET", k))
}
func (c *RedisConnWrapper) SetEx(k string, v interface{}, expire int) (string, error) {
	return redis.String(c.Do("SETEX", k, expire, v))
}
