package zestapi

import (
	"context"
	"net/http"
	"os"
	"testing"

	"github.com/zestlabs-io/zest-go-sdk/openapi"
)

func TestHMACAccount(t *testing.T) {
	ctx := context.TODO()
	client := NewHMACAPIClient(&openapi.Configuration{
		BasePath: "https://dev.zestlabs.cloud"},
		os.Getenv("ZEST_KEY"),
		os.Getenv("ZEST_SECRET"))
	res, httpResp, err := client.AuthServiceApi.GetOwnAccount(ctx)
	if err != nil {
		t.Fatalf("Error getting account : %v", err)
	}
	if res.Account.AccountID == "" {
		t.Errorf("Expected not empty account id")
	}

	if httpResp.StatusCode != http.StatusOK {
		t.Errorf("Expected OK response, but got %d", httpResp.StatusCode)
	}
	t.Logf("Account: %v", res.Account)
}
