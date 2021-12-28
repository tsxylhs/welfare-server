package rest

import (
	"fmt"
	"library/libraryDemo/model"
	"library/libraryDemo/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type lotteryStation int

var LotteryStation lotteryStation

func (lotteryStation) list(c *gin.Context) {
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
	listlotteryStations := &[]model.LotteryStation{}
	if err := service.LotteryStations.List(libraryid, page, listlotteryStations); err != nil {
		c.String(500, "id 参数错误")
		c.Abort()
		return
	}
	r := map[string]interface{}{}
	r["data"] = listlotteryStations
	c.JSON(200, r)

}
func (lotteryStation) get(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.String(400, "id 参数错误")
		c.Abort()
		return
	}
	form := &model.LotteryStation{}
	form.ID = id
	if err := service.LotteryStations.Get(form); err != nil {
		c.String(500, "id 参数错误")
		c.Abort()
		return
	}
	c.JSON(200, form)
}
func (lotteryStation) put(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.String(400, "id 参数错误")
		c.Abort()
		return
	}
	fmt.Print(id)
}
func (lotteryStation) delete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {

		c.String(400, "id 参数错误")
		c.Abort()
		return
	}
	fmt.Print(id)

}
func (lotteryStation) save(c *gin.Context) {

}
func (lotteryStation) updatelotteryStation(c *gin.Context) {

}
func (lotteryStation) Register(r *gin.RouterGroup) {
	r.POST("/v1/lotteryStation/update", LotteryStation.updatelotteryStation)
	r.GET("/v1/lotteryStation", LotteryStation.list)
	r.GET("/v1/lotteryStation/:id", LotteryStation.get)
	r.PUT("/v1/lotteryStation/:id", LotteryStation.put)
	r.DELETE("/v1/lotteryStation/:id", LotteryStation.delete)
	r.POST("/v1/lotteryStation", LotteryStation.save)
}
