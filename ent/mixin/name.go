package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// NameMixin defines a mixin for entities that require a name field.
type NameMixin struct {
	mixin.Schema
}

// Fields of the NameMixin.
func (NameMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("name_en").NotEmpty(),
		field.String("name_ja").Optional(),
	}
}