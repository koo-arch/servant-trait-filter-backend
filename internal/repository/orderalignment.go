package repository

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/koo-arch/servant-trait-filter-backend/ent"
	"github.com/koo-arch/servant-trait-filter-backend/ent/orderalignment"
	"github.com/koo-arch/servant-trait-filter-backend/internal/transaction"
)

type OrderAlignmentRepository interface {

}

type OrderAlignmentRepositoryImpl struct {
	client *ent.Client
}

func NewOrderAlignmentRepository(client *ent.Client) OrderAlignmentRepository {
	return &OrderAlignmentRepositoryImpl{
		client: client,
	}
}

func (r *OrderAlignmentRepositoryImpl) ListAll(ctx context.Context) ([]*ent.OrderAlignment, error) {
	return r.client.OrderAlignment.Query().
		Order(ent.Asc(orderalignment.FieldID)).
		All(ctx)
}

func (r *OrderAlignmentRepositoryImpl) UpsertBulk(ctx context.Context, orderAlignments []*ent.OrderAlignment) error {
	tx, err := r.client.Tx(ctx)
	if err != nil {
		return err
	}
	defer transaction.HandleTransaction(tx, &err)

	if len(orderAlignments) == 0 {
		return nil
	}
	builders := make([]*ent.OrderAlignmentCreate, 0, len(orderAlignments))
	for _, oa := range orderAlignments {
		builder := tx.OrderAlignment.Create().
			SetID(oa.ID).
			SetNameEn(oa.NameEn)
		builders = append(builders, builder)
	}
	err = tx.OrderAlignment.CreateBulk(builders...).
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

	return tx.Commit()
}