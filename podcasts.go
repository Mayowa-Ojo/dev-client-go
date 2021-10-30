package dev

import (
	"context"
	"fmt"

	"github.com/google/go-querystring/query"
)

type PodcastEpisode struct {
	TypeOf   string `json:"type_of"`
	ID       int32  `json:"id"`
	Path     string `json:"path"`
	ImageURL string `json:"image_url"`
	Title    string `json:"title"`
	Podcast  struct {
		Title    string `json:"title"`
		Slug     string `json:"slug"`
		ImageURL string `json:"image_url"`
	}
}

type PodcastQueryParams struct {
	Page     int32  `url:"page"`
	PerPage  int32  `url:"per_page"`
	Username string `url:"username"`
}

// GetPublishedPodcastEpisodes allows the client to retrieve a list of podcast episodes
func (c *Client) GetPublishedPodcastEpisodes(q PodcastQueryParams) ([]PodcastEpisode, error) {
	var podcasts []PodcastEpisode

	query, err := query.Values(q)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("/podcast_episodes?%s", query.Encode())

	req, err := c.NewRequest(context.Background(), "GET", path, nil)
	if err != nil {
		return nil, err
	}

	if err := c.SendHttpRequest(req, &podcasts); err != nil {
		return nil, err
	}

	return podcasts, nil
}
