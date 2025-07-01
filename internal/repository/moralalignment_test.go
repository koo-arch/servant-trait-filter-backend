package repository_test

import (
	"context"
	"testing"

	"github.com/koo-arch/servant-trait-filter-backend/ent/moralalignment"
	"github.com/koo-arch/servant-trait-filter-backend/ent/enttest"
	"github.com/koo-arch/servant-trait-filter-backend/internal/model"
	"github.com/koo-arch/servant-trait-filter-backend/internal/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	_ "github.com/mattn/go-sqlite3"
)

func TestMoralAlignmentRepository_UpsertBulk(t *testing.T) {
	ctx := context.Background()
	client := enttest.Open(t, "sqlite3", "file:memdb?mode=memory&_fk=1")
	defer client.Close()

	err := client.Schema.Create(ctx)
	require.NoError(t, err)

	repo := repository.NewMoralAlignmentRepository(client)

	// テストデータを準備
	in := []model.MoralAlignment{
		{ID: 1, Name: "good"},
		{ID: 2, Name: "neutral"},
		{ID: 3, Name: "evil"},
	}

	err = repo.UpsertBulk(ctx, in)
	require.NoError(t, err)

	// データが正しく挿入されたか確認
	morals, err := repo.ListAll(ctx)
	assert.NoError(t, err)
	assert.Len(t, morals, 3)
	for i, ma := range in {
		assert.Equal(t, ma.ID, morals[i].ID)
		assert.Equal(t, ma.Name, morals[i].NameEn)
	}

	// ２回目のアップサートで更新されるか確認
	in[0].Name = "summer"
	err = repo.UpsertBulk(ctx, in)
	assert.NoError(t, err)
	morals, err = repo.ListAll(ctx)
	assert.NoError(t, err)
	assert.Len(t, morals, 3)
	assert.Equal(t, "summer", morals[0].NameEn)

	// 更新したデータ以外が変更されていないか確認
	for i, ma := range in {
		if i == 0 {
			continue
		}
		assert.Equal(t, ma.ID, morals[i].ID)
		assert.Equal(t, ma.Name, morals[i].NameEn)
	}

	// 存在しないIDを含むデータをアップサート
	new_alignment := model.MoralAlignment{ID: 4, Name: "madness"}
	in = append(in, new_alignment) // 既存のデータに追加

	err = repo.UpsertBulk(ctx, []model.MoralAlignment{new_alignment})
	assert.NoError(t, err)
	// 新しいデータが追加されたか確認
	morals, err = repo.ListAll(ctx)
	assert.NoError(t, err)
	assert.Len(t, morals, 4)
	
	for i, ma := range in {
		assert.Equal(t, ma.ID, morals[i].ID)
		assert.Equal(t, ma.Name, morals[i].NameEn)
	}

	// 手動で日本語のデータを入力した時の挙動
	_, err = client.MoralAlignment.Update().
		SetNameJa("善").
		Where(moralalignment.ID(1)).
		Save(ctx)
	require.NoError(t, err)
	// 再度アップサートして日本語のデータが上書きされないことを確認
	err = repo.UpsertBulk(ctx, in)
	require.NoError(t, err)
	morals, err = repo.ListAll(ctx)
	assert.NoError(t, err)
	assert.Len(t, morals, 4)
	assert.Equal(t, "善", morals[0].NameJa) // 日本語のデータが保持されていることを確認
	for i, ma := range in {
		assert.Equal(t, ma.ID, morals[i].ID)
		assert.Equal(t, ma.Name, morals[i].NameEn)
	}

	// 名前を変更しても日本語のデータが保持されることを確認
	in[0].Name = "good"
	err = repo.UpsertBulk(ctx, in)
	require.NoError(t, err)
	morals, err = repo.ListAll(ctx)
	assert.NoError(t, err)
	assert.Len(t, morals, 4)
	assert.Equal(t, "善", morals[0].NameJa) // 日本語のデータが保持されていることを確認
	for i, ma := range in {
		assert.Equal(t, ma.ID, morals[i].ID)
		assert.Equal(t, ma.Name, morals[i].NameEn)
	}
}