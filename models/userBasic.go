package models

import (
	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
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
