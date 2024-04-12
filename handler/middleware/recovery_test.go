package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/TechBowl-japan/go-stations/handler/router"
	testutils "github.com/TechBowl-japan/go-stations/util/test"
)

func TestPanicRecovery(t *testing.T) {
	todoDB := testutils.SetUpTestDB(t)
	r := router.NewRouter(todoDB)

	server := httptest.NewServer(r)
	defer server.Close()
	resp, err := http.Get(server.URL + "/do-panic")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("Expected status code %d, got %d", http.StatusInternalServerError, resp.StatusCode)
	}
}
