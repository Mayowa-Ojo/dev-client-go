package dev

import (
	"context"
	"fmt"

	"github.com/google/go-querystring/query"
)

type Comment struct {
	TypeOf    string    `json:"type_of"`
	IDCode    string    `json:"id_code"`
	CreatedAt string    `json:"created_at"`
	BodyHTML  string    `json:"body_html"`
	User      *User     `json:"user"`
	Children  []Comment `json:"children"`
}

type CommentQueryParams struct {
	ArticleID int32 `url:"a_id"`
	PodcastID int32 `url:"p_id"`
}

// GetComments allows the client to retrieve all comments
// belonging to an article or podcast as threaded conversations
func (c *Client) GetComments(q CommentQueryParams) ([]Comment, error) {
	var comments []Comment

	query, err := query.Values(q)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("/comments?%s", query.Encode())

	req, err := c.NewRequest(context.Background(), "GET", path, nil)
	if err != nil {
		return nil, err
	}

	if err := c.SendHttpRequest(req, &comments); err != nil {
		return nil, err
	}

	return comments, nil
}

// GetComment allows the client to retrieve a comment alongside
// its descendants
func (c *Client) GetComment(commentID string) (*Comment, error) {
	path := fmt.Sprintf("/comments/%s", commentID)

	req, err := c.NewRequest(context.Background(), "GET", path, nil)
	if err != nil {
		return nil, err
	}

	comment := new(Comment)

	if err := c.SendHttpRequest(req, &comment); err != nil {
		return nil, err
	}

	return comment, nil
}
