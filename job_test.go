package skylabtech

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListJobs(t *testing.T) {
	mockResponse := `[{"id":1,"name":"Job 1"},{"id":2,"name":"Job 2"}]`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(mockResponse))
	}))
	defer server.Close()

	client := &Client{baseURL: server.URL, client: &http.Client{}}

	jobs, err := client.ListJobs()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(jobs) != 2 {
		t.Fatalf("Expected 2 jobs, got %d", len(jobs))
	}
}

func TestCreateJob(t *testing.T) {
	mockResponse := `{"id":1,"name":"New Job"}`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(mockResponse))
	}))
	defer server.Close()

	client := &Client{baseURL: server.URL, client: &http.Client{}}
	jobParams := JobParams{Name: "New Job"}

	job, err := client.CreateJob(jobParams)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if job.Name != "New Job" {
		t.Fatalf("Expected job name to be 'New Job', got %s", job.Name)
	}
}

func TestGetJob(t *testing.T) {
	mockResponse := `{"id":1,"name":"Existing Job"}`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(mockResponse))
	}))
	defer server.Close()

	client := &Client{baseURL: server.URL, client: &http.Client{}}

	job, err := client.GetJob(1)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if job.Name != "Existing Job" {
		t.Fatalf("Expected job name to be 'Existing Job', got %s", job.Name)
	}
}
