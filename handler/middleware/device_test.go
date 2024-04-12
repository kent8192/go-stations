package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/TechBowl-japan/go-stations/handler/middleware"
)

func TestRecoveryMiddlewareUserAgentContext(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		val := r.Context().Value(middleware.DeviceKey("User-Agent"))
		if val == nil {
			t.Error("User-Agentがコンテキストに含まれていません")
			return
		}
		if val != r.UserAgent() {
			t.Errorf("期待されるUser-Agentの値(%s)がコンテキストに含まれていませんでした。実際の値: %v", r.UserAgent(), val)
			return
		}
	})

	testServer := httptest.NewServer(middleware.Device(handler))
	defer testServer.Close()

	client := &http.Client{}
	req, _ := http.NewRequest("GET", testServer.URL, nil)
	userAgentString := "TestUserAgent"
	req.Header.Set("User-Agent", userAgentString)
	_, err := client.Do(req)
	if err != nil {
		t.Fatalf("リクエスト中にエラーが発生しました: %v", err)
	}
}
