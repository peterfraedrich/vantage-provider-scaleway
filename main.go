package main

import (
	"flag"

	"github.com/rs/zerolog"
)

var CONFIG *Config
var LOG zerolog.Logger

func main() {
	var period = flag.String("period", "2024-01", "Billing period in YYYY-MM format")
	flag.Parse()
	CONFIG = loadConfig("config.yaml")
	CONFIG.ChargePeriod = *period
	LOG = configureLogging(CONFIG.Env, CONFIG.LogLevel)
	LOG.Info().Msg("Configured logger.")

	s := ScalewayProvider{}
	projects, err := s.GetProjects(CONFIG.ScalewayOrgID)
	if err != nil {
		LOG.Fatal().Err(err).Send()
	}
	cons, err := s.GetConsumption(CONFIG.ScalewayOrgID, CONFIG.ChargePeriod)
	if err != nil {
		LOG.Fatal().Err(err).Send()
	}
	data, err := TransformData(cons, projects)
	if err != nil {
		LOG.Fatal().Err(err).Send()
	}
	txt, err := MakeCSV(data)
	if err != nil {
		LOG.Fatal().Err(err).Send()
	}
	v := VantageProvider{}
	err = v.PushBillingData(txt)
	if err != nil {
		LOG.Fatal().Err(err).Send()
	}
	LOG.Info().Msg("Complete. Exiting.")
}
