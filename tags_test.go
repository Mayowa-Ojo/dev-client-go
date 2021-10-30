package dev

import (
	"testing"
)

func TestGetFollowedTags(t *testing.T) {
	c, err := NewTestClient()
	if err != nil {
		t.Errorf("Failed to create TestClient: %s", err.Error())
	}

	tags, err := c.GetFollowedTags()

	if err != nil {
		t.Errorf("Error fetching podcasts: %s", err.Error())
	}

	if len(tags) < 1 {
		t.Errorf("Expected result to be a list of tags, instead got: '%d'", len(tags))
	}
}
