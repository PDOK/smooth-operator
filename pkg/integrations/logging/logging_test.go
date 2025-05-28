package logging

import (
	"testing"

	"go.uber.org/zap/zapcore"
)

func TestLogger(t *testing.T) {
	_ = t
	logger, err := SetupLogger("myOperator", "", zapcore.InfoLevel)
	if err != nil {
		panic(err)
	}

	// This should log a message to stdout
	// If slackWebhookUrl is provided, it will use the URL to send a Slack message
	logger.Info("Foo")
	logger.Error("Bar")
}
