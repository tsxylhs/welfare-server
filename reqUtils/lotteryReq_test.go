package reqUtils

import (
	"log"
	"lottery/welfare/model"
	"testing"
)

func Test_lotteryTypes(t *testing.T) {
	LotteryReq.LotteryTypes()
	t.Log()
}
func Test_lotteryBouns(t *testing.T) {
	selectlottery := &model.SelectLottery{}
	res := LotteryReq.LotteryBonus(selectlottery)
	log.Fatal(res.LotteryPrize...)
	t.Log()
}

func Test_LotteryQuery(t *testing.T) {
	selectQuery := &model.SelectQuery{}
	selectQuery.LotteryNo = "22001"
	selectQuery.LotteryId = "ssq"
	res := LotteryReq.LotteryQuery(selectQuery)
	//res.LotteryOpenQuery
	log.Print("res", res)
	t.Log()
}
