package service

import (
	"library/libraryDemo/cs"
	"library/libraryDemo/model"
)

var Messages message

type message int

func (message) Get(form *model.Message) error {
	// 更新数据库中的记录
	if _, err := cs.Sql.ID(form.ID).Get(form); err != nil {
		return err
	}

	return nil
}

// list 获取多个项目列表
func (message) List(form *model.Message, page *model.Page, list *[]model.Message) error {
	// 分页查询
	if cnt, err := cs.Sql.Limit(page.Limit(), page.Skip()).Desc("status", "created_at").FindAndCount(list, form); err != nil {
		return err
	} else {
		page = page.GetPager(cnt)
	}

	return nil
}

// Update 更新新的纪录
func (message) Update(form *model.Message) error {
	if _, err := cs.Sql.Update(form, form); err != nil {

		return err
	}

	return nil
}

// Delete 删除记录
func (t message) Delete(form *model.Message) error {
	// 删除记录
	if _, err := cs.Sql.ID(form.ID).Delete(form); err != nil {

		return err
	}

	return nil
}

// Receive 保存记录
func (t message) Save(form *model.Message) error {
	if form.ID == 0 {
	}
	if _, err := cs.Sql.Insert(form); err != nil {
		return err
	}

	return nil
}
