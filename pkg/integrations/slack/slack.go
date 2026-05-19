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

	err = response.Body.Close()
	if err != nil {
		return err
	}

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
