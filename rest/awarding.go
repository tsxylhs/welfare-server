package rest

// 兑奖
import (
	"fmt"
	"lottery/welfare/model"
	"lottery/welfare/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type awarding int

var Awarding awarding

func (awarding) list(c *gin.Context) {
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
	listawardings := &[]model.Awarding{}
	if err := service.Awardings.List(libraryid, page, listawardings); err != nil {
		c.String(500, "id 参数错误")
		c.Abort()
		return
	}
	r := map[string]interface{}{}
	r["data"] = listawardings
	c.JSON(200, r)

}
func (awarding) get(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.String(400, "id 参数错误")
		c.Abort()
		return
	}
	form := &model.Awarding{}
	form.ID = id
	if err := service.Awardings.Get(form); err != nil {
		c.String(500, "id 参数错误")
		c.Abort()
		return
	}
	c.JSON(200, form)
}
func (awarding) put(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.String(400, "id 参数错误")
		c.Abort()
		return
	}
	fmt.Print(id)
}
func (awarding) delete(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {

		c.String(400, "id 参数错误")
		c.Abort()
		return
	}
	fmt.Print(id)

}
func (awarding) save(c *gin.Context) {

}
func (awarding) updateawarding(c *gin.Context) {

}
func (awarding) Register(r *gin.RouterGroup) {
	r.POST("/v1/awarding/update", Awarding.updateawarding)
	r.GET("/v1/awarding", Awarding.list)
	r.GET("/v1/awarding/:id", Awarding.get)
	r.PUT("/v1/awarding/:id", Awarding.put)
	r.DELETE("/v1/awarding/:id", Awarding.delete)
	r.POST("/v1/awarding", Awarding.save)
}