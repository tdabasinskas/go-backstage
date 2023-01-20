package backstage

import (
	"context"
	"net/http"
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
	Spec struct {
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

		// Presence describes whether the presence of the location target is required and it should be considered an error if it
		// can not be found.
		Presence string `json:"presence,omitempty"`
	} `json:"spec"`
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
