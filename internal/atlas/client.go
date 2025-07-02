package atlas

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/koo-arch/servant-trait-filter-backend/internal/model"
)

type Client interface {
	FetchServants(ctx context.Context, region string) ([]Servant, error)
	FetchTraits(ctx context.Context, region string) ([]model.Trait, error)
}

type client struct {
	baseURL *url.URL
	client *http.Client
	userAgent string
	etagCache map[string]string
}

type Option func(*client)

var baseURL = "https://api.atlasacademy.io"

func NewClient(options ...Option) Client {
	parsedURL, err := url.Parse(baseURL)
	if err != nil {
		return nil
	}

	client := &client{
		baseURL: parsedURL,
		client:  &http.Client{
			Timeout: 15 * time.Second, 
		},
		userAgent: "servant-trait-filter/1.0",
		etagCache: make(map[string]string),
	}

	for _, option := range options {
		option(client)
	}

	return client
}

func WithTimeout(timeout time.Duration) Option {
	return func(c *client) {
		c.client.Timeout = timeout
	}
}

func WithHTTPClient(httpClient *http.Client) Option {
	return func(c *client) {
		c.client = httpClient
	}
}

func WithUserAgent(userAgent string) Option {
	return func(c *client) {
		c.userAgent = userAgent
	}
}

func (c *client) NewRequestAndDo(ctx context.Context, method string, apiURL *url.URL, body any) (*http.Response, error) {
	var reqBody []byte = nil
	var err error
	if body != nil {
		reqBody, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequestWithContext(ctx, method, c.baseURL.ResolveReference(apiURL).String(), bytes.NewReader(reqBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", c.userAgent)
	req.Header.Set("Accept-Encoding", "gzip")
	if etag, ok := c.etagCache[apiURL.String()]; ok {
		req.Header.Set("If-None-Match", etag)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	if etag := resp.Header.Get("ETag"); etag != "" {
		c.etagCache[apiURL.String()] = etag
	}

	return resp, nil
}

func (c *client) DoJSON(ctx context.Context, method string, apiURL *url.URL, in, out any) error {
	resp, err := c.NewRequestAndDo(ctx, method, apiURL, in)
	if err != nil {
		return err
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			fmt.Printf("Error closing response body: %v\n", err)
		}
	}()

	// 304: 差分なし　→ outはそのまま
	if resp.StatusCode == http.StatusNotModified {
		return nil
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		b, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("atlas: %s - %s", resp.Status, b)
	}

	reader := c.getReader(resp)

	return json.NewDecoder(reader).Decode(out)
}

func (c *client) getReader(resp *http.Response) io.Reader {
	if resp.Header.Get("Content-Encoding") != "gzip" {
		return resp.Body
	}
	
	gzipReader, err := gzip.NewReader(resp.Body)
	if err != nil {
		return resp.Body
	}

	return gzipReader
}