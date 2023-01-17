package backstage

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
