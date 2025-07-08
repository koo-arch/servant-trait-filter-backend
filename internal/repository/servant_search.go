package repository

import (
	"context"

	"github.com/koo-arch/servant-trait-filter-backend/ent"
	"github.com/koo-arch/servant-trait-filter-backend/ent/servant"
	"github.com/koo-arch/servant-trait-filter-backend/ent/ascension"
	"github.com/koo-arch/servant-trait-filter-backend/ent/predicate"
	"github.com/koo-arch/servant-trait-filter-backend/ent/trait"
	"github.com/koo-arch/servant-trait-filter-backend/ent/class"
)

type Expr struct {
    And  []*Expr `json:"and,omitempty"`
    Or   []*Expr `json:"or,omitempty"`
    Not  *Expr   `json:"not,omitempty"`

    // 原子条件
    TraitID         *int     `json:"trait,omitempty"`
    ClassID         *int     `json:"class,omitempty"`
    AttributeID     *int     `json:"attribute,omitempty"`
    OrderAlignID    *int     `json:"orderAlignment,omitempty"`
    MoralAlignID    *int     `json:"moralAlignment,omitempty"`
}

type ServantSearchQuery struct {
	Root Expr
	Limit int
	Offset int
}

func (r *ServantRepositoryImpl) Search(ctx context.Context, req ServantSearchQuery) ([]*ent.Servant, error) {
	query := r.client.Servant.Query().
		WithTraits().
		WithAscensions(func(q *ent.AscensionQuery) {
				q.WithAttribute().
				WithOrderAlignment().
				WithMoralAlignment()
		}).
		WithClass()
	
	if p := r.buildPredicate(&req.Root); p != nil {
		query = query.Where(p)
	}
	if req.Limit > 0 {
		query = query.Limit(req.Limit)
	}
	if req.Offset > 0 {
		query = query.Offset(req.Offset)
	}
	return query.Order(ent.Asc(servant.FieldID)).
		All(ctx)

}

func (r *ServantRepositoryImpl) buildPredicate(expr *Expr) predicate.Servant {
	if expr == nil {
		return nil // nilの場合は無視
	}
	var atoms []predicate.Servant
	if expr.TraitID != nil {
		atoms = append(atoms, servant.HasTraitsWith(
			trait.ID(*expr.TraitID),
		))
	}
	if expr.ClassID != nil {
		atoms = append(atoms, servant.HasClassWith(
			class.ID(*expr.ClassID),
		))
	}
	if expr.AttributeID != nil {
		atoms = append(atoms, servant.HasAscensionsWith(
			ascension.AttributeIDEQ(*expr.AttributeID),
		))
	}
	if expr.OrderAlignID != nil {
		atoms = append(atoms, servant.HasAscensionsWith(
			ascension.OrderAlignmentIDEQ(*expr.OrderAlignID),
		))
	}
	if expr.MoralAlignID != nil {
		atoms = append(atoms, servant.HasAscensionsWith(
			ascension.MoralAlignmentIDEQ(*expr.MoralAlignID),
		))
	}
	switch {
	case len(expr.And) > 0:
		var cs []predicate.Servant
		for _, e := range expr.And {
			if p := r.buildPredicate(e); p != nil {
				cs = append(cs, p)
			}
		}
		return servant.And(append(atoms, cs...)...)
	case len(expr.Or) > 0:
		var cs []predicate.Servant
		for _, e := range expr.Or {
			if p := r.buildPredicate(e); p != nil {
				cs = append(cs, p)
			}
		}
		if len(atoms) == 0 {
			return servant.Or(cs...)
		}
		// 原子がある場合は (AND(atoms…) AND OR(cs…)) で括弧を明示
		return servant.And(
			servant.And(atoms...),
			servant.Or(cs...),
		)
	case expr.Not != nil:
		if p := r.buildPredicate(expr.Not); p != nil {
			if len(atoms) == 0 {
				return servant.Not(p)
			}
			// 原子もある場合は (AND(atoms…) AND NOT(p))
			return servant.And(append(atoms, servant.Not(p))...)
		}
	default:
		if len(atoms) == 0 {
			return nil 
		}
		return servant.And(atoms...)
	}
	if len(atoms) == 0 {
		return nil // 原子条件がない場合は無視
	}
	return servant.And(atoms...)
}