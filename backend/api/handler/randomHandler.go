package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"react-go-practice/repository/random"
)

type RandomHandler struct {
	randomRepoFactory random.RandomRepoFactory
}

func NewRandomHandler() *RandomHandler {
	return &RandomHandler{}
}

// 乱数取得ハンドラ
func (h *RandomHandler) GetRandom(c *gin.Context) {
	// TODO: データベースからランダムなデータを取得する処理を実装
	repo := h.randomRepoFactory.Create()
	randoms, err := repo.GetAll(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": randoms,
	})
}
