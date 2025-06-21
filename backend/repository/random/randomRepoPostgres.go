package random

import (
	"context"
	"database/sql"
	"log"
	"os"
	"time"

	"react-go-practice/db"
	"react-go-practice/models"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type postgresRepo struct {
	queries *db.Queries
}

func (r *postgresRepo) Init() error {
	log.Printf("[DEBUG] Init called")

	dsn := os.Getenv("POSTGRES_DSN")
	// デフォルトDB接続先を設定
	if dsn == "" {
		dsn = "host=postgres port=5432 user=randomuser password=randompass dbname=randomdb sslmode=disable"
	}

	log.Printf("[DEBUG] Using DSN: %s", dsn)

	// DB接続
	database, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Printf("[ERROR] Failed to open database: %v", err)
		return err
	}

	// DB接続確認
	if err := database.Ping(); err != nil {
		log.Printf("[ERROR] Failed to ping database: %v", err)
		return err
	}

	log.Printf("[DEBUG] Database connection successful")

	// sqlcで生成されたQueriesを使用
	r.queries = db.New(database)
	log.Printf("[DEBUG] Queries initialized")

	return nil
}

func (r *postgresRepo) Create(ctx context.Context, data models.Random) error {
	// UUIDをパース
	id, err := uuid.Parse(data.UUID)
	if err != nil {
		return err
	}

	// タイムスタンプ生成
	now := time.Now()

	// sqlcで生成されたCreateRandomメソッドを使用
	arg := db.CreateRandomParams{
		ID:        id,
		RandomVal: data.Value,
		UserID:    data.UserID,
		CreatedAt: now,
		UpdatedAt: now,
	}

	return r.queries.CreateRandom(ctx, arg)
}

func (r *postgresRepo) Read(ctx context.Context, id string) (models.Random, error) {
	// UUIDをパース
	uuid, err := uuid.Parse(id)
	if err != nil {
		return models.Random{}, err
	}

	// sqlcで生成されたGetRandomメソッドを使用
	random, err := r.queries.GetRandom(ctx, uuid)
	if err != nil {
		return models.Random{}, err
	}

	// db.Randomをmodels.Randomに変換
	return models.Random{
		UUID:   random.ID.String(),
		Value:  random.RandomVal,
		UserID: random.UserID,
	}, nil
}

func (r *postgresRepo) Update(ctx context.Context, id string, random int64) error {
	// UUIDをパース
	uuid, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	// タイムスタンプ生成
	now := time.Now()

	// sqlcで生成されたUpdateRandomメソッドを使用
	arg := db.UpdateRandomParams{
		ID:        uuid,
		RandomVal: random,
		UpdatedAt: now,
	}

	return r.queries.UpdateRandom(ctx, arg)
}

func (r *postgresRepo) Delete(ctx context.Context, id string) error {
	// UUIDをパース
	uuid, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	// sqlcで生成されたDeleteRandomメソッドを使用
	return r.queries.DeleteRandom(ctx, uuid)
}

func (r *postgresRepo) GetAll(ctx context.Context) ([]models.Random, error) {
	log.Printf("[DEBUG] GetAll called")

	// sqlcで生成されたGetAllRandomsメソッドを使用
	randoms, err := r.queries.GetAllRandoms(ctx)
	if err != nil {
		log.Printf("[ERROR] GetAllRandoms failed: %v", err)
		return nil, err
	}

	// デバッグログ
	log.Printf("[DEBUG] GetAllRandoms result count: %d", len(randoms))
	log.Printf("[DEBUG] GetAllRandoms result: %+v", randoms)

	// db.Randomのスライスをmodels.Randomのスライスに変換
	result := make([]models.Random, len(randoms))
	for i, random := range randoms {
		log.Printf("[DEBUG] Converting random[%d]: ID=%s, Value=%d, UserID=%s",
			i, random.ID.String(), random.RandomVal, random.UserID)

		result[i] = models.Random{
			UUID:   random.ID.String(),
			Value:  random.RandomVal,
			UserID: random.UserID,
		}
	}

	// デバッグログ
	log.Printf("[DEBUG] Converted result count: %d", len(result))
	log.Printf("[DEBUG] Converted result: %+v", result)

	return result, nil
}

func (r *postgresRepo) QueryByUser(ctx context.Context, userId string) ([]models.Random, error) {
	// sqlcで生成されたGetRandomsByUserメソッドを使用
	randoms, err := r.queries.GetRandomsByUser(ctx, userId)
	if err != nil {
		return nil, err
	}

	// db.Randomのスライスをmodels.Randomのスライスに変換
	result := make([]models.Random, len(randoms))
	for i, random := range randoms {
		result[i] = models.Random{
			UUID:   random.ID.String(),
			Value:  random.RandomVal,
			UserID: random.UserID,
		}
	}

	return result, nil
}

func (r *postgresRepo) GetQuerier() db.Querier {
	return r.queries
}
