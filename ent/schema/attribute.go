package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"

	"github.com/koo-arch/servant-trait-filter-backend/ent/mixin"

)

// Attribute holds the schema definition for the Attribute entity.
type Attribute struct {
	ent.Schema
}

// Edges of the Attribute.
func (Attribute) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("servants", Servant.Type),
	}
}

func (Attribute) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		mixin.NameMixin{},
	}
}
