package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type RandomHandler struct {
	// 必要に応じて依存関係を注入
}

func NewRandomHandler() *RandomHandler {
	return &RandomHandler{}
}

// GetRandom はランダムなデータを取得するハンドラー
func (h *RandomHandler) GetRandom(c *gin.Context) {
	// TODO: データベースからランダムなデータを取得する処理を実装
	c.JSON(http.StatusOK, gin.H{
		"message": "ランダムなデータを取得しました",
		"data":    nil, // ここに実際のデータを設定
	})
}
