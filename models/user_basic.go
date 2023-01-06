package models

import (
	"ginchat/common"
	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	Name          string `json:"name"`
	PassWord      string `json:"pass_word"`
	Phone         string `json:"phone"`
	Email         string `json:"email"`
	Identity      string `json:"identity"`
	ClientIP      string `json:"client_ip"`
	ClientPort    string `json:"client_port"`
	LoginTime     uint64 `json:"login_time"`
	HeartbeatTime uint64 `json:"heartbeat_time"`
	LoginOutTime  uint64 `json:"login_out_time"`
	IsLogout      bool   `json:"is_logout"`
	DeviceInfo    string `json:"device_info"`
}

func (u *UserBasic) TableName() string {
	return "user_basic"
}
func (u UserBasic) GetUserList() ([]*UserBasic, error) {
	userList := make([]*UserBasic, 10)
	err := common.DB.Model(UserBasic{}).Where("id <> 0").Find(&userList).Error
	if err != nil {
		return nil, err
	}
	return userList, nil
}
func (u UserBasic) CreateUser(user UserBasic) error {
	return common.DB.Model(u).Create(&user).Error
}
