package backstage

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestNewCatalogService tests the creation of a new catalog service.
func TestNewCatalogService(t *testing.T) {
	baseURL, _ := url.Parse("http://localhost:7007/api")
	c := &Client{
		BaseURL: baseURL,
	}
	s := newCatalogService(c)

	if actual, expected := s.apiPath.String(), baseURL.String()+catalogApiPath; actual != expected {
		assert.Equal(t, expected, actual, "Catalog API path should match client base URL with catalog API path appended")
	}
}
