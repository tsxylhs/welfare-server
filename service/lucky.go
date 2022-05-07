package service

import (
	"encoding/json"
	"lottery/welfare/cs"
	"lottery/welfare/model"
	"math/rand"
	"time"
)

type lucky int

var Lucky lucky

func (lucky) InsertLuckyData(lucky *model.LuckyDataVo) (err error, val []int) {
	val = generateData(lucky.Ty)
	return nil, val
}
func (lucky) Get(id int64, lucky *[]model.LuckyDataV) error {
	if err := cs.Sql.Table("lucky_data").Where("id=?", id).Find(lucky); err != nil {
		return err
	}
	return nil
}
func (lucky) Update(lucky *model.LuckyDataV) error {
	lucky.WinningAmount = "-1"
	cs.Sql.ShowSQL(true)
	if _, err := cs.Sql.Table("lucky_data").ID(lucky.ID).Update(lucky); err != nil {
		return err
	}
	return nil
}

func (lucky) Save(lucky *model.LuckyData) (err error) {
	lucky.Base.BeforeInsert()
	jval, _ := json.Marshal(lucky.LuckyData)
	lucky.LuckyData = string(jval)
	if _, err := cs.Sql.Insert(lucky); err != nil {
		return err
	}
	return nil
}
func (lucky) List(page *model.Page, list *[]model.LuckyDataV) error {
	// 分页查询
	cs.Sql.ShowSQL(true)
	if cnt, err := cs.Sql.Table("lucky_data").Limit(page.Limit(), page.Skip()).Where("user_id=?", page.UserId).FindAndCount(list); err != nil {
		return err
	} else {
		page = page.GetPager(cnt)
	}

	return nil
}
func generateData(ty string) []int {
	switch ty {
	case "1":
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		val := generateRandomNumber(1, 33, 7)
		val[6] = r.Intn((16 - 0)) + 0
		return val
	case "2":
		return generateRandomNumber3D(0, 9, 3)
	case "3":
		return generateRandomNumber(1, 33, 7)
	}
	return nil
}
func generateRandomNumber3D(start int, end int, count int) []int {
	//范围检查
	if end < start || (end-start) < count {
		return nil
	}
	//存放结果的slice
	nums := make([]int, 0)
	//随机数生成器，加入时间戳保证每次生成的随机数不一样
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for len(nums) < count {
		//生成随机数
		num := r.Intn((end - start)) + start
		nums = append(nums, num)
	}
	return nums
}

func generateRandomNumber(start int, end int, count int) []int {
	//范围检查
	if end < start || (end-start) < count {
		return nil
	}
	//存放结果的slice
	nums := make([]int, 0)
	//随机数生成器，加入时间戳保证每次生成的随机数不一样
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for len(nums) < count {
		//生成随机数
		num := r.Intn((end - start)) + start
		//查重
		exist := false
		for _, v := range nums {
			if v == num {
				exist = true
				break
			}
		}
		if !exist {
			nums = append(nums, num)
		}
	}
	return nums
}
