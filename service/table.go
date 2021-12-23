package service

import (
	"library/libraryDemo/cs"
	"library/libraryDemo/model"
)

var Table table

type table int

func (table) Get(form *model.Table) error {
	// 更新数据库中的记录
	if _, err := cs.Sql.ID(form.ID).Get(form); err != nil {
		return err
	}

	return nil
}

// list 获取多个项目列表
func (table) List(form *model.Table, page *model.Page, list *[]model.Table) error {
	// 分页查询
	if cnt, err := cs.Sql.Limit(page.Limit(), page.Skip()).Desc("status", "created_at").FindAndCount(list, form); err != nil {
		return err
	} else {
		page = page.GetPager(cnt)
	}

	return nil
}

// Update 更新新的纪录
func (table) Update(form *model.Table) error {
	if _, err := cs.Sql.Update(form, form); err != nil {

		return err
	}

	return nil
}

// Delete 删除记录
func (t table) Delete(form *model.Table) error {
	// 删除记录
	if _, err := cs.Sql.ID(form.ID).Delete(form); err != nil {

		return err
	}

	return nil
}

// Receive 保存记录
func (t table) Save(form *model.Table) error {
	if form.ID == 0 {
	}
	if _, err := cs.Sql.Insert(form); err != nil {
		return err
	}

	return nil
}
