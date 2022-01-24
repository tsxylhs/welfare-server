package rest

import (
	"log"
	"lottery/welfare/model"
	"lottery/welfare/reqUtils"
	"lottery/welfare/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

//开奖查询
type lotteryOpenQuery int

var LotteryOpenQuery lotteryOpenQuery

func (lotteryOpenQuery) list(c *gin.Context) {
	page := &model.Page{}
	if err := c.Bind(page); err != nil {
		c.String(400, "id 参数错误")
		c.Abort()
		return
	}
	listlotteryOpenQuerys := &[]model.LotteryOpenQuery{}
	if err := service.LotteryOpenQuerys.List(page, listlotteryOpenQuerys); err != nil {
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
func (lotteryOpenQuery) open(c *gin.Context) {
	//开奖查询
	selectQuery := &model.SelectQuery{}
	if err := c.BindJSON(selectQuery); err != nil {
		log.Println("err", err.Error())
		c.String(400, "参数错误！")
	}
	err, form := service.LotteryOpenQuerys.Query(selectQuery)
	if err != nil {
		c.String(500, "查询错误！")
	}
	if form != nil {
		c.JSON(200, (*form)[0])
	} else {
		res := reqUtils.LotteryReq.LotteryQuery(selectQuery)

		if err := service.LotteryOpenQuerys.Save(&res.LotteryOpenQuery); err != nil {
			log.Println("err", err.Error())
			c.String(500, "查询接口入库失败")
		}
		c.JSON(200, res.LotteryOpenQuery)
	}
}

//开奖查询
func (lotteryOpenQuery) Register(r *gin.RouterGroup) {
	r.GET("/v1/lotteryOpenQuery", LotteryOpenQuery.list)
	r.GET("/v1/lotteryOpenQuery/:id", LotteryOpenQuery.get)
	r.POST("/v1/lotteryOpenQuery", LotteryOpenQuery.open)
}
