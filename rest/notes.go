package rest

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"library/libraryDemo/model"
	"library/libraryDemo/service"
	"strconv"
)

type notes int

var Notes notes

func (notes) list(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Query("userId"), 10, 64)
	if err != nil {
		c.String(400, "id 参数错误")
		c.Abort()
		return
	}
	page := &model.Page{}
	list := &[]model.Notes{}
	if err := service.Notes.List(userId, page, list); err != nil {
		c.String(500, "id 参数错误")
		c.Abort()
		return
	}
	r := map[string]interface{}{}
	r["data"] = list
	c.JSON(200, r)

}
func (notes) get(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.String(400, "id 参数错误")
		c.Abort()
		return
	}

	form := &model.Notes{}
	form.ID = id
	if err := service.Notes.Get(form); err != nil {
		c.String(500, "id 参数错误")
		c.Abort()
		return
	}
	c.JSON(200, form)
}
func (notes) put(c *gin.Context) {
	form := &model.Notes{}
	if err := c.Bind(form); err != nil {
		c.String(500, "错误")
		c.Abort()
		return
	}
	if err := service.Notes.Update(form); err != nil {
		c.String(500, "错误")
		c.Abort()
		return
	}
	c.JSON(200, "ok")
}
func (notes) delete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {

		c.String(400, "id 参数错误")
		c.Abort()
		return
	}
	fmt.Print(id)
	c.JSON(200, "ok")
}
func (notes) save(c *gin.Context) {
	notes := &model.Notes{}
	if err := c.Bind(notes); err != nil {
		c.String(400, "id 参数错误")
		c.Abort()
		return
	}
	if err := service.Notes.Save(notes); err != nil {
		c.String(500, "id 参数错误")
		c.Abort()
		return
	}
	c.JSON(200, "ok")
}
func (notes) Register(r *gin.RouterGroup) {
	r.GET("/v1/notes", Notes.list)
	r.GET("/v1/notes/:id", Notes.get)
	r.PUT("/v1/notes/:id", Notes.put)
	r.DELETE("/v1/notes/:id", Notes.delete)
	r.POST("/v1/notes", Notes.save)
}
