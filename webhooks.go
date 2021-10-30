package dev

import (
	"context"
	"fmt"
)

type Webhook struct {
	TypeOf    string   `json:"type_of"`
	ID        int64    `json:"id"`
	Source    string   `json:"source"`
	TargetURL string   `json:"target_url"`
	Events    []string `json:"events"`
	CreatedAt string   `json:"created_at"`
}

type WebhookBodySchema struct {
	WebhookEndpoint struct {
		Source    string   `json:"source"`
		TargetURL string   `json:"target_url"`
		Events    []string `json:"events"`
	} `json:"webhook_endpoint"`
}

// GetWebhooks allows the client to retrieve a list of
// webhooks they have previously registered.
func (c *Client) GetWebhooks() ([]Webhook, error) {
	var webhooks []Webhook

	path := "/webhooks"

	req, err := c.NewRequest(context.Background(), "GET", path, nil)
	if err != nil {
		return nil, err
	}

	if err := c.SendHttpRequest(req, &webhooks); err != nil {
		return nil, err
	}

	return webhooks, nil
}

// CreateWebhook allows the client to create a new webhook
func (c *Client) CreateWebhook(payload WebhookBodySchema) (*Webhook, error) {
	path := "/webhooks"

	req, err := c.NewRequest(context.Background(), "POST", path, payload)
	if err != nil {
		return nil, err
	}

	webhook := new(Webhook)

	if err := c.SendHttpRequest(req, &webhook); err != nil {
		return nil, err
	}

	return webhook, nil
}

// GetWebhookByID allows the client to retrieve a single webhook given its id
func (c *Client) GetWebhookByID(webhookID string) (*Webhook, error) {
	path := fmt.Sprintf("/webhooks/%s", webhookID)

	req, err := c.NewRequest(context.Background(), "GET", path, nil)
	if err != nil {
		return nil, err
	}

	webhook := new(Webhook)

	if err := c.SendHttpRequest(req, &webhook); err != nil {
		return nil, err
	}

	return webhook, nil
}

// DeleteWebhook allows the client to delete a single webhook given its id
func (c *Client) DeleteWebhook(webhookID string) error {
	path := fmt.Sprintf("/webhooks/%s", webhookID)

	req, err := c.NewRequest(context.Background(), "DELETE", path, nil)
	if err != nil {
		return err
	}

	if err := c.SendHttpRequest(req, nil); err != nil {
		return err
	}

	return nil
}
