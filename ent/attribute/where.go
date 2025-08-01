// Code generated by ent, DO NOT EDIT.

package attribute

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/koo-arch/servant-trait-filter-backend/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Attribute {
	return predicate.Attribute(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Attribute {
	return predicate.Attribute(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Attribute {
	return predicate.Attribute(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Attribute {
	return predicate.Attribute(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Attribute {
	return predicate.Attribute(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Attribute {
	return predicate.Attribute(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Attribute {
	return predicate.Attribute(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Attribute {
	return predicate.Attribute(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Attribute {
	return predicate.Attribute(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Attribute {
	return predicate.Attribute(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Attribute {
	return predicate.Attribute(sql.FieldEQ(FieldUpdatedAt, v))
}

// NameEn applies equality check predicate on the "name_en" field. It's identical to NameEnEQ.
func NameEn(v string) predicate.Attribute {
	return predicate.Attribute(sql.FieldEQ(FieldNameEn, v))
}

// NameJa applies equality check predicate on the "name_ja" field. It's identical to NameJaEQ.
func NameJa(v string) predicate.Attribute {
	return predicate.Attribute(sql.FieldEQ(FieldNameJa, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Attribute {
	return predicate.Attribute(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Attribute {
	return predicate.Attribute(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Attribute {
	return predicate.Attribute(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Attribute {
	return predicate.Attribute(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Attribute {
	return predicate.Attribute(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Attribute {
	return predicate.Attribute(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Attribute {
	return predicate.Attribute(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Attribute {
	return predicate.Attribute(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Attribute {
	return predicate.Attribute(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Attribute {
	return predicate.Attribute(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Attribute {
	return predicate.Attribute(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Attribute {
	return predicate.Attribute(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Attribute {
	return predicate.Attribute(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Attribute {
	return predicate.Attribute(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Attribute {
	return predicate.Attribute(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Attribute {
	return predicate.Attribute(sql.FieldLTE(FieldUpdatedAt, v))
}

// NameEnEQ applies the EQ predicate on the "name_en" field.
func NameEnEQ(v string) predicate.Attribute {
	return predicate.Attribute(sql.FieldEQ(FieldNameEn, v))
}

// NameEnNEQ applies the NEQ predicate on the "name_en" field.
func NameEnNEQ(v string) predicate.Attribute {
	return predicate.Attribute(sql.FieldNEQ(FieldNameEn, v))
}

// NameEnIn applies the In predicate on the "name_en" field.
func NameEnIn(vs ...string) predicate.Attribute {
	return predicate.Attribute(sql.FieldIn(FieldNameEn, vs...))
}

// NameEnNotIn applies the NotIn predicate on the "name_en" field.
func NameEnNotIn(vs ...string) predicate.Attribute {
	return predicate.Attribute(sql.FieldNotIn(FieldNameEn, vs...))
}

// NameEnGT applies the GT predicate on the "name_en" field.
func NameEnGT(v string) predicate.Attribute {
	return predicate.Attribute(sql.FieldGT(FieldNameEn, v))
}

// NameEnGTE applies the GTE predicate on the "name_en" field.
func NameEnGTE(v string) predicate.Attribute {
	return predicate.Attribute(sql.FieldGTE(FieldNameEn, v))
}

// NameEnLT applies the LT predicate on the "name_en" field.
func NameEnLT(v string) predicate.Attribute {
	return predicate.Attribute(sql.FieldLT(FieldNameEn, v))
}

// NameEnLTE applies the LTE predicate on the "name_en" field.
func NameEnLTE(v string) predicate.Attribute {
	return predicate.Attribute(sql.FieldLTE(FieldNameEn, v))
}

// NameEnContains applies the Contains predicate on the "name_en" field.
func NameEnContains(v string) predicate.Attribute {
	return predicate.Attribute(sql.FieldContains(FieldNameEn, v))
}

// NameEnHasPrefix applies the HasPrefix predicate on the "name_en" field.
func NameEnHasPrefix(v string) predicate.Attribute {
	return predicate.Attribute(sql.FieldHasPrefix(FieldNameEn, v))
}

// NameEnHasSuffix applies the HasSuffix predicate on the "name_en" field.
func NameEnHasSuffix(v string) predicate.Attribute {
	return predicate.Attribute(sql.FieldHasSuffix(FieldNameEn, v))
}

// NameEnEqualFold applies the EqualFold predicate on the "name_en" field.
func NameEnEqualFold(v string) predicate.Attribute {
	return predicate.Attribute(sql.FieldEqualFold(FieldNameEn, v))
}

// NameEnContainsFold applies the ContainsFold predicate on the "name_en" field.
func NameEnContainsFold(v string) predicate.Attribute {
	return predicate.Attribute(sql.FieldContainsFold(FieldNameEn, v))
}

// NameJaEQ applies the EQ predicate on the "name_ja" field.
func NameJaEQ(v string) predicate.Attribute {
	return predicate.Attribute(sql.FieldEQ(FieldNameJa, v))
}

// NameJaNEQ applies the NEQ predicate on the "name_ja" field.
func NameJaNEQ(v string) predicate.Attribute {
	return predicate.Attribute(sql.FieldNEQ(FieldNameJa, v))
}

// NameJaIn applies the In predicate on the "name_ja" field.
func NameJaIn(vs ...string) predicate.Attribute {
	return predicate.Attribute(sql.FieldIn(FieldNameJa, vs...))
}

// NameJaNotIn applies the NotIn predicate on the "name_ja" field.
func NameJaNotIn(vs ...string) predicate.Attribute {
	return predicate.Attribute(sql.FieldNotIn(FieldNameJa, vs...))
}

// NameJaGT applies the GT predicate on the "name_ja" field.
func NameJaGT(v string) predicate.Attribute {
	return predicate.Attribute(sql.FieldGT(FieldNameJa, v))
}

// NameJaGTE applies the GTE predicate on the "name_ja" field.
func NameJaGTE(v string) predicate.Attribute {
	return predicate.Attribute(sql.FieldGTE(FieldNameJa, v))
}

// NameJaLT applies the LT predicate on the "name_ja" field.
func NameJaLT(v string) predicate.Attribute {
	return predicate.Attribute(sql.FieldLT(FieldNameJa, v))
}

// NameJaLTE applies the LTE predicate on the "name_ja" field.
func NameJaLTE(v string) predicate.Attribute {
	return predicate.Attribute(sql.FieldLTE(FieldNameJa, v))
}

// NameJaContains applies the Contains predicate on the "name_ja" field.
func NameJaContains(v string) predicate.Attribute {
	return predicate.Attribute(sql.FieldContains(FieldNameJa, v))
}

// NameJaHasPrefix applies the HasPrefix predicate on the "name_ja" field.
func NameJaHasPrefix(v string) predicate.Attribute {
	return predicate.Attribute(sql.FieldHasPrefix(FieldNameJa, v))
}

// NameJaHasSuffix applies the HasSuffix predicate on the "name_ja" field.
func NameJaHasSuffix(v string) predicate.Attribute {
	return predicate.Attribute(sql.FieldHasSuffix(FieldNameJa, v))
}

// NameJaIsNil applies the IsNil predicate on the "name_ja" field.
func NameJaIsNil() predicate.Attribute {
	return predicate.Attribute(sql.FieldIsNull(FieldNameJa))
}

// NameJaNotNil applies the NotNil predicate on the "name_ja" field.
func NameJaNotNil() predicate.Attribute {
	return predicate.Attribute(sql.FieldNotNull(FieldNameJa))
}

// NameJaEqualFold applies the EqualFold predicate on the "name_ja" field.
func NameJaEqualFold(v string) predicate.Attribute {
	return predicate.Attribute(sql.FieldEqualFold(FieldNameJa, v))
}

// NameJaContainsFold applies the ContainsFold predicate on the "name_ja" field.
func NameJaContainsFold(v string) predicate.Attribute {
	return predicate.Attribute(sql.FieldContainsFold(FieldNameJa, v))
}

// HasAscensions applies the HasEdge predicate on the "ascensions" edge.
func HasAscensions() predicate.Attribute {
	return predicate.Attribute(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, AscensionsTable, AscensionsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAscensionsWith applies the HasEdge predicate on the "ascensions" edge with a given conditions (other predicates).
func HasAscensionsWith(preds ...predicate.Ascension) predicate.Attribute {
	return predicate.Attribute(func(s *sql.Selector) {
		step := newAscensionsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Attribute) predicate.Attribute {
	return predicate.Attribute(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Attribute) predicate.Attribute {
	return predicate.Attribute(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Attribute) predicate.Attribute {
	return predicate.Attribute(sql.NotPredicates(p))
}
