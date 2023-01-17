package backstage

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
	Spec struct {
		// Type of resource.
		Type string `json:"type"`

		// Owner is an entity reference to the owner of the resource.
		Owner string `json:"owner"`

		// DependsOn is an array of references to other entities that the resource depends on to function.
		DependsOn []string `json:"dependsOn,omitempty"`

		// System is an entity reference to the system that the resource belongs to.
		System string `json:"system,omitempty"`
	} `json:"spec"`
}
