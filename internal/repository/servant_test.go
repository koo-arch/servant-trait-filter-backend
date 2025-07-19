package repository_test

import (
	"context"
	"log"
	"testing"

	"github.com/koo-arch/servant-trait-filter-backend/ent"
	"github.com/koo-arch/servant-trait-filter-backend/ent/enttest"
	"github.com/koo-arch/servant-trait-filter-backend/internal/model"
	"github.com/koo-arch/servant-trait-filter-backend/internal/repository"
	"github.com/koo-arch/servant-trait-filter-backend/internal/search"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestServantRepository_UpsertBulk(t *testing.T) {
	ctx := context.Background()
	client := enttest.Open(t, "sqlite3", "file:memdb?mode=memory&_fk=1")
	defer client.Close()

	err := client.Schema.Create(ctx)
  	require.NoError(t, err)

	// マスターデータをシード
	err = seedMaster(ctx, client)
	require.NoError(t, err)

	repo := repository.NewServantRepository(client)

	// テストデータを準備
	in := []model.Servant{
		{ID: 1, Name: "アルトリア・ペンドラゴン", CollectionNo: 1, Face: "face1.png", ClassID: 1, Traits: []int{1, 2}},
		{ID: 2, Name: "ギルガメッシュ", CollectionNo: 2, Face: "face2.png", ClassID: 2, Traits: []int{2, 3, 4}},
		{ID: 3, Name: "エミヤ", CollectionNo: 3, Face: "face3.png", ClassID: 3, Traits: []int{1, 4}},
	}
	

	err = repo.UpsertBulk(ctx, in)
	require.NoError(t, err)

	// データが正しく挿入されたか確認
	servants, err := repo.ListAll(ctx)
	assert.NoError(t, err)
	assert.Len(t, servants, 3)
	for i, svt := range in {
		assert.Equal(t, svt.ID, servants[i].ID)
		assert.Equal(t, svt.Name, servants[i].Name)
		assert.Equal(t, svt.CollectionNo, servants[i].CollectionNo)
		assert.Equal(t, svt.Face, servants[i].Face)
		assert.Equal(t, svt.ClassID, servants[i].ClassID)
	}

	// ２回目のアップサートで更新されるか確認
	in[0].Name = "アルトリア・ペンドラゴン・オルタ"
	err = repo.UpsertBulk(ctx, in)
	assert.NoError(t, err)
	servants, err = repo.ListAll(ctx)
	assert.NoError(t, err)
	assert.Len(t, servants, 3)
	for i, svt := range in {
		assert.Equal(t, svt.ID, servants[i].ID)
		assert.Equal(t, svt.Name, servants[i].Name)
		assert.Equal(t, svt.CollectionNo, servants[i].CollectionNo)
		assert.Equal(t, svt.Face, servants[i].Face)
		assert.Equal(t, svt.ClassID, servants[i].ClassID)
	}
	
	// 件数が同じまま
	assert.Equal(t, 3, len(servants), "Should have 3 servants after upsert")
	s1, _ := client.Servant.Get(ctx, 1)
	assert.Equal(t, "アルトリア・ペンドラゴン・オルタ", s1.Name, "First servant's name should be updated")

	// 他は変更されていないことを確認
	for i, svt := range in {
		if i == 0 {
			assert.Equal(t, "アルトリア・ペンドラゴン・オルタ", servants[i].Name, "First servant's name should be updated")
		} else {
			assert.Equal(t, svt.Name, servants[i].Name, "Other servants' names should remain unchanged")
		}
		assert.Equal(t, svt.CollectionNo, servants[i].CollectionNo)
		assert.Equal(t, svt.Face, servants[i].Face)
		assert.Equal(t, svt.ClassID, servants[i].ClassID)
	}

	// 3回目のアップサートで新しいデータを追加
	in = append(in, model.Servant{
		ID: 4, Name: "クー・フーリン", CollectionNo: 4, Face: "face4.png", ClassID: 4, Traits: []int{5, 6},
	})
	err = repo.UpsertBulk(ctx, in)
	assert.NoError(t, err)
	servants, err = repo.ListAll(ctx)
	assert.NoError(t, err)
	assert.Len(t, servants, 4, "Should have 4 servants after upsert")

	// 新しいデータが正しく挿入されたか確認
	for i, svt := range in {
		assert.Equal(t, svt.ID, servants[i].ID)
		assert.Equal(t, svt.Name, servants[i].Name)
		assert.Equal(t, svt.CollectionNo, servants[i].CollectionNo)
		assert.Equal(t, svt.Face, servants[i].Face)
		assert.Equal(t, svt.ClassID, servants[i].ClassID)
	}

	// Traitsに更新がある場合
	in[0].Traits = []int{2, 3, 6} // Traitsを更新
	err = repo.UpsertBulk(ctx, in)
	assert.NoError(t, err)
	gotTraitIDs, _ := s1.QueryTraits().IDs(ctx)
	assert.ElementsMatch(t, []int{2, 3, 6}, gotTraitIDs, "Traits should be updated correctly")
	
}

func seedMaster(ctx context.Context, client *ent.Client) error {
	// クラスのデータを準備
	classes := []struct {
		ID   int
		NameEn string
		NameJa string
	}{
		{1, "Saber", "セイバー"},
		{2, "Archer", "アーチャー"},
		{3, "Caster", "キャスター"},
		{4, "Lancer", "ランサー"},
	}

	for _, cls := range classes {
		if _, err := client.Class.Create().SetID(cls.ID).SetNameEn(cls.NameEn).SetNameJa(cls.NameJa).Save(ctx); err != nil {
			return err
		}
	}

	// OrderAlignmentのデータを準備
	orderAlignments := []struct {
		ID     int
		NameEn string
		NameJa string
	}{
		{1, "Lawful", "秩序"},
		{2, "Neutral", "中立"},
		{3, "Chaotic", "混沌"},
	}

	for _, oa := range orderAlignments {
		if _, err := client.OrderAlignment.Create().SetID(oa.ID).SetNameEn(oa.NameEn).SetNameJa(oa.NameJa).Save(ctx); err != nil {
			return err
		}
	}

	// MoralAlignmentのデータを準備
	moralAlignments := []struct {
		ID     int
		NameEn string
		NameJa string
	}{
		{1, "Good", "善"},
		{2, "Neutral", "中庸"},
		{3, "Evil", "悪"},
		{4, "Mad", "狂"},
		{5, "Summer", "夏"},
	}
	for _, ma := range moralAlignments {
		if _, err := client.MoralAlignment.Create().SetID(ma.ID).SetNameEn(ma.NameEn).SetNameJa(ma.NameJa).Save(ctx); err != nil {
			return err
		}
	}

	// Attributeのデータを準備
	attributes := []struct {
		ID   int
		NameEn string
		NameJa string
	}{
		{1, "Earth", "地"},
		{2, "Sky", "天"},
		{3, "Human", "人"},
		{4, "Star", "星"},
		{5, "Beast", "獣"},
	}
	for _, attr := range attributes {
		if _, err := client.Attribute.Create().SetID(attr.ID).SetNameEn(attr.NameEn).SetNameJa(attr.NameJa).Save(ctx); err != nil {
			return err
		}
	}

	// Traitsのデータを準備
	traits := []struct {
		ID   int
		NameEn string
		NameJa string
	}{
		{1, "Trait1", "トレイト1"},
		{2, "Trait2", "トレイト2"},
		{3, "Trait3", "トレイト3"},
		{4, "Trait4", "トレイト4"},
		{5, "Trait5", "トレイト5"},
		{6, "Trait6", "トレイト6"},
	}
	for _, tr := range traits {
		if _, err := client.Trait.Create().SetID(tr.ID).SetNameEn(tr.NameEn).SetNameJa(tr.NameJa).Save(ctx); err != nil {
			return err
		}
	}

	return nil
}

func TestServantRepository_Search(t *testing.T) {
	ctx := context.Background()
	client := enttest.Open(t, "sqlite3", "file:memdb?mode=memory&_fk=1")
	defer client.Close()

	err := client.Schema.Create(ctx)
	require.NoError(t, err)

	// マスターデータをシード
	err = seedMaster(ctx, client)
	require.NoError(t, err)

	repo := repository.NewServantRepository(client)

	// テストデータを準備
	in := []model.Servant{
		{ID: 1, Name: "アルトリア・ペンドラゴン", CollectionNo: 1, Face: "face1.png", ClassID: 1, Traits: []int{1, 2}},
		{ID: 2, Name: "ギルガメッシュ", CollectionNo: 2, Face: "face2.png", ClassID: 2, Traits: []int{2, 3, 4}},
		{ID: 3, Name: "エミヤ", CollectionNo: 3, Face: "face3.png", ClassID: 3, Traits: []int{1, 4}},
	}
	err = repo.UpsertBulk(ctx, in)
	require.NoError(t, err)

	// 検索DTL
	req := search.ServantSearchQuery{
		Root: search.Expr{
			Or: []*search.Expr{
				{
					And: []*search.Expr{
						{ TraitID: ptr(1) }, // トレイト1
						{ TraitID: ptr(2) }, // トレイト2
					},
				},
				{
					And: []*search.Expr{
						{ ClassID: ptr(2) }, // クラス2 (アーチャー)
						{ TraitID: ptr(3) }, // トレイト3
					},
				},
			},
		},
		Limit: 10,
		Offset: 0,
	}

	// 検索実行
	result, err := repo.Search(ctx, req)
	require.NoError(t, err)
	log.Printf("Found %d servants", len(result.Servants))
}

func ptr(i int) *int {
	if i < 0 {
		return nil
	}
	return &i
}