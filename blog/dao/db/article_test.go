package db

import (
	"go_projects/blog/model"
	"testing"
	"time"
)

func init() {
	dns := "root:123456@tcp(localhost:3306)/blog?parseTime=true"
	err := Init(dns)
	if err != nil {
		panic(err)
	}
}

func TestInsertArticle(t *testing.T) {
	//article := &model.ArticleDetail{
	//	ArticleInfo: model.ArticleInfo{
	//		CategoryId:   1,
	//		CommentCount: 0,
	//		CreateTime:   time.Now(),
	//		Title:        "t1",
	//		Username:     "dave",
	//		Summary:      "s1",
	//		ViewCount:    1,
	//	},
	//	Content: "test content",
	//}
	article := &model.ArticleDetail{}
	article.ArticleInfo.CategoryId = 1
	article.ArticleInfo.CommentCount = 0
	article.ArticleInfo.CreateTime = time.Now()
	article.ArticleInfo.Title = "t1"
	article.ArticleInfo.Username = "dave"
	article.ArticleInfo.Summary = "s1"
	article.ArticleInfo.ViewCount = 1
	article.Content = "c1"
	articleId, err := InsertArticle(article)
	if err != nil {
		panic(err)
	}
	t.Logf("insert success article %v", articleId)
}

func TestGetArticleList(t *testing.T) {
	list, err := GetArticleList(0, 2)
	if err != nil {
		panic(err)
	}
	t.Logf("article %d\n", len(list))
}

func TestGetArticleDetail(t *testing.T) {
	detail, err := GetArticleDetail(1)
	if err != nil {
		panic(err)
	}
	t.Logf("article %#v", detail)
}

func TestGetArticleListByCategoryId(t *testing.T) {
	list, err := GetArticleListByCategoryId(1, 0, 1)
	if err != nil {
		panic(err)
	}
	t.Logf("article  %#v", list)
}
