package repository

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/koo-arch/servant-trait-filter-backend/ent"
	"github.com/koo-arch/servant-trait-filter-backend/ent/class"
	"github.com/koo-arch/servant-trait-filter-backend/internal/model"
)

// ClassRepository is the interface for the class repository.
type ClassRepository interface {
	UpsertBulk(ctx context.Context, classes []model.Class) error
	ListAll(ctx context.Context) ([]*ent.Class, error)
	WithTx(tx *ent.Tx) ClassRepository
}

type ClassRepositoryImpl struct {
	client *ent.Client
}

func NewClassRepository(client *ent.Client) ClassRepository {
	return &ClassRepositoryImpl{
		client: client,
	}
}

func (r *ClassRepositoryImpl) WithTx(tx *ent.Tx) ClassRepository {
	return &ClassRepositoryImpl{
		client: tx.Client(),
	}
}

func (r *ClassRepositoryImpl) ListAll(ctx context.Context) ([]*ent.Class, error) {
	return r.client.Class.Query().
		Order(ent.Asc(class.FieldID)).
		All(ctx)
}

func (r *ClassRepositoryImpl) UpsertBulk(ctx context.Context, classes []model.Class) error {

	// 一度に1000件ずつ処理する
	const batchSize = 1000
	for i := 0; i < len(classes); i += batchSize {
		end := min(i+batchSize, len(classes))
		builders := make([]*ent.ClassCreate, 0, end-i)

		for _, cls := range classes[i:end] {
			builder := r.client.Class.Create().
				SetID(cls.ID).
				SetNameEn(cls.Name)
			builders = append(builders, builder)
		}
		if len(builders) == 0 {
			continue
		}

		// 一括で作成
		err := r.client.Class.CreateBulk(builders...).
			OnConflict(
				sql.ConflictColumns(class.FieldID),
				sql.UpdateWhere(
					sql.ExprP(
						"classes.name_en IS DISTINCT FROM EXCLUDED.name_en",
					),
				),
			).
			UpdateNewValues().
			Exec(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}