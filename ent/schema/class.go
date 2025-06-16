package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/edge"

	"github.com/koo-arch/servant-trait-filter-backend/ent/mixin"
)

// Class holds the schema definition for the Class entity.
type Class struct {
	ent.Schema
}

// Fields of the Class.
func (Class) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Positive().Immutable(),
	}
}

// Edges of the Class.
func (Class) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("servants", Servant.Type).Annotations(entsql.OnDelete(entsql.Cascade)),
	}
}

func (Class) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		mixin.NameMixin{},
	}
}