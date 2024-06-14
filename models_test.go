package skylabtech

import (
	"encoding/json"
	"testing"
)

func TestJobStruct(t *testing.T) {
	job := Job{
		ID:        1,
		Name:      "Test Job",
		AccountID: 1,
		State:     "pending",
		ProfileID: 1,
	}

	if job.ID != 1 {
		t.Fatalf("Expected ID to be 1, got %d", job.ID)
	}

	if job.Name != "Test Job" {
		t.Fatalf("Expected Name to be 'Test Job', got %s", job.Name)
	}
}

func TestPhotoStruct(t *testing.T) {
	photo := Photo{
		ID:          1,
		Name:        "Test Photo",
		Key:         "photo-key",
		JobID:       1,
		OriginalURL: "http://example.com/photo.jpg",
	}

	if photo.ID != 1 {
		t.Fatalf("Expected ID to be 1, got %d", photo.ID)
	}

	if photo.Name != "Test Photo" {
		t.Fatalf("Expected Name to be 'Test Photo', got %s", photo.Name)
	}
}

func TestProfileStruct(t *testing.T) {
	profile := Profile{
		ID:        1,
		Name:      "Test Profile",
		AccountID: 1,
	}

	if profile.ID != 1 {
		t.Fatalf("Expected ID to be 1, got %d", profile.ID)
	}

	if profile.Name != "Test Profile" {
		t.Fatalf("Expected Name to be 'Test Profile', got %s", profile.Name)
	}
}

func TestJobParamsJSON(t *testing.T) {
	jobParams := JobParams{
		Name:      "Test Job",
		ProfileID: 1,
		Type:      "regular",
	}

	jsonData, err := json.Marshal(jobParams)
	if err != nil {
		t.Fatalf("Error marshaling JSON: %v", err)
	}

	var unmarshalledParams JobParams
	err = json.Unmarshal(jsonData, &unmarshalledParams)
	if err != nil {
		t.Fatalf("Error unmarshaling JSON: %v", err)
	}

	if unmarshalledParams.Name != "Test Job" {
		t.Fatalf("Expected Name to be 'Test Job', got %s", unmarshalledParams.Name)
	}

	if unmarshalledParams.ProfileID != 1 {
		t.Fatalf("Expected ProfileID to be 1, got %d", unmarshalledParams.ProfileID)
	}

	if unmarshalledParams.Type != "regular" {
		t.Fatalf("Expected Type to be 'regular', got %s", unmarshalledParams.Type)
	}
}

func TestPhotoParamsJSON(t *testing.T) {
	photoParams := PhotoParams{
		JobID: 1,
		Name:  "Test Photo",
		Key:   "photo-key",
	}

	jsonData, err := json.Marshal(photoParams)
	if err != nil {
		t.Fatalf("Error marshaling JSON: %v", err)
	}

	var unmarshalledParams PhotoParams
	err = json.Unmarshal(jsonData, &unmarshalledParams)
	if err != nil {
		t.Fatalf("Error unmarshaling JSON: %v", err)
	}

	if unmarshalledParams.JobID != 1 {
		t.Fatalf("Expected JobID to be 1, got %d", unmarshalledParams.JobID)
	}

	if unmarshalledParams.Name != "Test Photo" {
		t.Fatalf("Expected Name to be 'Test Photo', got %s", unmarshalledParams.Name)
	}

	if unmarshalledParams.Key != "photo-key" {
		t.Fatalf("Expected Key to be 'photo-key', got %s", unmarshalledParams.Key)
	}
}
