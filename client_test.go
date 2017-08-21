package restclient

import (
	"context"
	"os"
	"testing"
)

var user = os.Getenv("AKAMAI_USER")
var pass = os.Getenv("AKAMAI_PASSWORD")
var arl = os.Getenv("TEST_ARL")

func TestPurge(t *testing.T) {
	c, err := NewClient(user, pass, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	ctx := context.Background()
	res, err := c.Purge(ctx, arl)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if res == nil {
		t.Fatalf("unexpected error: response struct is empty")
	}
}

func TestGetQueueLength(t *testing.T) {
	c, err := NewClient(user, pass, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	ctx := context.Background()
	res, err := c.GetQueueLength(ctx)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if res == nil {
		t.Fatalf("unexpected error: response struct is empty")
	}
}
