package logging

import (
	"github.com/pdok/smooth-operator/pkg/integrations/slack"
	"go.uber.org/zap/zapcore"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
)

func GetBasicLoggingOptions() zap.Options {
	return zap.Options{
		Development:          false,
		Encoder:              nil,
		EncoderConfigOptions: nil,
		NewEncoder:           nil,
		DestWriter:           nil,
		Level:                zapcore.InfoLevel,
		StacktraceLevel:      nil,
		ZapOpts:              nil,
		TimeEncoder:          nil,
	}
}

func UpdateLoggingOptions(options *zap.Options, operatorName string, slackWebhookUrl string) {
	options.DestWriter = &slack.SlackZapWriter{
		OperatorName:    operatorName,
		SlackWebhookUrl: slackWebhookUrl,
	}
}
