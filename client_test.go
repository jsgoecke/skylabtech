package skylabtech

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewClient(t *testing.T) {
	apiKey := "test-api-key"
	client := NewClient(apiKey)

	if client.apiKey != apiKey {
		t.Fatalf("Expected apiKey to be %s, got %s", apiKey, client.apiKey)
	}

	if client.baseURL != "https://studio.skylabtech.ai/api/public/v1" {
		t.Fatalf("Expected baseURL to be 'https://studio.skylabtech.ai/api/public/v1', got %s", client.baseURL)
	}
}

func TestRequest(t *testing.T) {
	mockResponse := `{"message":"test"}`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(mockResponse))
	}))
	defer server.Close()

	client := &Client{baseURL: server.URL, client: &http.Client{}}

	var result map[string]string
	err := client.request("GET", "/", nil, &result)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if result["message"] != "test" {
		t.Fatalf("Expected message to be 'test', got %s", result["message"])
	}
}
