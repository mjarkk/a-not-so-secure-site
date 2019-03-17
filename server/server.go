package server

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mjarkk/a-not-so-secure-site/credentials"
	"github.com/mjarkk/a-not-so-secure-site/database"
	"github.com/mjarkk/a-not-so-secure-site/templates"
	"github.com/mjarkk/a-not-so-secure-site/utils"
)

// Init sets up the server
func Init() error {
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r.GET("/", templates.Overview)
	r.GET("/post/:id", templates.Post)
	r.GET("/login", login)
	r.GET("/create", templates.CreatePost)
	r.POST("/create", createPost)
	r.POST("/login", postLogin)
	r.POST("/register", register)
	fmt.Println("server running on localhost:8080")
	return r.Run()
}

func createPost(c *gin.Context) {
	title := c.PostForm("Title")
	content := c.PostForm("Content")

	key, err := c.Cookie("sessionKey")
	if err != nil {
		c.Redirect(302, utils.MKFullPath(c, "/"))
		return
	}
	user, ok := credentials.GetSession(key)
	if !ok {
		c.Redirect(302, utils.MKFullPath(c, "/"))
		return
	}

	database.NewPost(title, content, user.ID)

	c.Redirect(302, utils.MKFullPath(c, "/"))
}

func login(c *gin.Context) {
	templates.Login(c)
}

func register(c *gin.Context) {
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

	c.Redirect(302, utils.MKFullPath(c, "/"))
}

func postLogin(c *gin.Context) {
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
	if err != nil {
		templates.Login(c, templates.LoginErr{
			What: "login",
			Err:  err.Error(),
		})
		return
	}

	c.Redirect(302, utils.MKFullPath(c, "/"))
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
