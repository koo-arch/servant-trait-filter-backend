package repository

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"github.com/koo-arch/servant-trait-filter-backend/ent"
	"github.com/koo-arch/servant-trait-filter-backend/ent/ascension"
	"github.com/koo-arch/servant-trait-filter-backend/internal/model"
)

// AscensionRepository is the interface for the ascension repository.
type AscensionRepository interface {
	GetByServantAndStage(ctx context.Context, servantID int, stage int) (*ent.Ascension, error)
	ListAll(ctx context.Context) ([]*ent.Ascension, error)
	UpsertBulk(ctx context.Context, ascensions []model.Ascension) error
	WithTx(tx *ent.Tx) AscensionRepository
}

type AscensionRepositoryImpl struct {
	client *ent.Client
}

func NewAscensionRepository(client *ent.Client) AscensionRepository {
	return &AscensionRepositoryImpl{
		client: client,
	}
}

func (r *AscensionRepositoryImpl) WithTx(tx *ent.Tx) AscensionRepository {
	return &AscensionRepositoryImpl{
		client: tx.Client(),
	}
}

func (r AscensionRepositoryImpl) GetByServantAndStage(ctx context.Context, servantID int, stage int) (*ent.Ascension, error) {
	return r.client.Ascension.Query().
		Where(
			ascension.ServantID(servantID),
			ascension.Stage(stage),
		).
		Only(ctx)
}

func (r *AscensionRepositoryImpl) ListAll(ctx context.Context) ([]*ent.Ascension, error) {
	return r.client.Ascension.Query().
		Order(ent.Asc(ascension.FieldID)).
		All(ctx)
}

func (r *AscensionRepositoryImpl) UpsertBulk(ctx context.Context, ascensions []model.Ascension) error {
	if len(ascensions) == 0 {
		return nil
	}

	// 一度に1000件ずつ処理する
	const batchSize = 1000
	for i := 0; i < len(ascensions); i += batchSize {
		end := min(i+batchSize, len(ascensions))
		builders := make([]*ent.AscensionCreate, 0, end-i)

		for _, asc := range ascensions[i:end] {
			builder := r.client.Ascension.Create().
				SetServantID(asc.ServantID).
				SetStage(asc.Stage)
			if asc.AttributeID > 0 {
				builder.SetAttributeID(asc.AttributeID)
			}
			if asc.MoralAlignmentID > 0 {
				builder.SetMoralAlignmentID(asc.MoralAlignmentID)
			}
			if asc.OrderAlignmentID > 0 {
				builder.SetOrderAlignmentID(asc.OrderAlignmentID)
			}
			builders = append(builders, builder)
		}
		if len(builders) == 0 {
			continue
		}

		err := r.client.Ascension.CreateBulk(builders...).
			OnConflict(
				sql.ConflictColumns(ascension.FieldServantID, ascension.FieldStage),
			).
			UpdateNewValues().
			Exec(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}