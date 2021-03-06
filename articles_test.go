package dev

import (
	"math"
	"os"
	"strconv"
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

func TestGetPublishedArticlesSorted(t *testing.T) {
	c, err := NewTestClient()
	if err != nil {
		t.Errorf("Failed to create TestClient: %s", err.Error())
	}

	articles, err := c.GetPublishedArticlesSorted(
		ArticleQueryParams{
			Page:    1,
			PerPage: 10,
		},
	)

	if err != nil {
		t.Errorf("Error fetching articles: %s", err.Error())
	}

	t1, _ := parseUTCDate(articles[0].PublishedAt)
	t2, err := parseUTCDate(articles[1].PublishedAt)
	if err != nil {
		t.Errorf("Error parsing UTC date: %v", err.Error())
	}

	diff := t1.Sub(t2).Seconds()

	if math.Signbit(diff) {
		t.Errorf("Expected result to contain articles ordered by descending publish date")
	}
}

func TestGetPublishedArticleByID(t *testing.T) {
	c, err := NewTestClient()
	if err != nil {
		t.Errorf("Failed to create TestClient: %s", err.Error())
	}

	articleID := os.Getenv("TEST_PUBLISHED_ARTICLE_ID")

	article, err := c.GetPublishedArticleByID(articleID)

	if err != nil {
		t.Errorf("Error fetching article: %s", err.Error())
	}

	want, err := strconv.Atoi(articleID)
	if err != nil {
		t.Errorf("Error converting string to int: %s", err.Error())
	}
	got := article.ID

	if int32(want) != got {
		t.Errorf("Expected result to contain articles ordered by descending publish date")
	}
}

func TestUpdateArticle(t *testing.T) {
	c, err := NewTestClient()
	if err != nil {
		t.Errorf("Failed to create TestClient: %s", err.Error())
	}

	articleID := "880104"

	payload := ArticleBodySchema{}
	payload.Article.Title = "The crust of structs in Go 3"
	payload.Article.BodyMarkdown = ""
	payload.Article.Published = false
	payload.Article.Tags = []string{"golang", "discuss"}

	article, err := c.UpdateArticle(articleID, payload, "article_sample.md")
	if err != nil {
		t.Errorf("Error trying to update article: %s", err.Error())
	}

	want := "The crust of structs in Go 3"
	got := article.Title

	if got != want {
		t.Errorf("Expected article title to be '%s', got '%s'", want, got)
	}

	if len(article.Tags) != 2 {
		t.Errorf("Expected article to have two tags, got '%d'", len(article.Tags))
	}
}

func TestGetPublishedArticleByPath(t *testing.T) {
	c, err := NewTestClient()
	if err != nil {
		t.Errorf("Failed to create TestClient: %s", err.Error())
	}

	username := os.Getenv("TEST_USERNAME")
	slug := os.Getenv("TEST_PUBLISHED_ARTICLE_SLUG")

	article, err := c.GetPublishedArticleByPath(username, slug)

	if err != nil {
		t.Errorf("Error fetching article: %s", err.Error())
	}

	if article.Slug != slug {
		t.Errorf("Expected article slug to be '%s', got '%s'", slug, article.Slug)
	}
}

func TestGetUserArticles(t *testing.T) {
	c, err := NewTestClient()
	if err != nil {
		t.Errorf("Failed to create TestClient: %s", err.Error())
	}

	articles, err := c.GetUserArticles(
		ArticleQueryParams{
			Page:    1,
			PerPage: 2,
		},
	)

	username := os.Getenv("TEST_USERNAME")

	if err != nil {
		t.Errorf("Error fetching articles: %s", err.Error())
	}

	if len(articles) != 2 {
		t.Errorf("Expected result to contain 2 articles, got %d", len(articles))
	}

	if articles[0].User.Username != username {
		t.Error("Expected result to include articles for authenticated user")
	}
}

func TestGetUserPublishedArticles(t *testing.T) {
	c, err := NewTestClient()
	if err != nil {
		t.Errorf("Failed to create TestClient: %s", err.Error())
	}

	articles, err := c.GetUserPublishedArticles(
		ArticleQueryParams{
			Page:    1,
			PerPage: 5,
		},
	)

	username := os.Getenv("TEST_USERNAME")

	if err != nil {
		t.Errorf("Error fetching articles: %s", err.Error())
	}

	if articles[0].User.Username != username {
		t.Error("Expected result to include articles for authenticated user")
	}

	for _, v := range articles {
		if !v.Published {
			t.Error("Expected result to contain published articles")
		}
	}
}

func TestGetUserUnPublishedArticles(t *testing.T) {
	c, err := NewTestClient()
	if err != nil {
		t.Errorf("Failed to create TestClient: %s", err.Error())
	}

	articles, err := c.GetUserUnPublishedArticles(
		ArticleQueryParams{
			Page:    1,
			PerPage: 5,
		},
	)

	username := os.Getenv("TEST_USERNAME")

	if err != nil {
		t.Errorf("Error fetching articles: %s", err.Error())
	}

	if articles[0].User.Username != username {
		t.Error("Expected result to include articles for authenticated user")
	}

	for _, v := range articles {
		if v.Published {
			t.Error("Expected result to contain unpublished articles")
		}
	}
}

func TestGetArticlesWithVideo(t *testing.T) {
	c, err := NewTestClient()
	if err != nil {
		t.Errorf("Failed to create TestClient: %s", err.Error())
	}

	articles, err := c.GetArticlesWithVideo(
		ArticleQueryParams{
			Page:    1,
			PerPage: 5,
		},
	)

	if err != nil {
		t.Errorf("Error fetching articles: %s", err.Error())
	}

	for _, v := range articles {
		if v.TypeOf != "video_article" {
			t.Error("Expected result to include articles for authenticated user")
		}
	}
}
