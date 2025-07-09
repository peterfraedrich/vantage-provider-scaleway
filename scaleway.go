package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ScalewayProvider struct{}

func GetConsumption(orgID string, billingPeriod string) (*ScalewayConsumptionResponse, error) {

	return nil, nil
}

func getProjects(orgID string) (map[string]string, error) {
	projects := map[string]string{}
	uri := fmt.Sprintf("%s/account/v3/projects?organization_id=%s", CONFIG.VantageAPIUrl, CONFIG.ScalewayOrgID)
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("X-Auth-Token", CONFIG.ScalewayAPISecret)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		res.Body.Close()
		return nil, err
	}
	defer res.Body.Close()
	rawBody, err := io.ReadAll(res.Body)
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
