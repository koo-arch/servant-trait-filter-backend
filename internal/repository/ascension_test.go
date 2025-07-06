package repository_test

import (
	"context"
	"testing"

	"github.com/koo-arch/servant-trait-filter-backend/ent/enttest"
	"github.com/koo-arch/servant-trait-filter-backend/internal/model"
	"github.com/koo-arch/servant-trait-filter-backend/internal/repository"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAscensionRepository_UpsertBulk(t *testing.T) {
	ctx := context.Background()
	client := enttest.Open(t, "sqlite3", "file:memdb?mode=memory&_fk=1")
	defer client.Close()
	
	err := client.Schema.Create(ctx)
	require.NoError(t, err)

	// マスターデータをシード
	err = seedMaster(ctx, client)
	require.NoError(t, err)

	// サーヴァントリポジトリを作成
	servantRepo := repository.NewServantRepository(client)
	servants := []model.Servant{
		{ID: 1, Name: "アルトリア・ペンドラゴン", CollectionNo: 1, Face: "face1.png", ClassID: 1, Traits: []int{1, 2}},
		{ID: 2, Name: "ギルガメッシュ", CollectionNo: 2, Face: "face2.png", ClassID: 2, Traits: []int{2, 3, 4}},
		{ID: 3, Name: "エミヤ", CollectionNo: 3, Face: "face3.png", ClassID: 3, Traits: []int{1, 4}},
	}
	err = servantRepo.UpsertBulk(ctx, servants)
	require.NoError(t, err)

	// アセンションリポジトリを作成
	repo := repository.NewAscensionRepository(client)
	ascensions := []model.Ascension{
		{ServantID: 1, Stage: 1, AttributeID: 1, MoralAlignmentID: 1, OrderAlignmentID: 1},
		{ServantID: 1, Stage: 2, AttributeID: 2, MoralAlignmentID: 2, OrderAlignmentID: 2},
		{ServantID: 2, Stage: 1, AttributeID: 1, MoralAlignmentID: 1, OrderAlignmentID: 1},
		{ServantID: 2, Stage: 2, AttributeID: 2, MoralAlignmentID: 2, OrderAlignmentID: 2},
		{ServantID: 3, Stage: 1, AttributeID: 1, MoralAlignmentID: 1, OrderAlignmentID: 1},
		{ServantID: 3, Stage: 2, AttributeID: 2, MoralAlignmentID: 2, OrderAlignmentID: 2},
	}
	err = repo.UpsertBulk(ctx, ascensions)
	require.NoError(t, err)

	// データが正しく挿入されたか確認
	ascList, err := repo.ListAll(ctx)
	assert.NoError(t, err)
	assert.Len(t, ascList, 6)
	for i, asc := range ascList {
		assert.Equal(t, ascensions[i].ServantID, asc.ServantID)
		assert.Equal(t, ascensions[i].Stage, asc.Stage)
		assert.Equal(t, ascensions[i].AttributeID, asc.AttributeID)
		assert.Equal(t, ascensions[i].MoralAlignmentID, asc.MoralAlignmentID)
		assert.Equal(t, ascensions[i].OrderAlignmentID, asc.OrderAlignmentID)
	}
	// ２回目のアップサートで更新されるか確認
	ascensions[0].AttributeID = 3
	err = repo.UpsertBulk(ctx, ascensions)
	assert.NoError(t, err)
	ascList, err = repo.ListAll(ctx)
	assert.NoError(t, err)
	assert.Len(t, ascList, 6)
	assert.Equal(t, 3, ascList[0].AttributeID, "AttributeID should be updated")
	// 更新したデータ以外が変更されていないか確認
	for i, asc := range ascList {
		assert.Equal(t, ascensions[i].ServantID, asc.ServantID)
		assert.Equal(t, ascensions[i].Stage, asc.Stage)
		assert.Equal(t, ascensions[i].AttributeID, asc.AttributeID)
		assert.Equal(t, ascensions[i].MoralAlignmentID, asc.MoralAlignmentID)
		assert.Equal(t, ascensions[i].OrderAlignmentID, asc.OrderAlignmentID)
	}
	
	// AlignmentのIDが未設定の場合は更新されないことを確認
	ascensions[0].MoralAlignmentID = 0
	ascensions[0].OrderAlignmentID = 0
	err = repo.UpsertBulk(ctx, ascensions)
	for i, asc := range ascList {
		if i == 0 {
			assert.Equal(t, 1, asc.MoralAlignmentID, "MoralAlignmentID should not be updated")
			assert.Equal(t, 1, asc.OrderAlignmentID, "OrderAlignmentID should not be updated")
		} else {
			assert.Equal(t, ascensions[i].MoralAlignmentID, asc.MoralAlignmentID)
			assert.Equal(t, ascensions[i].OrderAlignmentID, asc.OrderAlignmentID)
		}
	}

}