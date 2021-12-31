package reqUtils

import (
	"encoding/json"
	"io/ioutil"
	"lottery/welfare/reqUtils/reqModel"
	"lottery/welfare/service"
	"net/http"
	"net/url"
)

var (
	openId = "JH7c50653df943461c58e0ceaa9eb6b607"

	lotteryKey = "5cef4e823e071004968fce4cf622d3c7"

	//彩票种类
	lotteryTyps = "http://apis.juhe.cn/lottery/types"

	//开奖结果
	lotteryQuery = "http://apis.juhe.cn/lottery/query"

	//历史开奖结果
	lotteryHistory = "http://apis.juhe.cn/lottery/history"

	//中奖计算器
	lotteryBonus = "http://apis.juhe.cn/lottery/bonus"
)

type lotteryReq int

var LotteryReq lotteryReq

//彩票类型获取
func (lotteryReq) LotteryTypes() {
	params := url.Values{}
	Url, err := url.Parse(lotteryTyps)
	if err != nil {
		panic(err.Error())
	}
	params.Set("key", lotteryKey)
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	resp, err := http.Get(urlPath)
	if err != nil {
		panic(err.Error())
	}
	defer resp.Body.Close()
	response := &reqModel.LotteryResp{}
	respByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}
	err = json.Unmarshal(respByte, response)
	if err != nil {
		panic(err.Error())
	}
	if response.Reason == "查询成功" {
		for _, lottery := range response.Result {
			service.Lotterys.Save(&lottery)
		}
	}

}
