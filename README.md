# ZEST Cloud Golang API

This project is based on OpenAPI generation, including additional functionallity for singing requests with HMAC key and secret.

## Prerequisites

You will need a valid `ACCESS_KEY` and `ACCESS_SECRET` for your user from the Zest Cloud.

## Usage

```sh
go get github.com/zestlabs-io/zest-go-sdk
```

```go
import (
  sdk "github.com/zestlabs-io/zest-go-sdk"
  "github.com/go-openapi/strfmt"
  "github.com/zestlabs-io/zest-go-sdk/api/client"
  "github.com/zestlabs-io/zest-go-sdk/api/client/auth_service"
  "context"
  "fmt"
)

func main() {

  cl := sdk.NewHMACAPIClient(strfmt.Default, 
      &client.TransportConfig{
        Host: "dev.zestlabs.cloud", 
        Schemes: []string{"https"}}, 
      os.Getenv("ZEST_KEY"), 
      os.Getenv("ZEST_SECRET"))

  // Call Auth service to fetch the current account details 
  res, err := cl.AuthService.AuthServiceGetOwnAccount(&auth_service.AuthServiceGetOwnAccountParams{Context: ctx})
  
  if err != nil {
    panic(fmt.Errorf("Error getting account : %v", err))
  }
  
  fmt.Printf("Fetched own account with ID: %s", res.Payload.Account.AccountID)
}
```

## License

MIT
