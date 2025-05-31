package random

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"react-go-practice/models"

	"github.com/google/uuid"
)

type postgresRepo struct {
}

type randomColumn struct {
	id         string
	random     int64
	user_id    string
	created_at string
	updated_at string
}

func (r *postgresRepo) Create(ctx context.Context, data models.Random) error {
	return nil
}

func (r *postgresRepo) Read(ctx context.Context, id int) (models.Random, error) {
	// TODO: データベースからデータを取得する処理を実装
	tmp := models.Random{
		UUID:   "1234567890",
		Value:  1234567890,
		UserID: "1234567890",
	}
	return tmp, nil
}

func (r *postgresRepo) Update(ctx context.Context, id int, random int64) error {
	// タイムスタンプ生成
	now := time.Now()

	// UUID生成
	uuidObj, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	uuid := uuidObj.String()

	// DB登録用の構造体作成
	data := randomColumn{
		id:         uuid,
		random:     random,
		user_id:    strconv.Itoa(id),
		created_at: now.String(),
		updated_at: now.String(),
	}
	// TODO: INSERT
	fmt.Printf("%+v\n", data)

	return nil
}

func (r *postgresRepo) Delete(ctx context.Context, id int) error {
	// TODO: DELETE
	return nil
}

func (r *postgresRepo) GetAll(ctx context.Context) ([]models.Random, error) {
	return nil, nil
}

func (r *postgresRepo) QueryByUser(ctx context.Context, id int) ([]models.Random, error) {
	return nil, nil
}
