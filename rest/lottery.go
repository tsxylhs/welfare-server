package rest

import (
	"fmt"
	"lottery/welfare/model"
	"lottery/welfare/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type lottery int

var Lottery lottery

func (lottery) list(c *gin.Context) {
	page := &model.Page{}
	if err := c.Bind(page); err != nil {
		c.String(400, "id 参数错误")
		c.Abort()
		return
	}
	listlotterys := &[]model.Lottery{}
	if err := service.Lotterys.List(page, listlotterys); err != nil {
		c.String(500, "id 参数错误")
		c.Abort()
		return
	}
	r := map[string]interface{}{}
	r["data"] = listlotterys
	c.JSON(200, r)

}
func (lottery) get(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.String(400, "id 参数错误")
		c.Abort()
		return
	}
	form := &model.Lottery{}
	form.ID = id
	if err := service.Lotterys.Get(form); err != nil {
		c.String(500, "id 参数错误")
		c.Abort()
		return
	}
	c.JSON(200, form)
}
func (lottery) put(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.String(400, "id 参数错误")
		c.Abort()
		return
	}
	fmt.Print(id)
}
func (lottery) delete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {

		c.String(400, "id 参数错误")
		c.Abort()
		return
	}
	fmt.Print(id)

}
func (lottery) save(c *gin.Context) {

}
func (lottery) updatelottery(c *gin.Context) {

}

//彩票店
func (lottery) Register(r *gin.RouterGroup) {
	r.POST("/v1/lottery/update", Lottery.updatelottery)
	r.GET("/v1/lottery", Lottery.list)
	r.GET("/v1/lottery/:id", Lottery.get)
	r.PUT("/v1/lottery/:id", Lottery.put)
	r.DELETE("/v1/lottery/:id", Lottery.delete)
	r.POST("/v1/lottery", Lottery.save)

}
