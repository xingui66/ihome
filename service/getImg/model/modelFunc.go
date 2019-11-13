package model

import (
	"github.com/gomodule/redigo/redis"
)

var RedisPool redis.Pool

func InitRedis(){
	RedisPool = redis.Pool{
		MaxIdle:20,
		MaxActive:50,
		IdleTimeout:60 * 5,
		Dial: func() (redis.Conn, error) {
			//return redis.Dial("tcp","192.168.11.47:6379")
			return redis.Dial("tcp","127.0.0.1:6379")
		},
	}

}

func SaveImgRnd(uuid,rnd string)error{
	conn:=RedisPool.Get()
	_,err := conn.Do("setex",uuid,60 * 5,rnd)
	return err
}


