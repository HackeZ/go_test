package router

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// Login Object By JSON
// User     string
// Password string
type Login struct {
	User     string `form:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// RunRouterWithModel ..
// @param port
// $ curl -d '{"user":"hackerz", "password":"123"}' 127.0.0.1:port/loginJSON
// $ curl --form user=hackerz --form password=123 127.0.0.1:8080/loginForm
func RunRouterWithModel(port string) {

	if port == "" || port[0] != ':' {
		log.Println("New Gin Server Failed! You need set up a Port Like this => :8080")
		log.Println("Not => ", port)
		os.Exit(1)
	}

	// Router Start.
	router := gin.Default()

	// Example for binding JSON ({"user": "hackerz", "password": "123"})
	router.POST("/loginJSON", func(c *gin.Context) {
		var loginJSON Login
		if c.BindJSON(&loginJSON) == nil {
			if loginJSON.User == "hackerz" && loginJSON.Password == "123" {
				c.JSON(http.StatusOK, gin.H{"status": "you are logged in by JSON"})
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			}
		}
	})

	// Example for binding a HTML form (user=hackerz&password=123)
	router.POST("/loginForm", func(c *gin.Context) {
		var loginForm Login
		// This will infer what binder to use depending on the content-type header.
		if c.Bind(&loginForm) == nil {
			if loginForm.User == "hackerz" && loginForm.Password == "123" {
				c.JSON(http.StatusOK, gin.H{"status": "you are logged in by From"})
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			}
		}
	})

	router.Run(port)
}
