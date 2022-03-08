package storage

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Storage struct {
	db    *pgxpool.Pool
	cache *redis.Client
}

func (s *Storage) DB() *pgxpool.Pool {
	return s.db
}

func (s *Storage) Cache() *redis.Client {
	return s.cache
}

func init() {
	ctx := context.Background()
	storage := New(ctx)
	storage.CreateTable(ctx)
}

func New(ctx context.Context) *Storage {
	postgresPassword := os.Getenv("POSTGRES_PASSWORD")
	postgresUsername := os.Getenv("POSTGRES_USERNAME")
	postgresIPAddr := os.Getenv("POSTGRES_IP")
	redisIPAddr := os.Getenv("REDIS_IP")

	log.Println(postgresPassword)

	conn, err := pgxpool.Connect(ctx, fmt.Sprintf("postgres://%s:%s@%s/postgres", postgresUsername, postgresPassword, postgresIPAddr))
	if err != nil {
		log.Fatalf("ERROR: unable to connect PSQL: %v", err)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     redisIPAddr,
		Password: "",
		DB:       0,
	})

	return &Storage{
		db:    conn,
		cache: rdb,
	}
}

func (s *Storage) CreateTable(ctx context.Context) {
	res, err := s.db.Query(ctx, CreateTable)
	if err != nil {
		log.Printf("ERROR: failed to create todos table: %v\n", err)
	}

	log.Printf("successfully created todos table: %v\n", res)
}
