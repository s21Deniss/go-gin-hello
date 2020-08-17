package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	isOK := true

	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	router := gin.New()

	// log request to stdout
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	router.Use(gin.Recovery())

	// Define handlers
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	})

	router.GET("/setOK", func(c *gin.Context) {
		isOK = true
		c.String(http.StatusOK, "now will be ok")
	})

	router.GET("/setFail", func(c *gin.Context) {
		isOK = false
		c.String(http.StatusOK, "now will fail")
	})

	router.GET("/health", func(c *gin.Context) {

		if isOK {
			c.JSON(http.StatusOK, gin.H{"status": "up"})
		} else {
			c.JSON(http.StatusServiceUnavailable, gin.H{"status": "down"})
		}

	})

	// Listen and serve on defined port
	log.Printf("Listening on port %s", port)
	router.Run(":" + port)
}
