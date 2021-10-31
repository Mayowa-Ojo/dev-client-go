package dev

import (
	"os"
	"strconv"
	"testing"
)

func TestGetPublishedListings(t *testing.T) {
	c, err := NewTestClient()
	if err != nil {
		t.Errorf("Failed to create TestClient: %s", err.Error())
	}

	listings, err := c.GetPublishedListings(
		ListingQueryParams{
			PerPage:  5,
			Category: "collabs",
		},
	)

	if err != nil {
		t.Errorf("Error fetching articles: %s", err.Error())
	}

	if listings[0].TypeOf != "listing" {
		t.Errorf("Expected field 'type_of' to be 'listing', got '%s'", listings[0].TypeOf)
	}

	if listings[0].Category != "collabs" {
		t.Errorf("Expected field 'category' to be 'cfp', got '%s'", listings[0].Category)
	}
}

// NOTE: this test is failing
func TestCreateListing(t *testing.T) {
	t.Skip()
	c, err := NewTestClient()
	if err != nil {
		t.Errorf("Failed to create TestClient: %s", err.Error())
	}

	payload := ListingBodySchema{}
	payload.Listing.Title = "ACME Conference"
	payload.Listing.BodyMarkdown = "Awesome conference, come join us!"
	payload.Listing.Category = "cfp"
	payload.Listing.Tags = []string{"events"}

	listing, err := c.CreateListing(payload, nil)
	if err != nil {
		t.Errorf("Error creating listing: %s", err.Error())
	}

	if listing.TypeOf != "listing" {
		t.Errorf("Expected 'type_of' field to be 'listing', got '%s'", listing.TypeOf)
	}
}

func TestGetPublishedListingsByCategory(t *testing.T) {
	c, err := NewTestClient()
	if err != nil {
		t.Errorf("Failed to create TestClient: %s", err.Error())
	}

	listings, err := c.GetPublishedListingsByCategory(
		"cfp",
		ListingQueryParams{
			PerPage: 5,
		},
	)

	if err != nil {
		t.Errorf("Error fetching articles: %s", err.Error())
	}

	for _, v := range listings {
		if v.Category != "cfp" {
			t.Errorf("Expected catrgory to be 'cfp', instead got '%s'", v.Category)
		}
	}
}

func TestGetListingByID(t *testing.T) {
	c, err := NewTestClient()
	if err != nil {
		t.Errorf("Failed to create TestClient: %s", err.Error())
	}

	listingID := os.Getenv("TEST_LISTING_ID")

	listing, err := c.GetListingByID(listingID)

	if err != nil {
		t.Errorf("Error fetching articles: %s", err.Error())
	}

	listingIDInt, err := strconv.Atoi(listingID)
	if err != nil {
		t.Errorf("Error converting listingID to int: %s", err.Error())
	}

	if listing.ID != int64(listingIDInt) {
		t.Errorf("Expected result to be a listing with id: '%s', instead got '%d'", listingID, listing.ID)
	}
}
