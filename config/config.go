package config

import (
	"ginchat/common"
)

type Mysql struct {
	Ip       string `json:"ip"`
	Port     string `json:"port"`
	Database string `json:"database"`
	User     string `json:"user"`
	PassWord string `json:"password"`
}

func GetMysqlConfig() (*Mysql, error) {
	mapConfig := ReadSection("mysql")
	mysqlConfig := Mysql{}
	mysqlConfig.Ip = mapConfig["ip"]
	mysqlConfig.Port = mapConfig["port"]
	mysqlConfig.Database = mapConfig["database"]
	mysqlConfig.User = mapConfig["user"]
	mysqlConfig.PassWord = mapConfig["password"]
	return &mysqlConfig, nil
}
func ReadSection(name string) map[string]string {
	return common.VP.GetStringMapString(name)
}
