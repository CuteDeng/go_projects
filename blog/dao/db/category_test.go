package db

import (
	"go_projects/blog/model"
	"testing"
)

func init() {
	dns := "root:123456@tcp(localhost:3306)/blog?parseTime=true"
	err := Init(dns)
	if err != nil {
		panic(err)
	}
}

func TestInsertCategory(t *testing.T) {
	category := &model.Category{
		CategoryName: "test",
		CategoryNo:   10,
	}
	id, err := InsertCategory(category)
	if err != nil {
		panic(err)
	}
	t.Logf("category id:%#v", id)
}

func TestGetCategoryById(t *testing.T) {
	category, err := GetCategoryById(1)
	if err != nil {
		panic(err)
	}
	t.Logf("category:%#v", category)
}

func TestGetCategoryList(t *testing.T) {
	var categoryIds []int64 = []int64{1, 2, 3}
	list, err := GetCategoryList(categoryIds)
	if err != nil {
		panic(err)
	}
	for _, v := range list {
		t.Logf("category:%#v", v)
	}
}

func TestGetAllCategoryList(t *testing.T) {
	list, err := GetAllCategoryList()
	if err != nil {
		panic(err)
	}
	for _, v := range list {
		t.Logf("category:%#v", v)
	}
}
