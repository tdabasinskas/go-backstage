package backstage

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// Entity represents the parts of the format that's common to all versions/kinds of entity.
// https://github.com/backstage/backstage/blob/master/packages/catalog-model/src/schema/Entity.schema.json
type Entity struct {
	// ApiVersion is the version of specification format for this particular entity that this is written against.
	ApiVersion string `json:"apiVersion" yaml:"apiVersion"`

	// Kind is the high level entity type being described.
	Kind string `json:"kind" yaml:"kind"`

	// Metadata is metadata related to the entity. Should always be "System".
	Metadata EntityMeta `json:"metadata" yaml:"metadata"`

	// Spec is the specification data describing the entity itself.
	Spec map[string]interface{} `json:"spec,omitempty" yaml:"spec,omitempty"`

	// Relations that this entity has with other entities.
	Relations []EntityRelation `json:"relations,omitempty" yaml:"relations,omitempty"`

	// The current status of the entity, as claimed by various sources.
	Status *EntityStatus `json:"status,omitempty" yaml:"status,omitempty"`
}

// EntityMeta represents metadata fields common to all versions/kinds of entity.
// https://github.com/backstage/backstage/blob/master/packages/catalog-model/src/schema/EntityMeta.schema.json
type EntityMeta struct {
	// UID A globally unique ID for the entity. This field can not be set by the user at creation time, and the server will reject
	// an attempt to do so. The field will be populated in read operations.
	UID string `json:"uid,omitempty" yaml:"uid,omitempty"`

	// Etag is an opaque string that changes for each update operation to any part of the entity, including metadata. This field
	// can not be set by the user at creation time, and the server will reject an attempt to do so. The field will be populated in read
	// operations.The field can (optionally) be specified when performing update or delete operations, and the server will then reject
	// the operation if it does not match the current stored value.
	Etag string `json:"etag,omitempty" yaml:"etag,omitempty"`

	// Name of the entity. Must be unique within the catalog at any given point in time, for any given namespace + kind pair.
	Name string `json:"name" yaml:"name"`

	// Namespace that the entity belongs to.
	Namespace string `json:"namespace,omitempty" yaml:"namespace,omitempty"`

	// Title is a display name of the entity, to be presented in user interfaces instead of the name property, when available.
	Title string `json:"title,omitempty" yaml:"title,omitempty"`

	// Description is a short (typically relatively few words, on one line) description of the entity.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Labels are key/value pairs of identifying information attached to the entity.
	Labels map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`

	// Annotations are key/value pairs of non-identifying auxiliary information attached to the entity.
	Annotations map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`

	// Tags is a list of single-valued strings, to for example classify catalog entities in various ways.
	Tags []string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Links is a list of external hyperlinks related to the entity. Links can provide additional contextual
	// information that may be located outside of Backstage itself. For example, an admin dashboard or external CMS page.
	Links []EntityLink `json:"links,omitempty" yaml:"links,omitempty"`
}

