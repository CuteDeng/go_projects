package service

import (
	"go_projects/blog/dao/db"
	"go_projects/blog/model"
)

func GetAllCategoryList() (categoryList []*model.Category, err error) {
	categoryList, err = db.GetAllCategoryList()
	if err != nil {
		return
	}
	return
}
