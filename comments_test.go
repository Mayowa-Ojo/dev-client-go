package dev

import "testing"

func TestGetComments(t *testing.T) {
	c, err := NewTestClient()
	if err != nil {
		t.Errorf("Failed to create TestClient: %s", err.Error())
	}

	comments, err := c.GetComments(
		CommentQueryParams{
			ArticleID: 721780,
		},
	)

	if err != nil {
		t.Errorf("Error fetching articles: %s", err.Error())
	}

	if comments[0].TypeOf != "comment" {
		t.Errorf("Expected field 'type_of' to be 'comment', got '%s'", comments[0].TypeOf)
	}
}

func TestGetComment(t *testing.T) {
	c, err := NewTestClient()
	if err != nil {
		t.Errorf("Failed to create TestClient: %s", err.Error())
	}

	commentID := "1f87e"

	comment, err := c.GetComment(commentID)
	if err != nil {
		t.Errorf("Error fetching articles: %s", err.Error())
	}

	if comment.TypeOf != "comment" {
		t.Errorf("Expected field 'type_of' to be 'comment', got '%s'", comment.TypeOf)
	}

	if comment.IDCode != commentID {
		t.Errorf("Expected comment id  to be '%s', got '%s'", commentID, comment.IDCode)
	}
}
