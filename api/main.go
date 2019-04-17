package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hatch-group/keywordss-api/api/controller"
	"github.com/jmoiron/sqlx"
)

func main() {
	r := gin.Default()

	api := r.Group("/api")
	{
		story := &controller.Story{}
		user := &controller.User{}
		dburl := os.Getenv("MYSQL_URL")
		db, err := sqlx.Connect("mysql", dburl)
		if err != nil {
			fmt.Println("mysql connect error")
		}
		story.DB = db
		user.DB = db
		api.GET("/stories", story.IndexGet)
		api.GET("/stories/:id", story.ShowItem)
		api.POST("/stories", story.Post)
		api.PUT("/stories", story.Edit)
		api.DELETE("/stories", story.Delete)

		api.GET("/user/:user_id/stories", story.IndexMyPost)
		api.Post("/users/signup", user.UserSignUp)
		api.Post("/users/signin", user.UserSignIn)
	}

	r.Run(":8080")
}
