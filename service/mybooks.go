package service

import (
	"library/libraryDemo/cs"
	"library/libraryDemo/model"
)

var Mybooks mybooks

type mybooks int

func (mybooks) Get(form *model.MyBook) error {
	// 更新数据库中的记录
	if _, err := cs.Sql.ID(form.ID).Get(form); err != nil {
		return err
	}

	return nil
}

// list 获取多个项目列表
func (mybooks) List(isqueue string, id int64, page *model.Page, list *[]model.MyBook) error {
	// 分页查询
	cs.Sql.ShowSQL(true)
	if isqueue != "true" {
		if cnt, err := cs.Sql.Limit(page.Limit(), page.Skip()).Where("user_id=? and site is null", id).Desc("created_at").FindAndCount(list); err != nil {
			return err
		} else {
			page = page.GetPager(cnt)
		}
	} else {
		if cnt, err := cs.Sql.Limit(page.Limit(), page.Skip()).Where("user_id=? and site is not null", id).Desc("created_at").FindAndCount(list); err != nil {
			return err
		} else {
			page = page.GetPager(cnt)
		}
	}
	for i := 0; i < len(*list); i++ {
		library := &model.Library{}

		if _, err := cs.Sql.ID((*list)[i].LibraryId).Get(library); err != nil {
			return err
		}
		book := &model.Books{}
		if _, err := cs.Sql.ID((*list)[i].BookId).Get(book); err != nil {
			return err
		}
		(*list)[i].Book = *book
		(*list)[i].Library = *library
	}

	return nil
}

// Update 更新新的纪录
func (mybooks) Update(form *model.MyBook) error {
	if _, err := cs.Sql.Update(form, form); err != nil {

		return err
	}

	return nil
}

// Delete 删除记录
func (t mybooks) Delete(form *model.MyBook) error {
	// 删除记录
	if _, err := cs.Sql.ID(form.ID).Delete(form); err != nil {

		return err
	}

	return nil
}

// Receive 保存记录
func (t mybooks) Save(form *model.MyBook) error {
	if form.ID == 0 {
	}
	if _, err := cs.Sql.Insert(form); err != nil {
		return err
	}

	return nil
}

func (t mybooks) UpdateBooks(form *model.Updatebook) error {
	book := &model.MyBook{}
	cs.Sql.ShowSQL(true)
	if _, err := cs.Sql.ID(form.BookId).Cols("status").Update(book); err != nil {
		return err
	}
	mybook := &model.MyBook{}
	if _, err := cs.Sql.Where("book_id=?", form.BookId).Get(mybook); err != nil {
		return err
	}
	if mybook.ID != 0 {
		//
	} else {
		mybook.Base.BeforeInsert()
		mybook.BookId = form.BookId
		mybook.NotesId = 0
		mybook.Status = 0
		mybook.LibraryId = form.LibraryId
		mybook.BorrowTime = form.StartTime
		mybook.ReturnTime = form.EndTime
		mybook.UserId = form.UserId
		if _, err := cs.Sql.Insert(mybook); err != nil {
			return err
		}
	}
	return nil
}
func (t mybooks) SaveApply(apply *model.Apply) error {
	apply.BeforeInsert()
	applys := &[]model.Apply{}
	if err := cs.Sql.Where("library_id=? and user_id=?", apply.LibraryId, apply.UserId).Find(applys); err != nil {
	}
	if len(*applys) <= 0 {
		//发送短信验证码
		if _, err := cs.Sql.Insert(apply); err != nil {
			return err
		}
	}
	return nil
}
func (t mybooks) ApplyLibrary(id int64, librarys *[]model.Library) error {
	cs.Sql.ShowSQL(true)
	if _, err := cs.Sql.FindAndCount(librarys); err != nil {
		return err
	}
	for i := 0; i < len(*librarys); i++ {
		apply := &model.Apply{}
		if _, err := cs.Sql.Where("user_id=? and library_id=?", id, (*librarys)[i].ID).Get(apply); err != nil {

		}
		(*librarys)[i].Apply = *apply
	}
	return nil
}
func (t mybooks) Applylist(id int64, applys *[]model.ApplysVo) error {
	if err := cs.Sql.Table("apply").Where("site is not null").Find(applys); err != nil {
		return err
	}
	for i := 0; i < len(*applys); i++ {
		library := &model.Library{}

		if _, err := cs.Sql.ID((*applys)[i].LibraryId).Get(library); err != nil {
			return err
		}
		(*applys)[i].Library = *library
	}
	return nil
}

func (t mybooks) ApplyDelete(id int64) error {
	form := &model.Apply{}
	if _, err := cs.Sql.Where("user_id=? and site is not null", id).Delete(form); err != nil {
		return err
	}
	return nil
}
