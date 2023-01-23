package backstage

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type service struct {
	client  *Client
	apiPath string
}

const (
	userAgent            = "go-backstage"
	contentTypeJSON      = "application/json"
	defaultNamespaceName = "default"
)

// Client manages communication with the Backstage API.
type Client struct {
	// Client is an HTTP client used to communicate with the API.
	client *http.Client

	// BaseURL for API requests, e.g. http://localhost:7007/api/.
	BaseURL *url.URL

	// User agent used when communicating with the Backstage API.
	UserAgent string

	// Name of the namespace to use by default when communicating with the Backstage API.
	DefaultNamespace string

	// Catalog service to handle communication with the Backstage Catalog API.
	Catalog *catalogService
}

// NewClient returns a new Backstage API client. If a nil httpClient is  provided, a new http.Client will be used.
// To use API methods which require authentication, provide a http.Client that will perform the authentication.
func NewClient(baseURL string, defaultNamespace string, httpClient *http.Client) (*Client, error) {
	baseEndpoint, err := url.Parse(strings.TrimSuffix(baseURL, "/"))
	if err != nil {
		return nil, err
	}

	if httpClient == nil {
		httpClient = &http.Client{}
	}

	ns := defaultNamespace
	if defaultNamespace == "" {
		ns = defaultNamespaceName
	}

	c := &Client{
		client:           httpClient,
		BaseURL:          baseEndpoint,
		UserAgent:        userAgent,
		DefaultNamespace: ns,
	}

	c.Catalog = newCatalogService(c)

	return c, nil
}

// newRequest creates an API request. A relative URL can be provided in urlStr, in which case it is resolved relative to the BaseURL.
func (c *Client) newRequest(method string, urlStr string, body interface{}) (*http.Request, error) {
	u, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	var resolvedURL string
	if c.BaseURL != nil {
		u.Path, _ = url.JoinPath(c.BaseURL.Path, u.Path)
		resolvedURL = c.BaseURL.ResolveReference(u).String()
	} else {
		resolvedURL = u.String()
	}

	var buf io.ReadWriter
	if body != nil {
		buf = &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		err = enc.Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, resolvedURL, buf)
	if err != nil {
		return nil, err
	}

	if body != nil {
		req.Header.Set("Content-Type", contentTypeJSON)
	}

	req.Header.Set("Accept", contentTypeJSON)

	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}

	return req, nil
}

// do send an API request and returns the API response. The API response is JSON decoded and stored in the value pointed to by v.
func (c *Client) do(ctx context.Context, req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.client.Do(req.WithContext(ctx))
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	decErr := json.NewDecoder(resp.Body).Decode(v)
	if decErr == io.EOF {
		decErr = nil
	}

	return resp, decErr
}
