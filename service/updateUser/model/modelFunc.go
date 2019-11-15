package model


func GetUserInfo(userName string )(User,error){
	var user User
	err:=GlobalDB.Where("name=?",userName).Find(&user).Error
	return user,err
}




