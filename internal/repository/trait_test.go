package repository_test

import (
	"context"
	"testing"
	"github.com/koo-arch/servant-trait-filter-backend/ent/enttest"
	"github.com/koo-arch/servant-trait-filter-backend/ent/trait"
	"github.com/koo-arch/servant-trait-filter-backend/internal/model"
	"github.com/koo-arch/servant-trait-filter-backend/internal/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	_ "github.com/mattn/go-sqlite3"
)

func TestTraitRepository_UpsertBulk(t *testing.T) {
	ctx := context.Background()
	client := enttest.Open(t, "sqlite3", "file:memdb?mode=memory&_fk=1")
	defer client.Close()

	err := client.Schema.Create(ctx)
	require.NoError(t, err)

	repo := repository.NewTraitRepository(client)

	// テストデータを準備
	in := []model.Trait{
		{ID: 1, Name: "Trait A"},
		{ID: 2, Name: "Trait B"},
		{ID: 3, Name: "Trait C"},
	}

	err = repo.UpsertBulk(ctx, in)
	require.NoError(t, err)

	// データが正しく挿入されたか確認
	traits, err := repo.ListAll(ctx)
	assert.NoError(t, err)
	assert.Len(t, traits, 3)
	for i, trt := range in {
		assert.Equal(t, trt.ID, traits[i].ID)
		assert.Equal(t, trt.Name, traits[i].NameEn)
	}

	// ２回目のアップサートで更新されるか確認
	in[0].Name = "Updated Trait A"
	err = repo.UpsertBulk(ctx, in)
	assert.NoError(t, err)
	traits, err = repo.ListAll(ctx)
	assert.NoError(t, err)
	assert.Len(t, traits, 3)
	assert.Equal(t, "Updated Trait A", traits[0].NameEn)

	// 更新したデータ以外が変更されていないか確認
	for i, trt := range in {
		if i == 0 {
			continue
		}
		assert.Equal(t, trt.ID, traits[i].ID)
		assert.Equal(t, trt.Name, traits[i].NameEn)
	}

	// 手動で日本語のデータを挿入
	_, err = client.Trait.Update().
		SetNameJa("特性").
		Where(trait.ID(1)).
		Save(ctx)
	require.NoError(t, err)

	// 再度アップサートして日本語のデータが上書きされないことを確認
	err = repo.UpsertBulk(ctx, in)
	assert.NoError(t, err)
	traits, err = repo.ListAll(ctx)
	assert.NoError(t, err)
	assert.Len(t, traits, 3)
	assert.Equal(t, "特性", traits[0].NameJa)

	for i, trt := range traits {
		assert.Equal(t, in[i].ID, trt.ID)
		assert.Equal(t, in[i].Name, trt.NameEn)
	}
	

}