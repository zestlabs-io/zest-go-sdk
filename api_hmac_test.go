package zestapi

import (
	"context"
	"os"
	"testing"

	"github.com/go-openapi/strfmt"
	"github.com/zestlabs-io/zest-go-sdk/api/client"
	"github.com/zestlabs-io/zest-go-sdk/api/client/auth_service"
)

func TestHMACAccount(t *testing.T) {
	ctx := context.TODO()
	// client := NewHMACAPIClient(&openapi.Configuration{
	// 	BasePath: "https://dev.zestlabs.cloud"},
	// 	os.Getenv("ZEST_KEY"),
	// 	os.Getenv("ZEST_SECRET"))
	// res, httpResp, err := client.AuthServiceApi.GetOwnAccount(ctx)
	// if err != nil {
	// 	t.Fatalf("Error getting account : %v", err)
	// }
	// if res.Account.AccountID == "" {
	// 	t.Errorf("Expected not empty account id")
	// }

	// if httpResp.StatusCode != http.StatusOK {
	// 	t.Errorf("Expected OK response, but got %d", httpResp.StatusCode)
	// }
	// t.Logf("Account: %v", res.Account)

	cl := NewHMACAPIClient(strfmt.Default, &client.TransportConfig{Host: "dev.zestlabs.cloud", Schemes: []string{"https"}}, os.Getenv("ZEST_KEY"), os.Getenv("ZEST_SECRET"))
	res, err := cl.AuthService.AuthServiceGetOwnAccount(&auth_service.AuthServiceGetOwnAccountParams{Context: ctx})

	if err != nil {
		t.Fatalf("Error getting account : %v", err)
	}
	if res.Payload.Account.AccountID == "" {
		t.Errorf("Expected not empty account id")
	}

	t.Logf("Account: %v", res.Payload.Account)

}
