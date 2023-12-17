package main

import (
	"auth/internal/api"
	"auth/internal/user"
	"auth/postgres"
	"auth/postgres/sqlc"
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
)

func main() {
	ctx := context.Background()
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	router := gin.Default()

	conn, err := pgx.Connect(ctx, "host=127.0.0.1 port=5432 user=postgres password=postgres dbname=auth sslmode=disable")
	if err != nil {
		logger.Fatal("error connecting database", zap.Error(err))
	}
	defer conn.Close(ctx)

	queries := sqlc.New(conn)
	repo := postgres.NewRepository(queries)
	srv := user.NewUserService(repo, logger)
	handler := api.NewHandler(srv, logger)

	router.POST("/signup", handler.SignUp)
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	errs := make(chan error, 1)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		errs <- router.Run(":8080")
	}()
	logger.Info("shutting down", zap.Error(<-errs))
}
