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
	"github.com/koo-arch/servant-trait-filter-backend/ent/predicate"
	"github.com/koo-arch/servant-trait-filter-backend/ent/servant"
	"github.com/koo-arch/servant-trait-filter-backend/ent/trait"
)

// TraitUpdate is the builder for updating Trait entities.
type TraitUpdate struct {
	config
	hooks    []Hook
	mutation *TraitMutation
}

// Where appends a list predicates to the TraitUpdate builder.
func (tu *TraitUpdate) Where(ps ...predicate.Trait) *TraitUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetUpdatedAt sets the "updated_at" field.
func (tu *TraitUpdate) SetUpdatedAt(t time.Time) *TraitUpdate {
	tu.mutation.SetUpdatedAt(t)
	return tu
}

// SetNameEn sets the "name_en" field.
func (tu *TraitUpdate) SetNameEn(s string) *TraitUpdate {
	tu.mutation.SetNameEn(s)
	return tu
}

// SetNillableNameEn sets the "name_en" field if the given value is not nil.
func (tu *TraitUpdate) SetNillableNameEn(s *string) *TraitUpdate {
	if s != nil {
		tu.SetNameEn(*s)
	}
	return tu
}

// SetNameJa sets the "name_ja" field.
func (tu *TraitUpdate) SetNameJa(s string) *TraitUpdate {
	tu.mutation.SetNameJa(s)
	return tu
}

// SetNillableNameJa sets the "name_ja" field if the given value is not nil.
func (tu *TraitUpdate) SetNillableNameJa(s *string) *TraitUpdate {
	if s != nil {
		tu.SetNameJa(*s)
	}
	return tu
}

// ClearNameJa clears the value of the "name_ja" field.
func (tu *TraitUpdate) ClearNameJa() *TraitUpdate {
	tu.mutation.ClearNameJa()
	return tu
}

// AddServantIDs adds the "servants" edge to the Servant entity by IDs.
func (tu *TraitUpdate) AddServantIDs(ids ...int) *TraitUpdate {
	tu.mutation.AddServantIDs(ids...)
	return tu
}

// AddServants adds the "servants" edges to the Servant entity.
func (tu *TraitUpdate) AddServants(s ...*Servant) *TraitUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return tu.AddServantIDs(ids...)
}

// Mutation returns the TraitMutation object of the builder.
func (tu *TraitUpdate) Mutation() *TraitMutation {
	return tu.mutation
}

// ClearServants clears all "servants" edges to the Servant entity.
func (tu *TraitUpdate) ClearServants() *TraitUpdate {
	tu.mutation.ClearServants()
	return tu
}

// RemoveServantIDs removes the "servants" edge to Servant entities by IDs.
func (tu *TraitUpdate) RemoveServantIDs(ids ...int) *TraitUpdate {
	tu.mutation.RemoveServantIDs(ids...)
	return tu
}

