package backstage

// ComponentEntityV1alpha1 describes a software component. It is typically intimately linked to the source code that constitutes the
// component, and should be what a developer may regard a "unit of software", usually with a distinct deployable or linkable artifact.
// https://github.com/backstage/backstage/blob/master/packages/catalog-model/src/schema/kinds/Component.v1alpha1.schema.json
type ComponentEntityV1alpha1 struct {
	Entity
	// ApiVersion is always "backstage.io/v1alpha1".
	ApiVersion string `json:"apiVersion"`

	// Kind is always "Component".
	Kind string `json:"kind"`

	// Spec is the specification data describing the component itself.
	Spec struct {
		// Type of component.
		Type string `json:"type"`

		// Lifecycle state of the component.
		Lifecycle string `json:"lifecycle"`

		// Owner is an entity reference to the owner of the component.
		Owner string `json:"owner"`

		// SubcomponentOf is an entity reference to another component of which the component is a part.
		SubcomponentOf string `json:"subcomponentOf,omitempty"`

		// ProvidesApis is an array of entity references to the APIs that are provided by the component.
		ProvidesApis []string `json:"providesApis,omitempty"`

		// ConsumesApis is an array of entity references to the APIs that are consumed by the component.
		ConsumesApis []string `json:"consumesApis,omitempty"`

		// DependsOn is an array of entity references to the components and resources that the component depends on.
		DependsOn []string `json:"dependsOn,omitempty"`

		// System is an array of references to other entities that the component depends on to function.
		System string `json:"system,omitempty"`
	} `json:"spec"`
}
