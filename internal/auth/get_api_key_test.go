package auth

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAPIKeyEmpty(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/data", nil)

	_, err := GetAPIKey(req.Header)

	if err == nil {
		t.Errorf("GetAPIKey should give auth error for no headers")
	}
}

func TestGetAPIKeySuccess(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/data", nil)

	req.Header.Set("Authorization", "ApiKey test_bearer_token")
	want, err := GetAPIKey(req.Header)

	if err != nil {
		t.Errorf("GetAPIKey should not give auth error")
	}

	if want != "test_bearer_token" {
		t.Errorf("GetAPIKey should return parsed api key")
	}
}
