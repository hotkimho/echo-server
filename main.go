package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/heartbeat", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "success",
		})
	})

	r.Run("127.0.0.1:8080")
}
