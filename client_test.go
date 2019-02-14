package restclient

import (
	"bufio"
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
		MaxBody:      1024 * 4,
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

func TestDeleteARLS(t *testing.T) {
	conf := edgegrid.Config{
		Host:         host,
		AccessToken:  accessToken,
		ClientSecret: secret,
		ClientToken:  clientToken,
		MaxBody:      1024 * 50,
		Debug:        false,
	}
	c, err := NewClient(Staging, conf)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	arls := make([]string,0)
	// 500で50kbyte超え
	fp, err := os.Open("./testdata/400_arls.txt")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	scan := bufio.NewScanner(fp)
	for scan.Scan() {
		arls = append(arls, scan.Text())
	}
	defer fp.Close()


	ctx := context.Background()
	res, err := c.Delete(ctx, arls...)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if res == nil {
		t.Fatalf("unexpected error: response struct is empty")
	}
}
