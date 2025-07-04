package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/edge"

	"github.com/koo-arch/servant-trait-filter-backend/ent/mixin"
)

// MoralAlignment holds the schema definition for the MoralAlignment entity.
type MoralAlignment struct {
	ent.Schema
}

// Fields of the MoralAlignment.
func (MoralAlignment) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Positive().Immutable(),
	}
}

// Edges of the MoralAlignment.
func (MoralAlignment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("ascensions", Ascension.Type).
			Annotations(entsql.OnDelete(entsql.Cascade)),
	}
}

func (MoralAlignment) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		mixin.NameMixin{},
	}
}
