![Vantage x Scaleway](docs/vantage%20x%20scaleway.png)

# Vantage-Provider-Scaleway
This connector allows you to export billing data for your Scaleway organization into Vantage for cost tracking purposes.

## Features
* Native currency -- uses the same currency as your Scaleway billing. Any conversion happens on the Vantage side.
* Custom tags -- add custom tags to your Scaleway resources.
* Billing period -- set the billing period easily with a flag, allowing you to easily import the data you need.

## Usage

### Quickstart
`$> ./vantage-provider-scaleway -period 2025-07`

### Flags
The application takes one flag:
* `-period` -- The billing period to pull data for. Format: `YYYY-MM`

### Env Vars
The application looks for a few env vars:
* `VANTAGE_API_KEY` -- your Vantage API key
* `SCALEWAY_API_KEY` -- the API key ID for Scaleway
* `SCALEWAY_API_SECRET` -- secret key for your Scaleway credentials

### config.yaml
The application sources `config.yaml` for configuration values. Here is an example config file:
```yaml
##### logging
# env: (dev | prod)
# dev - pretty logging with colors
# prod - log JSON
env: dev

# loglevel: (trace | debug | info | warn | error)
loglevel: debug

##### vantage
vantage_custom_provider_token: accss_crdntl_abcd1234
vantage_api_url: https://api.vantage.sh/v2

# scaleway
scaleway_api_url: https://api.scaleway.com
scaleway_org_id: abcd1234

#tags
tags:
  Product: MyProduct
  Portfolio: Marketing
  Platform: Scaleway
```

## Contributing
Contributors welcome! Feel free to submit PR's for new features, bug fixes, etc.