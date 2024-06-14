package skylabtech

import "fmt"

func (c *Client) ListJobs() ([]Job, error) {
	var jobs []Job
	if err := c.request("GET", "/jobs", nil, &jobs); err != nil {
		return nil, err
	}
	return jobs, nil
}

func (c *Client) CreateJob(params JobParams) (*Job, error) {
	var job Job
	if err := c.request("POST", "/jobs", params, &job); err != nil {
		return nil, err
	}
	return &job, nil
}

func (c *Client) GetJob(id int) (*Job, error) {
	var job Job
	if err := c.request("GET", fmt.Sprintf("/jobs/%d", id), nil, &job); err != nil {
		return nil, err
	}
	return &job, nil
}

func (c *Client) UpdateJob(id int, params JobParams) (*Job, error) {
	var job Job
	if err := c.request("PATCH", fmt.Sprintf("/jobs/%d", id), params, &job); err != nil {
		return nil, err
	}
	return &job, nil
}

func (c *Client) DeleteJob(id int) error {
	return c.request("DELETE", fmt.Sprintf("/jobs/%d", id), nil, nil)
}

func (c *Client) QueueJob(id int, callbackURL string, skipPhotoValidation bool) (*Job, error) {
	var job Job
	params := map[string]interface{}{
		"callback_url":          callbackURL,
		"skip_photo_validation": skipPhotoValidation,
	}
	if err := c.request("POST", fmt.Sprintf("/jobs/%d/queue", id), params, &job); err != nil {
		return nil, err
	}
	return &job, nil
}

func (c *Client) CancelJob(id int) (*Job, error) {
	var job Job
	if err := c.request("POST", fmt.Sprintf("/jobs/%d/cancel", id), nil, &job); err != nil {
		return nil, err
	}
	return &job, nil
}

func (c *Client) JobsInFront(id int) (*JobsInFront, error) {
	var jobsInFront JobsInFront
	if err := c.request("GET", fmt.Sprintf("/jobs/%d/jobs_in_front", id), nil, &jobsInFront); err != nil {
		return nil, err
	}
	return &jobsInFront, nil
}
