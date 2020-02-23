package main

import (
	"go_projects/blog/controller"
	"go_projects/blog/dao/db"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	dns := "root:123456@tcp(localhost:3306)/blog?parseTime=true"
	err := db.Init(dns)
	if err != nil {
		panic(err)
	}
	router.Static("/static/", "./static")
	router.LoadHTMLGlob("views/*")
	router.GET("/", controller.IndexHandle)
	router.GET("/category/", controller.CategoryListHandle)
	_ = router.Run(":8000")
}
