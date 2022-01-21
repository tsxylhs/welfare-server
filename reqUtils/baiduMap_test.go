package reqUtils

import (
	"fmt"
	"lottery/welfare/model"
	"testing"
)

func TestBaiduMap_BaiduMapAddrToLat(t *testing.T) {
	lotteryStation := &model.LotteryStation{}
	lotteryStation.Location = "上海市商汤科技大厦"
	BaiduMap.BaiduMapAddrToLat(lotteryStation)
	fmt.Print("lotterStationlat:", lotteryStation.Lat, "lng:", lotteryStation.Lng)
}
