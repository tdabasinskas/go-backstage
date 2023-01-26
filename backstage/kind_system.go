package backstage

import (
	"context"
	"net/http"
)

// KindSystem defines name for system kind.
const KindSystem = "System"

// SystemEntityV1alpha1 is a collection of resources and components. The system may expose or consume one or several APIs. It is viewed as
// abstraction level that provides potential consumers insights into exposed features without needing a too detailed view into the details
// of all components. This also gives the owning team the possibility to decide about published artifacts and APIs.
// https://github.com/backstage/backstage/blob/master/packages/catalog-model/src/schema/kinds/System.v1alpha1.schema.json
type SystemEntityV1alpha1 struct {
	Entity

	// ApiVersion is always "backstage.io/v1alpha1".
	ApiVersion string `json:"apiVersion"`

	// Kind is always "System".
	Kind string `json:"kind"`

	// Spec is the specification data describing the system itself.
	Spec struct {
		// Owner is an entity reference to the owner of the system.
		Owner string `json:"owner"`

		// Domain is an entity reference to the domain that the system belongs to.
		Domain string `json:"domain,omitempty"`
	} `json:"spec"`
}

// systemService handles communication with the system methods of the Backstage Catalog API.
type systemService typedEntityService[SystemEntityV1alpha1]

// newSystemService returns a new instance of system-type entityService.
func newSystemService(s *entityService) *systemService {
	return &systemService{
		client:  s.client,
		apiPath: s.apiPath,
	}
}

// Get returns a system entity identified by the name and the namespace ("default", if not specified) it belongs to.
func (s *systemService) Get(ctx context.Context, n string, ns string) (*SystemEntityV1alpha1, *http.Response, error) {
	cs := (typedEntityService[SystemEntityV1alpha1])(*s)
	return cs.get(ctx, KindSystem, n, ns)
}
