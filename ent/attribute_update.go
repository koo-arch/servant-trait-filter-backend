// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/koo-arch/servant-trait-filter-backend/ent/ascension"
	"github.com/koo-arch/servant-trait-filter-backend/ent/attribute"
	"github.com/koo-arch/servant-trait-filter-backend/ent/predicate"
)

// AttributeUpdate is the builder for updating Attribute entities.
type AttributeUpdate struct {
	config
	hooks    []Hook
	mutation *AttributeMutation
}

// Where appends a list predicates to the AttributeUpdate builder.
func (au *AttributeUpdate) Where(ps ...predicate.Attribute) *AttributeUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetUpdatedAt sets the "updated_at" field.
func (au *AttributeUpdate) SetUpdatedAt(t time.Time) *AttributeUpdate {
	au.mutation.SetUpdatedAt(t)
	return au
}

// SetNameEn sets the "name_en" field.
func (au *AttributeUpdate) SetNameEn(s string) *AttributeUpdate {
	au.mutation.SetNameEn(s)
	return au
}

// SetNillableNameEn sets the "name_en" field if the given value is not nil.
func (au *AttributeUpdate) SetNillableNameEn(s *string) *AttributeUpdate {
	if s != nil {
		au.SetNameEn(*s)
	}
	return au
}

// SetNameJa sets the "name_ja" field.
func (au *AttributeUpdate) SetNameJa(s string) *AttributeUpdate {
	au.mutation.SetNameJa(s)
	return au
}

// SetNillableNameJa sets the "name_ja" field if the given value is not nil.
func (au *AttributeUpdate) SetNillableNameJa(s *string) *AttributeUpdate {
	if s != nil {
		au.SetNameJa(*s)
	}
	return au
}

// ClearNameJa clears the value of the "name_ja" field.
func (au *AttributeUpdate) ClearNameJa() *AttributeUpdate {
	au.mutation.ClearNameJa()
	return au
}

// AddAscensionIDs adds the "ascensions" edge to the Ascension entity by IDs.
func (au *AttributeUpdate) AddAscensionIDs(ids ...int) *AttributeUpdate {
	au.mutation.AddAscensionIDs(ids...)
	return au
}

// AddAscensions adds the "ascensions" edges to the Ascension entity.
func (au *AttributeUpdate) AddAscensions(a ...*Ascension) *AttributeUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return au.AddAscensionIDs(ids...)
}

// Mutation returns the AttributeMutation object of the builder.
func (au *AttributeUpdate) Mutation() *AttributeMutation {
	return au.mutation
}

// ClearAscensions clears all "ascensions" edges to the Ascension entity.
func (au *AttributeUpdate) ClearAscensions() *AttributeUpdate {
	au.mutation.ClearAscensions()
	return au
}

// RemoveAscensionIDs removes the "ascensions" edge to Ascension entities by IDs.
func (au *AttributeUpdate) RemoveAscensionIDs(ids ...int) *AttributeUpdate {
	au.mutation.RemoveAscensionIDs(ids...)
	return au
}

