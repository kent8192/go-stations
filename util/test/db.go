package testutils

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/TechBowl-japan/go-stations/db"
)

func SetUpTestDB(t *testing.T) *sql.DB {
	dbPath := "./temp_test.db"
	if err := os.Setenv("DB_PATH", dbPath); err != nil {
		log.Fatalf("dbPathのセットに失敗しました。%v", err)
	}

	todoDB, err := db.NewDB(dbPath)
	if err != nil {
		log.Fatalf("データベースの作成に失敗しました: %v", err)
	}
	t.Cleanup(func() {
		if err := todoDB.Close(); err != nil {
			log.Fatalf("データベースのクローズに失敗しました: %v", err)
		}
		if err := os.Remove(dbPath); err != nil {
			log.Fatalf("テスト用のDBファイルの削除に失敗しました: %v", err)
		}
	})

	return todoDB
}
