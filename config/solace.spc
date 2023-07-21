connection "solace" {
  plugin = "SolaceLabs/solace"

  # Get your API key from https://console.solace.cloud/api-tokens
  # This can also be set via the `SOLACE_API_TOKEN` environment variable.
  # api_token = "XXXXXXXXX"

  # The API URL. By default it is pointed to "https://api.solace.cloud/"
  # If working with the AU region , use "https://api.solacecloud.com.au/"
  # This can also be set via the `SOLACE_API_URL` environment variable.
  api_url = "https://api.solace.cloud/"
}
