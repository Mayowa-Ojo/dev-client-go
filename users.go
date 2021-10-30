package dev

import (
	"context"
	"fmt"

	"github.com/google/go-querystring/query"
)

type User struct {
	TypeOf          string `json:"type_of"`
	ID              int32  `json:"id"`
	Username        string `json:"username"`
	Name            string `json:"name"`
	Summary         string `json:"summary,omitempty"`
	TwitterUsername string `json:"twitter_username,omitempty"`
	GithubUsername  string `json:"github_username,omitempty"`
	WebsiteURL      string `json:"website_url,omitempty"`
	Location        string `json:"location,omitempty"`
	JoinedAt        string `json:"joined_at"`
	ProfileImage    string `json:"profile_image"`
}

type UserQueryParams struct {
	Page    int32  `url:"page,omitempty"`
	PerPage int32  `url:"per_page,omitempty"`
	URL     string `url:"url"`
	Sort    string `url:"sort"`
}

type ReadingListStatus string

const (
	Valid        ReadingListStatus = "valid"
	InvalidValid ReadingListStatus = "invalid"
	Confirmed    ReadingListStatus = "confirmed"
	Archived     ReadingListStatus = "archived"
)

type ReadingList struct {
	TypeOf  string            `json:"type_of"`
	ID      int32             `json:"id"`
	Status  ReadingListStatus `json:"status"`
	Article *Article          `json:"article"`
}

type ReadingListQueryParams struct {
	Page    int32 `url:"page,omitempty"`
	PerPage int32 `url:"per_page,omitempty"`
}

// GetUserByID allows the client to retrieve a single user
// with the given id
func (c *Client) GetUserByID(userID string) (*User, error) {
	path := fmt.Sprintf("/users/%s", userID)

	req, err := c.NewRequest(context.Background(), "GET", path, nil)
	if err != nil {
		return nil, err
	}

	user := new(User)

	if err := c.SendHttpRequest(req, &user); err != nil {
		return nil, err
	}

	return user, nil
}

// GetUserByUsername allows the client to retrieve a single user
// with the given username
func (c *Client) GetUserByUsername(q UserQueryParams) (*User, error) {
	query, err := query.Values(q)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("/users/by_username?%s", query.Encode())

	req, err := c.NewRequest(context.Background(), "GET", path, nil)
	if err != nil {
		return nil, err
	}

	user := new(User)

	if err := c.SendHttpRequest(req, &user); err != nil {
		return nil, err
	}

	return user, nil
}

// GetAuthenticatedUser allows the client to retrieve information
// about the authenticated user
func (c *Client) GetAuthenticatedUser() (*User, error) {
	path := "/users/me"

	req, err := c.NewRequest(context.Background(), "GET", path, nil)
	if err != nil {
		return nil, err
	}

	user := new(User)

	if err := c.SendHttpRequest(req, &user); err != nil {
		return nil, err
	}

	return user, nil
}

// GetUserReadingList allows the client to retrieve a list of readinglist reactions
// along with the related article for the authenticated user.
func (c *Client) GetUserReadingList(q ReadingListQueryParams) ([]ReadingList, error) {
	var readinglist []ReadingList

	query, err := query.Values(q)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("/readinglist?%s", query.Encode())

	req, err := c.NewRequest(context.Background(), "GET", path, nil)
	if err != nil {
		return nil, err
	}

	if err := c.SendHttpRequest(req, &readinglist); err != nil {
		return nil, err
	}

	return readinglist, nil
}

// GetUserFollowers allows the client to retrieve a list of the followers they have.
func (c *Client) GetUserFollowers(q UserQueryParams) ([]User, error) {
	var followers []User

	query, err := query.Values(q)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("/followers/users?%s", query.Encode())

	req, err := c.NewRequest(context.Background(), "GET", path, nil)
	if err != nil {
		return nil, err
	}

	if err := c.SendHttpRequest(req, &followers); err != nil {
		return nil, err
	}

	return followers, nil
}
