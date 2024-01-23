package redistool

import (
	"time"

	"fclink.cn/grabblock/conf"
	"fclink.cn/grabblock/tools/log"
	"github.com/gomodule/redigo/redis"
)

var pool *redis.Pool

func Setup() {
	pool = poolInitRedis(&conf.Context.Redis)
}
func Get() redis.Conn {
	log.Infof("redis activecount=%d", pool.ActiveCount())
	return pool.Get()
}

// redis pool
func poolInitRedis(config *conf.Redis) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     config.Maxidle, //空闲数
		IdleTimeout: 30 * time.Second,
		MaxActive:   config.Maxactive, //最大数
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", config.Host,
				redis.DialConnectTimeout(time.Duration(3)*time.Second),
				redis.DialReadTimeout(time.Duration(3)*time.Second),
				redis.DialWriteTimeout(time.Duration(3)*time.Second),
			)
			if err != nil {
				return nil, err
			}
			if config.Password != "" {
				if _, err := c.Do("AUTH", config.Password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}
