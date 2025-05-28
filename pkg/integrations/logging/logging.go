package logging

import (
	"github.com/pdok/smooth-operator/pkg/integrations/slack"
	zap "go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type StdoutWriter struct {
}

func (s StdoutWriter) Write(p []byte) (n int, err error) {
	println(string(p))
	return len(p), nil
}

func SetupLogger(operatorName string, slackWebhookURL string, minLogLevel zapcore.LevelEnabler) (*zap.Logger, error) {
	// Standard output writer
	stdoutSyncer := zapcore.Lock(zapcore.AddSync(StdoutWriter{}))

	// Slack writer for errors
	slackSyncer := zapcore.Lock(zapcore.AddSync(&slack.ZapWriter{
		OperatorName:    operatorName,
		SlackWebhookURL: slackWebhookURL,
	}))

	// Encoder configuration for human-readable output
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	// Create a core for stdout (all levels)
	stdoutCore := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		stdoutSyncer,
		minLogLevel,
	)

	// Create a core for Slack (errors only)
	slackCore := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		slackSyncer,
		zapcore.ErrorLevel,
	)

	// Combine cores
	combinedCore := zapcore.NewTee(stdoutCore, slackCore)

	// Build the logger
	logger := zap.New(combinedCore, zap.AddCaller())
	return logger, nil
}
