package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"sct-backend-service/app/keys"
	"sct-backend-service/graph/model"
	"sct-backend-service/types"

	"go.uber.org/zap"
)

type GraphQLController interface {
	types.GraphQLService
}

type GraphQLControllerImpl struct {
	deps ControllerDeps
}

func CreateGraphQLController(deps ControllerDeps) GraphQLController {
	return &GraphQLControllerImpl{
		deps: deps,
	}
}

func (impl *GraphQLControllerImpl) SendContactInfo(ctx context.Context, input model.SendContactInfoRequest) (*model.SendContactInfoResponse, error) {
	for _, contact := range input.ContactInfo {
		slackBody := generateSlackMessage(input.Source, contact)
		// Convert the struct to JSON
		jsonData, err := json.Marshal(slackBody)
		if err != nil {
			impl.deps.Logger.Error("Error marshalling JSON", zap.Error(err))
			return nil, fmt.Errorf("error marshalling JSON: %w", err)
		}

		webhookURL := os.Getenv(keys.SlackWebhookEnvKey)
		if webhookURL == "" {
			impl.deps.Logger.Error("Slack webhook URL is not configured")
			return nil, fmt.Errorf("slack webhook url is not configured")
		}

		// Send the HTTP POST request
		resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			impl.deps.Logger.Error("Error sending message", zap.Error(err))
			return nil, fmt.Errorf("error sending message: %w", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			impl.deps.Logger.Error("failed to send notification", zap.Int("status code", resp.StatusCode))
			return nil, fmt.Errorf("failed to send notification, status code: %d", resp.StatusCode)
		}
	}
	return &model.SendContactInfoResponse{
		IsSuccess: true,
	}, nil
}

func generateSlackMessage(source model.WebsiteSource, input *model.ContactInfoInput) map[string]interface{} {
	slackBody := map[string]interface{}{
		"blocks": []map[string]interface{}{
			{
				"type": "header",
				"text": map[string]string{
					"type": "plain_text",
					"text": fmt.Sprintf("New Customer Enquiry from %s", source.String()),
				},
			},
			{
				"type": "section",
				"text": map[string]string{
					"type": "mrkdwn",
					"text": fmt.Sprintf("%s has filled out the enquiry form on the website", input.Name),
				},
			},
			{
				"type": "divider",
			},
			{
				"type": "section",
				"text": map[string]string{
					"type": "mrkdwn",
					"text": fmt.Sprintf("Name: %s", input.Name),
				},
			},
			{
				"type": "section",
				"text": map[string]string{
					"type": "mrkdwn",
					"text": fmt.Sprintf("Email: %s", input.Email),
				},
			},
			{
				"type": "section",
				"text": map[string]string{
					"type": "mrkdwn",
					"text": fmt.Sprintf("Phone Number: %s", input.PhoneNumber),
				},
			},
			{
				"type": "section",
				"text": map[string]string{
					"type": "mrkdwn",
					"text": fmt.Sprintf("Company Name: %s", input.CompanyName),
				},
			},
			{
				"type": "section",
				"text": map[string]string{
					"type": "mrkdwn",
					"text": fmt.Sprintf("Subject: %s", input.Subject),
				},
			},
			{
				"type": "section",
				"text": map[string]string{
					"type": "mrkdwn",
					"text": fmt.Sprintf("Message: %s", input.Message),
				},
			},
			{
				"type": "divider",
			},
			{
				"type": "actions",
				"elements": []map[string]interface{}{
					{
						"type": "button",
						"text": map[string]string{
							"type": "plain_text",
							"text": "Reply to Customer",
						},
						"url":   fmt.Sprintf("mailto:%s", input.Email),
						"style": "primary",
					},
				},
			},
		},
	}

	return slackBody
}
