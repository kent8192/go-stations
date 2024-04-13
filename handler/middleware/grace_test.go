package middleware

import (
	"net/http"
	"net/http/httptest"
	"os"
	"syscall"
	"testing"
	"time"
)

func TestGracefulShutdown(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	httpServer := &http.Server{
		Addr:    server.URL,
		Handler: server.Config.Handler,
	}

	go GracefulShutdown(httpServer)

	// サーバーが起動していることを確認
	resp, err := http.Get(server.URL)
	if err != nil {
		t.Fatalf("サーバーへのリクエストに失敗しました: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("期待するステータスコード200が返ってきませんでした: %v", resp.StatusCode)
	}

	// SIGINTシグナルを送信してGracefulShutdownをトリガー
	proc, err := os.FindProcess(os.Getpid())
	if err != nil {
		t.Fatalf("プロセスの取得に失敗しました: %v", err)
	}
	proc.Signal(syscall.SIGINT)

	// シャットダウンが完了するのを待つ
	time.Sleep(1 * time.Second)

	// シャットダウン後はサーバーへのリクエストが失敗することを確認
	_, err = http.Get(server.URL)
	if err == nil {
		t.Fatalf("サーバーがシャットダウン後もリクエストに応答しています")
	}
}
