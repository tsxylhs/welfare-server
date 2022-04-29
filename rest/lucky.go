package rest

import (
	"log"
	"lottery/welfare/model"
	"lottery/welfare/service"

	"github.com/gin-gonic/gin"
)

type lucky int

var Lucky lucky

func (lucky) generateLuckyData(c *gin.Context) {
	lucky := &model.LuckyDataVo{}
	if err := c.Bind(lucky); err != nil {
		log.Fatal("err:", err)
	}
	r := map[string]interface{}{}
	if err, val := service.Lucky.InsertLuckyData(lucky); err != nil {
		log.Fatal("error:", err)

	} else {
		r["data"] = val
	}

	c.JSON(200, r)
}

func (lucky) Register(r *gin.RouterGroup) {
	r.GET("/v1/generateLuckyData", Lucky.generateLuckyData)

}
