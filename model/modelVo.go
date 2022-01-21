package model

type LotteryStationVo struct {
	Base     `xorm:"extends"`
	Name     string  `xorm:"name" form:"name"`
	Location string  `xorm:"location" form:"location"`
	ImageUrl string  `xorm:"image_url" form:"image_url"`
	Mobile   string  `json:"mobile" form:"mobile"`
	Lng      float64 `xorm:"Lng" form:"Lng"` //经度
	Lat      float64 `xorm:"lat" form:"lat"`
	Distance int     `xorm:"distance" form:"distance"`
}
