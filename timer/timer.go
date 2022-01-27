package timer

import (
	"log"
	"lottery/welfare/model"
	"lottery/welfare/reqUtils"
	"lottery/welfare/service"
)

func Timer() {
	selectQuery := &model.SelectQuery{}
	selectQuery.LotteryId = "ssq"
	selectQuery.LotteryNo = "22001"
	resp := reqUtils.LotteryReq.LotteryQuery(selectQuery)
	if resp == nil {
		log.Fatal("查询失败！")
	}
	if err := service.LotteryOpenQuerys.Save(&resp.LotteryOpenQuery); err != nil {
		log.Fatal("保存失败", err.Error())
	}
}