// RemoveAscensions removes "ascensions" edges to Ascension entities.
func (au *AttributeUpdate) RemoveAscensions(a ...*Ascension) *AttributeUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return au.RemoveAscensionIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AttributeUpdate) Save(ctx context.Context) (int, error) {
	au.defaults()
	return withHooks(ctx, au.sqlSave, au.mutation, au.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (au *AttributeUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AttributeUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AttributeUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (au *AttributeUpdate) defaults() {
	if _, ok := au.mutation.UpdatedAt(); !ok {
		v := attribute.UpdateDefaultUpdatedAt()
		au.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (au *AttributeUpdate) check() error {
	if v, ok := au.mutation.NameEn(); ok {
		if err := attribute.NameEnValidator(v); err != nil {
			return &ValidationError{Name: "name_en", err: fmt.Errorf(`ent: validator failed for field "Attribute.name_en": %w`, err)}
		}
	}
	return nil
}

func (au *AttributeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := au.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(attribute.Table, attribute.Columns, sqlgraph.NewFieldSpec(attribute.FieldID, field.TypeInt))
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.UpdatedAt(); ok {
		_spec.SetField(attribute.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := au.mutation.NameEn(); ok {
		_spec.SetField(attribute.FieldNameEn, field.TypeString, value)
	}
	if value, ok := au.mutation.NameJa(); ok {
		_spec.SetField(attribute.FieldNameJa, field.TypeString, value)
	}
	if au.mutation.NameJaCleared() {
		_spec.ClearField(attribute.FieldNameJa, field.TypeString)
	}
	if au.mutation.AscensionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   attribute.AscensionsTable,
			Columns: []string{attribute.AscensionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ascension.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.RemovedAscensionsIDs(); len(nodes) > 0 && !au.mutation.AscensionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   attribute.AscensionsTable,
			Columns: []string{attribute.AscensionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ascension.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.AscensionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   attribute.AscensionsTable,
			Columns: []string{attribute.AscensionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ascension.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{attribute.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	au.mutation.done = true
	return n, nil
}

// AttributeUpdateOne is the builder for updating a single Attribute entity.
type AttributeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AttributeMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (auo *AttributeUpdateOne) SetUpdatedAt(t time.Time) *AttributeUpdateOne {
	auo.mutation.SetUpdatedAt(t)
	return auo
}

// SetNameEn sets the "name_en" field.
func (auo *AttributeUpdateOne) SetNameEn(s string) *AttributeUpdateOne {
	auo.mutation.SetNameEn(s)
	return auo
}

// SetNillableNameEn sets the "name_en" field if the given value is not nil.
func (auo *AttributeUpdateOne) SetNillableNameEn(s *string) *AttributeUpdateOne {
	if s != nil {
		auo.SetNameEn(*s)
	}
	return auo
}

// SetNameJa sets the "name_ja" field.
func (auo *AttributeUpdateOne) SetNameJa(s string) *AttributeUpdateOne {
	auo.mutation.SetNameJa(s)
	return auo
}

// SetNillableNameJa sets the "name_ja" field if the given value is not nil.
func (auo *AttributeUpdateOne) SetNillableNameJa(s *string) *AttributeUpdateOne {
	if s != nil {
		auo.SetNameJa(*s)
	}
	return auo
}

// ClearNameJa clears the value of the "name_ja" field.
func (auo *AttributeUpdateOne) ClearNameJa() *AttributeUpdateOne {
	auo.mutation.ClearNameJa()
	return auo
}

// AddAscensionIDs adds the "ascensions" edge to the Ascension entity by IDs.
func (auo *AttributeUpdateOne) AddAscensionIDs(ids ...int) *AttributeUpdateOne {
	auo.mutation.AddAscensionIDs(ids...)
	return auo
}

// AddAscensions adds the "ascensions" edges to the Ascension entity.
func (auo *AttributeUpdateOne) AddAscensions(a ...*Ascension) *AttributeUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return auo.AddAscensionIDs(ids...)
}

// Mutation returns the AttributeMutation object of the builder.
func (auo *AttributeUpdateOne) Mutation() *AttributeMutation {
	return auo.mutation
}

// ClearAscensions clears all "ascensions" edges to the Ascension entity.
func (auo *AttributeUpdateOne) ClearAscensions() *AttributeUpdateOne {
	auo.mutation.ClearAscensions()
	return auo
}

// RemoveAscensionIDs removes the "ascensions" edge to Ascension entities by IDs.
func (auo *AttributeUpdateOne) RemoveAscensionIDs(ids ...int) *AttributeUpdateOne {
	auo.mutation.RemoveAscensionIDs(ids...)
	return auo
}

// RemoveAscensions removes "ascensions" edges to Ascension entities.
func (auo *AttributeUpdateOne) RemoveAscensions(a ...*Ascension) *AttributeUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return auo.RemoveAscensionIDs(ids...)
}

// Where appends a list predicates to the AttributeUpdate builder.
func (auo *AttributeUpdateOne) Where(ps ...predicate.Attribute) *AttributeUpdateOne {
	auo.mutation.Where(ps...)
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AttributeUpdateOne) Select(field string, fields ...string) *AttributeUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Attribute entity.
func (auo *AttributeUpdateOne) Save(ctx context.Context) (*Attribute, error) {
	auo.defaults()
	return withHooks(ctx, auo.sqlSave, auo.mutation, auo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AttributeUpdateOne) SaveX(ctx context.Context) *Attribute {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AttributeUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AttributeUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (auo *AttributeUpdateOne) defaults() {
	if _, ok := auo.mutation.UpdatedAt(); !ok {
		v := attribute.UpdateDefaultUpdatedAt()
		auo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (auo *AttributeUpdateOne) check() error {
	if v, ok := auo.mutation.NameEn(); ok {
		if err := attribute.NameEnValidator(v); err != nil {
			return &ValidationError{Name: "name_en", err: fmt.Errorf(`ent: validator failed for field "Attribute.name_en": %w`, err)}
		}
	}
	return nil
}

func (auo *AttributeUpdateOne) sqlSave(ctx context.Context) (_node *Attribute, err error) {
	if err := auo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(attribute.Table, attribute.Columns, sqlgraph.NewFieldSpec(attribute.FieldID, field.TypeInt))
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Attribute.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, attribute.FieldID)
		for _, f := range fields {
			if !attribute.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != attribute.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.UpdatedAt(); ok {
		_spec.SetField(attribute.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := auo.mutation.NameEn(); ok {
		_spec.SetField(attribute.FieldNameEn, field.TypeString, value)
	}
	if value, ok := auo.mutation.NameJa(); ok {
		_spec.SetField(attribute.FieldNameJa, field.TypeString, value)
	}
	if auo.mutation.NameJaCleared() {
		_spec.ClearField(attribute.FieldNameJa, field.TypeString)
	}
	if auo.mutation.AscensionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   attribute.AscensionsTable,
			Columns: []string{attribute.AscensionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ascension.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.RemovedAscensionsIDs(); len(nodes) > 0 && !auo.mutation.AscensionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   attribute.AscensionsTable,
			Columns: []string{attribute.AscensionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ascension.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.AscensionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   attribute.AscensionsTable,
			Columns: []string{attribute.AscensionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ascension.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Attribute{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{attribute.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	auo.mutation.done = true
	return _node, nil
}
