package settings

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/urfave/cli"
)

var (
	ListenPort   int
	DynamoLocal  bool
	DynamoTable  string
	DynamoRegion string
	Logger       zerolog.Logger
)

// NewContext - app configuration
func NewContext() []cli.Flag {
	return []cli.Flag{
		cli.IntFlag{
			Name:        "listen_port, lp",
			Value:       1234,
			Destination: &ListenPort,
			EnvVar:      "LISTEN_PORT",
		},
		cli.BoolFlag{
			Name:        "dynamo_local, dl",
			Destination: &DynamoLocal,
			EnvVar:      "DYNAMO_LOCAL",
		},
		cli.StringFlag{
			Name:        "dynamo_table, dt",
			Value:       "prometheus",
			Destination: &DynamoTable,
			EnvVar:      "DYNAMO_TABLE",
		},
		cli.StringFlag{
			Name:        "dynamo_region, dr",
			Value:       "eu-west-1",
			Destination: &DynamoRegion,
			EnvVar:      "DYNAMO_REGION",
		},
	}
}

// LoggerContext - init logging function
func LoggerContext(ctx *cli.Context) {
	Logger = zerolog.
		New(os.Stderr).
		Output(zerolog.ConsoleWriter{Out: os.Stderr}).
		With().
		Timestamp().
		Str("app", ctx.App.Name).
		Logger()
}
