package repository_test

import (
	"context"
	"testing"

	"github.com/koo-arch/servant-trait-filter-backend/ent/class"
	"github.com/koo-arch/servant-trait-filter-backend/ent/enttest"
	"github.com/koo-arch/servant-trait-filter-backend/internal/model"
	"github.com/koo-arch/servant-trait-filter-backend/internal/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	_ "github.com/mattn/go-sqlite3"
)

func TestClassRepository_UpsertBulk(t *testing.T) {
	ctx := context.Background()
	client := enttest.Open(t, "sqlite3", "file:memdb?mode=memory&_fk=1")
	defer client.Close()

	err := client.Schema.Create(ctx)
	require.NoError(t, err)

	repo := repository.NewClassRepository(client)

	// テストデータを準備
	in := []model.Class{
		{ID: 1, Name: "Saber"},
		{ID: 2, Name: "Archer"},
		{ID: 3, Name: "Lancer"},
	}

	err = repo.UpsertBulk(ctx, in)
	require.NoError(t, err)

	// データが正しく挿入されたか確認
	classes, err := repo.ListAll(ctx)
	assert.NoError(t, err)
	assert.Len(t, classes, 3)
	for i, cls := range in {
		assert.Equal(t, cls.ID, classes[i].ID)
		assert.Equal(t, cls.Name, classes[i].NameEn)
		assert.Empty(t, classes[i].NameJa) // 日本語のデータは空であることを確認
	}

	// ２回目のアップサートで更新されるか確認
	in[0].Name = "Grand Saber"
	err = repo.UpsertBulk(ctx, in)
	assert.NoError(t, err)
	classes, err = repo.ListAll(ctx)
	assert.NoError(t, err)
	assert.Len(t, classes, 3)
	assert.Equal(t, "Grand Saber", classes[0].NameEn)

	// 更新したデータ以外が変更されていないか確認
	for i, cls := range in {
		if i == 0 {
			continue
		}
		assert.Equal(t, cls.ID, classes[i].ID)
		assert.Equal(t, cls.Name, classes[i].NameEn)
	}

	// 手動で日本語のデータを入力した時の挙動
	_, err = client.Class.Update().
		SetNameJa("セイバ-").
		Where(class.ID(1)).
		Save(ctx)
	require.NoError(t, err)
	// 再度アップサートして日本語のデータが上書きされないことを確認
	err = repo.UpsertBulk(ctx, in)
	require.NoError(t, err)
	classes, err = repo.ListAll(ctx)
	assert.NoError(t, err)
	assert.Len(t, classes, 3)
	assert.Equal(t, "セイバ-", classes[0].NameJa) // 日本語のデータが保持されていることを確認

	for _, cls := range classes {
		assert.NotEmpty(t, cls.NameEn) // 英語名が空でないことを確認
	}
	
	// 名前を変更しても日本語のデータが保持されることを確認
	in[0].Name = "Saber"
	err = repo.UpsertBulk(ctx, in)
	require.NoError(t, err)
	classes, err = repo.ListAll(ctx)
	assert.NoError(t, err)
	assert.Len(t, classes, 3)
	assert.Equal(t, "セイバ-", classes[0].NameJa) // 日本語のデータが保持されていることを確認
	for _, cls := range classes {
		assert.NotEmpty(t, cls.NameEn) // 英語名が空でないことを確認
	}
	assert.Equal(t, "Saber", classes[0].NameEn) // 英語名が更新されていることを確認
	assert.Equal(t, "Archer", classes[1].NameEn) // 他のクラスも確認
	assert.Equal(t, "Lancer", classes[2].NameEn) // 他のクラスも確認
}
