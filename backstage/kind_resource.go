package backstage

import (
	"context"
	"net/http"
)

// KindResource defines name for resource kind.
const KindResource = "Resource"

// ResourceEntityV1alpha1 describes the infrastructure a system needs to operate, like BigTable databases, Pub/Sub topics, S3 buckets
// or CDNs. Modelling them together with components and systems allows to visualize resource footprint, and create tooling around them.
// https://github.com/backstage/backstage/blob/master/packages/catalog-model/src/schema/kinds/Resource.v1alpha1.schema.json
type ResourceEntityV1alpha1 struct {
	Entity

	// ApiVersion is always "backstage.io/v1alpha1".
	ApiVersion string `json:"apiVersion"`

	// Kind is always "Resource".
	Kind string `json:"kind"`

	// Spec is the specification data describing the resource itself.
	Spec *ResourceEntityV1alpha1Spec `json:"spec"`
}

// ResourceEntityV1alpha1Spec describes the specification data describing the resource itself.
type ResourceEntityV1alpha1Spec struct {
	// Type of resource.
	Type string `json:"type"`

	// Owner is an entity reference to the owner of the resource.
	Owner string `json:"owner"`

	// DependsOn is an array of references to other entities that the resource depends on to function.
	DependsOn []string `json:"dependsOn,omitempty"`

	// System is an entity reference to the system that the resource belongs to.
	System string `json:"system,omitempty"`
}

// resourceService handles communication with the resource related methods of the Backstage Catalog API.
type resourceService typedEntityService[ResourceEntityV1alpha1]

// newResourceService returns a new instance of resource-type entityService.
func newResourceService(s *entityService) *resourceService {
	return &resourceService{
		client:  s.client,
		apiPath: s.apiPath,
	}
}

// Get returns a resource entity identified by the name and the namespace ("default", if not specified) it belongs to.
func (s *resourceService) Get(ctx context.Context, n string, ns string) (*ResourceEntityV1alpha1, *http.Response, error) {
	cs := (typedEntityService[ResourceEntityV1alpha1])(*s)
	return cs.get(ctx, KindResource, n, ns)
}
