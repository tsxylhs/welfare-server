package service

import (
	"library/libraryDemo/cs"
	"library/libraryDemo/model"
)

var LotteryStations lotteryStations

type lotteryStations int

func (lotteryStations) Get(form *model.LotteryStation) error {
	// 更新数据库中的记录
	if _, err := cs.Sql.ID(form.ID).Get(form); err != nil {
		return err
	}

	return nil
}

// list 获取多个项目列表
func (lotteryStations) List(id int64, page *model.Page, list *[]model.LotteryStation) error {
	// 分页查询
	cs.Sql.ShowSQL(true)
	if cnt, err := cs.Sql.Limit(page.Limit(), page.Skip()).Where("library_id=?", id).Desc("created_at").FindAndCount(list); err != nil {
		return err
	} else {
		page = page.GetPager(cnt)
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

// Receive 保存记录
func (t lotteryStations) Save(form *model.LotteryStation) error {
	if form.ID == 0 {
	}
	if _, err := cs.Sql.Insert(form); err != nil {
		return err
	}

	return nil
}
