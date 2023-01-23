package backstage

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/h2non/gock"
	"github.com/stretchr/testify/assert"
)

// TestNewClient tests the creation of a new Backstage client.
func TestNewClient(t *testing.T) {
	const baseURL = "http://localhost:7007/api"
	const nameSpace = "custom"

	c, err := NewClient(baseURL, nameSpace, nil)

	assert.NoError(t, err, "New client should not return an error")
	assert.Equal(t, baseURL, c.BaseURL.String(), "Base URL should match the one provided")
	assert.Equal(t, nameSpace, c.DefaultNamespace, "Default namespace should match the one provided")
}

// TestNewClient tests the creation of a new Backstage client from an existing HTTP client.
func TestNewClient_ExistingHTTPClient(t *testing.T) {
	const baseURL = "http://localhost:7007/api"
	const timeout = 1 * time.Second

	ec := &http.Client{
		Timeout: timeout,
	}
	c, err := NewClient(baseURL, "", ec)

	assert.NoError(t, err, "New client should not return an error")
	assert.Equal(t, timeout, c.client.Timeout, "Timeout should match the one from the existing HTTP client")
}

// TestNewClient_InvalidBaseURL tests if an error is returned when the base URL is invalid.
func TestNewClient_InvalidBaseURL(t *testing.T) {
	_, err := NewClient("\\foo:bar", "", nil)
	assert.Error(t, err, "New client should return an error when the base URL is invalid")
}

// TestNewClient_TrimBaseURLSuffix tests the creation of a new Backstage client with a base URL that has a trailing slash.
func TestNewClient_TrimBaseURLSuffix(t *testing.T) {
	c, err := NewClient("http://localhost:7007/api/", "", nil)

	assert.NoError(t, err, "New client should not return an error")
	assert.Equal(t, "http://localhost:7007/api", c.BaseURL.String(), "Base URL not contain a trailing slash")
}

// TestNewClient_DefaultNamespace tests that namespace is set to default if not provided.
func TestNewClient_DefaultNamespace(t *testing.T) {
	c, err := NewClient("http://localhost:7007/api/", "", nil)

	assert.NoError(t, err, "New client should not return an error")
	assert.Equal(t, defaultNamespaceName, c.DefaultNamespace, "Default namespace should be set to default if not provided")
}

// TestNewRequest_Get tests the creation of a new GET request.
func TestNewRequest_Get(t *testing.T) {
	const path = "/catalog/entities"
	const userAgent = "foo"

	baseURL, _ := url.Parse("http://localhost:7007/api")
	u, _ := url.JoinPath(baseURL.String(), path)
	c := &Client{
		UserAgent: userAgent,
		BaseURL:   baseURL,
	}
	req, err := c.newRequest(http.MethodGet, path, nil)

	assert.NoError(t, err, "New request should not return an error")
	assert.Equal(t, http.MethodGet, req.Method, "Request method should match the one provided")
	assert.Equal(t, u, req.URL.String(), "Request URL should match the one provided")
	assert.Equal(t, "application/json", req.Header.Get("Accept"), "Request should have an Accept header set to application/json")
	assert.Equal(t, userAgent, req.Header.Get("User-Agent"), "Request should have a User-Agent header set to the one provided")
}

// TestNewRequest_Post tests the creation of a new POST request.
func TestNewRequest_Post(t *testing.T) {
	const u = "http://localhost:7007/api/catalog/entities"

	c := &Client{}
	req, err := c.newRequest(http.MethodPost, u, struct {
		Foo string
	}{
		Foo: "Bar",
	})
	buf := new(strings.Builder)
	_, _ = io.Copy(buf, req.Body)

	assert.NoError(t, err, "New request should not return an error")
	assert.Equal(t, http.MethodPost, req.Method, "Request method should match the one provided")
	assert.Equal(t, u, req.URL.String(), "Request URL should match the one provided")
	assert.Equal(t, "application/json", req.Header.Get("Accept"), "Request should have an Accept set to application/json")
	assert.Equal(t, "application/json", req.Header.Get("Content-Type"), "Request should have a Content-Type set to application/json")
	assert.Equal(t, fmt.Sprintf("%s\n", `{"Foo":"Bar"}`), buf.String(), "Request body should match the one provided")
}

// TestNewRequest_InvalidURL tests if an error is returned when the URL of the request is invalid.
func TestNewRequest_InvalidURL(t *testing.T) {
	c := &Client{}
	_, err := c.newRequest(http.MethodGet, "\\foo:bar", nil)
	assert.Error(t, err, "New request should return an error when the URL is invalid")
}

// TestNewRequest_InvalidMethod tests if an error is returned when the method of the request is invalid.
func TestNewRequest_InvalidMethod(t *testing.T) {
	c := &Client{}
	_, err := c.newRequest("FOO:BAR", "http://localhost", nil)
	assert.Error(t, err, "New request should return an error when the method is invalid")
}

// TestNewRequest_InvalidBody tests if an error is returned when the body of the request is invalid.
func TestNewRequest_InvalidBody(t *testing.T) {
	c := &Client{}
	_, err := c.newRequest(http.MethodGet, "", math.Inf(1))
	assert.Error(t, err, "New request should return an error when the body is invalid")
}

// TestDo tests the execution of a request.
func TestDo(t *testing.T) {
	const path = "/foo/bar"
	const baseURL = "http://localhost:7007/api"

	defer gock.Off()
	gock.New(baseURL).
		Get(path).
		MatchHeader("Accept", "application/json").
		Reply(http.StatusOK).
		JSON(map[string]string{"foo": "bar"})

	u, _ := url.Parse(baseURL)
	c := &Client{
		BaseURL: u,
		client:  &http.Client{},
	}

	data := new(interface{})
	req, _ := c.newRequest(http.MethodGet, path, nil)
	resp, err := c.do(context.Background(), req, data)
	dataJSON, _ := json.Marshal(data)

	assert.NoError(t, err, "Do should not return an error")
	assert.Equal(t, http.StatusOK, resp.StatusCode, "Response status code should match the one from the server")
	assert.Equal(t, "application/json", resp.Header.Get("Content-Type"), "Response should have a Content-Type set to application/json")
	assert.Equal(t, `{"foo":"bar"}`, string(dataJSON), "Response body should match the one from the server")
}

// TestDo_EmptyBody tests the execution of a request when the response is empty.
func TestDo_EmptyBody(t *testing.T) {
	const baseURL = "http://localhost:7007/api"
	const path = "/foo/bar"

	defer gock.Off()
	gock.New(baseURL).
		Reply(http.StatusOK)

	u, _ := url.Parse(baseURL)
	c := &Client{
		BaseURL: u,
		client:  &http.Client{},
	}

	req, _ := c.newRequest(http.MethodGet, path, nil)
	resp, err := c.do(context.Background(), req, nil)
	buf := new(strings.Builder)
	_, _ = io.Copy(buf, resp.Body)

	assert.NoError(t, err, "Do should not return an error")
	assert.Empty(t, buf, "Response body should be empty")
}

// TestDo_EmptyBody tests the request that fails.
func TestDo_Fail(t *testing.T) {
	u, _ := url.Parse("http://localhost")
	c := &Client{
		BaseURL: u,
		client:  &http.Client{},
	}

	req, _ := c.newRequest(http.MethodGet, "/foo/bar", nil)
	_, err := c.do(context.Background(), req, nil)

	assert.Error(t, err, "Do should return an error when request fails")
}
