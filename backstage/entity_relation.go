package backstage

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
