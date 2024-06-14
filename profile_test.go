package skylabtech

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListProfiles(t *testing.T) {
	mockResponse := `[{"id":1,"name":"Profile 1"},{"id":2,"name":"Profile 2"}]`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(mockResponse))
	}))
	defer server.Close()

	client := &Client{baseURL: server.URL, client: &http.Client{}}

	profiles, err := client.ListProfiles()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(profiles) != 2 {
		t.Fatalf("Expected 2 profiles, got %d", len(profiles))
	}
}

func TestCreateProfile(t *testing.T) {
	mockResponse := `{"id":1,"name":"New Profile"}`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(mockResponse))
	}))
	defer server.Close()

	client := &Client{baseURL: server.URL, client: &http.Client{}}
	profileParams := ProfileParams{Name: "New Profile"}

	profile, err := client.CreateProfile(profileParams)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if profile.Name != "New Profile" {
		t.Fatalf("Expected profile name to be 'New Profile', got %s", profile.Name)
	}
}

func TestGetProfile(t *testing.T) {
	mockResponse := `{"id":1,"name":"Existing Profile"}`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(mockResponse))
	}))
	defer server.Close()

	client := &Client{baseURL: server.URL, client: &http.Client{}}

	profile, err := client.GetProfile(1)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if profile.Name != "Existing Profile" {
		t.Fatalf("Expected profile name to be 'Existing Profile', got %s", profile.Name)
	}
}

func TestUpdateProfile(t *testing.T) {
	mockResponse := `{"id":1,"name":"Updated Profile"}`
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(mockResponse))
	}))
	defer server.Close()

	client := &Client{baseURL: server.URL, client: &http.Client{}}
	profileParams := ProfileParams{Name: "Updated Profile"}

	profile, err := client.UpdateProfile(1, profileParams)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if profile.Name != "Updated Profile" {
		t.Fatalf("Expected profile name to be 'Updated Profile', got %s", profile.Name)
	}
}

func TestDeleteProfile(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client := &Client{baseURL: server.URL, client: &http.Client{}}

	err := client.DeleteProfile(1)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}
