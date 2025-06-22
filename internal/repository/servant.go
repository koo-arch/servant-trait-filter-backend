package repository

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/koo-arch/servant-trait-filter-backend/ent"
	"github.com/koo-arch/servant-trait-filter-backend/ent/servant"
	"github.com/koo-arch/servant-trait-filter-backend/internal/model"
	"github.com/koo-arch/servant-trait-filter-backend/internal/transaction"
)

type ServantRepository interface {
	ListAll(ctx context.Context) ([]*ent.Servant, error)
	UpsertBulk(ctx context.Context, servants []model.Servant) error
}

type ServantRepositoryImpl struct {
	client *ent.Client
}

func NewServantRepository(client *ent.Client) ServantRepository {
	return &ServantRepositoryImpl{
		client: client,
	}
}

func (r *ServantRepositoryImpl) ListAll(ctx context.Context) ([]*ent.Servant, error) {
	return r.client.Servant.Query().
		Order(ent.Asc(servant.FieldID)).
		All(ctx)
}

func (r *ServantRepositoryImpl) UpsertBulk(ctx context.Context, servants []model.Servant) error {
	tx, err := r.client.Tx(ctx)
	if err != nil {
		return err
	}

	defer transaction.HandleTransaction(tx, &err)

	// 一度に1000件ずつ処理する
	const batchSize = 1000
	for i := 0; i < len(servants); i += batchSize {
		end := min(i+batchSize, len(servants))
		builders := make([]*ent.ServantCreate, 0,  end-i)
		batchSvt := servants[i:end]

		for _, svt := range batchSvt {
			builder := tx.Servant.Create().
				SetID(svt.ID).
				SetName(svt.Name).
				SetCollectionNo(svt.CollectionNo).
				SetFace(svt.Face).
				SetClassID(svt.ClassID).
				SetOrderAlignmentID(svt.OrderAlignmentID).
				SetMoralAlignmentID(svt.MoralAlignmentID).
				SetAttributeID(svt.AttributeID)
			builders = append(builders, builder)
			}
		if len(builders) == 0 {
			continue
		}

		// 一括で作成
		if err := tx.Servant.CreateBulk(builders...).
			OnConflict(
				sql.ConflictColumns(
					servant.FieldID,
					servant.FieldCollectionNo,
				),
			).
			UpdateNewValues().
			Exec(ctx); err != nil {
				return err
			}
		
		// 同時にTraitsも更新
		for _, svt := range batchSvt {
			if err := r.syncServantTraits(ctx, tx, svt.ID, svt.Traits); err != nil {
				return err
			}
		}
	}

	return tx.Commit()
}

func (r *ServantRepositoryImpl) syncServantTraits(ctx context.Context, tx *ent.Tx, servantID int, newTraitIDs []int) error {
	currentTraitIDs, err := tx.Servant.Query().
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

	syncTraits:= tx.Servant.UpdateOneID(servantID)
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