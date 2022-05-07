package rest

import (
	"fmt"
	"log"
	"lottery/welfare/model"
	"lottery/welfare/service"
	"strconv"
	"strings"

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
		v.WinningAmount = val.WinningAmount
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
func (lucky) checkLucky(c *gin.Context) {
	checkLucky := &model.CheckLuckyVo{}
	if err := c.Bind(checkLucky); err != nil {
		c.String(400, "id 参数错误")
		c.Abort()
		return
	}
	//获取改期号码
	luckyData := &[]model.LuckyDataV{}
	val, _ := strconv.ParseInt(checkLucky.LuckyDataID, 10, 64)
	if err := service.Lucky.Get(val, luckyData); err != nil {
		c.String(500, "id 参数错误")
		c.Abort()
		return
	}
	if len(*luckyData) > 0 {
		//构造参数
		if (*luckyData)[0].Type == "1" {
			luckyDataVal := (*luckyData)[0].LuckyData[1 : len((*luckyData)[0].LuckyData)-1]
			strs := strings.Split(luckyDataVal, ",")
			temp_strs := ""
			for i := 0; i < len(strs); i++ {
				temp_str := ""
				v, _ := strconv.Atoi(strs[i])
				if v < 10 {
					temp_str = "0" + strs[i]
				} else {
					temp_str = strs[i]
				}
				if i >= (len(strs) - 1) {
					temp_strs = temp_strs + temp_str
				} else {
					temp_strs = temp_strs + temp_str + ","
				}

			}

			des_temp := strings.Replace(temp_strs, ",", "@", -1)
			des := strings.Replace(des_temp, "@", ",", 5)
			fmt.Print("des" + des)
			selectlottery := &model.SelectLottery{}
			selectlottery.LotteryType = "ssq"
			selectlottery.LotteryOpenNo = checkLucky.Issue
			selectlottery.LotteryNum = des
			// awarding := reqUtils.LotteryReq.LotteryBonus(selectlottery)
			// if awarding == nil {
			// 	c.String(500, "服务器错误")
			// 	c.Abort()
			// 	return
			// }
			// awarding.UserId = (*luckyData)[0].UserId
			// err := service.Awardings.Save(awarding)

			// if err != nil {
			// 	log.Fatal("insert error", err.Error())
			// }
			if err := service.Lucky.Update(&(*luckyData)[0]); err != nil {
				log.Fatal("insert error ", err.Error())
			}
			r := map[string]interface{}{}
			r["data"] = (*luckyData)[0]
			c.JSON(200, r)
		}
		c.String(500, "暂不支持该彩种")
		c.Abort()
		return

	}

}

func (lucky) Register(r *gin.RouterGroup) {
	r.GET("/v1/generateLuckyData", Lucky.generateLuckyData)
	r.POST("/v1/lucky", Lucky.save)
	r.GET("/v1/lucky", Lucky.list)
	r.POST("/v1/lucky/check", Lucky.checkLucky)

}
