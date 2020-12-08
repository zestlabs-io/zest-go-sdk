# ZEST Cloud Golang API

This project is based on OpenAPI generation, including additional functionallity for singing requests with HMAC key and secret.

## Prerequisites

You need to obtain `ACCESS_KEY` and `ACCESS_SECRET` for your user in Zest Cloud

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
  ctx := context.TODO()
  cl := NewHMACAPIClient(strfmt.Default, 
      &client.TransportConfig{
        Host: "dev.zestlabs.cloud", 
        Schemes: []string{"https"}}, 
      os.Getenv("ZEST_KEY"), os.Getenv("ZEST_SECRET"))
  // Call Auth service to fetch the Current account    
  res, err := cl.AuthService.AuthServiceGetOwnAccount(&auth_service.AuthServiceGetOwnAccountParams{Context: ctx})
  
  if err != nil {
		panic(fmt.Errorf("Error getting account : %v", err))
  }
  
  fmt.Printf("Fetched my current account: %s", res.Payload.Account.AccountID)
}
```

## License

MIT