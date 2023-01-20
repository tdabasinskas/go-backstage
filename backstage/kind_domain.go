package backstage

import (
	"context"
	"net/http"
)

// KindDomain defines name for domain kind.
const KindDomain = "Domain"

// DomainEntityV1alpha1 groups a collection of systems that share terminology, domain models, business purpose, or documentation,
// i.e. form a bounded context.
// https://github.com/backstage/backstage/blob/master/packages/catalog-model/src/schema/kinds/Domain.v1alpha1.schema.json
type DomainEntityV1alpha1 struct {
	Entity

	// ApiVersion is always "backstage.io/v1alpha1".
	ApiVersion string `json:"apiVersion"`

	// Kind is always "Domain".
	Kind string `json:"kind"`

	// Spec is the specification data describing the domain itself.
	Spec struct {
		// Owner is an entity reference to the owner of the component.
		Owner string `json:"owner"`
	} `json:"spec"`
}

// domainService handles communication with the domain related methods of the Backstage Catalog API.
type domainService typedEntityService[DomainEntityV1alpha1]

// newDomainService returns a new instance of domain-type entityService.
func newDomainService(s *entityService) *domainService {
	return &domainService{
		client:  s.client,
		apiPath: s.apiPath,
	}
}

// Get returns a domain entity identified by the name and the namespace ("default", if not specified) it belongs to.
func (s *domainService) Get(ctx context.Context, n string, ns string) (*DomainEntityV1alpha1, *http.Response, error) {
	cs := (typedEntityService[DomainEntityV1alpha1])(*s)
	return cs.get(ctx, KindDomain, n, ns)
}
