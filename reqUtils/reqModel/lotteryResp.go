package reqModel

import "lottery/welfare/model"

//彩票类型
type LotteryResp struct {
	Reason    string          `json:"reason"`
	Result    []model.Lottery `json:"result"`
	ErrorCode int             `json:"error_code"`
}

//兑奖记录
type Awarding struct {
	Reason    string         `json:"reason"`
	Awarding  model.Awarding `json:"result"`
	ErrorCode int            `json:"error_code"`
}

//开奖记录
type LotteryQuery struct {
	Reason           string                 `json:"reason"`
	LotteryOpenQuery model.LotteryOpenQuery `json:"result"`
	ErrorCode        int                    `json:"error_code"`
}
