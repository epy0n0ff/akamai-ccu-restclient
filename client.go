package restclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const (
	baseURL       = "https://api.ccu.akamai.com"
	purgeEndPoint = "/ccu/v2/queues/default"
)

type Client struct {
	baseURL       *url.URL
	basicUser     string
	basicPassword string
	client        *http.Client
}

// NewClient function returns akamai ccu rest client
func NewClient(basicUser, basicPassword string) (*Client, error) {
	base, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	return &Client{
		baseURL:       base,
		basicUser:     basicUser,
		basicPassword: basicPassword,
		client:        http.DefaultClient,
	}, nil
}

// Purge function does purge request with ARL objects. Not support cpcode type.
// If successful, this method will return a response that includes progress url and more.
// If unsuccessful, this will return  an error.
func (c *Client) Purge(ctx context.Context, objects ...string) (*Response, error) {
	purge := PurgeRequest{Objects: objects}
	body := new(bytes.Buffer)
	if err := c.encodeBody(body, &purge); err != nil {
		return nil, err
	}

	headerOps := map[string]string{"Content-Type": "application/json"}
	req, err := c.newRequest(ctx, "POST", c.getURL(purgeEndPoint, ""), body, headerOps)
	if err != nil {
		return nil, err
	}

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var purgeResponse Response
	if err := c.decodeBody(res, &purgeResponse); err != nil {
		return nil, err
	}

	return &purgeResponse, nil
}

func (c *Client) newRequest(ctx context.Context, method, url string, body io.Reader, headerOps map[string]string) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", fmt.Sprintf("akamai-ccu-restclient/%s", version))
	req.SetBasicAuth(c.basicUser, c.basicPassword)

	for name, value := range headerOps {
		req.Header.Set(name, value)
	}

	return req.WithContext(ctx), nil
}

func (c *Client) decodeBody(resp *http.Response, v interface{}) error {
	return json.NewDecoder(resp.Body).Decode(v)
}

func (c *Client) encodeBody(writer io.Writer, v interface{}) error {
	return json.NewEncoder(writer).Encode(v)
}

func (c *Client) getURL(path, query string) string {
	u := *c.baseURL
	u.Path = path
	u.RawQuery = query
	return u.String()
}
