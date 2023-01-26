package backstage

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"testing"

	"github.com/h2non/gock"
	"github.com/stretchr/testify/assert"
)

// TestEntityServiceGet tests the retrieval of a specific entity.
func TestEntityServiceGet(t *testing.T) {
	const dataFile = "testdata/entities_single.json"
	const uid = "uid"

	var expected Entity
	expectedData, _ := os.ReadFile(dataFile)
	err := json.Unmarshal(expectedData, &expected)

	assert.FileExists(t, dataFile, "Test data file should exist")
	assert.NoError(t, err, "Unmarshal should not return an error")

	baseURL, _ := url.Parse("https://foo:1234/api")
	defer gock.Off()
	gock.New(baseURL.String()).
		MatchHeader("Accept", "application/json").
		Get(fmt.Sprintf("/catalog/entities/by-uid/%s", uid)).
		Reply(200).
		File(dataFile)

	c, _ := NewClient(baseURL.String(), "", nil)
	s := newEntityService(newCatalogService(c))

	actual, resp, err := s.Get(context.Background(), uid)
	assert.NoError(t, err, "Get should not return an error")
	assert.NotEmpty(t, resp, "Response should not be empty")
	assert.Equal(t, http.StatusOK, resp.StatusCode, "Response status code should be 200")
	assert.EqualValues(t, &expected, actual, "Response body should match the one from the server")
}

// TestEntityServiceList tests the retrieval of a list of entities.
func TestEntityServiceList(t *testing.T) {
	const dataFile = "testdata/entities.json"

	var expected []Entity
	expectedData, _ := os.ReadFile(dataFile)
	err := json.Unmarshal(expectedData, &expected)

	assert.FileExists(t, dataFile, "Test data file should exist")
	assert.NoError(t, err, "Unmarshal should not return an error")

	baseURL, _ := url.Parse("https://foo:1234/api")
	defer gock.Off()
	gock.New(baseURL.String()).
		MatchHeader("Accept", "application/json").
		Get("/catalog/entities").
		Reply(200).
		File(dataFile)

	c, _ := NewClient(baseURL.String(), "", nil)
	s := newEntityService(newCatalogService(c))

	actual, resp, err := s.List(context.Background(), nil)
	assert.NoError(t, err, "Get should not return an error")
	assert.NotEmpty(t, resp, "Response should not be empty")
	assert.Equal(t, http.StatusOK, resp.StatusCode, "Response status code should be 200")
	assert.EqualValues(t, expected, actual, "Response body should match the one from the server")
}

// TestEntityServiceList_Filter tests the retrieval of a list of entities with a filter.
func TestEntityServiceList_Filter(t *testing.T) {
	const dataFile = "testdata/entities_filter.json"

	var expected []Entity
	expectedData, _ := os.ReadFile(dataFile)
	err := json.Unmarshal(expectedData, &expected)

	assert.FileExists(t, dataFile, "Test data file should exist")
	assert.NoError(t, err, "Unmarshal should not return an error")

	baseURL, _ := url.Parse("https://foo:1234/api")
	defer gock.Off()
	gock.New(baseURL.String()).
		MatchHeader("Accept", "application/json").
		Get("/catalog/entities").
		MatchParam("filter", "kind=User").
		Reply(200).
		File(dataFile)

	c, _ := NewClient(baseURL.String(), "", nil)
	s := newEntityService(newCatalogService(c))

	actual, resp, err := s.List(context.Background(), &ListEntityOptions{
		Filters: ListEntityFilter{
			"kind": "User",
		},
	})
	assert.NoError(t, err, "Get should not return an error")
	assert.NotEmpty(t, resp, "Response should not be empty")
	assert.EqualValues(t, expected, actual, "Response body should match the one from the server")
}

// TestEntityServiceList_Fields tests the retrieval of a list of entities containing only the specified fields.
func TestEntityServiceList_Fields(t *testing.T) {
	const dataFile = "testdata/entities_fields.json"

	var expected []Entity
	expectedData, _ := os.ReadFile(dataFile)
	err := json.Unmarshal(expectedData, &expected)

	assert.FileExists(t, dataFile, "Test data file should exist")
	assert.NoError(t, err, "Unmarshal should not return an error")

	baseURL, _ := url.Parse("https://foo:1234/api")
	defer gock.Off()
	gock.New(baseURL.String()).
		MatchHeader("Accept", "application/json").
		Get("/catalog/entities").
		MatchParam("fields", "metadata.name").
		Reply(200).
		File(dataFile)

	c, _ := NewClient(baseURL.String(), "", nil)
	s := newEntityService(newCatalogService(c))

	actual, resp, err := s.List(context.Background(), &ListEntityOptions{
		Filters: ListEntityFilter{},
		Fields: []string{
			"metadata.name",
		},
	})
	assert.NoError(t, err, "Get should not return an error")
	assert.NotEmpty(t, resp, "Response should not be empty")
	assert.EqualValues(t, expected, actual, "Response body should match the one from the server")
}

