package restclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/akamai/AkamaiOPEN-edgegrid-golang/edgegrid"
)

const (
	deleteByURL     = "/ccu/v3/delete/url/"
	invalidateByURL = "/ccu/v3/invalidate/url/"
)

const (
	Staging    Network = "staging"
	Production Network = "production"
)

type Network string

type Client struct {
	network Network
	config  edgegrid.Config
	client  *http.Client
}

type ClientOpts struct {
	HTTPClient *http.Client
}

// NewClient function returns akamai ccu v3 rest client
func NewClient(network Network, config edgegrid.Config, opts *ClientOpts) (*Client, error) {
	var client = http.DefaultClient
	if opts != nil {
		client = opts.HTTPClient
	}

	return &Client{
		network: network,
		config:  config,
		client:  client,
	}, nil
}

func (c *Client) purge(ctx context.Context, path string, objects ...string) (*DeleteResponse, error) {
	purge := PurgeRequest{Objects: objects}
	body := new(bytes.Buffer)
	if err := c.encodeBody(body, &purge); err != nil {
		return nil, err
	}

	req, err := c.newRequest(ctx, "POST", fmt.Sprintf("https://%s%s%s", c.config.Host, path, c.network), body, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	switch res.StatusCode {
	case http.StatusTooManyRequests:
		var resp RateLimitResponse
		if err := c.decodeBody(res, &resp); err != nil {
			return nil, err
		}
		return nil, resp
	case http.StatusCreated:
		var resp DeleteResponse
		if err := c.decodeBody(res, &resp); err != nil {
			return nil, err
		}
		return &resp, nil
	case http.StatusRequestEntityTooLarge:
		return nil, fmt.Errorf("request entity needed to be under the 50,000 byte limit")
	default:
		return nil, fmt.Errorf(res.Status)
	}
}

// Delete function does delete request with ARL objects. Not support cpcode type.
// If successful, this method will return a response that includes progress estimatedSeconds and more.
// If unsuccessful, this will return  an error.
func (c *Client) Delete(ctx context.Context, objects ...string) (*DeleteResponse, error) {
	return c.purge(ctx, deleteByURL, objects...)
}

// Invalidate function does invalidate request with ARL objects. Not support cpcode type.
// If successful, this method will return a response that includes progress estimatedSeconds and more.
// If unsuccessful, this will return  an error.
func (c *Client) Invalidate(ctx context.Context, objects ...string) (*DeleteResponse, error) {
	return c.purge(ctx, invalidateByURL, objects...)
}

// ExceededRateLimit function check error type.
// If exceeded rate limit, this method will return RateLimitResponse,
// else this will return a raw error.
func (c *Client) ExceededRateLimit(err error) (*RateLimitResponse, error) {
	switch err.(type) {
	case RateLimitResponse:
		if resp, ok := err.(RateLimitResponse); ok {
			return &resp, nil
		}

		return nil, err
	default:
		return nil, err
	}
}

func (c *Client) newRequest(ctx context.Context, method, url string, body io.Reader, headerOps map[string]string) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", fmt.Sprintf("akamai-ccu-restclient/%s", version))
	req.Header.Set("Content-Type", "application/json")

	for name, value := range headerOps {
		req.Header.Set(name, value)
	}

	req = edgegrid.AddRequestHeader(c.config, req)

	return req.WithContext(ctx), nil
}

func (c *Client) decodeBody(resp *http.Response, v interface{}) error {
	return json.NewDecoder(resp.Body).Decode(v)
}

func (c *Client) encodeBody(writer io.Writer, v interface{}) error {
	return json.NewEncoder(writer).Encode(v)
}
