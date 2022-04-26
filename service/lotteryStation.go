package service

import (
	"log"
	"lottery/welfare/cs"
	"lottery/welfare/model"

	"github.com/mitchellh/mapstructure"
	"github.com/xormplus/xorm"
)

type lotteryStations int

var LotteryStations lotteryStations
var (
	selectSql         = "select *, ROUND(6378.138*2*ASIN(SQRT(POW(SIN((31.17406175605199*PI()/180-lat*PI()/180)/2),2)+COS(31.17406175605199*PI()/180)*COS(lat*PI()/180)*POW(SIN((121.40638221320846*PI()/180-lng*PI()/180)/2),2)))*1000) AS distance from lottery_station having distance <50000 order by distance asc"
	selectSqlParam    = "select *, ROUND(6378.138*2*ASIN(SQRT(POW(SIN(( ? *PI()/180-lat*PI()/180)/2),2)+COS(? *PI()/180)*COS(lat*PI()/180)*POW(SIN((?*PI()/180-lng*PI()/180)/2),2)))*1000) AS distance from lottery_station  where  location like ? having distance <? order by distance asc"
	selectSqlParamNoD = "select *, ROUND(6378.138*2*ASIN(SQRT(POW(SIN(( ? *PI()/180-lat*PI()/180)/2),2)+COS(? *PI()/180)*COS(lat*PI()/180)*POW(SIN((?*PI()/180-lng*PI()/180)/2),2)))*1000) AS distance from lottery_station  where location like ?  order by distance asc"
)

func (lotteryStations) Get(form *model.LotteryStation) error {
	// 更新数据库中的记录
	if _, err := cs.Sql.ID(form.ID).Get(form); err != nil {
		return err
	}

	return nil
}

// list 获取多个项目列表
func (lotteryStations) List(params *model.Params, list *[]model.LotteryStationVo) error {
	// 分页查询
	listCount := &[]model.LotteryStationVo{}
	cs.Sql.ShowSQL(true)
	if params.Lng != 0 {
		result := &xorm.ResultMap{}
		session := &xorm.Session{}
		var count int64 = 0
		count, _ = cs.Sql.Table("lottery_station").Limit(params.Limit(), params.Skip()).Desc("created_at").FindAndCount(listCount)
		if params.Distance != -1 {
			session = cs.Sql.SQL(selectSqlParam, params.Lat, params.Lat, params.Lng, "%"+params.Name+"%", params.Distance).Limit(params.Limit(), params.Skip())
			//count, _ = cs.Sql.SQL(selectSqlParam, params.Lat, params.Lat, params.Lng, params.Distance).FindAndCount(in)
		} else {
			session = cs.Sql.SQL(selectSqlParamNoD, params.Lat, params.Lat, params.Lng, "%"+params.Name+"%").Limit(params.Limit(), params.Skip())
			//count, _ = cs.Sql.SQL(selectSqlParamNoD, params.Lat, params.Lat, params.Lng).FindAndCount(in)
		}
		result = session.Query()
		params.Page = *params.Page.GetPager(count)
		log.Println(result)
		mapVal, err := result.List()
		if err != nil {
			log.Print("err", err)
		}

		err = mapstructure.Decode(mapVal, list)
		if err != nil {
			log.Print("err", err)
		}
	} else {
		cs.Sql.ShowSQL(true)
		if cnt, err := cs.Sql.Table("lottery_station").Limit(params.Limit(), params.Skip()).Desc("created_at").FindAndCount(list); err != nil {
			return err
		} else {
			params.Page = *params.Page.GetPager(cnt)
		}
	}
	return nil
}

// Update 更新新的纪录
func (lotteryStations) Update(form *model.LotteryStation) error {
	if _, err := cs.Sql.Update(form, form); err != nil {

		return err
	}

	return nil
}

// Delete 删除记录
func (t lotteryStations) Delete(form *model.LotteryStation) error {
	// 删除记录
	if _, err := cs.Sql.ID(form.ID).Delete(form); err != nil {

		return err
	}

	return nil
}

// save 保存记录
func (t lotteryStations) Save(form *model.LotteryStation) error {
	if form.ID == 0 {
		form.BeforeInsert()
	}
	if _, err := cs.Sql.Insert(form); err != nil {
		return err
	}

	return nil
}
