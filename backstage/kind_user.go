package backstage

import (
	"context"
	"net/http"
)

// KindUser defines name for user kind.
const KindUser = "User"

// UserEntityV1alpha1 describes a person, such as an employee, a contractor, or similar. Users belong to Group entities in the catalog.
// These catalog user entries are connected to the way that authentication within the Backstage ecosystem works.
// https://github.com/backstage/backstage/blob/master/packages/catalog-model/src/schema/kinds/User.v1alpha1.schema.json
type UserEntityV1alpha1 struct {
	Entity

	// ApiVersion is always "backstage.io/v1alpha1".
	ApiVersion string `json:"apiVersion"`

	// Kind is always "User".
	Kind string `json:"kind"`

	// Spec is the specification data describing the user itself.
	Spec *UserEntityV1alpha1Spec `json:"spec"`
}

// UserEntityV1alpha1Spec describes the specification data describing the user itself.
type UserEntityV1alpha1Spec struct {
	// Profile information about the user, mainly for display purposes. All fields of this structure are also optional.
	// The email would be a primary email of some form, that the user may wish to be used for contacting them.
	// The picture is expected to be a URL pointing to an image that's representative of the user, and that a browser could
	// fetch and render on a profile page or similar.
	Profile struct {
		// DisplayName is a simple display name to present to users.
		DisplayName string `json:"displayName,omitempty"`

		// Email where this user can be reached.
		Email string `json:"email,omitempty"`

		// Picture is a URL of an image that represents this user.
		Picture string `json:"picture,omitempty"`
	} `json:"profile,omitempty"`

	// MemberOf is the list of groups that the user is a direct member of (i.e., no transitive memberships are listed here).
	// The list must be present, but may be empty if the user is not member of any groups. The items are not guaranteed to be
	// ordered in any particular way. The entries of this array are entity references.
	MemberOf []string `json:"memberOf,omitempty"`
}

// userService handles communication with the user methods of the Backstage Catalog API.
type userService typedEntityService[UserEntityV1alpha1]

// newUserService returns a new instance of user-type entityService.
func newUserService(s *entityService) *userService {
	return &userService{
		client:  s.client,
		apiPath: s.apiPath,
	}
}

// Get returns a user entity identified by the name and the namespace ("default", if not specified) it belongs to.
func (s *userService) Get(ctx context.Context, n string, ns string) (*UserEntityV1alpha1, *http.Response, error) {
	cs := (typedEntityService[UserEntityV1alpha1])(*s)
	return cs.get(ctx, KindUser, n, ns)
}
