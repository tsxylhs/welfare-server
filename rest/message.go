package rest

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type message int

var Message message

func (message) list(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.String(400, "id 参数错误")
		c.Abort()
		return
	}
	fmt.Print(id)

}
func (message) get(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.String(400, "id 参数错误")
		c.Abort()
		return
	}
	fmt.Print(id)
}
func (message) put(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.String(400, "id 参数错误")
		c.Abort()
		return
	}
	fmt.Print(id)
}
func (message) delete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {

		c.String(400, "id 参数错误")
		c.Abort()
		return
	}
	fmt.Print(id)

}
func (message) save(c *gin.Context) {

}
func (message) Register(r *gin.RouterGroup) {
	r.GET("/v1/message/:id/list", Message.list)
	r.GET("/v1/message/:id", Message.get)
	r.PUT("/v1/message/:id", Message.put)
	r.DELETE("/v1/message/:id", Message.delete)
	r.POST("/v1/message", Message.save)
}
