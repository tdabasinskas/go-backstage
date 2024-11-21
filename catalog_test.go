package backstage

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestNewCatalogService tests the creation of a new catalog service.
func TestNewCatalogService(t *testing.T) {
	baseURL, _ := url.Parse("http://localhost:7007/api")
	actual := newCatalogService(&Client{
		BaseURL: baseURL,
	})

	assert.Equal(t, "/catalog", actual.apiPath, "Catalog API path should match client base URL with catalog API path appended")
}
