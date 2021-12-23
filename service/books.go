package service

import (
	"library/libraryDemo/cs"
	"library/libraryDemo/model"
)

var Books books

type books int

func (books) Get(form *model.Books) error {
	// 更新数据库中的记录
	if _, err := cs.Sql.ID(form.ID).Get(form); err != nil {
		return err
	}

	return nil
}

// list 获取多个项目列表
func (books) List(id int64, page *model.Page, list *[]model.Books) error {
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
func (books) Update(form *model.Books) error {
	if _, err := cs.Sql.Update(form, form); err != nil {

		return err
	}

	return nil
}

// Delete 删除记录
func (t books) Delete(form *model.Books) error {
	// 删除记录
	if _, err := cs.Sql.ID(form.ID).Delete(form); err != nil {

		return err
	}

	return nil
}

// Receive 保存记录
func (t books) Save(form *model.Books) error {
	if form.ID == 0 {
	}
	if _, err := cs.Sql.Insert(form); err != nil {
		return err
	}

	return nil
}

func (t books) UpdateBooks(form *model.Updatebook) error {
	book := &model.Books{}
	book.Status = "1"
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
