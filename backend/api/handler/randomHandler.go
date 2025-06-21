package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"react-go-practice/models"
	randomRepo "react-go-practice/repository/random"
	randomService "react-go-practice/services/random"
)

type RandomHandler struct {
	repo randomRepo.RandomRepository
}

func NewRandomHandler(repo randomRepo.RandomRepository) *RandomHandler {
	return &RandomHandler{
		repo: repo,
	}
}

// 乱数全取得
func (h *RandomHandler) GetRandom(c *gin.Context) {
	randoms, err := h.repo.GetAll(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": randoms,
	})
}

func (h *RandomHandler) GetRandomByUser(c *gin.Context) {
	// クエリパラメータからユーザーIDを取得
	// GET /api/random/user?userId=123
	userId := c.Query("userId")
	if userId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "userId is required"})
		return
	}

	// ユーザーIDに紐づく乱数を取得
	randoms, err := h.repo.QueryByUser(c, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 1件もデータがない場合
	if len(randoms) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No random numbers found for this user"})
		return
	}

	// 検索結果返却
	c.JSON(http.StatusOK, gin.H{
		"data": randoms,
	})
}

// 乱数生成・登録
func (h *RandomHandler) CreateRundom(c *gin.Context) {
	// リクエストボディからユーザーIDを取得
	var requestBody struct {
		UserID string `json:"userId"`
	}
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	randomGen := randomService.NewRandomGenerator()
	random := randomGen.Generate()

	// 乱数データのID
	uuidObj, err := uuid.NewRandom()
	if err != nil {
		// サーバエラー
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	uuid := uuidObj.String()

	// 乱数モデル生成
	data := models.Random{
		UserID: requestBody.UserID,
		Value:  random,
		UUID:   uuid,
	}

	err = h.repo.Create(c, data)

	if err != nil {
		// サーバエラー
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 処理成功
	c.JSON(http.StatusOK, gin.H{})
}

// 乱数更新
func (h *RandomHandler) UpdateRandom(c *gin.Context) {
	// リクエストボディからID取得
	var requestBody struct {
		Uuid string `json:"id"`
	}
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// 更新用の新規乱数生成
	randomGen := randomService.NewRandomGenerator()
	random := randomGen.Generate()

	// 更新
	err := h.repo.Update(c, requestBody.Uuid, random)
	if err != nil {
		// サーバエラー
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 処理成功
	c.JSON(http.StatusOK, gin.H{})
}

// 乱数削除
func (h *RandomHandler) DeleteRandom(c *gin.Context) {
	// リクエストボディからID取得
	var requestBody struct {
		Uuid string `json:"id"`
	}
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// 削除
	err := h.repo.Delete(c, requestBody.Uuid)
	if err != nil {
		// サーバエラー
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 処理成功
	c.JSON(http.StatusOK, gin.H{})
}
