package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
        "os"
)

func main() {
	// Create a new Gin router
	router := gin.Default()

	// Define a route for the "Hello, World!" endpoint
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})

        router.GET("/printenv", func(c *gin.Context) {
                myEnv := os.Getenv("MySuperValue")
	
		if myEnv == "" {
		   myEnv = "ValueNotFound"
                }

                c.JSON(http.StatusOK, gin.H{
                        "envValue": myEnv,
                })
        })

	// Run the server on port 8080
	router.Run(":8080")
}
