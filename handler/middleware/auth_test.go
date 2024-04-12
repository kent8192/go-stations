package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/TechBowl-japan/go-stations/handler/middleware"
)

func Test_Name(t *testing.T) {
	tests := []struct {
		name string
		mes  string
	}{
		{
			name: "Valid id and pass",
			mes:  "ID unmatched",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
			testServer := httptest.NewServer(middleware.Auth(handler))
			defer testServer.Close()
			id := "id"
			pass := "password"
			os.Setenv("ID", id)
			os.Setenv("PASSWORD", pass)
			client := &http.Client{}
			req, _ := http.NewRequest("GET", testServer.URL, nil)
			req.SetBasicAuth(id, pass)
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
