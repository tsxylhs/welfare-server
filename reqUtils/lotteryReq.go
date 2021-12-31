package reqUtils

import (
	"encoding/json"
	"io/ioutil"
	"lottery/welfare/model"
	"lottery/welfare/reqUtils/reqModel"
	"lottery/welfare/service"
	"net/http"
	"net/url"
)

var (
	openId = "JH7c50653df943461c58e0ceaa9eb6b607"

	lotteryKey = "5cef4e823e*71**4968fce4cf622d3c7"

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
	params.Set("key", lotteryKey)
	respByte, err := GetReq(params, lotteryTyps)
	if err != nil {
		panic(err.Error())
	}

	response := &reqModel.LotteryResp{}
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

//中奖计算器
func (lotteryReq) LotteryBonus(selectlottery model.SelectLottery) (resp *reqModel.Awarding) {

	params := url.Values{}
	params.Set("key", lotteryKey)
	params.Set("lottery_no", selectlottery.LotteryOpenNo)
	params.Set("lottery_id", selectlottery.LotteryType)
	params.Set("lottery_res", selectlottery.LotteryNum)
	respByte, err := GetReq(params, lotteryBonus)
	if err != nil {
		panic(err.Error())
	}
	//fmt.Println(respByte)
	response := &reqModel.Awarding{}
	err = json.Unmarshal(respByte, response)
	if err != nil {
		panic(err.Error())
	}
	//todo  保存兑奖记录
	return response

}

func GetReq(params url.Values, path string) (respResult []byte, err error) {

	Url, err := url.Parse(path)
	if err != nil {
		panic(err.Error())
	}
	Url.RawQuery = params.Encode()
	urlPath := Url.String()
	resp, err := http.Get(urlPath)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	respByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return respByte, nil
}
