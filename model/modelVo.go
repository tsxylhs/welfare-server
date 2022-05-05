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
type LuckyDataVo struct {
	UserId string `json:"user_id" form:"user_id"`
	Ty     string `json:"ty" form:"ty"`
}
type LuckyDataV struct {
	Base          `xorm:"extends"`
	UserId        string `xorm:"user_id" json:"user_id"`
	Crt           string `xorm:"-" json:"crt"`
	TypeName      string `xorm:"-" json:"type_name"`
	Type          string `json:"type" form:"type"`
	LuckyData     string `xorm:"lucky_data" json:"lucky_data"`
	HitCount      int    `xorm:"hit_count" json:"hit_count"`
	WinningAmount string `xorm:"winning_amount" json:"winning_amount"`
}
