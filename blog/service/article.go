package service

import (
	"go_projects/blog/dao/db"
	"go_projects/blog/model"
)

func GetArticleRecordList(pageNum, pageSize int) (articleRecordList []*model.ArticleRecord, err error) {
	articleList, err := db.GetArticleList(pageNum, pageSize)
	if err != nil {
		return
	}
	if len(articleList) <= 0 {
		return
	}
	// 更加articleId找到对应的categoryId
	ids := getCategoryIds(articleList)
	categoryList, err := db.GetCategoryList(ids)
	if err != nil {
		return
	}
	for _, article := range articleList {
		articleRecord := &model.ArticleRecord{
			ArticleInfo: *article,
		}
		categoryId := article.CategoryId
		for _, category := range categoryList {
			if categoryId == category.CategoryId {
				articleRecord.Category = *category
				break
			}
		}
		articleRecordList = append(articleRecordList, articleRecord)
	}
	return
}

func GetArticleRecordListById(categoryId, pageNum, pageSize int) (articleRecordList []*model.ArticleRecord, err error) {
	articleList, err := db.GetArticleListByCategoryId(categoryId, pageNum, pageSize)
	if err != nil {
		return
	}
	if len(articleList) <= 0 {
		return
	}
	// 更加articleId找到对应的categoryId
	ids := getCategoryIds(articleList)
	categoryList, err := db.GetCategoryList(ids)
	if err != nil {
		return
	}
	for _, article := range articleList {
		articleRecord := &model.ArticleRecord{
			ArticleInfo: *article,
		}
		categoryId := article.CategoryId
		for _, category := range categoryList {
			if categoryId == category.CategoryId {
				articleRecord.Category = *category
				break
			}
		}
		articleRecordList = append(articleRecordList, articleRecord)
	}
	return
}

func getCategoryIds(articleList []*model.ArticleInfo) (categoryIds []int64) {
LABEL:
	for _, article := range articleList {
		articleId := article.CategoryId
		for _, id := range categoryIds {
			if id == articleId {
				continue LABEL
			}
		}
		categoryIds = append(categoryIds, articleId)
	}
	return
}
