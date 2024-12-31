package routes

import (
	"fmt"
	"net/http"
	"strconv"
	"todo/token"
	"todo/types"

	"github.com/gin-gonic/gin"
)

var Role string

func RoutesInitializer() {
	router := gin.Default()

	router.Static("../static", "./static")
	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html",
			gin.H{
				"Todos":    types.Todos,              // display the current todos available
				"LoggedIn": types.LoggedInUser != "", // return true if the loggedInUser is not an empty string
				"Username": types.LoggedInUser,
				"Role": Role,
			})
	})

	router.POST("/add", authenticateMiddleware, func(ctx *gin.Context) {
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

	// login route
	router.POST("/login", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		Role = username

		// this is dummy credential check, we dont have a database yet
		if (username == "employee" && password == "password") || (username == "senior" && password == "password") {
			tokenString, err := token.CreateToken(username)
			if err != nil {
				c.String(http.StatusInternalServerError, "error creating token")
				return
			}
			types.LoggedInUser = username
			fmt.Printf("token created %q\n", tokenString)
			c.SetCookie("token", tokenString, 9, "/", "localhost", false, true)
			c.Redirect(http.StatusSeeOther, "/") // redirect to home page
		} else {
			c.String(http.StatusUnauthorized, "Invalid credentials")
		}
	})

	router.GET("/logout", func(c *gin.Context) {
		types.LoggedInUser = ""
		c.SetCookie("token", "", -1, "/", "localhost", false, true)
		c.Redirect(http.StatusSeeOther, "/")
	})

	router.Run(":5555")
}

func toggleIndex(ind string) {
	i, _ := strconv.Atoi(ind)
	if i >= 0 && i < len(types.Todos) {
		// negate the value of Todo.Done when clicked
		types.Todos[i].Done = !types.Todos[i].Done
	}
}

func authenticateMiddleware(c *gin.Context) {
	// retrieve the token from the cookie
	tokenString, err := c.Cookie("token")
	if err != nil {
		fmt.Println("token missing in cookie")
		c.Redirect(http.StatusSeeOther, "/")
		c.Abort()
		return
	}

	// verify the token
	token, err := token.VerifyToken(tokenString)
	if err != nil {
		fmt.Printf("token verification failed %v\\n", err)
		c.Redirect(http.StatusSeeOther, "/login")
		c.Abort()
		return
	}

	fmt.Printf("token verfification successfully, claims %+v\\n", token.Claims)
	// continue with the next middleware or route handler
	c.Next()
}
