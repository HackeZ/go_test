package router

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// RunRouterMiddleware ...
// @param port string
func RunRouterMiddleware(port string) {

	if port == "" || port[0] != ':' {
		log.Println("New Gin Server Failed! You need set up a Port Like this => :8080")
		log.Println("Not => ", port)
		os.Exit(1)
	}

	// Create a router without any middleware by default.
	r := gin.New()

	// Global Middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// // Per router middleware, you can add as many as you want.
	r.GET("/benchmark", MyBenchLogger(), benchEndpoint)

	// Authorization group
	// authorized := r.Group("/", AuthRequired())
	// exactly the same as:
	authorized := r.Group("/")
	// per group middleware! in this case we use the custom created
	// AuthRequired() middleware just in the "authorized" group.
	authorized.Use(AuthRequired())
	{
		authorized.POST("/login", loginEndpoint)
		authorized.POST("/submit", submitEndpoint)
		authorized.POST("/read", readEndpoint)

		// nested group
		testing := authorized.Group("testing")
		testing.GET("/analytics", analyticsEndpoint)
	}

	r.Run(port)
}

func MyBenchLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "Benmark Middleware 1.\n")
	}
}
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "Auth Middleware 1.\n")
	}
}

func benchEndpoint(c *gin.Context) {
	c.String(http.StatusOK, "Benmark Middleware 2.")
}
func loginEndpoint(c *gin.Context) {
	log.Println("Auth Middleware Login Point.")
}
func submitEndpoint(c *gin.Context) {
	c.String(http.StatusOK, "Auth Middleware Submit Point.")
}
func readEndpoint(c *gin.Context) {
	c.String(http.StatusOK, "Auth Middleware Read Point.")
}
func analyticsEndpoint(c *gin.Context) {
	c.String(http.StatusOK, "Auth Middleware Analytics Point.")
}
