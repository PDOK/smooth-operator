package slack

import (
	"encoding/json"
	"net/http"
	"strings"
)

type SlackRequest struct {
	Attachments []SlackElement `json:"attachments"`
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
