package dev

import "testing"

func TestGetPublishedListings(t *testing.T) {
	c, err := NewTestClient()
	if err != nil {
		t.Errorf("Failed to create TestClient: %s", err.Error())
	}

	listings, err := c.GetPublishedListings(
		ListingQueryParams{
			PerPage:  5,
			Category: "cfp",
		},
	)

	if err != nil {
		t.Errorf("Error fetching articles: %s", err.Error())
	}

	if listings[0].TypeOf != "listing" {
		t.Errorf("Expected field 'type_of' to be 'listing', got '%s'", listings[0].TypeOf)
	}

	if listings[0].Category != "cfp" {
		t.Errorf("Expected field 'category' to be 'cfp', got '%s'", listings[0].Category)
	}
}

func TestCreateListing(t *testing.T) {
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
