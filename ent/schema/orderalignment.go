package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/edge"

	"github.com/koo-arch/servant-trait-filter-backend/ent/mixin"
)

// OrderAlignment holds the schema definition for the OrderAlignment entity.
type OrderAlignment struct {
	ent.Schema
}

// Fields of the OrderAlignment.
func (OrderAlignment) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Positive().Immutable(),
	}
}

// Edges of the OrderAlignment.
func (OrderAlignment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("ascensions", Ascension.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
	}
}

func (OrderAlignment) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		mixin.NameMixin{},
	}
}