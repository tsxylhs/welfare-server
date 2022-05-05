package rest

import (
	"log"
	"lottery/welfare/model"
	"lottery/welfare/service"

	"github.com/gin-gonic/gin"
)

type lucky int

var TypeMap = map[string]string{"1": "双色球", "2": "3D", "3": "七乐彩"}
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

func (lucky) save(c *gin.Context) {
	lucky := &model.LuckyData{}
	if err := c.Bind(lucky); err != nil {
		log.Fatal("err:", err)
	}
	if err := service.Lucky.Save(lucky); err != nil {
		log.Fatal("error:", err)

	}

	c.JSON(200, "插入成功")
}
func (lucky) list(c *gin.Context) {
	page := &model.Page{}
	if err := c.Bind(page); err != nil {
		c.String(400, "id 参数错误")
		c.Abort()
		return
	}
	list := &[]model.LuckyDataV{}
	if err := service.Lucky.List(page, list); err != nil {
		c.String(500, "id 参数错误")
		c.Abort()
		return
	}

	listv := []model.LuckyDataV{}
	for _, val := range *list {
		v := model.LuckyDataV{}
		v.Base = val.Base
		v.LuckyData = val.LuckyData
		v.Type = val.Type
		v.TypeName = typeName(val.Type)
		v.Crt = v.Base.CreatedAt.Format("2006-01-02 15:04:05")
		listv = append(listv, v)
	}

	r := map[string]interface{}{}
	r["data"] = listv
	c.JSON(200, r)

}
func typeName(ty string) string {
	if TypeMap[ty] != "" {
		return TypeMap[ty]
	}
	return ""
}

func (lucky) Register(r *gin.RouterGroup) {
	r.GET("/v1/generateLuckyData", Lucky.generateLuckyData)
	r.POST("/v1/lucky", Lucky.save)
	r.GET("/v1/lucky", Lucky.list)

}
