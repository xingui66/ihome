package model

import "github.com/garyburd/redigo/redis"

var RedisPool redis.Pool

func InitRedis(){
	RedisPool =redis.Pool{
		MaxIdle:20,
		MaxActive:50,
		IdleTimeout:60 * 5,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp","127.0.0.1:6379")
		},
	}
}

//获取图片验证码
func GetImgCode(uuid string)(string,error){
	//获取redis链接
	conn := RedisPool.Get()
	//获取数据
	return redis.String(conn.Do("get",uuid))
}

func SaveSmsCode(phone,vcode string)error{
	conn:=RedisPool.Get()
	_,err :=conn.Do("setex",phone+"_code",60*5,vcode)
	return err
}

