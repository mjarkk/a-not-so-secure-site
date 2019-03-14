package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mjarkk/a-not-so-secure-site/templates"
)

// Init sets up the server
func Init() error {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		templates.GetTemplate("overview", struct{}{})
	})
	r.GET("/post/:id", func(c *gin.Context) {
		templates.GetTemplate("post", struct{}{})
	})
	r.GET("/login", func(c *gin.Context) {
		templates.GetTemplate("login", struct{}{})
	})
	r.GET("/register", func(c *gin.Context) {
		templates.GetTemplate("register", struct{}{})
	})
	fmt.Println("server running on localhost:8080")
	return r.Run()
}
