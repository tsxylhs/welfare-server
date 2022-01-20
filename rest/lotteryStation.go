package rest

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"log"
	"lottery/welfare/model"
	"lottery/welfare/service"
	"path"
	"strconv"

	"github.com/gin-gonic/gin"
)

type lotteryStation int

var LotteryStation lotteryStation

func (lotteryStation) list(c *gin.Context) {

	page := &model.Page{}
	if err := c.BindJSON(page); err != nil {
		c.String(400, "id 参数错误")
		c.Abort()
		return
	}
	listlotteryStations := &[]model.LotteryStation{}
	if err := service.LotteryStations.List(page, listlotteryStations); err != nil {
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
		lotteryStation := &model.LotteryStation{}
		for i, val := range values {
			switch i {
			case 0:
				lotteryStation.Name = val
				break
			case 2:
				lotteryStation.Location = val
				break
			case 3:
				lotteryStation.ImageUrl = val
				break
			case 4:
				lotteryStation.Mobile = val
				break
			default:
				fmt.Println("i", i)
				fmt.Println("val", val)
			}

		}
		if err := service.LotteryStations.Save(lotteryStation); err != nil {
			log.Print(err.Error())
			log.Print("入库失败")
		}
	}
	c.JSON(200, "success")

}

func (lotteryStation) Register(r *gin.RouterGroup) {
	r.POST("/v1/lotteryStation/update", LotteryStation.updatelotteryStation)
	r.GET("/v1/lotteryStation", LotteryStation.list)
	r.GET("/v1/lotteryStation/:id", LotteryStation.get)
	r.PUT("/v1/lotteryStation/:id", LotteryStation.put)
	r.DELETE("/v1/lotteryStation/:id", LotteryStation.delete)
	r.POST("/v1/lotteryStation", LotteryStation.save)
	r.POST("/v1/lotteryStation/uploadFile", Lottery.uploadFile)
}
