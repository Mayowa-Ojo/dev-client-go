package dev

import (
	"strings"
	"testing"
)

func TestGetPublishedArticles(t *testing.T) {
	c, err := NewTestClient()
	if err != nil {
		t.Errorf("Failed to create TestClient: %s", err.Error())
	}

	t.Run("page limit", func(t *testing.T) {
		t.Skip()
		articles, err := c.GetPublishedArticles(
			ArticleQueryParams{
				Page:    1,
				PerPage: 10,
			},
		)

		if err != nil {
			t.Errorf("Error fetching articles: %s", err.Error())
		}

		if len(articles) != 10 {
			t.Errorf("Expected result to contain 10 articles, got %d", len(articles))
		}
	})

	t.Run("articles with tag", func(t *testing.T) {
		t.Skip()
		articles, err := c.GetPublishedArticles(
			ArticleQueryParams{
				PerPage: 1,
				Tag:     "golang",
			},
		)

		if err != nil {
			t.Errorf("Error fetching articles: %s", err.Error())
		}

		if !strings.Contains(articles[0].Tags, "go") {
			t.Errorf("Expected tags to contain given tag, got: %s", articles[0].Tags)
		}
	})

	t.Run("articles with tag", func(t *testing.T) {
		t.Skip()
		articles, err := c.GetPublishedArticles(
			ArticleQueryParams{
				PerPage:  1,
				Username: "unorthodev",
			},
		)

		if err != nil {
			t.Errorf("Error fetching articles: %s", err.Error())
		}

		want := "unorthodev"
		got := articles[0].User.Username

		if articles[0].User.Username != "unorthodev" {
			t.Errorf("Expected article user to be '%s', got '%s'", want, got)
		}
	})
}

func TestCreateArticle(t *testing.T) {
	t.Skip()
	c, err := NewTestClient()
	if err != nil {
		t.Errorf("Failed to create TestClient: %s", err.Error())
	}

	payload := ArticleBodySchema{}
	payload.Article.Title = "The crust of structs in Go"
	payload.Article.BodyMarkdown = ""
	payload.Article.Published = false
	payload.Article.Tags = []string{"golang"}

	article, err := c.CreateArticle(payload, "article_sample.md")
	if err != nil {
		t.Errorf("Error trying to create article: %s", err.Error())
	}

	want := "The crust of structs in Go"
	got := article.Title

	if article.Title != want {
		t.Errorf("Expected article title to be %s, got %s", want, got)
	}
}
