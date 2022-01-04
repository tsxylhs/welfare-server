package model

type SelectLottery struct {
	LotteryType   string `form:"lotteryType"`
	LotteryNum    string `form:"lotteryNum"`
	LotteryOpenNo string `form:"lotteryOpenNo"`
}

type SelectQuery struct {
	LotteryId string `form:"lotteryId"`
	LotteryNo string `form:"lotteryNo"`
}
