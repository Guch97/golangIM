package service

import (
	"fmt"
	"ginchat/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetUserList
// @Summary 所有用户
// @Tags 用户模块
// @Success 200 {string} json{"code","message"}
// @Router /user/GetUserList [get]
func GetUserList(c *gin.Context) {
	data := make([]*models.UserBasic, 10)
	data = models.GetUserList()
	c.JSON(http.StatusOK, gin.H{
		"msg":  "登录成功",
		"code": 200,
		"data": data,
	})
}

// CreateUser
// @Summary 新增用户
// @param  name query string  false "用户名"
// @param  passWord query string  false "密码"
// @Tags 用户模块
// @Success 200 {string} json{"code","message"}
// @Router /user/CreateUser [get]
func CreateUser(c *gin.Context) {
	user := models.UserBasic{}
	user.Name = c.Query("name")
	user.PassWord = c.Query("passWord")
	models.CreateUser(user)
	c.JSON(http.StatusOK, gin.H{
		"code":    "200",
		"message": "Success",
	})
}

// DeleteUser
// @Summary 删除用户
// @param  id query string  false "id"
// @Tags 用户模块
// @Success 200 {string} json{"code","message"}
// @Router /user/DeleteUser [delete]
func DeleteUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.Query("id"))
	fmt.Println(user, "user")
	user.ID = uint(id)
	models.DeleteUser(user)
	c.JSON(http.StatusOK, gin.H{
		"code":    "200",
		"message": "删除成功",
	})
}

// UpdateUser
// @Summary 更新用户
// @param  id formData string  false "id"
// @param  passWord formData string  false "passWord"
// @param  name formData string  false "name"
// @Tags 用户模块
// @Success 200 {string} json{"code","message"}
// @Router /user/UpdateUser [post]
func UpdateUser(c *gin.Context) {
	user := models.UserBasic{}
	ids := c.Request.FormValue("id")
	id, _ := strconv.Atoi(ids)
	user.ID = uint(id)
	user.Name = c.PostForm("name")
	user.PassWord = c.PostForm("passWord")
	fmt.Println(user, "user321321321")
	models.UpdateUser(user)

	c.JSON(http.StatusOK, gin.H{
		"code":    "200",
		"message": "修改成功",
	})
}
