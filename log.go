package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/diode"
)

func configureLogging(env string, level string) zerolog.Logger {
	var log zerolog.Logger
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	if env != "prod" {
		log = zerolog.New(zerolog.ConsoleWriter{
			Out:             os.Stderr,
			FormatLevel:     formatLevel,
			FormatFieldName: formatFieldName,
		}).With().Timestamp().Logger()
	} else {
		writer := diode.NewWriter(os.Stdout, 1000, 10*time.Millisecond, func(missed int) {
			fmt.Printf("Logger dropped %d messages", missed)
		})
		log = zerolog.New(writer).With().Timestamp().Logger()
	}
	lvl, err := zerolog.ParseLevel(level)
	if err != nil {
		log.Error().Err(err)
		lvl = zerolog.InfoLevel
	}
	zerolog.SetGlobalLevel(lvl)
	return log
}

func formatLevel(i interface{}) string {
	col := color.Bold
	switch i.(string) {
	case "trace":
		col = color.FgBlue
	case "debug":
		col = color.FgCyan
	case "info":
		col = color.FgGreen
	case "warn":
		col = color.FgYellow
	case "error":
		col = color.FgRed
	case "fatal", "panic":
		col = color.FgMagenta
	default:
		col = color.Bold
	}
	s := color.New(col).SprintFunc()
	return fmt.Sprintf("| %-8s |", s(strings.ToUpper(i.(string))))
}

func formatFieldName(i interface{}) string {
	return fmt.Sprintf("%s=", i)
}
