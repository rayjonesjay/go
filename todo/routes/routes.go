package routes

import (
	"net/http"
	"strconv"
	"todo/types"

	"github.com/gin-gonic/gin"
)

func RoutesInitializer() {
	router := gin.Default()

	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html",
			gin.H{
				"Todos":    types.Todos,              // display the current todos available
				"LoggedIn": types.LoggedInUser != "", // return true if the loggedInUser is not an empty string
				"Username": types.LoggedInUser,
			})
	})

	router.POST("/add", func(ctx *gin.Context) {
		text := ctx.PostForm("todo")
		todo := types.Todo{Text: text, Done: false}
		types.Todos = append(types.Todos, todo) // this acts as our database
		ctx.Redirect(http.StatusSeeOther, "/")
	})

	router.POST("/toggle", func(ctx *gin.Context) {
		index := ctx.PostForm("index")
		toggleIndex(index)
		ctx.Redirect(http.StatusSeeOther, "/")
	})
	router.Run(":5555")
}

func toggleIndex(ind string) {
	i, _ := strconv.Atoi(ind)
	if i >= 0 && i < len(types.Todos){
		// negate the value of Todo.Done when clicked
		types.Todos[i].Done = !types.Todos[i].Done
	}
}
