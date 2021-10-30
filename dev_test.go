package dev

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestNewClient(t *testing.T) {
	if err := godotenv.Load(); err != nil {
		t.Errorf("Error loading env file: %s", err.Error())
	}

	token := os.Getenv("DEV_API_KEY")

	t.Run("invalid token", func(t *testing.T) {
		client, err := NewClient("")

		if client != nil {
			t.Errorf("expected client to be nil, got %+v", client)
		}
		if err == nil {
			t.Errorf("expected {NewClient} to return invalid token error")
		}
	})

	t.Run("valid api", func(t *testing.T) {
		// t.Skip()
		client, err := NewClient(token)

		if client == nil {
			t.Errorf("expected client to not be nil")
		}
		if err != nil {
			t.Errorf("expected {NewClient} to run without error, got:\n %v", err)
		}
	})

	t.Run("base-url", func(t *testing.T) {
		// t.Skip()
		c, _ := NewClient(token)

		if c.BaseUrl == nil {
			t.Errorf("expected baseUrl to be defined")
		}
	})
}
