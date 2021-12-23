package rest

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"library/libraryDemo/model"
	"library/libraryDemo/service"
	"strconv"
)

type mybooks int

var Mybooks mybooks

func (mybooks) list(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Query("userId"), 10, 64)
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
	listbooks := &[]model.MyBook{}
	if err := service.Mybooks.List("salse", userId, page, listbooks); err != nil {
		c.String(500, "id 参数错误")
		c.Abort()
		return
	}
	r := map[string]interface{}{}
	r["data"] = listbooks
	c.JSON(200, r)

}
func (mybooks) get(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.String(400, "id 参数错误")
		c.Abort()
		return
	}
	form := &model.MyBook{}
	form.ID = id
	if err := service.Mybooks.Get(form); err != nil {
		c.String(500, "id 参数错误")
		c.Abort()
		return
	}
	c.JSON(200, form)
}
func (mybooks) put(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.String(400, "id 参数错误")
		c.Abort()
		return
	}
	fmt.Print(id)
}
func (mybooks) delete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {

		c.String(400, "id 参数错误")
		c.Abort()
		return
	}
	form := &model.MyBook{}
	form.ID = id
	if err := service.Mybooks.Delete(form); err != nil {
		c.String(500, "id 参数错误")
		c.Abort()
		return
	}
	c.JSON(200, form)

}
func (mybooks) save(c *gin.Context) {

}
func (mybooks) updatebook(c *gin.Context) {
	form := &model.Updatebook{}
	if err := c.Bind(form); err != nil {
		fmt.Println(err)
		c.String(400, "id 参数错误")
		c.Abort()
		return
	}
	if err := service.Mybooks.UpdateBooks(form); err != nil {
		c.String(500, "id 参数错误")
		c.Abort()
		return
	}
	r := map[string]interface{}{}
	r["data"] = "ok"
	c.JSON(200, r)
}
func (mybooks) applySave(c *gin.Context) {
	apply := &model.Apply{}
	if err := c.Bind(apply); err != nil {
		c.String(400, "id 参数错误")
		c.Abort()
		return
	}
	if err := service.Mybooks.SaveApply(apply); err != nil {
		c.String(500, "id 参数错误")
		c.Abort()
		return
	}
	c.JSON(200, "ok")
}
func (mybooks) applylibrary(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.String(400, "id 参数错误")
		c.Abort()
		return
	}
	librarys := &[]model.Library{}
	if err := service.Mybooks.ApplyLibrary(id, librarys); err != nil {
		c.String(500, "id 参数错误")
		c.Abort()
		return
	}
	r := map[string]interface{}{}
	r["data"] = librarys
	c.JSON(200, r)
}
func (mybooks) applylist(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Query("userId"), 10, 64)
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
	applys := &[]model.ApplysVo{}
	if err := service.Mybooks.Applylist(userId, applys); err != nil {
		c.String(400, "id 参数错误")
		c.Abort()
		return
	}
	r := map[string]interface{}{}
	r["data"] = applys
	c.JSON(200, r)
}
func (mybooks) deleteApply(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {

		c.String(400, "id 参数错误")
		c.Abort()
		return
	}
	if err := service.Mybooks.ApplyDelete(id); err != nil {
		c.String(400, "id 参数错误")
		c.Abort()
		return
	}
	c.JSON(200, "ok")
}
func (mybooks) Register(r *gin.RouterGroup) {
	r.POST("/v1/mybooks/update", Mybooks.updatebook)
	r.GET("/v1/mybooks", Mybooks.list)
	r.GET("/v1/mybooks/:id", Mybooks.get)
	r.PUT("/v1/mybooks/:id", Mybooks.put)
	r.DELETE("/v1/mybooks/:id", Mybooks.delete)
	r.POST("/v1/mybooks", Mybooks.save)
	r.POST("/v1/apply", Mybooks.applySave)
	r.GET("/v1/apply", Mybooks.applylist)
	r.DELETE("/v1/apply/:id", Mybooks.deleteApply)
	r.GET("/v1/applylibrary/:id", Mybooks.applylibrary)
}
