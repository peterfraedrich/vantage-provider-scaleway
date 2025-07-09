package main

import (
	"bytes"
	"fmt"
)

type VantageProvider struct{}

func (s VantageProvider) PushBillingData(csv string) error {
	uri := fmt.Sprintf("%s/integrations/%s/costs.csv", CONFIG.VantageAPIUrl, CONFIG.VantageCustomProviderToken)
	headers := map[string]string{
		"Accept":        "application/json",
		"Content-Type":  "multipart/form-data",
		"Authorization": fmt.Sprintf("Bearer %s", CONFIG.VantageAPIKey),
	}
	buf := bytes.NewBuffer(nil)
	_, err := buf.WriteString(csv)
	if err != nil {
		return err
	}
	res, err := HTTPRequest(uri, "POST", buf, headers)
	if err != nil {
		return err
	}
	fmt.Println(string(res))
	return nil
}
