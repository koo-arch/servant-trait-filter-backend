package repository

import (
	"context"
	"entgo.io/ent/dialect/sql"
	"github.com/koo-arch/servant-trait-filter-backend/ent"
	"github.com/koo-arch/servant-trait-filter-backend/ent/trait"
)

// Repository is the interface for the repository.
type TraitRepository interface {

}

type TraitRepositoryImpl struct {
	client *ent.Client
}

func NewTraitRepository(client *ent.Client) TraitRepository {
	return &TraitRepositoryImpl{
		client: client,
	}
}

func (r *TraitRepositoryImpl) UpsertBulk(ctx context.Context, traits []*ent.Trait) error {
	tx, err := r.client.Tx(ctx)
	if err != nil {
		return err
	}

	// 一度に1000件ずつ処理する
	const batchSize = 1000
	for i := 0; i < len(traits); i += batchSize {
		end := min(i+batchSize, len(traits))
		builders := make([]*ent.TraitCreate, 0, end-i)

		for _, trt := range traits[i:end] {
			builder := tx.Trait.Create().
				SetID(trt.ID).
				SetNameEn(trt.NameEn)
			builders = append(builders, builder)
		}
		if len(builders) == 0 {
			continue
		}

		// 一括で作成
		err = tx.Trait.CreateBulk(builders...).
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

	return tx.Commit()
}