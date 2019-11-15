package model

import "github.com/gomodule/redigo/redis"

var RedisPool redis.Pool
//校验短信验证码是否正确
func GetSmsCode(phone string)(string,error){
	//获取redis链接
	conn := RedisPool.Get()
	//获取数据
	return redis.String(conn.Do("get",phone+"_code"))
}

func SaveUser(mobile,password string)error{
	var user User
	user.Mobile=mobile
	user.Password_hash = password
	user.Name=mobile
	return GlobalDB.Create(&user).Error
}
//校验登录信息
func CheckUser(mobile,pwd_hash string)(User,error){
	//连接数据库

	var user User
	err := GlobalDB.Where("mobile = ?",mobile).Where("password_hash = ?",pwd_hash).Find(&user).Error

	return user,err
}



