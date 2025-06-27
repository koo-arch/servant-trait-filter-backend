package repository_test

import (
	"context"
	"testing"

	"github.com/koo-arch/servant-trait-filter-backend/ent/attribute"
	"github.com/koo-arch/servant-trait-filter-backend/ent/enttest"
	"github.com/koo-arch/servant-trait-filter-backend/internal/model"
	"github.com/koo-arch/servant-trait-filter-backend/internal/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	_ "github.com/mattn/go-sqlite3"
)

func TestAttributeRepository_UpsertBulk(t *testing.T) {
	ctx := context.Background()
	client := enttest.Open(t, "sqlite3", "file:memdb?mode=memory&_fk=1")
	defer client.Close()

	err := client.Schema.Create(ctx)
	require.NoError(t, err)

	repo := repository.NewAttributeRepository(client)

	// テストデータを準備
	in := []model.Attribute{
		{ID: 1, Name: "天"},
		{ID: 2, Name: "地"},
		{ID: 3, Name: "人"},
		{ID: 4, Name: "星"},
	}

	err = repo.UpsertBulk(ctx, in)
	require.NoError(t, err)

	// データが正しく挿入されたか確認
	attributes, err := repo.ListAll(ctx)
	assert.NoError(t, err)
	assert.Len(t, attributes, 4)
	for i, attr := range in {
		assert.Equal(t, attr.ID, attributes[i].ID)
		assert.Equal(t, attr.Name, attributes[i].NameEn)
	}
	
	// ２回目のアップサートで更新されるか確認
	in[3].Name = "獣"
	err = repo.UpsertBulk(ctx, in)
	assert.NoError(t, err)
	attributes, err = repo.ListAll(ctx)
	assert.NoError(t, err)
	assert.Len(t, attributes, 4)
	assert.Equal(t, "獣", attributes[3].NameEn)
	// 更新したデータ以外が変更されていないか確認
	for i, attr := range in {
		if i == 3 {
			continue
		}
		assert.Equal(t, attr.ID, attributes[i].ID)
		assert.Equal(t, attr.Name, attributes[i].NameEn)
	}

	// 存在しないIDをアップサートしてもエラーにならないことを確認
	new_attr := model.Attribute{ID: 5, Name: "星"}
	in = append(in, new_attr) // 既存のデータに追加
	
	
	err = repo.UpsertBulk(ctx, []model.Attribute{new_attr})
	assert.NoError(t, err)
	attributes, err = repo.ListAll(ctx)
	assert.NoError(t, err)
	assert.Len(t, attributes, 5)
	assert.Equal(t, "星", attributes[4].NameEn)

	// 手動で日本語のデータを入力した時の挙動
	_, err = client.Attribute.Update().
		SetNameJa("天").
		Where(attribute.ID(1)).
		Save(ctx)
	require.NoError(t, err)

	// 再度アップサートして日本語のデータが上書きされないことを確認
	err = repo.UpsertBulk(ctx, in)
	assert.NoError(t, err)
	attributes, err = repo.ListAll(ctx)
	assert.NoError(t, err)
	assert.Len(t, attributes, 5)
	assert.Equal(t, "天", attributes[0].NameJa) // 日本語のデータが保持されていることを確認

	// 他のレコードが変更されていないことを確認
	for i, attr := range in {
		if i == 0 {
			continue
		}
		assert.Equal(t, attr.ID, attributes[i].ID)
		assert.Equal(t, attr.Name, attributes[i].NameEn)
	}

	// 名前を変更しても日本語のデータが保持されることを確認
	in[0].Name = "天空"
	err = repo.UpsertBulk(ctx, in)
	assert.NoError(t, err)
	attributes, err = repo.ListAll(ctx)
	assert.NoError(t, err)
	assert.Len(t, attributes, 5)
	assert.Equal(t, "天", attributes[0].NameJa) // 日本語のデータが保持されていることを確認
	for _, attr := range attributes {
		assert.NotEmpty(t, attr.NameEn) // 英語名が空でないことを確認
	}
}