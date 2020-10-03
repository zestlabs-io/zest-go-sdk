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
  "github.com/zestlabs-io/zest-go-sdk/openapi"
  "context"
  "fmt"
)

func main() {
  ctx := context.TODO()
  client := NewHMACAPIClient(&openapi.Configuration{
    BasePath: "https://dev.zestlabs.cloud"},
    os.Getenv("ZEST_KEY"),
    os.Getenv("ZEST_SECRET"))
  _, _, _ = client.AuthServiceApi.GetOwnAccount(ctx)
}
```

## License

MIT