package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"time"

	"seg_service/config"
	"seg_service/internal/handler"
	"seg_service/internal/repository"
	"seg_service/internal/service"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	cfg := config.LoadConfig()
	db, err := sql.Open("postgres", cfg.PostgresDSN)
	if err != nil {
		log.Fatalf("failed to connect to postgres: %v", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("failed to ping postgres: %v", err)
	}
	log.Println("postgres connected successfully")
	defer db.Close()

	rdb := redis.NewClient(&redis.Options{
		Addr: cfg.RedisAddr,
	})
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		log.Fatalf("failed to connect to redis: %v", err)
	}

	userRepo := repository.NewPostgresRepository(db)
	segmentRepo := userRepo
	cache := repository.NewRedisCache(rdb)

	service := service.NewSegmentService(userRepo, segmentRepo, cache)
	h := handler.NewHandler(service)

	mux := http.NewServeMux()
	handler.RegisterRoutes(mux, h)

	server := &http.Server{
		Addr:         cfg.HTTPPort,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	log.Printf("Server started at %s, bd = %s", cfg.HTTPPort, cfg.PostgresDSN)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server error: %v", err)
	}
}
