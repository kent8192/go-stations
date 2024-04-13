package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/110y/run"
	"github.com/TechBowl-japan/go-stations/db"
	. "github.com/TechBowl-japan/go-stations/handler/middleware"
	"github.com/TechBowl-japan/go-stations/handler/router"
	"github.com/joho/godotenv"
)

func main() {
	err := loadEnv()
	if err != nil {
		log.Fatalln("Failed to load .env", err)
	}

	run.Run(func(ctx context.Context) int {
		if err := setUpServer(ctx); err != nil {
			log.Fatalln("main: failed to exit successfully, err =", err)
			return 1
		}
		return 0
	})
}

func loadEnv() error {
	err := godotenv.Load(".env")
	return err
}

func setUpServer(ctx context.Context) error {
	// config values
	const (
		defaultPort   = ":8080"
		defaultDBPath = ".sqlite3/todo.db"
	)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = defaultDBPath
	}

	// set time zone
	var err error
	time.Local, err = time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return err
	}

	// set up sqlite3
	todoDB, err := db.NewDB(dbPath)
	if err != nil {
		return err
	}
	defer todoDB.Close()

	// NOTE: 新しいエンドポイントの登録はrouter.NewRouterの内部で行うようにする
	mux := router.NewRouter(todoDB)

	// setUpServer関数内の一部
	server := &http.Server{
		Addr:    port,
		Handler: mux,
	}

	// Graceful Shutdownの設定
	GracefulShutdown(server)

	go func() {
		<-ctx.Done()
	}()

	log.Println("Starting server on", port)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}
