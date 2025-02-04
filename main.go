package main

import (
	"fmt"
	"golang_practice/controller"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// Set example variable
		c.Set("example", "12345")

		// before request

		c.Next()

		// after request
		latency := time.Since(t)
		fmt.Println("latency: ", latency)

		// access the status we are sending
		status := c.Writer.Status()
		fmt.Println("status: ", status)
	}
}

func main() {
	r := gin.New()

	r.Use(Logger())

	controller.InitPing(r)

	r.Run()
}

// test
