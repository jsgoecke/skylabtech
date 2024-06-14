package skylabtech

type Job struct {
	ID                   int               `json:"id"`
	AccountID            int               `json:"accountId"`
	Name                 string            `json:"name"`
	State                string            `json:"state"`
	ProfileID            int               `json:"profileId"`
	Photos               []Photo           `json:"photos"`
	Profile              Profile           `json:"profile"`
	CreatedAt            string            `json:"createdAt"`
	UpdatedAt            string            `json:"updatedAt"`
	SubmittedAt          string            `json:"submittedAt"`
	CompletedAt          string            `json:"completedAt"`
	PresetType           string            `json:"presetType"`
	Type                 string            `json:"type"`
	OutputLog            string            `json:"outputLog"`
	ExpiryDate           string            `json:"expiryDate"`
	PhotoDownloadUrls    map[string]string `json:"photoDownloadUrls"`
	Errors               []string          `json:"errors"`
	CreditAmountPerPhoto float64           `json:"creditAmountPerPhoto"`
	DollarAmountPerPhoto float64           `json:"dollarAmountPerPhoto"`
	Processed            map[string]int    `json:"processed"`
}

type JobParams struct {
	Name       string `json:"name"`
	ProfileID  int    `json:"profile_id,omitempty"`
	PresetType string `json:"preset_type,omitempty"`
	Type       string `json:"type,omitempty"`
}

type Photo struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Key          string `json:"key"`
	JobID        int    `json:"job_id"`
	OriginalURL  string `json:"original_url"`
	RetouchedURL string `json:"retouched_url"`
	CreatedAt    string `json:"createdAt"`
	UpdatedAt    string `json:"updatedAt"`
}

type PhotoParams struct {
	JobID     int    `json:"job_id,omitempty"`
	ProfileID int    `json:"profile_id,omitempty"`
	Name      string `json:"name"`
	Key       string `json:"key"`
}

type PhotoUploadUrlParams struct {
	PhotoID        int    `json:"photo_id,omitempty"`
	Type           string `json:"type,omitempty"`
	UseCacheUpload bool   `json:"use_cache_upload,omitempty"`
}

type PhotoUploadUrlResponse struct {
	URL string `json:"url"`
	Key string `json:"key"`
}

type Profile struct {
	ID          int    `json:"id"`
	AccountID   int    `json:"accountId"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

type ProfileParams struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

type JobsInFront struct {
	JobsInFront int    `json:"jobsInFront"`
	Message     string `json:"message"`
}
