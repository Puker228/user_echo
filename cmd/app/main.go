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

	"github.com/Puker228/user_echo/internal/database/postgresql"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

type UserStats struct {
	Android_version string `json:"android_version"`
	Device_model    string `json:"device_model"`
	Manufacturer    string `json:"manufacturer"`
	Total_ram_gb    int    `json:"total_ram_gb"`
	App_version     string `json:"app_version"`
}

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

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	err = postgresql.InitDB(db)
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()
	router.POST("/stats", func(c *gin.Context) {
		var stats UserStats

		if err := c.ShouldBindJSON(&stats); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		_, err := db.ExecContext(c.Request.Context(),
			`INSERT INTO users (android_version, device_model, manufacturer, total_ram_gb, app_version)
			 VALUES ($1, $2, $3, $4, $5)`,
			stats.Android_version, stats.Device_model, stats.Manufacturer, stats.Total_ram_gb, stats.App_version,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "saved"})
	})

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
