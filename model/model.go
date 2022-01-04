package model

import (
	"fmt"
	"lottery/welfare/cs"
	"time"

	"github.com/bwmarrin/snowflake"
)

var idgen *snowflake.Node

type Base struct {
	ID        int64     `xorm:"pk  'id'" json:"id,string" form:"id"`
	UpdatedAt time.Time `xorm:"updated" json:"updatedAt"`
	CreatedAt time.Time `xorm:"created" json:"createdAt"`
	DeletedAt time.Time `xorm:"deleted" json:"-"`
}

//彩票类型
type Lottery struct {
	Base          `xorm:"extends"`
	LotteryId     string `json:"lottery_id" form:"lottery_id"`
	LotteryName   string `json:"lottery_name" form:"lottery_name"`
	LotteryTypeId string `json:"lottery_type_id" form:"lottery_name"`
	Remarks       string `json:"remarks" form:"remarks"`
}

//彩票站
type LotteryStation struct {
	Base     `xorm:"extends"`
	Name     string  `xorm:"name" form:"name"`
	Location string  `xorm:"location" form:"location"`
	ImageUrl string  `xorm:"image_url" form:"image_url"`
	Mobile   string  `json:"mobile" form:"mobile"`
	Lot      float64 `xorm:"lot" form:"lot"` //经度
	Lat      float64 `xorm:"lat" form:"lat"`
}

//兑奖记录
type LotteryOpenMessage struct {
	Base              `xorm:"extends"`
	LotteryId         string       `json:"lottery_id" form:"lottery_id"`
	LotteryName       string       `json: "lottery_name" form: "lottery_name"`
	LotteryRes        string       `json:"lottery_res" form:"lottery_res"`
	LotteryNo         string       `json:"lottery_no" form:"lottery_no"`
	LotteryDate       string       `json:"lottery_date" form:"lottery_date"`
	LotteryExdate     string       `json:"lottery_exdate" form:"lottery_exdate"`
	LotterySaleAmount string       `json:"lottery_sale_amount" form:"lottery_sale_amount"`
	LotteryPoolAmount string       `json:"lottery_pool_amount" form:"lottery_pool_amount"`
	LotteryPrize      LotteryPrize `json:"lottery_prize" form:"lottery_prize"`
}

//开奖记录查询
type LotteryOpenQuery struct {
	Base              `xorm:"extends"`
	LotteryId         string         `json:"lottery_id" form:"lottery_id"`
	LotteryName       string         `json: "lottery_name" form: "lottery_name"`
	LotteryRes        string         `json:"lottery_res" form:"lottery_res"`
	LotteryNo         string         `json:"lottery_no" form:"lottery_no"`
	LotteryDate       string         `json:"lottery_date" form:"lottery_date"`
	LotteryExdate     string         `json:"lottery_exdate" form:"lottery_exdate"`
	LotterySaleAmount string         `json:"lottery_sale_amount" form:"lottery_sale_amount"`
	LotteryPoolAmount string         `json:"lottery_pool_amount" form:"lottery_pool_amount"`
	LotteryPrize      []LotteryPrize `xorm:"lottery_prize" json:"lottery_prize" form:"lottery_prize"`
}

type LotteryPrize struct {
	PrizeName    string `json:"prize_name" form:"prize_name"`
	PirzeNum     string `json:"prize_num" form:"prize_num"`
	PirzeAmount  string `json:"prize_amount" form:"prize_amount"`
	PirzeRequire string `json:"prize_require" form:"prize_require"`
}

//我的彩票
type MyLottery struct {
	Base      `xorm:"extends"`
	UserId    int64  `xorm:"user_id" form:"userId"`
	LotteryId string `json:"lottery_id" form:"lottery_id"`
	Number    string `xorm:"number" form:"number"` //购彩记录
	Tag       string `xorm:"tag" form:"tag"`
}

