package backstage

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
