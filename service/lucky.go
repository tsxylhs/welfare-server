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
	luckyData := &model.LuckyData{}
	luckyData.Base.BeforeInsert()
	luckyData.UserId = lucky.UserId
	//生成数字
	val = generateData(lucky.Ty)

	luckyData.Type = lucky.Ty
	jval, _ := json.Marshal(val)
	luckyData.LuckyData = string(jval)
	if _, err := cs.Sql.Insert(luckyData); err != nil {
		return err, val
	}
	return nil, val
}
func generateData(ty string) []int {
	switch ty {
	case "1":
		return generateRandomNumber(1, 33, 7)
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
