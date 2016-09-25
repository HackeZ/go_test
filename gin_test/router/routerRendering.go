package router

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// RunRouterRendering ..
// @param port string
// $ curl http://127.0.0.1:port/someJSON
// $ curl http://127.0.0.1:port/moreJSON
// $ curl http://127.0.0.1:port/someXML
// $ curl http://127.0.0.1:port/someYAML
// $ curl http://127.0.0.1:port/index
func RunRouterRendering(port string) {

	if port == "" || port[0] != ':' {
		log.Println("New Gin Server Failed! You need set up a Port Like this => :8080")
		log.Println("Not => ", port)
		os.Exit(1)
	}

	r := gin.Default()

	r.GET("/someJSON", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	r.GET("/moreJSON", func(c *gin.Context) {
		// you also can use a struct
		var msg struct {
			Name    string `json:"user"`
			Message string
			Number  int
		}
		msg.Name = "HackerZ"
		msg.Message = "Hey"
		msg.Number = 123
		// Note that msg.Name becomes "user" in the JSON
		// Will output  :   {"user": "HackerZ", "Message": "Hey", "Number": 123}
		c.JSON(http.StatusOK, msg)
	})

	r.GET("/someXML", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{"message": "Hey I am XML", "status": http.StatusOK})
	})

	r.GET("/someYAML", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{"message": "Hey I am YAML", "status": http.StatusOK})
	})

	// HTML Rendering
	r.LoadHTMLGlob("templates/*/*")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main",
		})
	})
	r.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
			"title": "Posts",
		})
	})
	r.GET("/users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
			"title": "Users",
		})
	})

	r.Run(port)
}
