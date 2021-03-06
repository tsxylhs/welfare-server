package service

import (
	"lottery/welfare/cs"
	"lottery/welfare/model"
)

type lotterys int

var Lotterys lotterys

func (lotterys) Get(form *model.Lottery) error {
	// 更新数据库中的记录
	if _, err := cs.Sql.ID(form.ID).Get(form); err != nil {
		return err
	}

	return nil
}

// list 获取多个项目列表
func (lotterys) List(page *model.Page, list *[]model.Lottery) error {
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
func (lotterys) Update(form *model.Lottery) error {
	if _, err := cs.Sql.Update(form, form); err != nil {

		return err
	}

	return nil
}

// Delete 删除记录
func (lotterys) Delete(form *model.Lottery) error {
	// 删除记录
	if _, err := cs.Sql.ID(form.ID).Delete(form); err != nil {

		return err
	}

	return nil
}

// Receive 保存记录
func (lotterys) Save(form *model.Lottery) error {
	if form.ID == 0 {
	}
	if _, err := cs.Sql.Insert(form); err != nil {
		return err
	}

	return nil
}
