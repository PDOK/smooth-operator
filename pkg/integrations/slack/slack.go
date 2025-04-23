package slack

import (
	"encoding/json"
	"fmt"
	"net/http"
	"slices"
	"strings"
)

type SlackRequest struct {
	Attachments []SlackElement `json:"attachments"`
	Text        *string        `json:"text,omitempty"`
}

type SlackElement struct {
	Color  string       `json:"color"`
	Blocks []SlackBlock `json:"blocks"`
}

type SlackBlock struct {
	Type   string            `json:"type"`
	Text   SlackValueBlock   `json:"text"`
	Fields []SlackValueBlock `json:"fields"`
}

type SlackValueBlock struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

func GetSlackBlock(generalMessage string, jobUrl string, color string, datasetName string) SlackRequest {
	var slackBlock = SlackBlock{
		Type: "section",
		Text: SlackValueBlock{
			Type: "mrkdwn",
			Text: generalMessage,
		},
		Fields: []SlackValueBlock{{
			Type: "mrkdwn",
			Text: "*Dataset*",
		}, {
			Type: "mrkdwn",
			Text: "*Job*",
		}, {
			Type: "mrkdwn",
			Text: datasetName,
		}, {
			Type: "mrkdwn",
			Text: jobUrl,
		}},
	}

	slackElement := SlackElement{
		Color:  color,
		Blocks: []SlackBlock{slackBlock},
	}

	return SlackRequest{
		Attachments: []SlackElement{slackElement},
	}
}

func GetSlackErrorMessage(message string, bundle string, color string) SlackRequest {
	var slackBlock = SlackBlock{
		Type: "section",
		Text: SlackValueBlock{
			Type: "mrkdwn",
			Text: message,
		},
		Fields: []SlackValueBlock{{
			Type: "mrkdwn",
			Text: "*Bundle*",
		}, {
			Type: "mrkdwn",
			Text: bundle,
		}},
	}

	slackElement := SlackElement{
		Color:  color,
		Blocks: []SlackBlock{slackBlock},
	}

	return SlackRequest{
		Attachments: []SlackElement{slackElement},
	}
}

func GetSimpleSlackErrorMessage(message string) SlackRequest {
	return SlackRequest{
		Attachments: nil,
		Text:        &message,
	}
}

func SendSlackRequest(slackRequest SlackRequest, slackUrl string) error {
	marshalled, err := json.Marshal(slackRequest)
	if err != nil {
		return err
	}
	response, err := http.DefaultClient.Post(slackUrl, "application/json", strings.NewReader(string(marshalled)))
	if err != nil {
		return err
	}
	defer response.Body.Close()
	return nil
}

type SlackZapWriter struct {
	OperatorName    string
	SlackWebhookUrl string
}

func (slackWriter *SlackZapWriter) Write(p []byte) (n int, err error) {
	messageMap := make(map[string]string)
	json.Unmarshal(p, &messageMap)
	level := messageMap["level"]
	errorMessage := messageMap["error"]
	if slices.Contains([]string{"error", "panic", "fatal"}, level) {
		print(string(p))
		if slackWriter.SlackWebhookUrl != "" {
			slackRequest := GetSimpleSlackErrorMessage(fmt.Sprintf("%s: %s", slackWriter.OperatorName, errorMessage))
			SendSlackRequest(slackRequest, slackWriter.SlackWebhookUrl)
		}
	} else {
		print(string(p))
	}
	return len(p), nil
}
