package db

import (
	"go_projects/blog/model"
)

func InsertArticle(article *model.ArticleDetail) (articleId int64, err error) {
	if article == nil {
		return
	}
	sqlStr := `insert into article(content,summary,title,username,category_id,view_count,comment_count) values(?,?,?,?,?,?,?)`
	result, err := DB.Exec(sqlStr, article.Content, article.Summary, article.Title, article.Username, article.ArticleInfo.CategoryId, article.ViewCount, article.CommentCount)
	if err != nil {
		return
	}
	articleId, err = result.LastInsertId()
	return
}

func GetArticleList(pageNum, pageSize int) (articleList []*model.ArticleInfo, err error) {
	if pageNum < 0 || pageSize <= 0 {
		return
	}
	sqlStr := `select id,summary,title,view_count,create_time,comment_count,username,category_id from article where status = 1 order by create_time desc limit ?,?`
	err = DB.Select(&articleList, sqlStr, pageNum, pageSize)
	return
}

func GetArticleDetail(articleId int64) (articleDetail *model.ArticleDetail, err error) {
	if articleId < 0 {
		return
	}
	articleDetail = &model.ArticleDetail{}
	sqlStr := `select id,summary,content,title,view_count,create_time,comment_count,username,category_id from article where id = ? and status = 1`
	err = DB.Get(articleDetail, sqlStr, articleId)
	return
}

func GetArticleListByCategoryId(categoryId, pageNum, pageSize int) (articleList []*model.ArticleInfo, err error) {
	if categoryId < 0 || pageNum < 0 || pageSize <= 0 {
		return
	}
	sqlStr := `select id,summary,title,view_count,create_time,comment_count,username,category_id from article where status = 1 and category_id = ? order by create_time desc limit ?,?`
	err = DB.Select(&articleList, sqlStr, categoryId, pageNum, pageSize)
	return
}
