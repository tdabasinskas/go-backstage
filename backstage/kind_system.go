package backstage

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
		// Owner is an entity reference to the owner of the component.
		Owner string `json:"owner"`

		// Domain is an entity reference to the domain that the system belongs to.
		Domain string `json:"domain,omitempty"`
	} `json:"spec"`
}
