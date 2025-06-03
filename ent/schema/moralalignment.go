package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"

	"github.com/koo-arch/servant-trait-filter-backend/ent/mixin"
)

// MoralAlignment holds the schema definition for the MoralAlignment entity.
type MoralAlignment struct {
	ent.Schema
}

// Edges of the MoralAlignment.
func (MoralAlignment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("servants", Servant.Type),
	}
}

func (MoralAlignment) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		mixin.NameMixin{},
	}
}
