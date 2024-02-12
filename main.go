package main

import (
	db "hello_gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("views/*.html")
	router.GET("/", func(ctx *gin.Context) {
		todos := db.SelectAllTodo()
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"todos": todos,
		})
	})

	router.POST("/", func(ctx *gin.Context) {
		text := ctx.PostForm("text")
		status := ctx.PostForm("status")
		db.AddTodo(text, status)
		ctx.Redirect(http.StatusFound, "/")
	})

	router.Run()
}
