package repository

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/koo-arch/servant-trait-filter-backend/ent"
	"github.com/koo-arch/servant-trait-filter-backend/ent/moralalignment"
)

type MoralAlignmentRepository interface {

}

type MoralAlignmentRepositoryImpl struct {
	client *ent.Client
}

func NewMoralAlignmentRepository(client *ent.Client) MoralAlignmentRepository {
	return &MoralAlignmentRepositoryImpl{
		client: client,
	}
}

func (r *MoralAlignmentRepositoryImpl) ListAll(ctx context.Context) ([]*ent.MoralAlignment, error) {
	return r.client.MoralAlignment.Query().
		Order(ent.Asc(moralalignment.FieldID)).
		All(ctx)
}

func (r *MoralAlignmentRepositoryImpl) UpsertBulk(ctx context.Context, moralAlignments []*ent.MoralAlignment) error {
	tx, err := r.client.Tx(ctx)
	if err != nil {
		return err
	}

	if len(moralAlignments) == 0 {
		return nil
	}

	builders := make([]*ent.MoralAlignmentCreate, 0, len(moralAlignments))
	for _, ma := range moralAlignments {
		builder := tx.MoralAlignment.Create().
			SetID(ma.ID).
			SetNameEn(ma.NameEn)
		builders = append(builders, builder)
	}

	err = tx.MoralAlignment.CreateBulk(builders...).
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
	return tx.Commit()
}


	

