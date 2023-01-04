package utils

import (
	"fmt"
	"ginchat/common"
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
	fmt.Println(common.VP.Get("mysql"))
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("root:root@tcp(127.0.0.1:3306)/ginchat")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&models.UserBasic{})
	//user := models.UserBasic{}
	//user.Name = "张三"
	//db.Create(&user)
	userres := models.UserBasic{}
	db.First(&userres, 2)
	fmt.Println(userres)
	db.Model(&models.UserBasic{}).Update("PassWord", "1234")
}
