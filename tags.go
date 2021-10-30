package dev

import (
	"context"
)

type Tag struct {
	ID     int64   `json:"id"`
	Name   string  `json:"name"`
	Points float64 `json:"points"`
}

// GetFollowedTags allows the client to retrieve a list of the tags they follow
func (c *Client) GetFollowedTags() ([]Tag, error) {
	var tags []Tag

	path := "/follows/tags"

	req, err := c.NewRequest(context.Background(), "GET", path, nil)
	if err != nil {
		return nil, err
	}

	if err := c.SendHttpRequest(req, &tags); err != nil {
		return nil, err
	}

	return tags, nil
}
