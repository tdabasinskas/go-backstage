package backstage

// ApiEntityV1alpha1 describes an interface that can be exposed by a component. The API can be defined in different formats,
// like OpenAPI, AsyncAPI, GraphQL, gRPC, or other formats.
// https://github.com/backstage/backstage/blob/master/packages/catalog-model/src/schema/kinds/API.v1alpha1.schema.json
type ApiEntityV1alpha1 struct {
	// ApiVersion is always "backstage.io/v1alpha1".
	ApiVersion string `json:"apiVersion"`

	// Kind is always "API".
	Kind string `json:"kind"`

	// Spec is the specification data describing the API itself.
	Spec struct {
		// Type of the API definition.
		Type string `json:"type"`

		// Lifecycle state of the API.
		Lifecycle string `json:"lifecycle"`

		// Owner is entity reference to the owner of the API.
		Owner string `json:"owner"`

		// Definition of the API, based on the format defined by the type.
		Definition string `json:"definition"`

		// System is entity reference to the system that the API belongs to.
		System string `json:"system,omitempty"`
	} `json:"spec"`
}
