package main

import "time"

type Config struct {
	VantageCustomProviderToken string `yaml:"vantage_custom_provider_token"`
	VantageAPIUrl              string `yaml:"vantage_api_url"`
	VantageAPIKey              string `yaml:"vantage_api_key"`
	ScalewayAPIUrl             string `yaml:"scaleway_api_url"`
	ScalewayOrgID              string `yaml:"scaleway_org_id"`
	ScalewayAPIKey             string `yaml:"scaleway_api_key"`
	ScalewayAPISecret          string `yaml:"scaleway_api_secret"`
}

type FOCUS struct {
	BillingCurrency   string
	ChargeCategory    string // Credit, Refund, Discount, Tax, Usage
	ChargePeriodStart time.Time
	ChargePeriodEnd   time.Time
	ConsumedQuantity  int64
	ConsumedUnit      string
	BilledCost        float64
	RegionId          string
	ResourceId        string
	ResourceType      string
	ServiceCategory   string
	ServiceName       string
	Tags              map[string]string
}

type ScalewayConsumptionResponse struct {
	TotalCount                int64     `json:"total_count"`
	TotalDiscountUntaxedValue int64     `json:"total_discount_untaxed_value"`
	UpdatedAt                 time.Time `json:"updated_at"`
	Consumptions              []struct {
		Value struct {
			CurrencyCode string `json:"currency_code"`
			Units        int64  `json:"units"`
			Nanos        int64  `json:"nanos"`
		}
		ProductName    string `json:"product_name"`
		ResourceName   string `json:"resource_name"`
		SKU            string `json:"sku"`
		ProjectID      string `json:"project_id"`
		CategoryName   string `json:"category_name"`
		Unit           string `json:"unit"`
		BilledQuantity string `json:"billed_quantity"`
	}
}

type ScalewayListProjectsResponse struct {
	TotalCount uint64 `json:"total_count"`
	Projects   []struct {
		ID             string    `json:"id"`
		Name           string    `json:"name"`
		OrganizationID string    `json:"organization_id"`
		CreatedAt      time.Time `json:"created_at"`
		UpdatedAt      time.Time `json:"updated_at"`
		Description    string    `json:"description"`
		Qualification  struct {
			ArchitectureType string `json:"architecture_type"`
		}
	}
}
