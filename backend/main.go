package main

import (
	"log"
	"react-go-practice/api/handler"
	randomRepo "react-go-practice/repository/random"
	userRepo "react-go-practice/repository/user"
	userService "react-go-practice/services/user"

	"github.com/gin-gonic/gin"
)

func main() {
	// ランダムリポジトリの初期化
	randomRepoFactory := randomRepo.New()
	randomRepo := randomRepoFactory.Create()

	// リポジトリの初期化（データベース接続）
	if err := randomRepo.Init(); err != nil {
		log.Fatalf("Failed to initialize repository: %v", err)
	}

	// ユーザーリポジトリの初期化
	userRepoFactory := userRepo.New(randomRepo.GetQuerier())
	userRepository := userRepoFactory.Create()

	// ユーザーサービスの初期化
	userServiceInstance := userService.NewUserService(userRepository)

	// ハンドラーの初期化
	randomHandler := handler.NewRandomHandler(randomRepo)
	userHandler := handler.NewUserHandler(userServiceInstance)

	r := gin.Default()

	// CORSミドルウェアの設定
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// ルーティングの設定
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// ランダム関連のルート
	r.GET("/api/random", randomHandler.GetRandom)
	r.GET("/api/random/user", randomHandler.GetRandomByUser)
	r.POST("/api/create_random", randomHandler.CreateRundom)
	r.POST("/api/update_random", randomHandler.UpdateRandom)
	r.POST("/api/delete_random", randomHandler.DeleteRandom)

	// ユーザー関連のルート
	userRoutes := r.Group("/api/users")
	{
		userRoutes.POST("/register", userHandler.Register)
		userRoutes.POST("/login", userHandler.Login)
		userRoutes.GET("/", userHandler.ListUsers)
		userRoutes.GET("/:id", userHandler.GetUserByID)
		userRoutes.GET("/username/:username", userHandler.GetUserByUsername)
		userRoutes.PUT("/:id/password", userHandler.ChangePassword)
		userRoutes.DELETE("/:id", userHandler.DeleteUser)
	}

	// 8080ポートで起動
	r.Run(":8080")
}
