package skylabtech

import "fmt"

func (c *Client) CreatePhoto(params PhotoParams) (*Photo, error) {
	var photo Photo
	if err := c.request("POST", "/photos", params, &photo); err != nil {
		return nil, err
	}
	return &photo, nil
}

func (c *Client) GetPhoto(id int) (*Photo, error) {
	var photo Photo
	if err := c.request("GET", fmt.Sprintf("/photos/%d", id), nil, &photo); err != nil {
		return nil, err
	}
	return &photo, nil
}

func (c *Client) DeletePhoto(id int) error {
	return c.request("DELETE", fmt.Sprintf("/photos/%d", id), nil, nil)
}

func (c *Client) GetPhotoUploadURL(params PhotoUploadUrlParams) (*PhotoUploadUrlResponse, error) {
	var response PhotoUploadUrlResponse
	if err := c.request("GET", "/photos/upload_url", params, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *Client) ListPhotosForJob(jobID int) ([]Photo, error) {
	var photos []Photo
	params := map[string]interface{}{
		"job_id": jobID,
	}
	if err := c.request("GET", "/photos/list_for_job", params, &photos); err != nil {
		return nil, err
	}
	return photos, nil
}
