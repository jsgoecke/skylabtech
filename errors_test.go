package skylabtech

import (
	"encoding/json"
	"testing"
)

func TestErrorStruct(t *testing.T) {
	apiError := APIError{
		StatusCode: 404,
		Err: Error{
			Code:    404,
			Message: "Not Found",
			Errors: map[string][]string{
				"field": {"error1", "error2"},
			},
		},
	}

	if apiError.StatusCode != 404 {
		t.Fatalf("Expected StatusCode to be 404, got %d", apiError.StatusCode)
	}

	if apiError.Err.Code != 404 {
		t.Fatalf("Expected Error Code to be 404, got %d", apiError.Err.Code)
	}

	if apiError.Err.Message != "Not Found" {
		t.Fatalf("Expected Error Message to be 'Not Found', got %s", apiError.Err.Message)
	}

	if len(apiError.Err.Errors["field"]) != 2 {
		t.Fatalf("Expected 2 errors for 'field', got %d", len(apiError.Err.Errors["field"]))
	}
}

func TestAPIErrorJSON(t *testing.T) {
	apiErrorJSON := `{"code": 404, "message": "Not Found", "errors": {"field": ["error1", "error2"]}}`
	var apiError Error
	err := json.Unmarshal([]byte(apiErrorJSON), &apiError)
	if err != nil {
		t.Fatalf("Error unmarshaling JSON: %v", err)
	}

	if apiError.Code != 404 {
		t.Fatalf("Expected Error Code to be 404, got %d", apiError.Code)
	}

	if apiError.Message != "Not Found" {
		t.Fatalf("Expected Error Message to be 'Not Found', got %s", apiError.Message)
	}

	if len(apiError.Errors["field"]) != 2 {
		t.Fatalf("Expected 2 errors for 'field', got %d", len(apiError.Errors["field"]))
	}
}
