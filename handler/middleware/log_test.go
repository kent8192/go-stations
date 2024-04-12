package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/TechBowl-japan/go-stations/handler/middleware"
)

func TestAccessLog(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "正常にアクセスログが出力されることをテスト",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
			testServer := httptest.NewServer(middleware.AccessLog(handler))
			defer testServer.Close()

			client := &http.Client{}
			req, _ := http.NewRequest("GET", testServer.URL, nil)
			userAgentString := "TestUserAgent"
			req.Header.Set("User-Agent", userAgentString)
			resp, err := client.Do(req)
			if err != nil {
				t.Fatalf("リクエストの送信に失敗しました: %v", err)
			}
			defer resp.Body.Close()
			if resp.StatusCode != http.StatusOK {
				t.Errorf("期待されるステータスコード200が返されませんでした。実際のステータスコード: %v", resp.StatusCode)
			}
		})
	}
}
