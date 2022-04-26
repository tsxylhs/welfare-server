package model

type LotteryStationVo struct {
	Base      `xorm:"extends"`
	Name      string  `xorm:"name" form:"name" json:"name"`
	Location  string  `xorm:"location" form:"location" json:"location"`
	Image_url string  `xorm:"image_url" form:"image_url" json:"image_url"`
	Mobile    string  `json:"mobile" form:"mobile" json:"mobile"`
	Lng       float64 `xorm:"Lng" form:"Lng" json:"lng"` //经度
	Lat       float64 `xorm:"lat" form:"lat" json:"lat"`
	Distance  int     `xorm:"distance" form:"distance" json:"distance"`
}
