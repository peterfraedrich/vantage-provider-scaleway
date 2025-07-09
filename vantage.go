package main

import (
	"fmt"
	"strings"
)

type VantageProvider struct{}

func (s VantageProvider) PushBillingData(csv string) error {
	LOG.Debug().Msg("Pushing data to Vantage")
	uri := fmt.Sprintf("%s/integrations/%s/costs.csv", CONFIG.VantageAPIUrl, CONFIG.VantageCustomProviderToken)
	headers := map[string]string{
		"Accept":        "application/json",
		"Content-Type":  "multipart/form-data; boundary=---011000010111000001101001",
		"Authorization": fmt.Sprintf("Bearer %s", CONFIG.VantageAPIKey),
	}
	fmtString := fmt.Sprintf(
		"-----011000010111000001101001\r\nContent-Disposition: form-data; name=\"csv\"; filename=\"scaleway-%s.csv\"\r\nContent-Type: text/csv\r\n\r\n%s\r\n-----011000010111000001101001--",
		CONFIG.ChargePeriod, csv,
	)
	LOG.Trace().Str("payload", fmtString).Msg("Formatted Vantage payload")
	payload := strings.NewReader(fmtString)
	res, err := HTTPRequest(uri, "POST", payload, headers)
	if err != nil {
		return err
	}
	LOG.Trace().Bytes("body", res).Msg("Vantage response")
	return nil
}
