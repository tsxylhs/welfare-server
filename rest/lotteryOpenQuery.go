package rest

import (
	"fmt"
	"lottery/welfare/model"
	"lottery/welfare/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type lotteryOpenQuery int

var LotteryOpenQuery lotteryOpenQuery

func (lotteryOpenQuery) list(c *gin.Context) {
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
	listlotteryOpenQuerys := &[]model.LotteryOpenQuery{}
	if err := service.LotteryOpenQuerys.List(libraryid, page, listlotteryOpenQuerys); err != nil {
		c.String(500, "id 参数错误")
		c.Abort()
		return
	}
	r := map[string]interface{}{}
	r["data"] = listlotteryOpenQuerys
	c.JSON(200, r)

}
func (lotteryOpenQuery) get(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.String(400, "id 参数错误")
		c.Abort()
		return
	}
	form := &model.LotteryOpenQuery{}
	form.ID = id
	if err := service.LotteryOpenQuerys.Get(form); err != nil {
		c.String(500, "id 参数错误")
		c.Abort()
		return
	}
	c.JSON(200, form)
}
func (lotteryOpenQuery) put(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.String(400, "id 参数错误")
		c.Abort()
		return
	}
	fmt.Print(id)
}
func (lotteryOpenQuery) delete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {

		c.String(400, "id 参数错误")
		c.Abort()
		return
	}
	fmt.Print(id)

}
func (lotteryOpenQuery) save(c *gin.Context) {

}
func (lotteryOpenQuery) updatelotteryOpenQuery(c *gin.Context) {

}
func (lotteryOpenQuery) Register(r *gin.RouterGroup) {
	r.POST("/v1/lotteryOpenQuery/update", LotteryOpenQuery.updatelotteryOpenQuery)
	r.GET("/v1/lotteryOpenQuery", LotteryOpenQuery.list)
	r.GET("/v1/lotteryOpenQuery/:id", LotteryOpenQuery.get)
	r.PUT("/v1/lotteryOpenQuery/:id", LotteryOpenQuery.put)
	r.DELETE("/v1/lotteryOpenQuery/:id", LotteryOpenQuery.delete)
	r.POST("/v1/lotteryOpenQuery", LotteryOpenQuery.save)
}
