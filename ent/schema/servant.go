package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"

	"github.com/koo-arch/servant-trait-filter-backend/ent/mixin"
)

// Servant holds the schema definition for the Servant entity.
type Servant struct {
	ent.Schema
}

// Edges of the Servant.
func (Servant) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("class", Class.Type).Ref("servants").Unique(),
		edge.From("attribute", Attribute.Type).Ref("servants").Unique(),
		edge.From("order_alignment", OrderAlignment.Type).Ref("servants").Unique(),
		edge.From("moral_alignment", MoralAlignment.Type).Ref("servants").Unique(),
		edge.From("traits", Trait.Type).Ref("servants"),
	}
}

func (Servant) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		mixin.NameMixin{},
	}
}