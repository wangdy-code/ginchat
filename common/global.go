package common

import (
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	VP *viper.Viper
	DB *gorm.DB
)
