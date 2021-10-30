package dev

import (
	"context"
	"fmt"
)

type ProfileImage struct {
	TypeOf         string `json:"type_of"`
	ImageOf        string `json:"image_of"`
	ProfileImage   string `json:"profile_image"`
	ProfileImage90 string `json:"profile_image_90"`
}

// GetProfileImage allows the client to retrieve a user or organization profile
// image information by its corresponding username
func (c *Client) GetProfileImage(username string) (*ProfileImage, error) {
	path := fmt.Sprintf("/profile_images/%s", username)

	req, err := c.NewRequest(context.Background(), "GET", path, nil)
	if err != nil {
		return nil, err
	}

	profileImage := new(ProfileImage)

	if err := c.SendHttpRequest(req, &profileImage); err != nil {
		return nil, err
	}

	return profileImage, nil
}
