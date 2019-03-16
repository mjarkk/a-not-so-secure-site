package server

import (
	"errors"
	"fmt"

	"github.com/mjarkk/a-not-so-secure-site/credentials"
	"github.com/mjarkk/a-not-so-secure-site/database"
	"github.com/mjarkk/a-not-so-secure-site/utils"

	"github.com/gin-gonic/gin"
	"github.com/mjarkk/a-not-so-secure-site/templates"
)

// Init sets up the server
func Init() error {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r.GET("/", templates.Overview)
	r.GET("/post/:id", templates.Post)
	r.GET("/login", func(c *gin.Context) {
		templates.Login(c)
	})
	r.POST("/login", func(c *gin.Context) {
		username := c.PostForm("Username")
		password := c.PostForm("Password")
		if len(username) == 0 || len(password) == 0 {
			templates.Login(c, templates.LoginErr{
				What: "login",
				Err:  "Username and/or Password can't be empty",
			})
			return
		}

		err := logedinCookies(c, username, password)
		if len(username) == 0 || len(password) == 0 {
			templates.Login(c, templates.LoginErr{
				What: "login",
				Err:  err.Error(),
			})
			return
		}

		templates.Overview(c)
	})
	r.POST("/register", func(c *gin.Context) {
		username := c.PostForm("Username")
		password := c.PostForm("Password")
		if len(username) == 0 || len(password) == 0 {
			templates.Login(c, templates.LoginErr{
				What: "register",
				Err:  "Username and/or Password can't be empty",
			})
			return
		}

		user, err := database.NewUser(username, password)
		if err != nil {
			templates.Login(c, templates.LoginErr{
				What: "register",
				Err:  err.Error(),
			})
			return
		}

		err = logedinCookies(c, user.Username, user.Password)
		if err != nil {
			templates.Login(c, templates.LoginErr{
				What: "register",
				Err:  err.Error(),
			})
			return
		}

		templates.Overview(c)
	})
	fmt.Println("server running on localhost:8080")
	return r.Run()
}

func logedinCookies(c *gin.Context, username, password string) error {
	user, err := database.User("username", username)
	if err != nil {
		return err
	}

	if user.Password != password {
		return errors.New("password wrong")
	}

	key, err := credentials.CreateSession(user)
	if err != nil {
		return err
	}

	second := 1
	minute := 60
	hour := 60
	day := 24
	mounth := 32
	c.SetCookie("sessionKey", key, second*minute*hour*day*mounth, "/", utils.GetDomain(c.GetHeader("Origin"), true), false, true)
	return nil
}
