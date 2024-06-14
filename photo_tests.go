package skylabtech

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreatePhoto(t *testing.T) {
	mockResponse := `{"id":1,"name":"New Photo","job_id":1}`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(mockResponse))
	}))
	defer server.Close()

	client := &Client{baseURL: server.URL, client: &http.Client{}}
	photoParams := PhotoParams{JobID: 1, Name: "New Photo", Key: "photo-key"}

	photo, err := client.CreatePhoto(photoParams)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if photo.Name != "New Photo" {
		t.Fatalf("Expected photo name to be 'New Photo', got %s", photo.Name)
	}
}

func TestGetPhoto(t *testing.T) {
	mockResponse := `{"id":1,"name":"Existing Photo","job_id":1}`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(mockResponse))
	}))
	defer server.Close()

	client := &Client{baseURL: server.URL, client: &http.Client{}}

	photo, err := client.GetPhoto(1)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if photo.Name != "Existing Photo" {
		t.Fatalf("Expected photo name to be 'Existing Photo', got %s", photo.Name)
	}
}

func TestGetPhotoUploadURL(t *testing.T) {
	mockResponse := `{"url":"https://s3.amazonaws.com/upload","key":"photo-key"}`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(mockResponse))
	}))
	defer server.Close()

	client := &Client{baseURL: server.URL, client: &http.Client{}}
	params := PhotoUploadUrlParams{PhotoID: 1, Type: "jpg"}

	uploadURL, err := client.GetPhotoUploadURL(params)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if uploadURL.URL != "https://s3.amazonaws.com/upload" {
		t.Fatalf("Expected upload URL to be 'https://s3.amazonaws.com/upload', got %s", uploadURL.URL)
	}
}

func TestListPhotosForJob(t *testing.T) {
	mockResponse := `[{"id":1,"name":"Photo 1","job_id":1},{"id":2,"name":"Photo 2","job_id":1}]`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(mockResponse))
	}))
	defer server.Close()

	client := &Client{baseURL: server.URL, client: &http.Client{}}

	photos, err := client.ListPhotosForJob(1)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(photos) != 2 {
		t.Fatalf("Expected 2 photos, got %d", len(photos))
	}
}
