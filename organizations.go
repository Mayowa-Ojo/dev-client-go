package dev

import (
	"context"
	"fmt"

	"github.com/google/go-querystring/query"
)

type Organization struct {
	TypeOf          string `json:"type_of"`
	Name            string `json:"name"`
	Username        string `json:"username"`
	Summary         string `json:"summary"`
	TwitterUsername string `json:"twitter_username,omitempty"`
	GithubUsername  string `json:"github_username,omitempty"`
	URL             string `json:"url"`
	Location        string `json:"location,omitempty"`
	TechStack       string `json:"tech_stack,omitempty"`
	TagLine         string `json:"tag_line,omitempty"`
	Story           string `json:"story,omitempty"`
	Slug            string `json:"slug"`
	JoinedAt        string `json:"joined_at"`
	ProfileImage    string `json:"profile_image"`
	ProfileImage90  string `json:"profile_image_90"`
}

type OrganizationQueryParams struct {
	Page     int32           `url:"page,omitempty"`
	PerPage  int32           `url:"per_page,omitempty"`
	Category ListingCategory `json:"category,omitempty"`
}

// GetOrganization allows the client retrieve a single organization
// by their username
func (c *Client) GetOrganization(orgname string) (*Organization, error) {
	path := fmt.Sprintf("/organizations/%s", orgname)

	req, err := c.NewRequest(context.Background(), "GET", path, nil)
	if err != nil {
		return nil, err
	}

	organization := new(Organization)

	if err := c.SendHttpRequest(req, &organization); err != nil {
		return nil, err
	}

	return organization, nil
}

// GetOrganizationUsers allows the client to retrieve a list of users belonging
// to the organization
func (c *Client) GetOrganizationUsers(orgname string, q OrganizationQueryParams) ([]User, error) {
	var users []User

	query, err := query.Values(q)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("/organizations/%s/users?%s", orgname, query.Encode())

	req, err := c.NewRequest(context.Background(), "GET", path, nil)
	if err != nil {
		return nil, err
	}

	if err := c.SendHttpRequest(req, &users); err != nil {
		return nil, err
	}

	return users, nil
}

// GetOrganizationListings allows the client to retrieve a list of
// listings belonging to the organization
func (c *Client) GetOrganizationListings(orgname string, q OrganizationQueryParams) ([]Listing, error) {
	var listings []Listing

	query, err := query.Values(q)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("/organizations/%s/listings?%s", orgname, query.Encode())

	req, err := c.NewRequest(context.Background(), "GET", path, nil)
	if err != nil {
		return nil, err
	}

	if err := c.SendHttpRequest(req, &listings); err != nil {
		return nil, err
	}

	return listings, nil
}

// GetOrganizationArticles allows the client to retrieve a list of
// Articles belonging to the organization
func (c *Client) GetOrganizationArticles(orgname string, q OrganizationQueryParams) ([]Article, error) {
	var articles []Article

	query, err := query.Values(q)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("/organizations/%s/articles?%s", orgname, query.Encode())

	req, err := c.NewRequest(context.Background(), "GET", path, nil)
	if err != nil {
		return nil, err
	}

	if err := c.SendHttpRequest(req, &articles); err != nil {
		return nil, err
	}

	return articles, nil
}
