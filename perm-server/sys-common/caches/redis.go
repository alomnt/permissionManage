package caches

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/pelletier/go-toml"
	"permissionManage/perm-server/sys-common/conf"
	"time"
)

// NewRedisPool 返回redis连接池
func NewRedisPool() *redis.Pool {
	configTree := conf.Conf.Get("redis").(*toml.Tree)
	var (
		redisURL            = "redis://" + configTree.Get("Addr").(string)
		redisMaxIdle        = int(configTree.Get("RedisMaxIdle").(int64)) //最大空闲连接数
		redisIdleTimeoutSec = int64(configTree.Get("RedisIdleTimeoutSec").(int64))
		redisPassword       = configTree.Get("Password").(string)
		dB                  = int(configTree.Get("DB").(int64))
	)
	return &redis.Pool{
		MaxIdle: redisMaxIdle,
		//IdleTimeout: redisIdleTimeoutSec * time.Second,
		Dial: func() (redis.Conn, error) {
			redis.DialConnectTimeout(time.Duration(redisIdleTimeoutSec))
			redis.DialDatabase(dB)
			c, err := redis.DialURL(redisURL)
			if err != nil {
				return nil, fmt.Errorf("redis connection error: %s", err)
			}
			//验证redis密码
			if _, authErr := c.Do("AUTH", redisPassword); authErr != nil {
				return nil, fmt.Errorf("redis auth password error: %s", authErr)
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			if err != nil {
				return fmt.Errorf("ping redis error: %s", err)
			}
			return nil
		},
	}
}

func Set(k, v string) {
	c := NewRedisPool().Get()
	defer c.Close()
	_, err := c.Do("SET", k, v)
	if err != nil {
		fmt.Println("set error", err.Error())
	}
}
func SetKVExpire(k, v string, ex int) {
	c := NewRedisPool().Get()
	defer c.Close()
	_, err := c.Do("EXPIRE", k, v, ex)
	if err != nil {
		fmt.Println("set error", err.Error())
	}
}

func GetStringValue(k string) string {
	c := NewRedisPool().Get()
	defer c.Close()
	username, err := redis.String(c.Do("GET", k))
	if err != nil {
		fmt.Println("Get Error: ", err.Error())
		return ""
	}
	return username
}

func SetKeyExpire(k string, ex int) {
	c := NewRedisPool().Get()
	defer c.Close()
	_, err := c.Do("EXPIRE", k, ex)
	if err != nil {
		fmt.Println("set error", err.Error())
	}
}

func CheckKey(k string) bool {
	c := NewRedisPool().Get()
	defer c.Close()
	exist, err := redis.Bool(c.Do("EXISTS", k))
	if err != nil {
		fmt.Println(err)
		return false
	} else {
		return exist
	}
}

func DelKey(k string) error {
	c := NewRedisPool().Get()
	defer c.Close()
	_, err := c.Do("DEL", k)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func SetJson(k string, data interface{}) error {
	c := NewRedisPool().Get()
	defer c.Close()
	value, _ := json.Marshal(data)
	n, _ := c.Do("SETNX", k, value)
	if n != int64(1) {
		return errors.New("set failed")
	}
	return nil
}

func GetJsonByte(k string) ([]byte, error) {
	c := NewRedisPool().Get()
	jsonGet, err := redis.Bytes(c.Do("GET", k))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return jsonGet, nil
}
