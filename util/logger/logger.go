package logger

import (
	"context"
	"fmt"
	"go-boilerplate/model/constants"

	"github.com/rs/zerolog/log"
)

// Please use this logger library to log in handler, usecase, repo
// and client layers, so we can get the correlation_id in the log

func parseMsg(ctx context.Context, msg string) string {
	return msg + " - " + fmt.Sprint(ctx.Value(constants.CORRELATION_ID))
}

func Info(ctx context.Context, msg string) {
	log.Info().Msg(parseMsg(ctx, msg))
}

func Warn(ctx context.Context, msg string) {
	log.Warn().Msg(parseMsg(ctx, msg))
}

func Error(ctx context.Context, msg string) {
	log.Error().Msg(parseMsg(ctx, msg))
}

func Fatal(ctx context.Context, msg string) {
	log.Fatal().Msg(parseMsg(ctx, msg))
}

func Infof(ctx context.Context, msg string, v ...interface{}) {
	log.Info().Msgf(parseMsg(ctx, msg), v...)
}

func Warnf(ctx context.Context, msg string, v ...interface{}) {
	log.Warn().Msgf(parseMsg(ctx, msg), v...)
}

func Errorf(ctx context.Context, msg string, v ...interface{}) {
	log.Error().Msgf(parseMsg(ctx, msg), v...)
}

func Fatalf(ctx context.Context, msg string, v ...interface{}) {
	log.Fatal().Msgf(parseMsg(ctx, msg), v...)
}
