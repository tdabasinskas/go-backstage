package backstage

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"testing"

	"github.com/h2non/gock"
	"github.com/stretchr/testify/assert"
)

// TestKindSystemGet tests functionality of getting a system.
func TestKindSystemGet(t *testing.T) {
	const dataFile = "testdata/system.json"
	const system = "example"

	expected := SystemEntityV1alpha1{}
	expectedData, _ := os.ReadFile(dataFile)
	err := json.Unmarshal(expectedData, &expected)

	assert.FileExists(t, dataFile, "Test data file should exist")
	assert.NoError(t, err, "Unmarshal should not return an error")

	baseURL, _ := url.Parse("https://foo:1234/api")
	defer gock.Off()
	gock.New(baseURL.String()).
		MatchHeader("Accept", "application/json").
		Get(fmt.Sprintf("/catalog/entities/by-name/system/default/%s", system)).
		Reply(200).
		File(dataFile)

	c, _ := NewClient(baseURL.String(), "", nil)
	s := newSystemService(&entityService{
		client:  c,
		apiPath: "/catalog/entities",
	})

	actual, resp, err := s.Get(context.Background(), system, "")
	assert.NoError(t, err, "Get should not return an error")
	assert.NotEmpty(t, resp, "Response should not be empty")
	assert.EqualValues(t, &expected, actual, "Response body should match the one from the server")
}
