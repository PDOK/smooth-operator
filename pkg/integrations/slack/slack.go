package slack

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Request struct {
	Attachments []Element `json:"attachments"`
	Text        *string   `json:"text,omitempty"`
}

type Element struct {
	Color  string  `json:"color"`
	Blocks []Block `json:"blocks"`
}

type Block struct {
	Type   string       `json:"type"`
	Text   ValueBlock   `json:"text"`
	Fields []ValueBlock `json:"fields"`
}

type ValueBlock struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

func GetSlackBlock(generalMessage string, jobURL string, color string, datasetName string) Request {
	var slackBlock = Block{
		Type: "section",
		Text: ValueBlock{
			Type: "mrkdwn",
			Text: generalMessage,
		},
		Fields: []ValueBlock{{
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
			Text: jobURL,
		}},
	}

	slackElement := Element{
		Color:  color,
		Blocks: []Block{slackBlock},
	}

	return Request{
		Attachments: []Element{slackElement},
	}
}

func GetSlackErrorMessage(message string, bundle string, color string) Request {
	var slackBlock = Block{
		Type: "section",
		Text: ValueBlock{
			Type: "mrkdwn",
			Text: message,
		},
		Fields: []ValueBlock{{
			Type: "mrkdwn",
			Text: "*Bundle*",
		}, {
			Type: "mrkdwn",
			Text: bundle,
		}},
	}

	slackElement := Element{
		Color:  color,
		Blocks: []Block{slackBlock},
	}

	return Request{
		Attachments: []Element{slackElement},
	}
}

func GetSimpleSlackErrorMessage(message string) Request {
	return Request{
		Attachments: nil,
		Text:        &message,
	}
}

func SendSlackRequest(slackRequest Request, slackURL string) error {
	marshalled, err := json.Marshal(slackRequest)
	if err != nil {
		return err
	}
	response, err := http.DefaultClient.Post(slackURL, "application/json", strings.NewReader(string(marshalled)))
	if err != nil {
		return err
	}
	defer response.Body.Close()
	return nil
}

type ZapWriter struct {
	OperatorName    string
	SlackWebhookURL string
}

func (slackWriter *ZapWriter) Sync() error {
	return nil
}

func (slackWriter *ZapWriter) Write(p []byte) (n int, err error) {
	if slackWriter.SlackWebhookURL != "" {
		slackRequest := GetSimpleSlackErrorMessage(fmt.Sprintf("%s: %s", slackWriter.OperatorName, string(p)))
		err = SendSlackRequest(slackRequest, slackWriter.SlackWebhookURL)
		if err != nil {
			return
		}
	}
	return len(p), nil
}
