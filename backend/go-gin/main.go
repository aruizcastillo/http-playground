package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type MessageRequest struct {
	Message string `json:"message"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	port := os.Getenv("PORT")
	frontendOrigin := os.Getenv("FRONTEND_ORIGIN")

	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", frontendOrigin)
		c.Header("Access-Control-Allow-Headers", "Content-Type")
		c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	})

	router.GET("/api/message", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello from GO-GIN backend",
		})
	})

	router.POST("/api/message", func(c *gin.Context) {
		var body MessageRequest

		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid JSON",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"received": body.Message,
		})
	})

	err = router.Run(":" + port)
	if err != nil {
		panic(err)
	}
}