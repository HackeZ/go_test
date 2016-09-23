package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// You Should visit
// http://127.0.0.1:8080/welcome/name/{yourname}/action/{whatyoudoing}?mess={andthemessage}

func main() {
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	// ===-                   -===
	// ===- Handle GET Method -===
	// ===-                   -===
	// Redirect to /welcome
	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/welcome/hackerz")
	})
	// Handle Status File
	router.StaticFile("/favicon.ico", "./upload/z.ico")
	// This handler will match /welcome/john but will not match neither /welcome/ or /welcome
	router.GET("/welcome/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s! How nice to meet you!", name)
	})

	// This router will match /welcome/john/ and also /welcome/john/send
	// If no other routers match /welcome/john, it will redirect to /welcome/john/
	router.GET("/welcome/:name/*action", func(c *gin.Context) {
		// Handle Param
		name := c.Param("name")
		action := c.Param("action")
		fmt.Println(action)
		// Handle Query
		mess := c.DefaultQuery("mess", "nothing")
		// lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("firstname")
		c.String(http.StatusOK, "Hei %s, you are %sing now which message is \"%s\"", name, action, mess)
	})
	// ===-                   -===
	// ===- Handle GET Method -===
	// ===-                   -===

	// ===-                    -===
	// ===- Handle POST Method -===
	// ===-                    -===
	router.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		username := c.DefaultPostForm("username", "HackerZ")

		c.JSON(200, gin.H{
			"status":   "posted",
			"username": username,
			"message":  message,
		})
	})

	router.POST("/upload", func(c *gin.Context) {
		file, header, err := c.Request.FormFile("upload")
		if err != nil {
			log.Fatal(err)
		}

		fileName := header.Filename
		fmt.Println("There I Get a File which is ", fileName)

		out, err := os.Create("./upload/" + fileName + ".png")
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()
		_, err = io.Copy(out, file)
		if err != nil {
			log.Fatal(err)
		}
	})
	// ===-                    -===
	// ===- Handle POST Method -===
	// ===-                    -===

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	router.Run()
}
