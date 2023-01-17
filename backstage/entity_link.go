package backstage

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
