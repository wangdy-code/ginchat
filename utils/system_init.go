package utils

import (
	"fmt"
	"ginchat/common"
	"ginchat/config"
	"ginchat/models"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitConfig() error {
	common.VP = viper.GetViper()
	common.VP.SetConfigName("app")
	common.VP.AddConfigPath("config")
	return common.VP.ReadInConfig()

}

func InitMysql() {
	mysqlConfig, err := config.GetMysqlConfig()
	if err != nil {
		panic("failed to read mysqlConfig")
	}

	DB, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", mysqlConfig.User, mysqlConfig.PassWord,
		mysqlConfig.Ip, mysqlConfig.Port, mysqlConfig.Database)), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB.AutoMigrate(&models.UserBasic{})
	//user := models.UserBasic{}
	//user.Name = "张三"
	//DB.Create(&user)
	//userres := models.UserBasic{}
	//DB.First(&userres, 2)
	//fmt.Println(userres)
	//DB.Model(&models.UserBasic{}).Update("PassWord", "1234")
}