//微信用户
type User struct {
	Base          `xorm:"extends"`
	Username      string `json:"username" form:"username"`
	Password      string `form:"password" json:"password"`
	Email         string `json:"email"`
	Mobile        string `json:"mobile"`
	Restaurant    string `json:"restaurant"`
	Status        int
	OpenId        string  `json:"openid" form:"openid"`     // open id
	Token         string  `json:"token" form:"token"`       // token access token
	NickName      string  `json:"nickName" form:"nickName"` // 昵称
	AvatarUrl     string  `json:"avatarUrl"`                // 头像
	Code          string  `json:"code" form:"code"`
	Province      string  `json:"province" form:"province"`
	Gender        int     `json:"gender" form:"gender"`
	City          string  `json:"city" form:"city"`
	SessionKey    string  `json:"session_key" form:"session_key"` // session key
	Region        string  `json:"region"`
	Url           string  `json:"url"`
	UnionId       string  `json:"unionid"`
	ShopId        int64   `json:"shopId,string"`
	Times         int     `json:"times"`       // 消费次数
	Expenditure   float64 `json:"expenditure"` // 消费支出
	Save          float64 `json:"save"`        // 节省了多少钱
	EncryptedData string  `json:"encryptedData" xorm:"-"`
	ErrMsg        string  `json:"errMsg"  xorm:"-"`
	Iv            string  `json:"iv"  xorm:"-"`
	RowData       string  `json:"rowData"  xorm:"-"`
	Signature     string  `json:"signature"  xorm:"-"`
}
type Message struct {
	Base      `xorm:"extends"`
	Title     string
	LibraryId int64
	Message   string
}

type Awarding struct {
	Base           `xorm:"extends"`
	LotteryID      string        `json:"lottery_id"`
	LotteryName    string        `json:"lottery_name"`
	LotteryNo      string        `json:"lottery_no"`
	LotteryDate    string        `json:"lottery_date"`
	RealLotteryRes string        `json:"real_lottery_res"`
	LotteryRes     string        `json:"lottery_res"`
	InMoney        string        `json:"in_money"`
	BuyRedBallNum  string        `json:"buy_red_ball_num"`
	BuyBlueBallNum string        `json:"buy_blue_ball_num"`
	HitRedBallNum  string        `json:"hit_red_ball_num"`
	HitBlueBallNum string        `json:"hit_blue_ball_num"`
	IsPrize        string        `json:"is_prize"`
	PrizeMsg       string        `json:"prize_msg"`
	LotteryPrize   []interface{} `json:"lottery_prize"`
}

func NewBD() {
	if err := cs.Sql.Sync(
		new(User),
		new(Lottery),
		new(LotteryStation),
		new(LotteryOpenMessage),
		new(LotteryOpenQuery),
		new(LotteryPrize),
		new(Awarding),
		new(MyLottery),
		new(Message)); err != nil {
		fmt.Print("初始化失败", err)
	}
}

type Page struct {
	P   int    `json:"p" form:"p"`
	Ps  int    `json:"ps" form:"ps"`
	Cnt int64  `json:"cnt"`
	K   string `josn:"k"`
	Pc  int    `json:"pc"`
	Od  string `json:"od,omitempty"`
}

func (page *Page) GetPage() *Page {
	return page
}

func (page *Page) GetPager(count int64) *Page {
	page.Cnt = count
	if page.P < 1 {
		page.P = 1
	}
	if page.Ps < 1 {
		page.Ps = 10
	}
	page.Pc = int(page.Cnt)/page.Ps + 1
	return page
}

func (page *Page) Skip() int {
	if page.Ps > 0 {
		return (page.P - 1) * page.Ps
	}

	return (page.P - 1) * 10
}

func (page *Page) Limit() int {
	if page.Ps > 0 {
		return page.Ps
	}

	return 10
}

func (b *Base) BeforeInsert() {
	b.ID, _ = Next()
	now := time.Now()
	b.CreatedAt = now
	b.UpdatedAt = now
}

var node *snowflake.Node

func Next() (int64, error) {
	return int64(node.Generate()), nil
}
func init() {
	var err error
	node, err = snowflake.NewNode(1)
	if err != nil {
		panic(err)
	}
}
