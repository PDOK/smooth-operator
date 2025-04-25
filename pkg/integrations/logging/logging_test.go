package logging

import (
	"go.uber.org/zap/zapcore"
	"testing"
)

func TestLogger(t *testing.T) {
	logger, err := SetupLogger("myOperator", "", zapcore.InfoLevel)
	if err != nil {
		panic(err)
	}

	// This should log a message to stdout
	// If slackWebhookUrl is provided, it will use the URL to send a Slack message
	logger.Info("Foo")
	logger.Error("Bar")
}
