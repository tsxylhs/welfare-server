package model

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"library/libraryDemo/cs"
	"time"
)

var idgen *snowflake.Node

type Base struct {
	ID        int64     `xorm:"pk  'id'" json:"id,string" form:"id"`
	UpdatedAt time.Time `xorm:"updated" json:"updatedAt"`
	CreatedAt time.Time `xorm:"created" json:"createdAt"`
	DeletedAt time.Time `xorm:"deleted" json:"-"`
}
type Apply struct {
	Base      `xorm:"extends"`
	Num       string `json:"num"`
	Name      string `json:"name"`
	Class     string `json:"class"`
	School    string `json:"school"`
	Phone     string `json:"phone"`
	LibraryId int64  `json:"libraryId,string"`
	UserId    int64  `json:"userId,string"`
	Site      string `json:"site"`
	Time      string `json:"time"`
}
type MyBook struct {
	Base       `xorm:"extends"`
	Status     int     `json:"status,string" form:"status,string"`
	UserId     int64   `json:"userId" form:"userId"`
	BookId     int64   `json:"bookId" form:"bookId"`
	BorrowTime string  `json:"borrowTime" form:"borrowTime"`
	ReturnTime string  `json:"returnTime" form:"returnTime"`
	LibraryId  int64   `json:"libraryId" form:"libraryId"`
	NotesId    int64   `json:"notesId" form:"notesId"`
	Library    Library `xorm:"-" json:"library"`
	Book       Books   `xorm:"-" json:"book"`
}
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
type Library struct {
	Base    `xorm:"extends"`
	Name    string `json:"name" form:"name"`
	Lat     string `json:"lat" from:"lat"`
	Lng     string `json:"lng" from:"lng"`
	Address string `json:"address" from:"address"`
	Image   string `json:"image" form:"image"`
	Apply   Apply  `xorm:"-" json:"apply"`
}
type Books struct {
	Base `xorm:"extends"`

	Name         string `json:"name" form:"name"`
	LibraryId    int64  `json:"libraryId" form:"libraryId"`
	Type         string `json:"type" form:"type "`
	Introduction string `json:"introduction" form:"introduction"`
	Location     string `json:"location" form:"location"`
	Author       string `json:"author" form:"author"`
	Image        string `json:"image" form:"image"`
	Status       string `json:"status" form:"status"` // 借阅状态
}
type Table struct {
	Base      `xorm:"extends"`
	Name      string `json:"name"`
	LibraryId int64  `json:"libraryId"`
	Location  string `json:"location"`
	Status    string `json:"status"`
}

type Cabinet struct {
	Base      `xorm:"extends"`
	Name      string
	LibraryId int64
	Location  string
	Status    string
}
type Message struct {
	Base      `xorm:"extends"`
	Title     string
	LibraryId int64
	Message   string
}
type Notes struct {
	Base       `xorm:"extends"`
	Title      string `json:"title"`
	BookId     int64  `json:"bookId"`
	BookName   string `json:"bookName" form:"bookName"`
	BookAuthor string `json:"bookAuthor" form:"bookAuthor"`
	Desc       string `json:"desc" form:"desc"`
	Notes      string
	UserId     int64 `json:"userId,string" form:"userId,string"`
}

func NewBD() {
	if err := cs.Sql.Sync(
		new(User),
		new(Library),
		new(Table),
		new(Cabinet),
		new(Message),
		new(Books),
		new(MyBook),
		new(Apply),
		new(Notes)); err != nil {
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
