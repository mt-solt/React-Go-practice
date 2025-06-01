package main

import (
	"react-go-practice/api/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// CORSミドルウェアの設定
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// ハンドラーの初期化
	randomHandler := handler.NewRandomHandler()

	// ルーティングの設定
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/api/random", randomHandler.GetRandom)
	// 例: GET http://localhost:8080/api/random/user?userId=123
	r.GET("/api/random/user", randomHandler.GetRandomByUser)
	r.POST("/api/create_random", randomHandler.CreateRundom)
	r.POST("/api/update_random", randomHandler.UpdateRandom)
	r.POST("/api/delete_random", randomHandler.DeleteRandom)

	// 8080ポートで起動
	r.Run(":8080")
}
