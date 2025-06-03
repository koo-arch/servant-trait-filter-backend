package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"

	"github.com/koo-arch/servant-trait-filter-backend/ent/mixin"
)

// Trait holds the schema definition for the Trait entity.
type Trait struct {
	ent.Schema
}


// Edges of the Trait.
func (Trait) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("servants", Servant.Type),
	}
}

func (Trait) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		mixin.NameMixin{},
	}
}