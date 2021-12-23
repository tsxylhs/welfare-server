package rest

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"library/libraryDemo/model"
	"library/libraryDemo/service"
	"strconv"
)

type table int

var Table table

func (table) list(c *gin.Context) {
	id, err := strconv.ParseInt(c.Query("id"), 10, 64)
	if err != nil {
		c.String(400, "id 参数错误")
		c.Abort()
		return
	}
	form := &model.Table{}
	form.ID = id
	list := &[]model.Table{}
	page := &model.Page{}
	if err := service.Table.List(form, page, list); err != nil {
		c.String(500, "id 参数错误")
		c.Abort()
		return
	}
	r := map[string]interface{}{}
	r["data"] = list
	c.JSON(200, r)
}
func (table) get(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.String(400, "id 参数错误")
		c.Abort()
		return
	}
	fmt.Print(id)
}
func (table) put(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.String(400, "id 参数错误")
		c.Abort()
		return
	}
	fmt.Print(id)
}
func (table) delete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {

		c.String(400, "id 参数错误")
		c.Abort()
		return
	}
	fmt.Print(id)

}
func (table) save(c *gin.Context) {

}
func (table) Register(r *gin.RouterGroup) {
	r.GET("/v1/table", Table.list)
	r.GET("/v1/table/:id", Table.get)
	r.PUT("/v1/table/:id", Table.put)
	r.DELETE("/v1/table/:id", Table.delete)
	r.POST("/v1/table", Table.save)
}
