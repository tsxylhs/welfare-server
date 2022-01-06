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
