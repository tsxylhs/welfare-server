package reqModel

import "lottery/welfare/model"

type LotteryResp struct {
	Reason    string          `json:"reason"`
	Result    []model.Lottery `json:"result"`
	ErrorCode int             `json:"error_code"`
}
