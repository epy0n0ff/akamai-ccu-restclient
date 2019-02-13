package restclient

import (
	"context"
	"os"
	"testing"

	"github.com/akamai/AkamaiOPEN-edgegrid-golang/edgegrid"
)

var secret = os.Getenv("AKA_CCU_CLIENT_SECRET")
var host = os.Getenv("AKA_CCU_HOST")
var accessToken = os.Getenv("AKA_CCU_ACCESS_TOKEN")
var clientToken = os.Getenv("AKA_CCU_CLIENT_TOKEN")
var arl = os.Getenv("TEST_ARL")

func TestDelete(t *testing.T) {
	conf := edgegrid.Config{
		Host:         host,
		AccessToken:  accessToken,
		ClientSecret: secret,
		ClientToken:  clientToken,
		MaxBody:      1024 * 8,
		Debug:        false,
	}
	c, err := NewClient(Staging, conf)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	ctx := context.Background()
	res, err := c.Delete(ctx, arl)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if res == nil {
		t.Fatalf("unexpected error: response struct is empty")
	}
}
