package router

import (
	"ginchat/docs"
	"ginchat/service"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.GET("/index", service.GetIndex)
	r.GET("/user/getUserList", service.GetUserList)
	r.POST("/user/create", service.CreateUser)
	r.PUT("/user/update", service.UpdateUser)
	r.DELETE("/user/delete", service.DeleteUser)

	return r
}
