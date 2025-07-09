package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/gocarina/gocsv"
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

func HTTPRequest(uri string, method string, body io.Reader, headers map[string]string) ([]byte, error) {
	LOG.Debug().Str("uri", uri).Str("method", method).Any("body", body).Any("headers", headers).Msg("Starting HTTP request")
	req, err := http.NewRequest(method, uri, body)
	if err != nil {
		return nil, err
	}
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	res, err := http.DefaultClient.Do(req)
	LOG.Debug().Int("status_code", res.StatusCode).Msg("Got HTTP status code")
	if err != nil || res.StatusCode >= 400 {
		defer res.Body.Close()
		if err == nil {
			rawBody, _ := io.ReadAll(res.Body)
			return nil, fmt.Errorf("%s", (string(rawBody)))

		}
		return nil, err
	}
	defer res.Body.Close()
	rawBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	LOG.Trace().Bytes("response_body", rawBody).Msg("HTTP request complete.")
	return rawBody, nil
}

func TransformData(consumption *ScalewayConsumptionResponse, projects map[string]string) ([]*FOCUS, error) {
	LOG.Debug().Msg("Transforming data")
	f := []*FOCUS{}
	for _, item := range consumption.Consumptions {
		tags := CONFIG.Tags
		tags["Project"] = projects[item.ProjectID]
		t, err := json.Marshal(tags)
		if err != nil {
			LOG.Debug().Err(err).Any("tags", tags).Msg("Error parsing tags, using empty tags")
			t = []byte{}
		}
		consumed, err := strconv.Atoi(item.BilledQuantity)
		if err != nil {
			LOG.Debug().Err(err).Str("billed_qty", item.BilledQuantity).Int("consumed", consumed).Msg("Error converting string to int, using 0")
			consumed = 0
		}
		line := &FOCUS{
			ChargePeriodStart: fmt.Sprintf("%s-%s", CONFIG.ChargePeriod, "01"),
			BillingCurrency:   item.Value.CurrencyCode,
			ChargeCategory:    "Usage",
			ConsumedQuantity:  int64(consumed),
			ConsumedUnit:      item.Unit,
			BilledCost:        float64(item.Value.Units) + (float64(item.Value.Nanos) * 0.000000001),
			RegionId:          GetRegion(item.ResourceName),
			ResourceId:        item.ResourceName,
			ResourceType:      item.SKU,
			ServiceCategory:   item.ProductName,
			ServiceName:       item.CategoryName,
			Tags:              string(t),
		}
		f = append(f, line)
		LOG.Trace().Any("line", line).Msg("Converted line to FOCUS format")
	}
	LOG.Debug().Int("lines", len(f)).Msg("Converted consumption to FOCUS format")
	return f, nil
}

func MakeCSV(lines []*FOCUS) (string, error) {
	LOG.Debug().Msg("Converting lines to CSV")
	txt, err := gocsv.MarshalString(lines)
	if err != nil {
		return "", err
	}
	return txt, nil
}
