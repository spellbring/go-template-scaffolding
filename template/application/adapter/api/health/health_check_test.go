package health

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthCheck(t *testing.T) {
	t.Parallel()

	req, err := http.NewRequest(http.MethodGet, "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	var (
		rr      = httptest.NewRecorder()
		handler = http.NewServeMux()
	)

	handler.HandleFunc("/health", HealthCheck)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Http Status Failed '%v' expected '%v'",
			status,
			http.StatusOK,
		)
	}
}
