package service

import (
	"lottery/welfare/cs"
	"lottery/welfare/model"
)

type lotteryOpenQuerys int

var LotteryOpenQuerys lotteryOpenQuerys

func (lotteryOpenQuerys) Get(form *model.LotteryOpenQuery) error {
	// 更新数据库中的记录
	if _, err := cs.Sql.ID(form.ID).Get(form); err != nil {
		return err
	}

	return nil
}

// list 获取多个项目列表
func (lotteryOpenQuerys) List(page *model.Page, list *[]model.LotteryOpenQuery) error {
	// 分页查询
	cs.Sql.ShowSQL(true)
	if cnt, err := cs.Sql.Limit(page.Limit(), page.Skip()).Desc("created_at").FindAndCount(list); err != nil {
		return err
	} else {
		page = page.GetPager(cnt)
	}

	return nil
}

// Update 更新新的纪录
func (lotteryOpenQuerys) Update(form *model.LotteryOpenQuery) error {
	if _, err := cs.Sql.Update(form, form); err != nil {

		return err
	}

	return nil
}

// Delete 删除记录
func (t lotteryOpenQuerys) Delete(form *model.LotteryOpenQuery) error {
	// 删除记录
	if _, err := cs.Sql.ID(form.ID).Delete(form); err != nil {

		return err
	}

	return nil
}

// Receive 保存记录
func (t lotteryOpenQuerys) Save(form *model.LotteryOpenQuery) error {
	if form.ID == 0 {
		form.BeforeInsert()
	}
	if _, err := cs.Sql.Insert(form); err != nil {
		return err
	}
	return nil
}

func (lotteryOpenQuerys) Query(form *model.SelectQuery) (error error, lotteryOpenQuerys *[]model.LotteryOpenQuery) {
	// 更新数据库中的记录
	lotterys := &[]model.LotteryOpenQuery{}
	count, err := cs.Sql.Where("lottery_no=?", form.LotteryNo).And("lottery_id=?", form.LotteryId).FindAndCount(lotterys)
	if err != nil {
		return err, nil
	}
	if count > 0 {
		return nil, lotterys
	}
	return nil, nil
}
