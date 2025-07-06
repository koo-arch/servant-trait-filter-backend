package repository

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/koo-arch/servant-trait-filter-backend/ent"
	"github.com/koo-arch/servant-trait-filter-backend/ent/servant"
	"github.com/koo-arch/servant-trait-filter-backend/internal/model"
)

type ServantRepository interface {
	Get(ctx context.Context, id int) (*ent.Servant, error)
	ListAll(ctx context.Context) ([]*ent.Servant, error)
	UpsertBulk(ctx context.Context, servants []model.Servant) error
	WithTx(tx *ent.Tx) ServantRepository
}

type ServantRepositoryImpl struct {
	client *ent.Client
}

func NewServantRepository(client *ent.Client) ServantRepository {
	return &ServantRepositoryImpl{
		client: client,
	}
}

func (r *ServantRepositoryImpl) WithTx(tx *ent.Tx) ServantRepository {
	return &ServantRepositoryImpl{
		client: tx.Client(),
	}
}

func (r *ServantRepositoryImpl) Get(ctx context.Context, id int) (*ent.Servant, error) {
	return r.client.Servant.Query().
		Where(servant.ID(id)).
		WithTraits().
		Only(ctx)
}

func (r *ServantRepositoryImpl) ListAll(ctx context.Context) ([]*ent.Servant, error) {
	return r.client.Servant.Query().
		Order(ent.Asc(servant.FieldID)).
		All(ctx)
}

func (r *ServantRepositoryImpl) UpsertBulk(ctx context.Context, servants []model.Servant) error {
	if len(servants) == 0 {
		return nil
	}

	// 一度に1000件ずつ処理する
	const batchSize = 1000
	for i := 0; i < len(servants); i += batchSize {
		end := min(i+batchSize, len(servants))
		builders := make([]*ent.ServantCreate, 0,  end-i)
		batchSvt := servants[i:end]

		for _, svt := range batchSvt {
			builder := r.client.Servant.Create().
				SetID(svt.ID).
				SetName(svt.Name).
				SetCollectionNo(svt.CollectionNo).
				SetFace(svt.Face).
				SetClassID(svt.ClassID)
			builders = append(builders, builder)
			}
		if len(builders) == 0 {
			continue
		}

		// 一括で作成
		if err := r.client.Servant.CreateBulk(builders...).
			OnConflict(
				sql.ConflictColumns(
					servant.FieldCollectionNo,
				),
			).
			UpdateNewValues().
			Exec(ctx); err != nil {
				return err
			}
		
		// 同時にTraitsも更新
		for _, svt := range batchSvt {
			if err := r.syncServantTraits(ctx, svt.ID, svt.Traits); err != nil {
				return err
			}
		}
	}

	return nil
}

func (r *ServantRepositoryImpl) syncServantTraits(ctx context.Context, servantID int, newTraitIDs []int) error {
	currentTraitIDs, err := r.client.Servant.Query().
		Where(servant.ID(servantID)).
		QueryTraits().
		IDs(ctx)
	if err != nil {
		return err
	}

	// Diffを計算して追加と削除のTrait IDを取得
	toAdd, toDel := diffIntSets(currentTraitIDs, newTraitIDs)

	if len(toAdd) == 0 && len(toDel) == 0 {
		return nil
	}

	syncTraits:= r.client.Servant.UpdateOneID(servantID)
	if len(toAdd) > 0 {
		syncTraits.AddTraitIDs(toAdd...)
	}
	if len(toDel) > 0 {
		syncTraits.RemoveTraitIDs(toDel...)
	}
	if err := syncTraits.Exec(ctx); err != nil {
		return err
	}
	return nil

}

func diffIntSets(have, want []int) (toAdd, toDel []int) {
	inHave := make(map[int]struct{}, len(have))
	for _, v := range have {
		inHave[v] = struct{}{}
	}

	for _, v := range want {
		if _, exists := inHave[v]; !exists {
			toAdd = append(toAdd, v)
		}
		delete(inHave, v)
	}

	for v := range inHave {
		toDel = append(toDel, v)
	}

	return toAdd, toDel
}