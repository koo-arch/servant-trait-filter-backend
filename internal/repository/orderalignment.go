package repository

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/koo-arch/servant-trait-filter-backend/ent"
	"github.com/koo-arch/servant-trait-filter-backend/ent/orderalignment"
	"github.com/koo-arch/servant-trait-filter-backend/internal/model"
)

type OrderAlignmentRepository interface {
	ListAll(ctx context.Context) ([]*ent.OrderAlignment, error)
	UpsertBulk(ctx context.Context, orderAlignments []model.OrderAlignment) error
	WithTx(tx *ent.Tx) OrderAlignmentRepository
}

type OrderAlignmentRepositoryImpl struct {
	client *ent.Client
}

func NewOrderAlignmentRepository(client *ent.Client) OrderAlignmentRepository {
	return &OrderAlignmentRepositoryImpl{
		client: client,
	}
}

func (r *OrderAlignmentRepositoryImpl) WithTx(tx *ent.Tx) OrderAlignmentRepository {
	return &OrderAlignmentRepositoryImpl{
		client: tx.Client(),
	}
}

func (r *OrderAlignmentRepositoryImpl) ListAll(ctx context.Context) ([]*ent.OrderAlignment, error) {
	return r.client.OrderAlignment.Query().
		Order(ent.Asc(orderalignment.FieldID)).
		All(ctx)
}

func (r *OrderAlignmentRepositoryImpl) UpsertBulk(ctx context.Context, orderAlignments []model.OrderAlignment) error {
	if len(orderAlignments) == 0 {
		return nil
	}

	builders := make([]*ent.OrderAlignmentCreate, 0, len(orderAlignments))
	for _, oa := range orderAlignments {
		builder := r.client.OrderAlignment.Create().
			SetID(oa.ID).
			SetNameEn(oa.Name)
		builders = append(builders, builder)
	}
	err := r.client.OrderAlignment.CreateBulk(builders...).
		OnConflict(
			sql.ConflictColumns(orderalignment.FieldID),
			sql.UpdateWhere(
				sql.ExprP("order_alignments.name_en IS DISTINCT FROM EXCLUDED.name_en"),
			),
		).
		UpdateNewValues().
		Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}