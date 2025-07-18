package repository

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/koo-arch/servant-trait-filter-backend/ent"
	"github.com/koo-arch/servant-trait-filter-backend/ent/moralalignment"
	"github.com/koo-arch/servant-trait-filter-backend/internal/model"
)

type MoralAlignmentRepository interface {
	ListAll(ctx context.Context) ([]*ent.MoralAlignment, error)
	UpsertBulk(ctx context.Context, moralAlignments []model.MoralAlignment) error
	WithTx(tx *ent.Tx) MoralAlignmentRepository
}

type MoralAlignmentRepositoryImpl struct {
	client *ent.Client
}

func NewMoralAlignmentRepository(client *ent.Client) MoralAlignmentRepository {
	return &MoralAlignmentRepositoryImpl{
		client: client,
	}
}

func (r *MoralAlignmentRepositoryImpl) WithTx(tx *ent.Tx) MoralAlignmentRepository {
	return &MoralAlignmentRepositoryImpl{
		client: tx.Client(),
	}
}

func (r *MoralAlignmentRepositoryImpl) ListAll(ctx context.Context) ([]*ent.MoralAlignment, error) {
	return r.client.MoralAlignment.Query().
		Order(ent.Asc(moralalignment.FieldID)).
		All(ctx)
}

func (r *MoralAlignmentRepositoryImpl) UpsertBulk(ctx context.Context, moralAlignments []model.MoralAlignment) error {
	if len(moralAlignments) == 0 {
		return nil
	}

	builders := make([]*ent.MoralAlignmentCreate, 0, len(moralAlignments))
	for _, ma := range moralAlignments {
		builder := r.client.MoralAlignment.Create().
			SetID(ma.ID).
			SetNameEn(ma.Name)
		builders = append(builders, builder)
	}

	err := r.client.MoralAlignment.CreateBulk(builders...).
		OnConflict(
			sql.ConflictColumns(moralalignment.FieldID),
			sql.UpdateWhere(
				sql.ExprP("moral_alignments.name_en IS DISTINCT FROM EXCLUDED.name_en"),
			),
		).
		UpdateNewValues().
		Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}


	

