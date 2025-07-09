package main

import (
	"flag"
	"fmt"
	"os"
)

var CONFIG *Config

func main() {
	var period = flag.String("period", "2024-01", "Billing period in YYYY-MM format")
	flag.Parse()
	CONFIG = loadConfig("config.yaml")
	CONFIG.ChargePeriod = *period

	s := ScalewayProvider{}
	projects, err := s.GetProjects(CONFIG.ScalewayOrgID)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	cons, err := s.GetConsumption(CONFIG.ScalewayOrgID, CONFIG.ChargePeriod)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	data, err := TransformData(cons, projects)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	txt, err := MakeCSV(data)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	v := VantageProvider{}
	err = v.PushBillingData(txt)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