// RemoveServants removes "servants" edges to Servant entities.
func (tu *TraitUpdate) RemoveServants(s ...*Servant) *TraitUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return tu.RemoveServantIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TraitUpdate) Save(ctx context.Context) (int, error) {
	tu.defaults()
	return withHooks(ctx, tu.sqlSave, tu.mutation, tu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TraitUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TraitUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TraitUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tu *TraitUpdate) defaults() {
	if _, ok := tu.mutation.UpdatedAt(); !ok {
		v := trait.UpdateDefaultUpdatedAt()
		tu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tu *TraitUpdate) check() error {
	if v, ok := tu.mutation.NameEn(); ok {
		if err := trait.NameEnValidator(v); err != nil {
			return &ValidationError{Name: "name_en", err: fmt.Errorf(`ent: validator failed for field "Trait.name_en": %w`, err)}
		}
	}
	return nil
}

func (tu *TraitUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := tu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(trait.Table, trait.Columns, sqlgraph.NewFieldSpec(trait.FieldID, field.TypeInt))
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.UpdatedAt(); ok {
		_spec.SetField(trait.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := tu.mutation.NameEn(); ok {
		_spec.SetField(trait.FieldNameEn, field.TypeString, value)
	}
	if value, ok := tu.mutation.NameJa(); ok {
		_spec.SetField(trait.FieldNameJa, field.TypeString, value)
	}
	if tu.mutation.NameJaCleared() {
		_spec.ClearField(trait.FieldNameJa, field.TypeString)
	}
	if tu.mutation.ServantsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   trait.ServantsTable,
			Columns: trait.ServantsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(servant.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedServantsIDs(); len(nodes) > 0 && !tu.mutation.ServantsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   trait.ServantsTable,
			Columns: trait.ServantsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(servant.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.ServantsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   trait.ServantsTable,
			Columns: trait.ServantsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(servant.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{trait.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tu.mutation.done = true
	return n, nil
}

// TraitUpdateOne is the builder for updating a single Trait entity.
type TraitUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TraitMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (tuo *TraitUpdateOne) SetUpdatedAt(t time.Time) *TraitUpdateOne {
	tuo.mutation.SetUpdatedAt(t)
	return tuo
}

// SetNameEn sets the "name_en" field.
func (tuo *TraitUpdateOne) SetNameEn(s string) *TraitUpdateOne {
	tuo.mutation.SetNameEn(s)
	return tuo
}

// SetNillableNameEn sets the "name_en" field if the given value is not nil.
func (tuo *TraitUpdateOne) SetNillableNameEn(s *string) *TraitUpdateOne {
	if s != nil {
		tuo.SetNameEn(*s)
	}
	return tuo
}

// SetNameJa sets the "name_ja" field.
func (tuo *TraitUpdateOne) SetNameJa(s string) *TraitUpdateOne {
	tuo.mutation.SetNameJa(s)
	return tuo
}

// SetNillableNameJa sets the "name_ja" field if the given value is not nil.
func (tuo *TraitUpdateOne) SetNillableNameJa(s *string) *TraitUpdateOne {
	if s != nil {
		tuo.SetNameJa(*s)
	}
	return tuo
}

// ClearNameJa clears the value of the "name_ja" field.
func (tuo *TraitUpdateOne) ClearNameJa() *TraitUpdateOne {
	tuo.mutation.ClearNameJa()
	return tuo
}

// AddServantIDs adds the "servants" edge to the Servant entity by IDs.
func (tuo *TraitUpdateOne) AddServantIDs(ids ...int) *TraitUpdateOne {
	tuo.mutation.AddServantIDs(ids...)
	return tuo
}

// AddServants adds the "servants" edges to the Servant entity.
func (tuo *TraitUpdateOne) AddServants(s ...*Servant) *TraitUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return tuo.AddServantIDs(ids...)
}

// Mutation returns the TraitMutation object of the builder.
func (tuo *TraitUpdateOne) Mutation() *TraitMutation {
	return tuo.mutation
}

// ClearServants clears all "servants" edges to the Servant entity.
func (tuo *TraitUpdateOne) ClearServants() *TraitUpdateOne {
	tuo.mutation.ClearServants()
	return tuo
}

// RemoveServantIDs removes the "servants" edge to Servant entities by IDs.
func (tuo *TraitUpdateOne) RemoveServantIDs(ids ...int) *TraitUpdateOne {
	tuo.mutation.RemoveServantIDs(ids...)
	return tuo
}

// RemoveServants removes "servants" edges to Servant entities.
func (tuo *TraitUpdateOne) RemoveServants(s ...*Servant) *TraitUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return tuo.RemoveServantIDs(ids...)
}

// Where appends a list predicates to the TraitUpdate builder.
func (tuo *TraitUpdateOne) Where(ps ...predicate.Trait) *TraitUpdateOne {
	tuo.mutation.Where(ps...)
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TraitUpdateOne) Select(field string, fields ...string) *TraitUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Trait entity.
func (tuo *TraitUpdateOne) Save(ctx context.Context) (*Trait, error) {
	tuo.defaults()
	return withHooks(ctx, tuo.sqlSave, tuo.mutation, tuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TraitUpdateOne) SaveX(ctx context.Context) *Trait {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TraitUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TraitUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tuo *TraitUpdateOne) defaults() {
	if _, ok := tuo.mutation.UpdatedAt(); !ok {
		v := trait.UpdateDefaultUpdatedAt()
		tuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tuo *TraitUpdateOne) check() error {
	if v, ok := tuo.mutation.NameEn(); ok {
		if err := trait.NameEnValidator(v); err != nil {
			return &ValidationError{Name: "name_en", err: fmt.Errorf(`ent: validator failed for field "Trait.name_en": %w`, err)}
		}
	}
	return nil
}

func (tuo *TraitUpdateOne) sqlSave(ctx context.Context) (_node *Trait, err error) {
	if err := tuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(trait.Table, trait.Columns, sqlgraph.NewFieldSpec(trait.FieldID, field.TypeInt))
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Trait.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, trait.FieldID)
		for _, f := range fields {
			if !trait.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != trait.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.UpdatedAt(); ok {
		_spec.SetField(trait.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := tuo.mutation.NameEn(); ok {
		_spec.SetField(trait.FieldNameEn, field.TypeString, value)
	}
	if value, ok := tuo.mutation.NameJa(); ok {
		_spec.SetField(trait.FieldNameJa, field.TypeString, value)
	}
	if tuo.mutation.NameJaCleared() {
		_spec.ClearField(trait.FieldNameJa, field.TypeString)
	}
	if tuo.mutation.ServantsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   trait.ServantsTable,
			Columns: trait.ServantsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(servant.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedServantsIDs(); len(nodes) > 0 && !tuo.mutation.ServantsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   trait.ServantsTable,
			Columns: trait.ServantsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(servant.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.ServantsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   trait.ServantsTable,
			Columns: trait.ServantsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(servant.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Trait{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{trait.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tuo.mutation.done = true
	return _node, nil
}
