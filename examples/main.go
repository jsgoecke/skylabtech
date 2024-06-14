// main.go
package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/jsgoecke/skylabtech"
)

func main() {
	// Get the API key from environment variables
	apiKey := os.Getenv("SKYLABTECH_API_KEY")
	if apiKey == "" {
		log.Fatal("Environment variable SKYLABTECH_API_KEY is required")
	}

	// Create a new SkylabTech client
	client := skylabtech.NewClient(apiKey)

	// Example: Create a Job
	jobParams := skylabtech.JobParams{
		Name:      "Test Job",
		ProfileID: 1, // Ensure this is a valid profile ID
	}

	job, err := client.CreateJob(jobParams)
	if err != nil {
		log.Fatalf("Error creating job: %v", err)
	}
	fmt.Printf("Created Job: %+v\n", job)

	// Example: Get Photo Upload URL
	photoUploadParams := skylabtech.PhotoUploadUrlParams{
		PhotoID:        0,
		Type:           "jpg",
		UseCacheUpload: false,
	}

	uploadResponse, err := client.GetPhotoUploadURL(photoUploadParams)
	if err != nil {
		log.Fatalf("Error getting photo upload URL: %v", err)
	}
	fmt.Printf("Photo Upload URL: %+v\n", uploadResponse)

	// Example: Upload a Photo
	filePath := "path/to/your/photo.jpg" // Replace with the actual file path
	err = uploadPhoto(uploadResponse.URL, filePath)
	if err != nil {
		log.Fatalf("Error uploading photo: %v", err)
	}
	fmt.Println("Photo uploaded successfully")

	// Example: Create a Photo record
	photoParams := skylabtech.PhotoParams{
		JobID: job.ID,
		Name:  filepath.Base(filePath),
		Key:   uploadResponse.Key,
	}

	photo, err := client.CreatePhoto(photoParams)
	if err != nil {
		log.Fatalf("Error creating photo: %v", err)
	}
	fmt.Printf("Created Photo: %+v\n", photo)

	// Example: List Photos for Job
	photos, err := client.ListPhotosForJob(job.ID)
	if err != nil {
		log.Fatalf("Error listing photos for job: %v", err)
	}
	fmt.Printf("List of Photos for Job: %+v\n", photos)

	// Example: Get Photo by ID
	photoDetails, err := client.GetPhoto(photo.ID)
	if err != nil {
		log.Fatalf("Error getting photo by ID: %v", err)
	}
	fmt.Printf("Photo Details: %+v\n", photoDetails)

	// Example: Delete Photo
	err = client.DeletePhoto(photo.ID)
	if err != nil {
		log.Fatalf("Error deleting photo: %v", err)
	}
	fmt.Println("Deleted Photo")

	// Example: Delete Job
	err = client.DeleteJob(job.ID)
	if err != nil {
		log.Fatalf("Error deleting job: %v", err)
	}
	fmt.Println("Deleted Job")
}

// uploadPhoto uploads a photo to the given URL
func uploadPhoto(url, filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("could not open file: %v", err)
	}
	defer file.Close()

	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)
	part, err := writer.CreateFormFile("file", filepath.Base(filePath))
	if err != nil {
		return fmt.Errorf("could not create form file: %v", err)
	}
	_, err = io.Copy(part, file)
	if err != nil {
		return fmt.Errorf("could not copy file to form: %v", err)
	}
	writer.Close()

	req, err := http.NewRequest("PUT", url, &buf)
	if err != nil {
		return fmt.Errorf("could not create request: %v", err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("could not upload file: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("upload failed: %v", string(body))
	}

	return nil
}
