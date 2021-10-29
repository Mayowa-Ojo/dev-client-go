package dev

import (
	"context"
	"fmt"

	"github.com/google/go-querystring/query"
)

type Article struct {
	TypeOf                 string              `json:"type_of"`
	ID                     int32               `json:"id"`
	Title                  string              `json:"title"`
	Description            string              `json:"description"`
	CoverImage             string              `json:"cover_image,omitempty"`
	Published              bool                `json:"published"`
	ReadablePublishDate    string              `json:"readable_publish_date"`
	SocialImage            string              `json:"social_image"`
	BodyMarkdown           string              `json:"body_markdown"`
	BodyHTML               string              `json:"body_html"`
	TagList                []string            `json:"tag_list"`
	Tags                   string              `json:"tags"`
	Slug                   string              `json:"slug"`
	Path                   string              `json:"path"`
	URL                    string              `json:"url"`
	CanonicalURL           string              `json:"canonical_url"`
	CommentsCount          int32               `json:"comments_count"`
	PositiveReactionsCount int32               `json:"positive_reactions_count"`
	PublicReactionsCount   int32               `json:"public_reactions_count"`
	CreatedAt              string              `json:"created_at"`
	EditedAt               string              `json:"edited_at,omitempty"`
	CrosspostedAt          string              `json:"crossposted_at,omitempty"`
	PublishedAt            string              `json:"published_at"`
	LastCommentAt          string              `json:"last_comment_at"`
	PublishedTimestamp     string              `json:"published_timestamp"`
	User                   *User               `json:"user"`
	ReadingTimeMinutes     int32               `json:"reading_time_minutes"`
	Organization           *SharedOrganization `json:"organization,omitempty"`
	FlareTag               *ArticleFlareTag    `json:"flare_tag,omitempty"`
}

type SharedOrganization struct {
	Name           string `json:"name"`
	Username       string `json:"username"`
	Slug           string `json:"slug"`
	ProfileImage   string `json:"profile_image"`
	ProfileImage90 string `json:"profile_image_90"`
}

type ArticleFlareTag struct {
	Name         string `json:"name"`
	BGColorHEX   string `json:"bg_color_hex"`
	TextColorHEX string `json:"text_color_hex"`
}

// There are some inconsistencies with regards to the article schema
// returned when fetching articles and the schema returned when creating
// articles. This structure fixes that
type ArticleVariant struct {
	Article
	Tags    []string `json:"tags"`
	TagList string   `json:"tag_list"`
}

type ArticleBodySchema struct {
	Article struct {
		Title          string   `json:"title"`
		BodyMarkdown   string   `json:"body_markdown"`
		Published      bool     `json:"published"`
		Series         string   `json:"series"`
		MainImage      string   `json:"main_image"`
		CanonicalURL   string   `json:"canonical_url"`
		Description    string   `json:"description"`
		Tags           []string `json:"tags"`
		OrganizationID int32    `json:"organization_id"`
	} `json:"article"`
}

type State string

const (
	Fresh  State = "fresh"
	Rising State = "rising"
	All    State = "all"
)

type ArticleQueryParams struct {
	Page         int32  `url:"page,omitempty"`
	PerPage      int32  `url:"per_page,omitempty"`
	Tag          string `url:"tag,omitempty"`
	Tags         string `url:"tags,omitempty"`
	TagsExclude  string `url:"tags_exclude,omitempty"`
	Username     string `url:"username,omitempty"`
	State        State  `url:"state,omitempty"`
	Top          int32  `url:"top,omitempty"`
	CollectionID int32  `url:"collection_id,omitempty"`
}

// GetPublishedArticles allows client to retrieve a list of articles
func (c *Client) GetPublishedArticles(q ArticleQueryParams) ([]Article, error) {
	var articles []Article

	query, err := query.Values(q)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("/articles?%s", query.Encode())

	req, err := c.NewRequest(context.Background(), "GET", path, nil)
	if err != nil {
		return nil, err
	}

	if err := c.SendHttpRequest(req, &articles); err != nil {
		return nil, err
	}

	return articles, nil
}

