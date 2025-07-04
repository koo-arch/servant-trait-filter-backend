package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/index"

	"github.com/koo-arch/servant-trait-filter-backend/ent/mixin"
)

// Ascension holds the schema definition for the Ascension entity.
type Ascension struct {
	ent.Schema
}

// Fields of the Ascension.
func (Ascension) Fields() []ent.Field {
	return []ent.Field{
		field.Int("servant_id"),
		field.Int("stage").Positive(),
		field.Int("attribute_id").Positive().Optional(),
		field.Int("order_alignment_id").Positive().Optional(),
		field.Int("moral_alignment_id").Positive().Optional(),
	}
}

// Edges of the Ascension.
func (Ascension) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("servant", Servant.Type).
			Ref("ascensions").
			Field("servant_id").
			Unique().
			Required(),
		edge.From("attribute", Attribute.Type).
			Ref("ascensions").
			Field("attribute_id").
			Unique(),
		edge.From("order_alignment", OrderAlignment.Type).
			Ref("ascensions").
			Field("order_alignment_id").
			Unique(),
		edge.From("moral_alignment", MoralAlignment.Type).
			Ref("ascensions").
			Field("moral_alignment_id").
			Unique(),
	}
}

func (Ascension) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

func (Ascension) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("servant_id", "stage").Unique(),
	}
}
