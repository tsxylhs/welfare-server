package wechat

import (
	"github.com/gin-gonic/gin"
	"library/libraryDemo/model"
	"library/libraryDemo/service"
	"strconv"
)

type user int

var User user

func (user) login(c *gin.Context) {
	form := &model.User{}
	if err := c.Bind(form); err != nil {
		c.String(400, "参数错误")
		c.Abort()
	}
	if err := service.User.Login(form); err != nil {

		c.String(500, "内部服务器错误")
		c.Abort()
		return
	}
	c.JSON(200, form)
}
func (user) get(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.String(400, "id 参数错误")
		c.Abort()
		return
	}
	form := &model.User{}
	form.ID = id
	if err := service.User.Get(form); err != nil {
		c.String(400, "id 参数错误")
		c.Abort()
		return
	}
	c.JSON(200, form)
}
func (user) put(c *gin.Context) {
	form := &model.User{}
	err := c.Bind(form)
	if err != nil {
		c.String(400, "id 参数错误")
		c.Abort()
		return
	}
	if err := service.User.Update(form); err != nil {
		c.String(400, "id 参数错误")
		c.Abort()
		return
	}
	c.JSON(200, form)
}
func (user) Register(r *gin.RouterGroup) {
	r.POST("/v1/user/login", User.login)
	r.PUT("/v1/user", User.put)
	r.GET("/v1/user/:id", User.get)
}
