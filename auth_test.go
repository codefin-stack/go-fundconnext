package fundconnext

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestLoginWithTimeout_AllowsAuthResponseWithinConfiguredTimeout(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(40 * time.Millisecond)
		_ = json.NewEncoder(w).Encode(map[string]string{
			"access_token": "eyJhbGciOiJSUzI1NiJ9.eyJzdWIiOiJ0ZXN0In0.signature",
			"username":     "test",
			"sa_code":      "SA",
		})
	}))
	defer server.Close()

	if _, err := LoginWithTimeout(server.URL, "user", "password", "", 10*time.Millisecond); err == nil {
		t.Fatal("expected auth request to time out")
	}

	if _, err := LoginWithTimeout(server.URL, "user", "password", "", 100*time.Millisecond); err != nil {
		t.Fatalf("expected auth request within configured timeout to succeed: %v", err)
	}
}
