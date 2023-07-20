connection "solace" {
  plugin = "SolaceLabs/solace"

  # Get your API key from https://console.solace.cloud/api-tokens
  # This can also be set via the `SOLACE_API_TOKEN` environment variable.
  # api_token = "eABCGciOiJSUzI1NiIsImtpMQITm1hYXNfcHJvZF8yMDIwMDMyNiIsInR5cCI6Ikp123J9.eyJvcmciOiJ5aG11aWU1azB5cCIsIm9yZ1R5cGUiOiJFTlRFUlBSSVNFIiwic3ViIjoiMmtmMGZ6ZjVjaDciLCJwZXJtaXNzaW9ucyI6IkFBQUFBSUFQQUFBQWZ6Z0E0QUVBQUFBQUFBQUFBQUFBQUlDeHpvY2hJQWpnTC8vL2c1WGZCZDREV01NRDQ0ZS9NUT09IiwiYXBpVG9rZW5JZCI6ImNkcWdpeHo5bHNzIiwiaXNzIjoiU29sYWNlIENvcnBvcmF0aW9uIiwiaWF0IjoxNjg5Nzc0NjMxfQ.Ry_uuVdcb4A_eFmK8LNnpFq1234W3wBwDxdebal8iz5CVWC1M1P48ao4w64MH1234RXpm9ZewguzINXO8hrAbcAa41EqRWrRK5V1P2Rpa78_z2YyCaO98ah5TtPbWdpytga6BxKJsy2I2sawdD8PIm7wbJpty9e4UG6fZEVSXDpcVA2QWERTHvIPV5PwOuT7WDa-IvIDIDeQpD8Dy44QvavpcqDaQhafW5B_P4EU715RVYuepVDgFtMLVXOuyz2m5DVab6BuizQ6zoi9adb1hO46j0bRi9M-eyG3B20ho5N_h0Jhd7GrDdqIaxoLRoKbqu4QWn6GnjH4prGd5H1mrrPQR"

  # The API URL. By default it is pointed to "https://api.solace.cloud/"
  # If working with the AU region , use "https://api.solacecloud.com.au/"
  # This can also be set via the `SOLACE_API_URL` environment variable.
  api_url = "https://api.solace.cloud/"
}
