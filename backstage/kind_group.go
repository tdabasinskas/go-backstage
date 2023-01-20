package backstage

import (
	"context"
	"net/http"
)

// KindGroup defines name for group kind.
const KindGroup = "Group"

// GroupEntityV1alpha1 describes an organizational entity, such as for example a team, a business unit, or a loose collection of people in
// an interest group. Members of these groups are modeled in the catalog as kind User.
// https://github.com/backstage/backstage/blob/master/packages/catalog-model/src/schema/kinds/Group.v1alpha1.schema.json
type GroupEntityV1alpha1 struct {
	Entity

	// ApiVersion is always "backstage.io/v1alpha1".
	ApiVersion string `json:"apiVersion"`

	// Kind is always "Group".
	Kind string `json:"kind"`

	// Spec is the specification data describing the group itself.
	Spec struct {
		// Type of group. There is currently no enforced set of values for this field, so it is left up to the adopting
		// organization to choose a nomenclature that matches their org hierarchy.
		Type string `json:"type"`

		// Profile information about the group, mainly for display purposes. All fields of this structure are also optional.
		// The email would be a group email of some form, that the group may wish to be used for contacting them.
		// The picture is expected to be a URL pointing to an image that's representative of the group, and that a browser could
		// fetch and render on a group page or similar.
		Profile struct {
			// DisplayName to present to users.
			DisplayName string `json:"displayName,omitempty"`

			// Email where this entity can be reached.
			Email string `json:"email,omitempty"`

			// Picture is a URL of an image that represents this entity.
			Picture string `json:"picture,omitempty"`
		} `json:"profile,omitempty"`

		// Parent is the immediate parent group in the hierarchy, if any. Not all groups must have a parent; the catalog supports
		// multi-root hierarchies. Groups may however not have more than one parent. This field is an entity reference.
		Parent string `json:"parent,omitempty"`

		// Children contains immediate child groups of this group in the hierarchy (whose parent field points to this group).
		// The list must be present, but may be empty if there are no child groups. The items are not guaranteed to be ordered in
		// any particular way. The entries of this array are entity references.
		Children []string `json:"children"`

		// Members contains users that are members of this group. The entries of this array are entity references.
		Members []string `json:"members,omitempty"`
	} `json:"spec"`
}

// groupService handles communication with the group related methods of the Backstage Catalog API.
type groupService typedEntityService[GroupEntityV1alpha1]

// newGroupService returns a new instance of group-type entityService.
func newGroupService(s *entityService) *groupService {
	return &groupService{
		client:  s.client,
		apiPath: s.apiPath,
	}
}

// Get returns a group entity identified by the name and the namespace ("default", if not specified) it belongs to.
func (s *groupService) Get(ctx context.Context, n string, ns string) (*GroupEntityV1alpha1, *http.Response, error) {
	cs := (typedEntityService[GroupEntityV1alpha1])(*s)
	return cs.get(ctx, KindGroup, n, ns)
}
