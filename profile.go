package skylabtech

import "fmt"

func (c *Client) ListProfiles() ([]Profile, error) {
	var profiles []Profile
	if err := c.request("GET", "/profiles", nil, &profiles); err != nil {
		return nil, err
	}
	return profiles, nil
}

func (c *Client) CreateProfile(params ProfileParams) (*Profile, error) {
	var profile Profile
	if err := c.request("POST", "/profiles", params, &profile); err != nil {
		return nil, err
	}
	return &profile, nil
}

func (c *Client) GetProfile(id int) (*Profile, error) {
	var profile Profile
	if err := c.request("GET", fmt.Sprintf("/profiles/%d", id), nil, &profile); err != nil {
		return nil, err
	}
	return &profile, nil
}

func (c *Client) UpdateProfile(id int, params ProfileParams) (*Profile, error) {
	var profile Profile
	if err := c.request("PATCH", fmt.Sprintf("/profiles/%d", id), params, &profile); err != nil {
		return nil, err
	}
	return &profile, nil
}

func (c *Client) DeleteProfile(id int) error {
	return c.request("DELETE", fmt.Sprintf("/profiles/%d", id), nil, nil)
}
