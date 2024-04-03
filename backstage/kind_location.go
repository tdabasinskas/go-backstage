package backstage

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

// KindLocation defines name for location kind.
const KindLocation = "Location"

// LocationEntityV1alpha1 is a marker that references other places to look for catalog data.
// https://github.com/backstage/backstage/blob/master/packages/catalog-model/src/schema/kinds/Location.v1alpha1.schema.json
type LocationEntityV1alpha1 struct {
	Entity

	// ApiVersion is always "backstage.io/v1alpha1".
	ApiVersion string `json:"apiVersion"`

	// Kind is always "Location".
	Kind string `json:"kind"`

	// Spec is the specification data describing the location itself.
	Spec *LocationEntityV1alpha1Spec `json:"spec"`
}

// LocationEntityV1alpha1Spec describes the specification data describing the location itself.
type LocationEntityV1alpha1Spec struct {
	// Type is the single location type, that's common to the targets specified in the spec. If it is left out, it is inherited
	// from the location type that originally read the entity data.
	Type string `json:"type,omitempty"`

	// Target as a string. Can be either an absolute path/URL (depending on the type), or a relative path
	// such as./details/catalog-info.yaml which is resolved relative to the location of this Location entity itself.
	Target string `json:"target,omitempty"`

	// Targets contains a list of targets as strings. They can all be either absolute paths/URLs (depending on the type),
	// or relative paths such as ./details/catalog-info.yaml which are resolved relative to the location of this Location
	// entity itself.
	Targets []string `json:"targets,omitempty"`

	// Presence describes whether the presence of the location target is required, and it should be considered an error if it
	// can not be found.
	Presence string `json:"presence,omitempty"`
}

// LocationCreateResponse defines POST response from location endpoints.
type LocationCreateResponse struct {
	// Exists is only set in dryRun mode.
	Exists bool `json:"exists,omitempty"`
	// Location contains details of created location.
	Location *LocationResponse `json:"location,omitempty"`
	// Entities is a list of entities that were discovered from the created location.
	Entities []Entity `json:"entities"`
}

// LocationResponse defines GET response to get single location from location endpoints.
type LocationResponse struct {
	// ID of the location.
	ID string `json:"id"`
	// Type of the location.
	Type string `json:"type"`
	// Target of the location.
	Target string `json:"target"`
}

// LocationListResponse defines GET response to get all locations from location endpoints.
type LocationListResponse struct {
	Data *LocationResponse `json:"data"`
}

// locationService handles communication with the location related methods of the Backstage Catalog API.
type locationService typedEntityService[LocationEntityV1alpha1]

// newLocationService returns a new instance of location-type entityService.
func newLocationService(s *entityService) *locationService {
	return &locationService{
		client:  s.client,
		apiPath: s.apiPath,
	}
}

// Get returns a location entity identified by the name and the namespace ("default", if not specified) it belongs to.
func (s *locationService) Get(ctx context.Context, n string, ns string) (*LocationEntityV1alpha1, *http.Response, error) {
	cs := (typedEntityService[LocationEntityV1alpha1])(*s)
	return cs.get(ctx, KindLocation, n, ns)
}

// Create creates a new location.
func (s *locationService) Create(ctx context.Context, target string, dryRun bool) (*LocationCreateResponse, *http.Response, error) {
	if target == "" {
		return nil, nil, errors.New("target cannot be empty")
	}

	path, _ := url.JoinPath(s.apiPath, "../locations")
	req, _ := s.client.newRequest(http.MethodPost, fmt.Sprintf("%s?dryRun=%t", path, dryRun), struct {
		Target string `json:"target"`
		Type   string `json:"type"`
	}{
		Target: target,
		Type:   "url",
	})

	var entity *LocationCreateResponse
	resp, err := s.client.do(ctx, req, &entity)

	return entity, resp, err

}

// List returns all locations.
func (s *locationService) List(ctx context.Context) ([]LocationListResponse, *http.Response, error) {
	path, _ := url.JoinPath(s.apiPath, "../locations")
	req, _ := s.client.newRequest(http.MethodGet, path, nil)

	var entities []LocationListResponse
	resp, err := s.client.do(ctx, req, &entities)

	return entities, resp, err
}

// GetByID returns a location identified by its ID.
func (s *locationService) GetByID(ctx context.Context, id string) (*LocationResponse, *http.Response, error) {
	path, _ := url.JoinPath(s.apiPath, "../locations", id)
	req, _ := s.client.newRequest(http.MethodGet, path, nil)

	var entity *LocationResponse
	resp, err := s.client.do(ctx, req, &entity)

	return entity, resp, err
}

// DeleteByID deletes a location identified by its ID.
func (s *locationService) DeleteByID(ctx context.Context, id string) (*http.Response, error) {
	if id == "" {
		return nil, errors.New("id cannot be empty")
	}

	path, _ := url.JoinPath(s.apiPath, "../locations", id)
	req, _ := s.client.newRequest(http.MethodDelete, path, nil)

	return s.client.do(ctx, req, nil)
}
