package backstage

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// Entity represents the parts of the format that's common to all versions/kinds of entity.
// https://github.com/backstage/backstage/blob/master/packages/catalog-model/src/schema/Entity.schema.json
type Entity struct {
	// ApiVersion is the version of specification format for this particular entity that this is written against.
	ApiVersion string `json:"apiVersion"`

	// Kind is the high level entity type being described.
	Kind string `json:"kind"`

	// Metadata is metadata related to the entity. Should always be "System".
	Metadata EntityMeta `json:"metadata"`

	// Spec is the specification data describing the entity itself.
	Spec map[string]interface{} `json:"spec,omitempty"`

	// Relations that this entity has with other entities.
	Relations []EntityRelation `json:"relations,omitempty"`
}

// EntityMeta represents metadata fields common to all versions/kinds of entity.
// https://github.com/backstage/backstage/blob/master/packages/catalog-model/src/schema/EntityMeta.schema.json
type EntityMeta struct {
	// UID A globally unique ID for the entity. This field can not be set by the user at creation time, and the server will reject
	// an attempt to do so. The field will be populated in read operations.
	UID string `json:"uid,omitempty"`

	// Etag is an opaque string that changes for each update operation to any part of the entity, including metadata. This field
	// can not be set by the user at creation time, and the server will reject an attempt to do so. The field will be populated in read
	// operations.The field can (optionally) be specified when performing update or delete operations, and the server will then reject
	// the operation if it does not match the current stored value.
	Etag string `json:"etag,omitempty"`

	// Name of the entity. Must be unique within the catalog at any given point in time, for any given namespace + kind pair.
	Name string `json:"name"`

	// Namespace that the entity belongs to.
	Namespace string `json:"namespace,omitempty"`

	// Title is a display name of the entity, to be presented in user interfaces instead of the name property, when available.
	Title string `json:"title,omitempty"`

	// Description is a short (typically relatively few words, on one line) description of the entity.
	Description string `json:"description,omitempty"`

	// Labels are key/value pairs of identifying information attached to the entity.
	Labels map[string]string `json:"labels,omitempty"`

	// Annotations are key/value pairs of non-identifying auxiliary information attached to the entity.
	Annotations map[string]string `json:"annotations,omitempty"`

	// Tags is a list of single-valued strings, to for example classify catalog entities in various ways.
	Tags []string `json:"tags,omitempty"`

	// Links is a list of external hyperlinks related to the entity. Links can provide additional contextual
	// information that may be located outside of Backstage itself. For example, an admin dashboard or external CMS page.
	Links []EntityLink `json:"links,omitempty"`
}

// EntityLink represents a link to external information that is related to the entity.
// https://github.com/backstage/backstage/blob/master/packages/catalog-model/src/schema/EntityMeta.schema.json
type EntityLink struct {
	// URL in a standard uri format.
	URL string `json:"url"`

	// Title is a user-friendly display name for the link.
	Title string `json:"title,omitempty"`

	// Icon is a key representing a visual icon to be displayed in the UI.
	Icon string `json:"icon,omitempty"`

	// Type is an optional value to categorize links into specific groups.
	Type string `json:"type,omitempty"`
}

// EntityRelation is a directed relation from one entity to another.
// https://github.com/backstage/backstage/blob/master/packages/catalog-model/src/schema/shared/common.schema.json
type EntityRelation struct {
	// Type of the relation.
	Type string `json:"type"`

	// TargetRef is the entity ref of the target of this relation.
	TargetRef string `json:"targetRef"`

	// Target is the entity of the target of this relation.
	Target EntityRelationTarget `json:"target"`
}

// EntityRelationTarget describes the target of an entity relation.
type EntityRelationTarget struct {
	// Name of the target entity. Must be unique within the catalog at any given point in time, for any given namespace + kind pair.
	Name string `json:"name"`

	// Kind is the high level target entity type being described.
	Kind string `json:"kind"`

	// Namespace that the target entity belongs to.
	Namespace string `json:"namespace"`
}

// ListEntityFilter defines a condition that can be used to filter entities.
type ListEntityFilter map[string]string

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
	Filters ListEntityFilter

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

	// entitiesApiPath is the path to the entities API.
	entitiesApiPath = "/entities"
)

// newEntityService returns a new instance of entityService.
func newEntityService(s *catalogService) *entityService {
	fp, _ := url.JoinPath(s.client.BaseURL.Path, s.apiPath.Path, entitiesApiPath)
	p, _ := s.apiPath.Parse(fp)

	return &entityService{
		client:  s.client,
		apiPath: p,
	}
}

// List returns a list of entities. It can optionally be filtered by a set of conditions and limited to a set of fields.
func (s *entityService) List(ctx context.Context, options *ListEntityOptions) ([]Entity, *http.Response, error) {
	path := fmt.Sprintf("%s?", s.apiPath)

	if options != nil {
		if options.Filters != nil && len(options.Filters) > 0 {
			path += fmt.Sprintf("filter=%s&", options.Filters.string())
		}

		if options.Fields != nil && len(options.Fields) > 0 {
			path += fmt.Sprintf("fields=%s&", strings.Join(options.Fields, ","))
		}

		if options.Order != nil && len(options.Order) > 0 {
			for _, o := range options.Order {
				if order, err := o.string(); err != nil {
					return nil, nil, err
				} else {
					path += fmt.Sprintf("order=%s&", order)
				}
			}
		}
	}

	req, err := s.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var entities []Entity
	resp, err := s.client.do(ctx, req, &entities)

	return entities, resp, err
}

// Get returns a single entity by its UID.
func (s *entityService) Get(ctx context.Context, uid string) (*Entity, *http.Response, error) {
	path, err := url.JoinPath(s.apiPath.Path, "/by-uid/", uid)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var entity *Entity
	resp, err := s.client.do(ctx, req, &entity)

	return entity, resp, err

}

// Delete deletes an orphaned entity by its UID.
func (s *entityService) Delete(ctx context.Context, uid string) (*http.Response, error) {
	path, err := url.JoinPath(s.apiPath.Path, "/by-uid/", uid)
	if err != nil {
		return nil, err
	}

	req, err := s.client.newRequest(http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	return s.client.do(ctx, req, nil)
}

// get returns n specific type entity identified by the name and the namespace ("default", if not specified) it belongs to.
func (s *typedEntityService[T]) get(ctx context.Context, t string, n string, ns string) (*T, *http.Response, error) {
	if ns == "" {
		ns = s.client.DefaultNamespace
	}

	path, err := url.JoinPath(s.apiPath.Path, "/by-name/", t, ns, n)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.newRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	var entity *T
	resp, err := s.client.do(ctx, req, &entity)

	return entity, resp, err
}

// string returns a string representation of the ListEntityFilter.
func (f *ListEntityFilter) string() string {
	var b bytes.Buffer

	for k, v := range *f {
		if v != "" {
			b.WriteString(fmt.Sprintf("%s=%s,", k, v))
		} else {
			b.WriteString(fmt.Sprintf("%s,", k))
		}
	}

	return strings.TrimSuffix(b.String(), ",")
}

// string returns a string representation of the ListEntityOrder.
func (o *ListEntityOrder) string() (string, error) {
	if o.Direction != OrderAscending && o.Direction != OrderDescending {
		return "", fmt.Errorf("invalid order direction: %s", o.Direction)
	}

	return fmt.Sprintf("%s:%s", o.Direction, o.Field), nil
}
