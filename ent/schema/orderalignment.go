package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"

	"github.com/koo-arch/servant-trait-filter-backend/ent/mixin"
)

// OrderAlignment holds the schema definition for the OrderAlignment entity.
type OrderAlignment struct {
	ent.Schema
}

// Edges of the OrderAlignment.
func (OrderAlignment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("servants", Servant.Type),
	}
}

func (OrderAlignment) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		mixin.NameMixin{},
	}
}