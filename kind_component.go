package backstage

import (
	"context"
	"net/http"
)

// KindComponent defines name for component kind.
const KindComponent = "Component"

// ComponentEntityV1alpha1 describes a software component. It is typically intimately linked to the source code that constitutes the
// component, and should be what a developer may regard a "unit of software", usually with a distinct deployable or linkable artifact.
// https://github.com/backstage/backstage/blob/master/packages/catalog-model/src/schema/kinds/Component.v1alpha1.schema.json
type ComponentEntityV1alpha1 struct {
	Entity

	// ApiVersion is always "backstage.io/v1alpha1".
	ApiVersion string `json:"apiVersion" yaml:"apiVersion"`

	// Kind is always "Component".
	Kind string `json:"kind" yaml:"kind"`

	// Spec is the specification data describing the component itself.
	Spec *ComponentEntityV1alpha1Spec `json:"spec"  yaml:"spec"`
}

// ComponentEntityV1alpha1Spec describes the specification data describing the component itself.
type ComponentEntityV1alpha1Spec struct {
	// Type of component.
	Type string `json:"type" yaml:"type"`

	// Lifecycle state of the component.
	Lifecycle string `json:"lifecycle" yaml:"lifecycle"`

	// Owner is an entity reference to the owner of the component.
	Owner string `json:"owner" yaml:"owner"`

	// SubcomponentOf is an entity reference to another component of which the component is a part.
	SubcomponentOf string `json:"subcomponentOf,omitempty" yaml:"subcomponentOf,omitempty"`

	// ProvidesApis is an array of entity references to the APIs that are provided by the component.
	ProvidesApis []string `json:"providesApis,omitempty" yaml:"providesApis,omitempty"`

	// ConsumesApis is an array of entity references to the APIs that are consumed by the component.
	ConsumesApis []string `json:"consumesApis,omitempty" yaml:"consumesApis,omitempty"`

	// DependsOn is an array of entity references to the components and resources that the component depends on.
	DependsOn []string `json:"dependsOn,omitempty" yaml:"dependsOn,omitempty"`

	// System is an array of references to other entities that the component depends on to function.
	System string `json:"system,omitempty" yaml:"system,omitempty"`
}

// componentService handles communication with the component related methods of the Backstage Catalog API.
type componentService typedEntityService[ComponentEntityV1alpha1]

// newComponentService returns a new instance of component-type entityService.
func newComponentService(s *entityService) *componentService {
	return &componentService{
		client:  s.client,
		apiPath: s.apiPath,
	}
}

// Get returns a component entity identified by the name and the namespace ("default", if not specified) it belongs to.
func (s *componentService) Get(ctx context.Context, n string, ns string) (*ComponentEntityV1alpha1, *http.Response, error) {
	cs := (typedEntityService[ComponentEntityV1alpha1])(*s)
	return cs.get(ctx, KindComponent, n, ns)
}
