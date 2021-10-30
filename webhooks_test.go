package dev

import (
	"os"
	"strconv"
	"testing"
)

func TestGetWebhooks(t *testing.T) {
	c, err := NewTestClient()
	if err != nil {
		t.Errorf("Failed to create TestClient: %s", err.Error())
	}

	webhooks, err := c.GetWebhooks()

	if err != nil {
		t.Errorf("Error fetching webhooks: %s", err.Error())
	}

	for _, v := range webhooks {
		if v.TypeOf != "webhook_endpoint" {
			t.Error("Expected result to include webhooks registered by the user")
		}
	}
}

func TestCreateWebhook(t *testing.T) {
	c, err := NewTestClient()
	if err != nil {
		t.Errorf("Failed to create TestClient: %s", err.Error())
	}

	targetURL := os.Getenv("TEST_WEBHOOK_TARGET_URL")

	payload := WebhookBodySchema{}
	payload.WebhookEndpoint.TargetURL = targetURL
	payload.WebhookEndpoint.Source = "DEV"
	payload.WebhookEndpoint.Events = []string{"article_created"}

	webhook, err := c.CreateWebhook(payload)
	if err != nil {
		t.Errorf("Error trying to create article: %s", err.Error())
	}

	if webhook.TypeOf != "webhook_endpoint" {
		t.Errorf("Expected 'type_of' field to be 'webhook_endpoint', instead got %s", webhook.TypeOf)
	}

	if webhook.Events[0] != "article_created" {
		t.Errorf("Expected webhook events to include 'article_created', instead got %s", webhook.Events[0])
	}
}

func TestGetWebhookByID(t *testing.T) {
	c, err := NewTestClient()
	if err != nil {
		t.Errorf("Failed to create TestClient: %s", err.Error())
	}

	webhookID := os.Getenv("TEST_WEBHOOK_ID")

	webhook, err := c.GetWebhookByID(webhookID)

	if err != nil {
		t.Errorf("Error fetching webhook: %s", err.Error())
	}

	if webhook.TypeOf != "webhook_endpoint" {
		t.Errorf("Expected 'type_of' field to be 'webhook_endpoint', instead got '%s'", webhook.TypeOf)
	}

	want, err := strconv.Atoi(webhookID)
	if err != nil {
		t.Errorf("Error converting string to int: %s", err.Error())
	}

	if webhook.ID != int64(want) {
		t.Errorf("Expected webhook id to be '%d', instead got '%d'", want, webhook.ID)
	}
}

func TestDeleteWebhook(t *testing.T) {
	c, err := NewTestClient()
	if err != nil {
		t.Errorf("Failed to create TestClient: %s", err.Error())
	}

	webhookID := os.Getenv("TEST_WEBHOOK_ID")

	if err := c.DeleteWebhook(webhookID); err != nil {
		t.Errorf("Error deleting webhook: %s", err.Error())
	}
}
