package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Puker228/user_echo/internal/handler"
	"github.com/Puker228/user_echo/internal/repository/postgresql"
	"github.com/Puker228/user_echo/internal/usecase"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

func main() {
	cfg := pq.Config{
		Host:     "localhost",
		Port:     5433,
		User:     "postgres",
		Database: "postgres",
		Password: "postgres",
		SSLMode:  "disable",
	}

	c, err := pq.NewConnectorConfig(cfg)
	if err != nil {
		log.Fatal(err)
	}

	db := sql.OpenDB(c)
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	if err = postgresql.InitDB(db); err != nil {
		log.Fatal(err)
	}

	repo := postgresql.NewStatsRepository(db)
	uc := usecase.NewStatsUseCase(repo)
	h := handler.NewStatsHandler(uc)

	router := gin.Default()
	h.RegisterRoutes(router)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router.Handler(),
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Println("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
