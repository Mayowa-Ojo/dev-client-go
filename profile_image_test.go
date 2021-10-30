package dev

import (
	"os"
	"testing"
)

func TestGetProfileImage(t *testing.T) {
	c, err := NewTestClient()
	if err != nil {
		t.Errorf("Failed to create TestClient: %s", err.Error())
	}

	t.Run("user profile_image", func(t *testing.T) {
		username := os.Getenv("TEST_USERNAME")

		image, err := c.GetProfileImage(username)

		if err != nil {
			t.Errorf("Error fetching profile image: %s", err.Error())
		}

		if image.TypeOf != "profile_image" {
			t.Errorf("Expected 'type_of' field to be 'profile_image', instead got '%s'", image.TypeOf)
		}

		if image.ImageOf != "user" {
			t.Errorf("Expected 'image_of' field to be 'user', instead got '%s'", image.ImageOf)
		}
	})

	t.Run("organization profile_image", func(t *testing.T) {
		orgname := os.Getenv("TEST_ORGANIZATION_USERNAME")

		image, err := c.GetProfileImage(orgname)

		if err != nil {
			t.Errorf("Error fetching profile image: %s", err.Error())
		}

		if image.TypeOf != "profile_image" {
			t.Errorf("Expected 'type_of' field to be 'profile_image', instead got '%s'", image.TypeOf)
		}

		if image.ImageOf != "organization" {
			t.Errorf("Expected 'image_of' field to be 'organization', instead got '%s'", image.ImageOf)
		}
	})
}
