package dev

import (
	"os"
	"testing"
)

func TestGetPublishedPodcastEpisodes(t *testing.T) {
	c, err := NewTestClient()
	if err != nil {
		t.Errorf("Failed to create TestClient: %s", err.Error())
	}

	podcastSlug := os.Getenv("TEST_PODCAST_SLUG")

	podcasts, err := c.GetPublishedPodcastEpisodes(
		PodcastQueryParams{
			PerPage:  5,
			Username: podcastSlug,
		},
	)

	if err != nil {
		t.Errorf("Error fetching podcasts: %s", err.Error())
	}

	expected := "podcast_episodes"

	for _, v := range podcasts {
		if v.TypeOf != expected {
			t.Errorf("Expected 'type_of' field to be '%s', instead got '%s'", expected, v.TypeOf)
		}

		if v.Podcast.Slug != podcastSlug {
			t.Errorf("Expected podcast slug to be '%s', instead got '%s'", podcastSlug, v.Podcast.Slug)
		}
	}
}
