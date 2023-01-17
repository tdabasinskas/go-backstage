package backstage

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
