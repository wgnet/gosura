package gosura

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"net/http"
	"strings"
)

const (
	DEFAULT_URL string = `http://localhost:8080`
)

type Client struct {
	headers  http.Header
	client   *http.Client
	url      string
	endpoint string
}

func NewHasuraClient() *Client {
	headers := make(http.Header)
	headers.Add("Content-Type", "application/json")

	return &Client{
		headers:  headers,
		client:   &http.Client{},
		endpoint: DEFAULT_ENDPOINT_PATH,
		url:      DEFAULT_URL,
	}
}

func (c *Client) AddHeader(key, value string) *Client {
	c.headers.Add(key, value)
	return c
}

func (c *Client) SetAdminSecret(secret string) *Client {
	c.AddHeader("X-Hasura-Admin-Secret", secret)
	return c
}

func (c *Client) SkipTLSVerify(skip bool) *Client {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: skip,
		},
	}
	c.client.Transport = transport
	return c
}

func (c *Client) URL(url string) *Client {
	c.url = strings.TrimRight(url, "/")
	return c
}

func (c *Client) getUrl() string {
	return strings.Join([]string{c.url, c.endpoint}, "")
}

func (c *Client) Endpoint(endpoint string) *Client {
	c.endpoint = endpoint
	return c
}

func (c *Client) Do(query Query) (interface{}, error) {
	data, err := query.Byte()
	if err != nil {
		return nil, fmt.Errorf("Can't get bytes from query: %w", err)
	}

	reader := bytes.NewReader(data)
	req, err := http.NewRequest(query.Method(), c.getUrl(), reader)
	if err != nil {
		return nil, fmt.Errorf("Can't make a request: %w", err)
	}
	req.Header = c.headers

	response, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error in Hasura call: %w", err)
	}
	defer response.Body.Close()

	return query.CheckResponse(response, err)
}
