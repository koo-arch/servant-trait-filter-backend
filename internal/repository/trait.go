package repository

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"github.com/koo-arch/servant-trait-filter-backend/ent"
	"github.com/koo-arch/servant-trait-filter-backend/ent/trait"
	"github.com/koo-arch/servant-trait-filter-backend/internal/model"
)

// Repository is the interface for the repository.
type TraitRepository interface {
	ListAll(ctx context.Context) ([]*ent.Trait, error)
	UpsertBulk(ctx context.Context, traits []model.Trait) error
	WithTx(tx *ent.Tx) TraitRepository
}

type TraitRepositoryImpl struct {
	client *ent.Client
}

func NewTraitRepository(client *ent.Client) TraitRepository {
	return &TraitRepositoryImpl{
		client: client,
	}
}

func (r *TraitRepositoryImpl) WithTx(tx *ent.Tx) TraitRepository {
	return &TraitRepositoryImpl{
		client: tx.Client(),
	}
}

func (r *TraitRepositoryImpl) ListAll(ctx context.Context) ([]*ent.Trait, error) {
	return r.client.Trait.Query().
		Order(ent.Asc(trait.FieldID)).
		All(ctx)
}


func (r *TraitRepositoryImpl) UpsertBulk(ctx context.Context, traits []model.Trait) error {
	if len(traits) == 0 {
		return nil
	}

	// 一度に1000件ずつ処理する
	const batchSize = 1000
	for i := 0; i < len(traits); i += batchSize {
		end := min(i+batchSize, len(traits))
		builders := make([]*ent.TraitCreate, 0, end-i)

		for _, trt := range traits[i:end] {
			builder := r.client.Trait.Create().
				SetID(trt.ID).
				SetNameEn(trt.Name)
			builders = append(builders, builder)
		}
		if len(builders) == 0 {
			continue
		}

		// 一括で作成
		err := r.client.Trait.CreateBulk(builders...).
			OnConflict(
				sql.ConflictColumns(trait.FieldID),
				sql.UpdateWhere(
					sql.ExprP(
						"traits.name_en IS DISTINCT FROM EXCLUDED.name_en",
					),
				),
			).
			UpdateNewValues().
			Exec(ctx);
		if err != nil {
			return err
		}
	}

	return nil
}