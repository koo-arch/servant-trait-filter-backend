package schema

import (
	"entgo.io/ent"
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
		field.Int("attribute_id").Positive(),
		field.Int("order_alignment_id").Positive().Optional(),
		field.Int("moral_alignment_id").Positive().Optional(),
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
		edge.From("attribute", Attribute.Type).
			Ref("servants").
			Field("attribute_id").
			Unique().
			Required(),
		edge.From("order_alignment", OrderAlignment.Type).
			Ref("servants").
			Field("order_alignment_id").
			Unique(),
		edge.From("moral_alignment", MoralAlignment.Type).
			Ref("servants").
			Field("moral_alignment_id").
			Unique(),
		edge.From("traits", Trait.Type).Ref("servants"),
	}
}

func (Servant) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}