package main

import (
	"ginchat/router"
	"ginchat/utils"
	"log"
)

func init() {
	err := utils.InitConfig()
	if err != nil {
		log.Fatalf("Init fail err: %s", err)
		return
	}
	utils.InitMysql()
}
func main() {
	r := router.Router()
	r.Run()
}