// CreateArticle allows the client to create a new article
// @filepath - article body can be set on the payload as a string
//            or passed via the path to a markdown file
func (c *Client) CreateArticle(payload ArticleBodySchema, filepath interface{}) (*ArticleVariant, error) {
	path := "/articles"

	if filepath != nil {
		content, err := ParseMarkdownFile(filepath.(string))
		if err != nil {
			return nil, err
		}

		payload.Article.BodyMarkdown = content
	}

	req, err := c.NewRequest(context.Background(), "POST", path, payload)
	if err != nil {
		return nil, err
	}

	article := new(ArticleVariant)

	if err := c.SendHttpRequest(req, &article); err != nil {
		return nil, err
	}

	return article, nil
}

// GetPublishedArticlesSorted allows the client to recieve a list
// of articles ordered by descending publish date
func (c *Client) GetPublishedArticlesSorted(q ArticleQueryParams) ([]Article, error) {
	var articles []Article

	query, err := query.Values(q)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("/articles?%s", query.Encode())

	req, err := c.NewRequest(context.Background(), "GET", path, nil)
	if err != nil {
		return nil, err
	}

	if err := c.SendHttpRequest(req, &articles); err != nil {
		return nil, err
	}

	return articles, nil
}

// GetPublishedArticleByID allows the client to retrieve a single published
// article by the specified ID
func (c *Client) GetPublishedArticleByID(articleID string) (*ArticleVariant, error) {
	path := fmt.Sprintf("/articles/%s", articleID)

	req, err := c.NewRequest(context.Background(), "GET", path, nil)
	if err != nil {
		return nil, err
	}

	article := new(ArticleVariant)

	if err := c.SendHttpRequest(req, &article); err != nil {
		return nil, err
	}

	return article, nil
}

// UpdateArticle allows the client to update an existing article
// This method is rate-limited (30req/30sec)
func (c *Client) UpdateArticle(articleID string, payload ArticleBodySchema, filepath interface{}) (*ArticleVariant, error) {
	path := fmt.Sprintf("/articles/%s", articleID)

	if filepath != nil {
		content, err := ParseMarkdownFile(filepath.(string))
		if err != nil {
			return nil, err
		}

		payload.Article.BodyMarkdown = content
	}

	req, err := c.NewRequest(context.Background(), "PUT", path, payload)
	if err != nil {
		return nil, err
	}

	article := new(ArticleVariant)

	if err := c.SendHttpRequest(req, &article); err != nil {
		return nil, err
	}

	return article, nil
}

// GetPublishedArticleByPath allows the client to retrieve a single published
// article given its path (slug)
func (c *Client) GetPublishedArticleByPath(username, slug string) (*ArticleVariant, error) {
	path := fmt.Sprintf("/articles/%s/%s", username, slug)

	req, err := c.NewRequest(context.Background(), "GET", path, nil)
	if err != nil {
		return nil, err
	}

	article := new(ArticleVariant)

	if err := c.SendHttpRequest(req, &article); err != nil {
		return nil, err
	}

	return article, nil
}

// GetUserArticles allows the client to retrieve a list of articles
// on behalf of an authenticated user
func (c *Client) GetUserArticles(q ArticleQueryParams) ([]Article, error) {
	var articles []Article

	query, err := query.Values(q)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("/articles/me/all?%s", query.Encode())

	req, err := c.NewRequest(context.Background(), "GET", path, nil)
	if err != nil {
		return nil, err
	}

	if err := c.SendHttpRequest(req, &articles); err != nil {
		return nil, err
	}

	return articles, nil
}

// GetUserArticles allows the client to retrieve a list of published
// articles on behalf of an authenticated user
func (c *Client) GetUserPublishedArticles(q ArticleQueryParams) ([]Article, error) {
	var articles []Article

	query, err := query.Values(q)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("/articles/me/published?%s", query.Encode())

	req, err := c.NewRequest(context.Background(), "GET", path, nil)
	if err != nil {
		return nil, err
	}

	if err := c.SendHttpRequest(req, &articles); err != nil {
		return nil, err
	}

	return articles, nil
}

// GetUserArticles allows the client to retrieve a list of unpublished
// articles on behalf of an authenticated user
func (c *Client) GetUserUnPublishedArticles(q ArticleQueryParams) ([]Article, error) {
	var articles []Article

	query, err := query.Values(q)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("/articles/me/unpublished?%s", query.Encode())

	req, err := c.NewRequest(context.Background(), "GET", path, nil)
	if err != nil {
		return nil, err
	}

	if err := c.SendHttpRequest(req, &articles); err != nil {
		return nil, err
	}

	return articles, nil
}