// EntityLink represents a link to external information that is related to the entity.
// https://github.com/backstage/backstage/blob/master/packages/catalog-model/src/schema/EntityMeta.schema.json
type EntityLink struct {
	// URL in a standard uri format.
	URL string `json:"url" yaml:"url"`

	// Title is a user-friendly display name for the link.
	Title string `json:"title,omitempty" yaml:"title,omitempty"`

	// Icon is a key representing a visual icon to be displayed in the UI.
	Icon string `json:"icon,omitempty" yaml:"icon,omitempty"`

	// Type is an optional value to categorize links into specific groups.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}

// EntityRelation is a directed relation from one entity to another.
// https://github.com/backstage/backstage/blob/master/packages/catalog-model/src/schema/shared/common.schema.json
type EntityRelation struct {
	// Type of the relation.
	Type string `json:"type" yaml:"type"`

	// TargetRef is the entity ref of the target of this relation.
	TargetRef string `json:"targetRef" yaml:"targetRef"`

	// Target is the entity of the target of this relation.
	Target EntityRelationTarget `json:"target" yaml:"target"`
}

// EntityRelationTarget describes the target of an entity relation.
type EntityRelationTarget struct {
	// Name of the target entity. Must be unique within the catalog at any given point in time, for any given namespace + kind pair.
	Name string `json:"name" yaml:"name"`

	// Kind is the high level target entity type being described.
	Kind string `json:"kind" yaml:"kind"`

	// Namespace that the target entity belongs to.
	Namespace string `json:"namespace" yaml:"namespace"`
}

// EntityStatus informs current status of the entity, as claimed by various sources.
// https://github.com/backstage/backstage/blob/master/packages/catalog-model/src/schema/shared/common.schema.json
type EntityStatus struct {
	// A specific status item on a well known format.
	Items []EntityStatusItem `json:"items,omitempty" yaml:"items,omitempty"`
}

// EntityStatusItem contains a specific status item on a well known format.
// https://github.com/backstage/backstage/blob/master/packages/catalog-model/src/schema/shared/common.schema.json
type EntityStatusItem struct {
	// The item type
	Type string `json:"type" yaml:"type"`

	// The status level / severity of the status item.
	// Either ["info", "warning", "error"]
	Level string `json:"level" yaml:"level"`

	// A brief message describing the status, intended for human consumption.
	Message string `json:"message" yaml:"message"`

	// An optional serialized error object related to the status.
	Error *EntityStatusItemError `json:"error" yaml:"error"`
}

// EntityStatusItemError has aA serialized error object.
// https://github.com/backstage/backstage/blob/master/packages/catalog-model/src/schema/shared/common.schema.json
type EntityStatusItemError struct {
	// The type name of the error"
	Name string `json:"name" yaml:"name"`

	// The message of the error
	Message string `json:"message" yaml:"message"`

	// An error code associated with the error
	Code *string `json:"code" yaml:"code"`

	// An error stack trace
	Stack *string `json:"stack" yaml:"stack"`
}

// ListEntityOrder defines a condition that can be used to order entities.
type ListEntityOrder struct {
	// Direction is the direction to order by.
	Direction string

	// Field is the field to order by.
	Field string
}

// ListEntityOptions specifies the optional parameters to the catalogService.List method.
type ListEntityOptions struct {
	// Filters is a set of conditions that can be used to filter entities.
	Filters []string

	// Fields is a set of fields that can be used to limit the response.
	Fields []string

	// Order is a set of conditions that can be used to order entities.
	Order []ListEntityOrder
}

// entityConstraint defines constrains for entity types.
type entityConstraint interface {
	ApiEntityV1alpha1 | ComponentEntityV1alpha1 | DomainEntityV1alpha1 | GroupEntityV1alpha1 | LocationEntityV1alpha1 |
		ResourceEntityV1alpha1 | SystemEntityV1alpha1 | UserEntityV1alpha1
}

// entityService handles communication with the Backstage entities endpoints in Backstage Catalog API.
type entityService service

// typedEntityService handles communication with the Backstage entities endpoints in Backstage Catalog API, for a specific type of entity.
type typedEntityService[T entityConstraint] service

const (
	// OrderAscending is used to order entities in ascending order.
	OrderAscending = "asc"

	// OrderDescending is used to order entities in descending order.
	OrderDescending = "desc"
)

// newEntityService returns a new instance of entityService.
func newEntityService(s *catalogService) *entityService {
	const entitiesApiPath = "/entities"

	apiPath, _ := url.JoinPath(s.apiPath, entitiesApiPath)

	return &entityService{
		client:  s.client,
		apiPath: apiPath,
	}
}

// List returns a list of entities. It can optionally be filtered by a set of conditions and limited to a set of fields.
func (s *entityService) List(ctx context.Context, options *ListEntityOptions) ([]Entity, *http.Response, error) {
	u := url.URL{
		Path: s.apiPath,
	}

	values := u.Query()
	if options != nil {
		for _, f := range options.Filters {
			values.Add("filter", f)
		}

		if len(options.Fields) > 0 {
			values.Add("fields", strings.Join(options.Fields, ","))
		}

		if len(options.Order) > 0 {
			for _, o := range options.Order {
				if order, err := o.string(); err != nil {
					return nil, nil, err
				} else {
					values.Add("order", order)
				}
			}
		}
	}

	req, _ := s.client.newRequest(http.MethodGet, fmt.Sprintf("%s?%s", u.Path, values.Encode()), nil)

	var entities []Entity
	resp, err := s.client.do(ctx, req, &entities)

	return entities, resp, err
}

// Get returns a single entity by its UID.
func (s *entityService) Get(ctx context.Context, uid string) (*Entity, *http.Response, error) {
	path, _ := url.JoinPath(s.apiPath, "/by-uid/", uid)
	req, _ := s.client.newRequest(http.MethodGet, path, nil)

	var entity *Entity
	resp, err := s.client.do(ctx, req, &entity)

	return entity, resp, err

}

// Delete deletes an orphaned entity by its UID.
func (s *entityService) Delete(ctx context.Context, uid string) (*http.Response, error) {
	if uid == "" {
		return nil, errors.New("uid cannot be empty")
	}

	path, _ := url.JoinPath(s.apiPath, "/by-uid/", uid)
	req, _ := s.client.newRequest(http.MethodDelete, path, nil)

	return s.client.do(ctx, req, nil)
}

// get returns n specific type entity identified by the name and the namespace ("default", if not specified) it belongs to.
func (s *typedEntityService[T]) get(ctx context.Context, t string, n string, ns string) (*T, *http.Response, error) {
	if ns == "" {
		ns = s.client.DefaultNamespace
	}

	path, _ := url.JoinPath(s.apiPath, "/by-name/", strings.ToLower(t), ns, n)
	req, _ := s.client.newRequest(http.MethodGet, path, nil)

	var entity *T
	resp, err := s.client.do(ctx, req, &entity)

	return entity, resp, err
}

// string returns a string representation of the ListEntityOrder.
func (o *ListEntityOrder) string() (string, error) {
	if o.Direction != OrderAscending && o.Direction != OrderDescending {
		return "", fmt.Errorf("invalid order direction: %s", o.Direction)
	}

	return fmt.Sprintf("%s:%s", o.Direction, o.Field), nil
}
