package service

//兑奖
import (
	"lottery/welfare/cs"
	"lottery/welfare/model"
)

type awardings int

var Awardings awardings

func (awardings) Get(form *model.Awarding) error {
	// 更新数据库中的记录
	if _, err := cs.Sql.ID(form.ID).Get(form); err != nil {
		return err
	}
	return nil
}

// list 获取多个项目列表
func (awardings) List(id int64, page *model.Page, list *[]model.Awarding) error {
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
func (awardings) Update(form *model.Awarding) error {
	if _, err := cs.Sql.Update(form, form); err != nil {

		return err
	}

	return nil
}

// Delete 删除记录
func (t awardings) Delete(form *model.Awarding) error {
	// 删除记录
	if _, err := cs.Sql.ID(form.ID).Delete(form); err != nil {

		return err
	}

	return nil
}

// Receive 保存记录
func (t awardings) Save(form *model.Awarding) error {
	if form == nil || form.ID == 0 {
		form.BeforeInsert()
	}
	if _, err := cs.Sql.Insert(form); err != nil {
		return err
	}

	return nil
}
