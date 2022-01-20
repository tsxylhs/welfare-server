package rest

import (
	"fmt"
	"log"
	"lottery/welfare/model"
	"lottery/welfare/service"
	"path"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

type lottery int

var Lottery lottery

func (lottery) list(c *gin.Context) {
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
	listlotterys := &[]model.Lottery{}
	if err := service.Lotterys.List(libraryid, page, listlotterys); err != nil {
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
func (lottery) uploadFile(c *gin.Context) {
	_, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		c.String(400, "文件读取失败")
	}
	log.Print("上传的文件：", fileHeader.Filename)
	//ex
	dst := path.Join("./upload", fileHeader.Filename)
	err = c.SaveUploadedFile(fileHeader, dst)
	if err != nil {
		c.String(500, "文件存储失败！")
	}

	excelizeFile, err := excelize.OpenFile(dst)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := excelizeFile.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	rows, err := excelizeFile.GetRows("sheet1")
	if err != nil {
		log.Fatal(err)
	}
	for index, values := range rows {
		if index == 0 {
			continue
		}
		for i, val := range values {
			fmt.Println("i", i)
			fmt.Println("val", val)
		}
	}
	c.JSON(200, "success")

}

//彩票店
func (lottery) Register(r *gin.RouterGroup) {
	r.POST("/v1/lottery/update", Lottery.updatelottery)
	r.GET("/v1/lottery", Lottery.list)
	r.GET("/v1/lottery/:id", Lottery.get)
	r.PUT("/v1/lottery/:id", Lottery.put)
	r.DELETE("/v1/lottery/:id", Lottery.delete)
	r.POST("/v1/lottery", Lottery.save)
	r.POST("/v1/lottery/uploadFile", Lottery.uploadFile)

}
