connection "solace" {
  plugin = "local/solace"

  # Get your API key from https://console.solace.cloud/api-tokens
  api_token = ""

  # The API URL. By default it is pointed to "https://api.solace.cloud/"
  # If working with the AU region , use "https://api.solacecloud.com.au/"
  api_url = "https://api.solace.cloud/"

  # Rate limiting
  # Solace Cloud REST API limits the number of requests you can send to the Cloud REST API API.
  # Solace Cloud REST API sets the rate limits based on your organization plan:
  # - Core: 60 per minute
  # - Pro: 120 per minute
  # - Teams: 240 per minute
  # - Enterprise: 1 000 per minute
  # We recommend to set a value below (or at most at) 80% of your total limit.
  # The default value is 50 if you don't override it here.
  # rate_limit = 50
}
