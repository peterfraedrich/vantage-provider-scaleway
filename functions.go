package main

import (
	"os"

	"gopkg.in/yaml.v3"
)

func loadConfig(filename string) *Config {
	c := &Config{}
	raw, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(raw, &c)
	if err != nil {
		panic(err)
	}
	c.VantageAPIKey = os.Getenv("VANTAGE_API_KEY")
	c.ScalewayAPIKey = os.Getenv("SCALEWAY_API_KEY")
	c.ScalewayAPISecret = os.Getenv("SCALEWAY_API_SECRET")
	return c

}
