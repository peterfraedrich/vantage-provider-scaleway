package main

import "testing"

func TestGetRegion(t *testing.T) {
	testStrings := map[string]string{
		"Multi-AZ - PAR": "PAR",
		"Private Network 10G per month - FR-PAR-2": "FR-PAR-2",
		"Zonal Ipv6 - FR-PAR-2":                    "FR-PAR-2",
		"Scaleway's logs - NL-AMS":                 "NL-AMS",
		"Run - PAR2":                               "PAR2",
		"Foo":                                      "no-region",
		"Foo - Bar":                                "no-region",
	}
	for k, v := range testStrings {
		res := GetRegion(k)
		if res != v {
			t.Errorf("Expected %s but got %s", res, v)
		}
	}
}
