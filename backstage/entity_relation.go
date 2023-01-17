package backstage

// EntityRelation is a directed relation from one entity to another.
// https://github.com/backstage/backstage/blob/master/packages/catalog-model/src/schema/shared/common.schema.json
type EntityRelation struct {
	// Type of the relation.
	Type string `json:"type"`

	// TargetRef is the entity ref of the target of this relation.
	TargetRef string `json:"targetRef"`
}
