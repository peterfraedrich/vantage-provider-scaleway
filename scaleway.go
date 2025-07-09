package main

import (
	"encoding/json"
	"fmt"
	"regexp"
)

type ScalewayProvider struct{}

func (s ScalewayProvider) GetConsumption(orgID string, billingPeriod string) (*ScalewayConsumptionResponse, error) {
	consumption := &ScalewayConsumptionResponse{}
	uri := fmt.Sprintf("%s/billing/v2beta1/consumptions?organization_id=%s&billing_period=%s", CONFIG.ScalewayAPIUrl, orgID, billingPeriod)
	rawBody, err := HTTPRequest(uri, "GET", nil, map[string]string{"X-Auth-Token": CONFIG.ScalewayAPISecret})
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(rawBody, &consumption)
	if err != nil {
		return nil, err
	}
	return consumption, nil
}

func (s ScalewayProvider) GetProjects(orgID string) (map[string]string, error) {
	projects := map[string]string{}
	uri := fmt.Sprintf("%s/account/v3/projects?organization_id=%s", CONFIG.ScalewayAPIUrl, CONFIG.ScalewayOrgID)
	rawBody, err := HTTPRequest(uri, "GET", nil, map[string]string{"X-Auth-Token": CONFIG.ScalewayAPISecret})
	if err != nil {
		return nil, err
	}
	obj := &ScalewayListProjectsResponse{}
	err = json.Unmarshal(rawBody, &obj)
	if err != nil {
		return nil, err
	}
	for _, p := range obj.Projects {
		projects[p.ID] = p.Name
	}
	return projects, nil
}

func GetRegion(resourceName string) string {
	patt, _ := regexp.Compile(`(?mi).*\s\-\s(?P<region>.*(PAR|WAW|AMS).*)`)
	matches := patt.FindStringSubmatch(resourceName)
	regionIdx := patt.SubexpIndex("region")
	if len(matches) > 0 {
		return matches[regionIdx]
	}
	return "no-region"
}
