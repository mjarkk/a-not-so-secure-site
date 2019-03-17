package templates

import (
	"io/ioutil"
	"net/http"

	"github.com/mjarkk/a-not-so-secure-site/credentials"

	"github.com/gin-gonic/gin"
	"github.com/mjarkk/a-not-so-secure-site/database"
)

func renderOut(c *gin.Context, data string, err error) {
	if err != nil {
		c.Data(http.StatusInternalServerError, "text/html", []byte(err.Error()))
		return
	}
	c.Data(http.StatusOK, "text/html", []byte(data))
}

func genMeta(c *gin.Context) Meta {
	toReturn := Meta{}

	cssData, err := ioutil.ReadFile("./templates/inline.css")
	if err == nil {
		toReturn.CSS = "<style>" + string(cssData) + "</style>"
	}

	key, err := c.Cookie("sessionKey")
	if err == nil {
		if user, ok := credentials.GetSession(key); ok {
			toReturn.LogedIn = true
			toReturn.User = user
		}
	}

	header, err := GetTemplate("header", toReturn)
	if err == nil {
		toReturn.Header = header
	} else {
		toReturn.Header = "<div class=\"header\">Error while rendering the header: " + err.Error() + "</div>"
	}

	return toReturn
}

// Meta contains meta information about the site
type Meta struct {
	CSS     string
	LogedIn bool
	User    database.UserT
	Header  string
}

// OverViewT is the data for the overview template
type OverViewT struct {
	Meta     Meta
	Posts    []database.PostT
	HasError bool
	Error    string
}

// Overview generates the overview page
func Overview(c *gin.Context) {
	meta := genMeta(c)

	posts, err := database.Posts()

	data := OverViewT{
		Meta:  meta,
		Posts: posts,
	}

	if err != nil {
		data.HasError = true
		data.Error = err.Error()
	}

	out, err := GetTemplate("overview", data)
	renderOut(c, out, err)
}

// PostT is the data type for the /post/:id route
type PostT struct {
	Meta     Meta
	Post     database.PostT
	HasError bool
	Error    string
}

// Post generates the post page
func Post(c *gin.Context) {
	meta := genMeta(c)
	post, err := database.Post("ID", c.Param("id"))

	data := PostT{
		Meta: meta,
		Post: post,
	}

	if err != nil {
		data.HasError = true
		data.Error = err.Error()
	}

	out, err := GetTemplate("post", data)
	renderOut(c, out, err)
}

// LoginT is the data type for the /login route
type LoginT struct {
	Meta   Meta
	Errors []LoginErr
}

// LoginErr the content of a login err
type LoginErr struct {
	What string
	Err  string
}

// Login generates the login page
func Login(c *gin.Context, loginErrors ...LoginErr) {
	out, err := GetTemplate("login", LoginT{
		Meta:   genMeta(c),
		Errors: loginErrors,
	})
	renderOut(c, out, err)
}

// CreatePostT is the data type for the /create route
type CreatePostT struct {
	Meta Meta
}

// CreatePost is the handeler for the /craete route
func CreatePost(c *gin.Context) {
	meta := genMeta(c)

	out, err := GetTemplate("create", CreatePostT{
		Meta: meta,
	})
	renderOut(c, out, err)
}
