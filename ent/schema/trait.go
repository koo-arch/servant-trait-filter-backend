package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/edge"

	"github.com/koo-arch/servant-trait-filter-backend/ent/mixin"
)

// Trait holds the schema definition for the Trait entity.
type Trait struct {
	ent.Schema
}

// Fields of the Trait.
func (Trait) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Positive().Immutable(),
	}
}

// Edges of the Trait.
func (Trait) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("servants", Servant.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
	}
}

func (Trait) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		mixin.NameMixin{},
	}
}