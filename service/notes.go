package service

import (
	"library/libraryDemo/cs"
	"library/libraryDemo/model"
)

var Notes notes

type notes int

func (notes) Get(form *model.Notes) error {
	// 更新数据库中的记录
	if _, err := cs.Sql.ID(form.ID).Get(form); err != nil {
		return err
	}

	return nil
}

// list 获取多个项目列表
func (notes) List(userId int64, page *model.Page, list *[]model.Notes) error {
	// 分页查询
	if cnt, err := cs.Sql.Where("user_id=?", userId).Desc("created_at").FindAndCount(list); err != nil {
		return err
	} else {
		page = page.GetPager(cnt)
	}

	return nil
}

// Update 更新新的纪录
func (notes) Update(form *model.Notes) error {
	if _, err := cs.Sql.ID(form.ID).Update(form); err != nil {

		return err
	}

	return nil
}

// Delete 删除记录
func (t notes) Delete(form *model.Notes) error {
	// 删除记录
	if _, err := cs.Sql.ID(form.ID).Delete(form); err != nil {

		return err
	}

	return nil
}

// Receive 保存记录
func (t notes) Save(form *model.Notes) error {
	form.BeforeInsert()
	if _, err := cs.Sql.Insert(form); err != nil {
		return err
	}

	return nil
}
