package rest

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"library/libraryDemo/model"
	"library/libraryDemo/service"
	"strconv"
)

type library int

var Library library

func (library) list(c *gin.Context) {
	library := &model.Library{}
	librarys := &[]model.Library{}
	page := &model.Page{}
	if err := c.Bind(page); err != nil {
		fmt.Println(err)
	}

	service.Library.List(library, page, librarys)
	r := map[string]interface{}{}
	r["data"] = librarys
	c.JSON(200, r)
}
func (library) get(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.String(400, "id 参数错误")
		c.Abort()
		return
	}
	library := &model.Library{}
	library.ID = id
	if err := service.Library.Get(library); err != nil {
		c.String(500, "id 参数错误")
		c.Abort()
		return
	}
	c.JSON(200, library)
}
func (library) put(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.String(400, "id 参数错误")
		c.Abort()
		return
	}
	fmt.Print(id)
}
func (library) delete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {

		c.String(400, "id 参数错误")
		c.Abort()
		return
	}
	fmt.Print(id)

}
func (library) save(c *gin.Context) {

}
func (library) Register(r *gin.RouterGroup) {
	r.GET("/v1/library", Library.list)
	r.GET("/v1/library/:id", Library.get)
	r.PUT("/v1/library/:id", Library.put)
	r.DELETE("/v1/library/:id", Library.delete)
	r.POST("/v1/library", Library.save)
}
