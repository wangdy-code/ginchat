package service

import (
	"ginchat/models"
	"github.com/gin-gonic/gin"
)

// GetUserList
// @Tags 获取用户列表
// @Success 200 {json} json{"code","message"}
// @Router /user/getUserList [get]
func GetUserList(c *gin.Context) {
	userList := make([]*models.UserBasic, 10)
	userList, err := models.UserBasic{}.GetUserList()
	if err != nil {
		c.JSON(500, gin.H{
			"message": "获取失败",
		})
	}
	c.JSON(200, gin.H{
		"message": userList,
	})
}

// Create
// @Tags 创建用户
// @Success 200 {json} json{"code","message"}
// params
// @Router /user/create [post]
func Create(c *gin.Context) {
	user := models.UserBasic{}
	err := c.ShouldBind(&user)
	userList, err := models.UserBasic{}.GetUserList()
	if err != nil {
		c.JSON(500, gin.H{
			"message": "获取失败",
		})
	}
	c.JSON(200, gin.H{
		"message": userList,
	})
}
