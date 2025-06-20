package random

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	"react-go-practice/models"

	_ "github.com/lib/pq"
)

var db *sql.DB

type postgresRepo struct {
}

type randomColumn struct {
	id         string
	random     int64
	user_id    string
	created_at string
	updated_at string
}

func (r *postgresRepo) Init() error {
	dsn := os.Getenv("POSTGRES_DSN")
	if dsn == "" {
		dsn = "host=postgres port=5432 user=randomuser password=randompass dbname=randomdb sslmode=disable"
	}
	database, err := sql.Open("postgres", dsn)
	if err != nil {
		return err
	}
	if err := database.Ping(); err != nil {
		return err
	}
	db = database
	return nil
}

func (r *postgresRepo) Create(ctx context.Context, data models.Random) error {
	// タイムスタンプ生成
	now := time.Now()

	// DB登録用の構造体作成
	ins := randomColumn{
		id:         data.UUID,
		random:     data.Value,
		user_id:    data.UserID,
		created_at: now.String(),
		updated_at: now.String(),
	}
	// TODO: INSERT
	fmt.Printf("%+v\n", ins)
	return nil
}

func (r *postgresRepo) Read(ctx context.Context, uuid string) (models.Random, error) {
	// TODO: データベースからデータを取得する処理を実装
	tmp := models.Random{
		UUID:   "1234567890",
		Value:  1234567890,
		UserID: "1234567890",
	}
	return tmp, nil
}

func (r *postgresRepo) Update(ctx context.Context, uuid string, random int64) error {
	return nil
}

func (r *postgresRepo) Delete(ctx context.Context, uuid string) error {
	// TODO: DELETE
	return nil
}

func (r *postgresRepo) GetAll(ctx context.Context) ([]models.Random, error) {
	// TODO: データベースからデータを取得する処理を実装
	tmp := models.Random{
		UUID:   "1234567890",
		Value:  1234567890,
		UserID: "1234567890",
	}
	return []models.Random{tmp}, nil
}

func (r *postgresRepo) QueryByUser(ctx context.Context, userId string) ([]models.Random, error) {
	// TODO: データベースからデータを取得する処理を実装
	tmp := models.Random{
		UUID:   "1234567890",
		Value:  1234567890,
		UserID: "1234567890",
	}
	return []models.Random{tmp}, nil
}
