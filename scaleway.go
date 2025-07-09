package main

import (
	"encoding/json"
	"fmt"
	"regexp"
)

type ScalewayProvider struct{}

func (s ScalewayProvider) GetConsumption(orgID string, billingPeriod string) (*ScalewayConsumptionResponse, error) {
	LOG.Debug().Str("org_id", orgID).Str("billing_period", billingPeriod).Msg("Getting Scaleway consumption")
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
	LOG.Debug().Int("returned_objects", len(consumption.Consumptions)).Msg("Fetched consumption from Scaleway API")
	return consumption, nil
}

func (s ScalewayProvider) GetProjects(orgID string) (map[string]string, error) {
	LOG.Debug().Str("org_id", orgID).Msg("Getting Scaleway projects")
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
	LOG.Debug().Int("returned_projects", len(projects)).Msg("Fetched projects from Scaleway API")
	return projects, nil
}

func GetRegion(resourceName string) string {
	LOG.Trace().Str("resource_name", resourceName).Msg("Getting region from resource name")
	patt, _ := regexp.Compile(`(?mi).*\s\-\s(?P<region>.*(PAR|WAW|AMS).*)`)
	matches := patt.FindStringSubmatch(resourceName)
	regionIdx := patt.SubexpIndex("region")
	if len(matches) > 0 {
		region := matches[regionIdx]
		LOG.Trace().Str("region", region).Msg("Found region identifier")
		return region
	}
	LOG.Trace().Str("region", "no-region").Msg("No matches for region, using fallback")
	return "no-region"
}
