package rest

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"library/libraryDemo/model"
	"library/libraryDemo/service"
	"strconv"
)

type book int

var Book book

func (book) list(c *gin.Context) {
	libraryid, err := strconv.ParseInt(c.Query("libraryId"), 10, 64)
	if err != nil {
		c.String(400, "id 参数错误")
		c.Abort()
		return
	}
	page := &model.Page{}
	if err := c.Bind(page); err != nil {
		c.String(400, "id 参数错误")
		c.Abort()
		return
	}
	listbooks := &[]model.Books{}
	if err := service.Books.List(libraryid, page, listbooks); err != nil {
		c.String(500, "id 参数错误")
		c.Abort()
		return
	}
	r := map[string]interface{}{}
	r["data"] = listbooks
	c.JSON(200, r)

}
func (book) get(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.String(400, "id 参数错误")
		c.Abort()
		return
	}
	form := &model.Books{}
	form.ID = id
	if err := service.Books.Get(form); err != nil {
		c.String(500, "id 参数错误")
		c.Abort()
		return
	}
	c.JSON(200, form)
}
func (book) put(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.String(400, "id 参数错误")
		c.Abort()
		return
	}
	fmt.Print(id)
}
func (book) delete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {

		c.String(400, "id 参数错误")
		c.Abort()
		return
	}
	fmt.Print(id)

}
func (book) save(c *gin.Context) {

}
func (book) updatebook(c *gin.Context) {
	form := &model.Updatebook{}
	if err := c.Bind(form); err != nil {
		fmt.Println(err)
		c.String(400, "id 参数错误")
		c.Abort()
		return
	}
	if err := service.Books.UpdateBooks(form); err != nil {
		c.String(500, "id 参数错误")
		c.Abort()
		return
	}
	r := map[string]interface{}{}
	r["data"] = "ok"
	c.JSON(200, r)
}
func (book) Register(r *gin.RouterGroup) {
	r.POST("/v1/book/update", Book.updatebook)
	r.GET("/v1/book", Book.list)
	r.GET("/v1/book/:id", Book.get)
	r.PUT("/v1/book/:id", Book.put)
	r.DELETE("/v1/book/:id", Book.delete)
	r.POST("/v1/book", Book.save)
}
