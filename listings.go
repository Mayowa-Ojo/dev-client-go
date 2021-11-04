package dev

import (
	"context"
	"fmt"

	"github.com/google/go-querystring/query"
)

type Listing struct {
	TypeOf        string          `json:"type_of"`
	ID            int64           `json:"id"`
	Title         string          `json:"title"`
	Slug          string          `json:"slug"`
	BodyMarkdown  string          `json:"body_markdown"`
	TagList       string          `json:"tag_list"`
	Tags          []string        `json:"tags"`
	Category      ListingCategory `json:"category"`
	ProcessedHTML string          `json:"processed_html"`
	Published     bool            `json:"published"`
	User          *User           `json:"user"`
	Organization  *Organization   `json:"organization"`
}

type ListingBodySchema struct {
	Listing struct {
		Title             string          `json:"title"`
		BodyMarkdown      string          `json:"body_markdown"`
		Category          ListingCategory `json:"category"`
		Tags              []string        `json:"tags"`
		TagList           string          `json:"tag_list"`
		ExpiresAt         string          `json:"expires_at"`
		ContactViaConnect bool            `json:"contact_via_connect"`
		Location          string          `json:"location"`
		OrganizationID    int64           `json:"organization_id"`
		Action            Action          `json:"action"`
	} `json:"listing"`
}

type ListingCategory string

const (
	ListingCategoryCfp       = ListingCategory("cfp")
	ListingCategoryForhire   = ListingCategory("forhire")
	ListingCategoryCollabs   = ListingCategory("collabs")
	ListingCategoryEducation = ListingCategory("education")
	ListingCategoryJobs      = ListingCategory("jobs")
	ListingCategoryMentors   = ListingCategory("mentors")
	ListingCategoryProducts  = ListingCategory("products")
	ListingCategoryMentees   = ListingCategory("mentees")
	ListingCategoryForsale   = ListingCategory("forsale")
	ListingCategoryEvents    = ListingCategory("events")
	ListingCategoryMisc      = ListingCategory("misc")
)

type Action string

const (
	Bump      = Action("bump")
	Publish   = Action("publish")
	Unpublish = Action("unpublish")
)

type ListingQueryParams struct {
	Page     int32  `url:"page"`
	PerPage  int32  `url:"per_page"`
	Category string `url:"category"`
}

// GetPublishedListings allows the client retrieve a list of listings
func (c *Client) GetPublishedListings(q ListingQueryParams) ([]Listing, error) {
	var listings []Listing

	query, err := query.Values(q)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("/listings?%s", query.Encode())

	req, err := c.NewRequest(context.Background(), "GET", path, nil)
	if err != nil {
		return nil, err
	}

	if err := c.SendHttpRequest(req, &listings); err != nil {
		return nil, err
	}

	return listings, nil
}

// CreateListing allows the client to create a listing.
// Listings are classified as ads that users create on DEV
func (c *Client) CreateListing(payload ListingBodySchema, filepath interface{}) (*Listing, error) {
	path := "/listings"

	if filepath != nil {
		content, err := parseMarkdownFile(filepath.(string))
		if err != nil {
			return nil, err
		}

		payload.Listing.BodyMarkdown = content
	}

	req, err := c.NewRequest(context.Background(), "POST", path, payload)
	if err != nil {
		return nil, err
	}

	listing := new(Listing)

	if err := c.SendHttpRequest(req, &listing); err != nil {
		return nil, err
	}

	return listing, nil
}

// GetPublishedListingsByCategory allows the client to retrieve a list
// of listings belonging to the given category
func (c *Client) GetPublishedListingsByCategory(category string, q ListingQueryParams) ([]Listing, error) {
	var listings []Listing

	query, err := query.Values(q)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("/listings/category/%s?%s", category, query.Encode())

	req, err := c.NewRequest(context.Background(), "GET", path, nil)
	if err != nil {
		return nil, err
	}

	if err := c.SendHttpRequest(req, &listings); err != nil {
		return nil, err
	}

	return listings, nil
}

// GetListingByID allows the client to retrieve a single listing given
// its Id
func (c *Client) GetListingByID(listingID string) (*Listing, error) {
	path := fmt.Sprintf("/listings/%s", listingID)

	req, err := c.NewRequest(context.Background(), "GET", path, nil)
	if err != nil {
		return nil, err
	}

	listing := new(Listing)

	if err := c.SendHttpRequest(req, &listing); err != nil {
		return nil, err
	}

	return listing, nil
}

func (c *Client) UpdateListing(listingID string, payload ListingBodySchema, filepath interface{}) (*Listing, error) {
	path := fmt.Sprintf("/listings/%s", listingID)

	if filepath != nil {
		content, err := parseMarkdownFile(filepath.(string))
		if err != nil {
			return nil, err
		}

		payload.Listing.BodyMarkdown = content
	}

	req, err := c.NewRequest(context.Background(), "PUT", path, payload)
	if err != nil {
		return nil, err
	}

	listing := new(Listing)

	if err := c.SendHttpRequest(req, &listing); err != nil {
		return nil, err
	}

	return listing, nil
}
