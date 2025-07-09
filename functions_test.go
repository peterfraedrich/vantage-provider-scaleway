package main

import (
	"testing"
	"time"
)

func TestHTTPRequest(t *testing.T) {
	uri := "https://github.com/peterfraedrich/vantage-provider-scaleway"
	method := "GET"
	headers := map[string]string{}
	b, err := HTTPRequest(uri, method, nil, headers)
	if err != nil {
		t.Error(err)
	}
	if len(b) == 0 {
		t.Errorf("got zero-length reponse body")
	}
}

func TestTransformData(t *testing.T) {
	CONFIG = loadConfig("config.yaml")
	c := &ScalewayConsumptionResponse{
		TotalCount:                1,
		TotalDiscountUntaxedValue: 0.01,
		UpdatedAt:                 time.Now(),
		Consumptions: []*ScalewayConsumptionItem{
			{
				Value: &ScalewayValue{
					CurrencyCode: "USD",
					Units:        20000,
					Nanos:        200000000,
				},
				ProductName:    "foo",
				ResourceName:   "bar",
				SKU:            "/foo/bar",
				ProjectID:      "abcd1234",
				CategoryName:   "baz",
				Unit:           "flerbs",
				BilledQuantity: "256.2",
			},
		},
	}
	p := map[string]string{
		"abcd1234": "project-foo",
	}
	f, err := TransformData(c, p)
	if err != nil {
		t.Error(err)
	}
	if len(f) != 1 {
		t.Errorf("Got zero length for transform result")
	}
}

func TestMakeCSV(t *testing.T) {
	lines := []*FOCUS{
		{
			ChargePeriodStart: "2025-01-01",
			BillingCurrency:   "USD",
			ChargeCategory:    "Usage",
			ConsumedQuantity:  123545,
			ConsumedUnit:      "flerbs",
			BilledCost:        11.22,
			RegionId:          "FOO",
			ResourceId:        "bar",
			ResourceType:      "/foo/bar",
			ServiceCategory:   "baz",
			ServiceName:       "flop",
			Tags:              "",
		},
	}
	s, err := MakeCSV(lines)
	if err != nil {
		t.Error(err)
	}
	if len(s) == 0 {
		t.Errorf("Got zero-length string; expected more")
	}
}
