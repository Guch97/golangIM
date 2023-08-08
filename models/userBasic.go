package models

import (
	"fmt"
	"ginchat/utils"
	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	ID           uint
	Name         string
	PassWord     string
	Phone        string
	Email        string
	Identity     string
	ClientIp     string
	ClientPort   string
	LoadingTime  uint64
	HearTime     uint64
	LogOutTime   uint64
	IsLogOutTime bool
	DeviceInfo   string
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}
func GetUserList() []*UserBasic {
	data := make([]*UserBasic, 10)
	fmt.Println(data, "dsadsadsaddada")
	utils.DB.Find(&data)
	for _, v := range data {
		fmt.Println(v)
	}
	return data
}

func CreateUser(user UserBasic) *gorm.DB {
	result := utils.DB.Create(&user)
	return result
}

func DeleteUser(user UserBasic) *gorm.DB {
	result := utils.DB.Delete(&user)
	return result
}

func UpdateUser(user UserBasic) *gorm.DB {
	result := utils.DB.Updates(UserBasic{Name: user.Name, PassWord: user.PassWord})
	fmt.Println(result, "resultresultresultresult")
	return result
}