// TestEntityServiceList_Order tests the retrieval of a list of entities sorted by the specified field.
func TestEntityServiceList_Order(t *testing.T) {
	const dataFile = "testdata/entities_order.json"

	var expected []Entity
	expectedData, _ := os.ReadFile(dataFile)
	err := json.Unmarshal(expectedData, &expected)

	assert.FileExists(t, dataFile, "Test data file should exist")
	assert.NoError(t, err, "Unmarshal should not return an error")

	baseURL, _ := url.Parse("https://foo:1234/api")
	defer gock.Off()
	gock.New(baseURL.String()).
		MatchHeader("Accept", "application/json").
		Get("/catalog/entities").
		MatchParam("order", "desc:metadata.name").
		Reply(200).
		File(dataFile)

	c, _ := NewClient(baseURL.String(), "", nil)
	s := newEntityService(newCatalogService(c))

	actual, resp, err := s.List(context.Background(), &ListEntityOptions{
		Filters: ListEntityFilter{},
		Order: []ListEntityOrder{
			{
				Direction: OrderDescending,
				Field:     "metadata.name",
			},
		},
	})
	assert.NoError(t, err, "Get should not return an error")
	assert.NotEmpty(t, resp, "Response should not be empty")
	assert.EqualValues(t, expected, actual, "Response body should match the one from the server")
}

// TestEntityServiceList_InvalidOrder tests the retrieval of a list when invalid order is provided.
func TestEntityServiceList_InvalidOrder(t *testing.T) {
	c, _ := NewClient("", "", nil)
	s := newEntityService(newCatalogService(c))

	_, _, err := s.List(context.Background(), &ListEntityOptions{
		Filters: ListEntityFilter{},
		Order: []ListEntityOrder{
			{
				Direction: "InvalidOrder",
				Field:     "metadata.name",
			},
		},
	})
	assert.Error(t, err, "Get should return an error when the order is invalid")
}

// TestEntityServiceDelete tests the deletion of an entity.
func TestEntityServiceDelete(t *testing.T) {
	const uid = "uid"

	baseURL, _ := url.Parse("https://foo:1234/api")
	defer gock.Off()
	gock.New(baseURL.String()).
		MatchHeader("Accept", "application/json").
		Delete(fmt.Sprintf("/catalog/entities/by-uid/%s", uid)).
		Reply(http.StatusNoContent)

	c, _ := NewClient(baseURL.String(), "", nil)
	s := newEntityService(newCatalogService(c))

	resp, err := s.Delete(context.Background(), uid)
	assert.NoError(t, err, "Delete should not return an error")
	assert.NotEmpty(t, resp, "Response should not be empty")
	assert.EqualValues(t, http.StatusNoContent, resp.StatusCode, "Response status code should match the one from the server")
}

// TestListEntityFilterString tests if list entity filter string is correctly generated.
func TestListEntityFilterString(t *testing.T) {
	tests := []struct {
		name     string
		filter   ListEntityFilter
		expected string
	}{
		{
			name:     "empty filter",
			filter:   ListEntityFilter{},
			expected: "",
		},
		{
			name: "filter with one key-value pair",
			filter: ListEntityFilter{
				"key1": "value1",
			},
			expected: "key1=value1",
		},
		{
			name: "filter with one key and no value",
			filter: ListEntityFilter{
				"key1": "",
			},
			expected: "key1",
		},
		{
			name: "filter with multiple key-value pairs",
			filter: ListEntityFilter{
				"key1": "value1",
				"key2": "value2",
				"key3": "",
			},
			expected: "key1=value1,key2=value2,key3",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := test.filter.string()
			assert.Equal(t, test.expected, actual, "List entity filter string should match expected value")
		})
	}
}

// TestListEntityOrderString tests if list entity order string is correctly generated.
func TestListEntityOrderString(t *testing.T) {
	tests := []struct {
		name      string
		order     ListEntityOrder
		expected  string
		shouldErr bool
	}{
		{
			name: "valid order",
			order: ListEntityOrder{
				Direction: OrderAscending,
				Field:     "field1",
			},
			expected:  "asc:field1",
			shouldErr: false,
		},
		{
			name: "valid descending order",
			order: ListEntityOrder{
				Direction: OrderDescending,
				Field:     "field2",
			},
			expected:  "desc:field2",
			shouldErr: false,
		},
		{
			name: "invalid order direction",
			order: ListEntityOrder{
				Direction: "invalid",
				Field:     "field3",
			},
			expected:  "",
			shouldErr: true,
		},
		{
			name:      "empty order",
			order:     ListEntityOrder{},
			expected:  "",
			shouldErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual, err := test.order.string()
			if test.shouldErr {
				assert.Error(t, err, "Expected error but got nil")
			} else {
				assert.NoError(t, err, "Unexpected error")
				assert.Equal(t, test.expected, actual, "List entity order string should match expected value")
			}
		})
	}
}
