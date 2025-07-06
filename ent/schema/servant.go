package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/edge"

	"github.com/koo-arch/servant-trait-filter-backend/ent/mixin"
)

// Servant holds the schema definition for the Servant entity.
type Servant struct {
	ent.Schema
}

// Fields of the Servant.
func (Servant) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Positive().Immutable(),
		field.Int("collection_no").Positive().Unique(),
		field.String("name").NotEmpty(),
		field.String("face").NotEmpty(),
		field.Int("class_id").Positive(),
	}
}

// Edges of the Servant.
func (Servant) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("class", Class.Type).
			Ref("servants").
			Field("class_id").
			Unique().
			Required(),
		edge.From("traits", Trait.Type).Ref("servants"),
		edge.To("ascensions", Ascension.Type).
			Annotations(entsql.OnDelete(entsql.Cascade)),
	}
}

func (Servant) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}