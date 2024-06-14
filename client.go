package skylabtech

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	apiKey  string
	baseURL string
	client  *http.Client
}

func NewClient(apiKey string) *Client {
	return &Client{
		apiKey:  apiKey,
		baseURL: "https://studio.skylabtech.ai/api/public/v1",
		client:  &http.Client{},
	}
}

func (c *Client) request(method, path string, body interface{}, result interface{}) error {
	var reqBody []byte
	var err error
	if body != nil {
		reqBody, err = json.Marshal(body)
		if err != nil {
			return err
		}
	}

	req, err := http.NewRequest(method, fmt.Sprintf("%s%s", c.baseURL, path), bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-SLT-API-KEY", c.apiKey)

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		var apiErr Error
		if err := json.Unmarshal(bodyBytes, &apiErr); err != nil {
			return fmt.Errorf("status: %d, error: %s, response: %s", resp.StatusCode, http.StatusText(resp.StatusCode), string(bodyBytes))
		}
		return &APIError{StatusCode: resp.StatusCode, Err: apiErr}
	}

	if result != nil {
		return json.NewDecoder(resp.Body).Decode(result)
	}

	return nil
}
