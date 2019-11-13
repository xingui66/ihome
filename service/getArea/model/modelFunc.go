package model

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"

)

func QueryArea()([]Area,error){
	var areas [] Area

	//从redis中获取数据
	conn:=GlobalRedis.Get()
	defer conn.Close()

	//获取redis里的字节切片
	areasByte,_:= redis.Bytes(conn.Do("get","areaData"))
	if len(areasByte)==0{
		//从mysql数据库中获取数据
		if err := GlobalDB.Find(&areas).Error;err != nil {
			return areas,err
		}
        //序列化
        areasJson,err := json.Marshal(areas);
        if err !=nil {
        	return nil,err
		}
		//把序列的字节切片  保存到到redis里
		_,err = conn.Do("set","areasData",areasJson);
		if err !=nil{
			fmt.Println("conn.do 序列的字节切片,没有保存到redis里,有错误  err:",err)
		}

	}else {
		err :=json.Unmarshal(areasByte,&areas)
		if err!=nil {
			fmt.Println("将redis获取到的字节切片,没有反序列成areas json.Unmarshal,有错误 err :",err)
		}

	}

	return areas,nil
}


