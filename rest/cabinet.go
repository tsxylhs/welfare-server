package rest

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type cabinet int

var Cabinet cabinet

func (cabinet) list(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.String(400, "id 参数错误")
		c.Abort()
		return
	}
	fmt.Print(id)

}
func (cabinet) get(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.String(400, "id 参数错误")
		c.Abort()
		return
	}
	fmt.Print(id)
}
func (cabinet) put(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.String(400, "id 参数错误")
		c.Abort()
		return
	}
	fmt.Print(id)
}
func (cabinet) delete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {

		c.String(400, "id 参数错误")
		c.Abort()
		return
	}
	fmt.Print(id)

}
func (cabinet) save(c *gin.Context) {

}
func (cabinet) Register(r *gin.RouterGroup) {
	r.GET("/v1/cabinet/:id/list", Cabinet.list)
	r.GET("/v1/cabinet/:id", Cabinet.get)
	r.PUT("/v1/cabinet/:id", Cabinet.put)
	r.DELETE("/v1/cabinet/:id", Cabinet.delete)
	r.POST("/v1/cabinet", Cabinet.save)
}
