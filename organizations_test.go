package dev

import (
	"os"
	"testing"
)

func TestGetOrganization(t *testing.T) {
	c, err := NewTestClient()
	if err != nil {
		t.Errorf("Failed to create TestClient: %s", err.Error())
	}

	username := os.Getenv("TEST_ORGANIZATION_USERNAME")

	org, err := c.GetOrganization(username)

	if err != nil {
		t.Errorf("Error fetching organization: %s", err.Error())
	}

	if org.TypeOf != "organization" {
		t.Errorf("Expected 'type_of' field to be 'organization', instead got '%s'", org.TypeOf)
	}
}

func TestGetOrganizationUsers(t *testing.T) {
	c, err := NewTestClient()
	if err != nil {
		t.Errorf("Failed to create TestClient: %s", err.Error())
	}

	username := os.Getenv("TEST_ORGANIZATION_USERNAME")

	users, err := c.GetOrganizationUsers(
		username,
		OrganizationQueryParams{
			Page:    1,
			PerPage: 5,
		},
	)

	if err != nil {
		t.Errorf("Error fetching users: %s", err.Error())
	}

	for _, v := range users {
		if v.TypeOf != "user" {
			t.Errorf("Expected 'type_of' field to be 'user', instead got '%s'", v.TypeOf)
		}
	}
}

func TestGetOrganizationListings(t *testing.T) {
	c, err := NewTestClient()
	if err != nil {
		t.Errorf("Failed to create TestClient: %s", err.Error())
	}

	username := os.Getenv("TEST_ORGANIZATION_USERNAME")

	users, err := c.GetOrganizationListings(
		username,
		OrganizationQueryParams{
			Page:    1,
			PerPage: 5,
		},
	)

	if err != nil {
		t.Errorf("Error fetching listings: %s", err.Error())
	}

	for _, v := range users {
		if v.TypeOf != "listing" {
			t.Errorf("Expected 'type_of' field to be 'listing', instead got '%s'", v.TypeOf)
		}
	}
}

func TestGetOrganizationArticles(t *testing.T) {
	c, err := NewTestClient()
	if err != nil {
		t.Errorf("Failed to create TestClient: %s", err.Error())
	}

	username := os.Getenv("TEST_ORGANIZATION_USERNAME")

	articles, err := c.GetOrganizationArticles(
		username,
		OrganizationQueryParams{
			Page:    1,
			PerPage: 5,
		},
	)

	if err != nil {
		t.Errorf("Error fetching articles: %s", err.Error())
	}

	for _, v := range articles {
		if v.TypeOf != "article" {
			t.Errorf("Expected 'type_of' field to be 'article', instead got '%s'", v.TypeOf)
		}
	}
}
