package dev

import (
	"os"
	"strconv"
	"testing"
)

func TestGetUserByID(t *testing.T) {
	c, err := NewTestClient()
	if err != nil {
		t.Errorf("Failed to create TestClient: %s", err.Error())
	}

	userID := os.Getenv("TEST_USER_ID")

	user, err := c.GetUserByID(userID)

	if err != nil {
		t.Errorf("Error fetching user: %s", err.Error())
	}

	expected, err := strconv.Atoi(userID)
	if err != nil {
		t.Errorf("Error converting string to int: %s", err.Error())
	}
	got := user.ID

	if int32(expected) != got {
		t.Errorf("Expected 'id' field to be '%d', instead got '%d'", expected, got)
	}
}

func TestGetUserByUsername(t *testing.T) {
	c, err := NewTestClient()
	if err != nil {
		t.Errorf("Failed to create TestClient: %s", err.Error())
	}

	username := os.Getenv("TEST_USERNAME")

	user, err := c.GetUserByUsername(
		UserQueryParams{
			URL: username,
		},
	)

	if err != nil {
		t.Errorf("Error fetching user: %s", err.Error())
	}

	if user.Username != username {
		t.Errorf("Expected 'username' field to be '%s', instead got '%s'", username, user.Username)
	}
}

func TestGetAuthenticatedUser(t *testing.T) {
	c, err := NewTestClient()
	if err != nil {
		t.Errorf("Failed to create TestClient: %s", err.Error())
	}

	username := os.Getenv("TEST_AUTH_USERNAME")

	user, err := c.GetAuthenticatedUser()

	if err != nil {
		t.Errorf("Error fetching user: %s", err.Error())
	}

	if user.Username != username {
		t.Errorf("Expected 'username' field to be '%s', instead got '%s'", username, user.Username)
	}
}

func TestGetUserReadingList(t *testing.T) {
	c, err := NewTestClient()
	if err != nil {
		t.Errorf("Failed to create TestClient: %s", err.Error())
	}

	readinglist, err := c.GetUserReadingList(
		ReadingListQueryParams{
			PerPage: 5,
		},
	)

	if err != nil {
		t.Errorf("Error fetching reading list: %s", err.Error())
	}

	for _, v := range readinglist {
		if v.TypeOf != "readinglist" {
			t.Errorf("Expected 'type_of' field to be 'readinglist', instead got '%s'", v.TypeOf)
		}
	}
}

func TestGetUserFollowers(t *testing.T) {
	c, err := NewTestClient()
	if err != nil {
		t.Errorf("Failed to create TestClient: %s", err.Error())
	}

	followers, err := c.GetUserFollowers(
		UserQueryParams{
			PerPage: 5,
		},
	)

	if err != nil {
		t.Errorf("Error fetching followers: %s", err.Error())
	}

	for _, v := range followers {
		if v.TypeOf != "user_follower" {
			t.Errorf("Expected 'type_of' field to be 'user_follower', instead got '%s'", v.TypeOf)
		}
	}
}
