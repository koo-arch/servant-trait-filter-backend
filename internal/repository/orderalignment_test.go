package repository_test

import (
	"context"
	"testing"

	"github.com/koo-arch/servant-trait-filter-backend/ent/orderalignment"
	"github.com/koo-arch/servant-trait-filter-backend/ent/enttest"
	"github.com/koo-arch/servant-trait-filter-backend/internal/model"
	"github.com/koo-arch/servant-trait-filter-backend/internal/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	_ "github.com/mattn/go-sqlite3"
)

func TestOrderAlignmentRepository_UpsertBulk(t *testing.T) {
	ctx := context.Background()
	client := enttest.Open(t, "sqlite3", "file:memdb?mode=memory&_fk=1")
	defer client.Close()

	err := client.Schema.Create(ctx)
	require.NoError(t, err)

	repo := repository.NewOrderAlignmentRepository(client)

	// テストデータを準備
	in := []model.OrderAlignment{
		{ID: 1, Name: "Lawful"},
		{ID: 2, Name: "Neutral"},
		{ID: 3, Name: "Chaotic"},
	}

	err = repo.UpsertBulk(ctx, in)
	require.NoError(t, err)

	// データが正しく挿入されたか確認
	alignments, err := repo.ListAll(ctx)
	assert.NoError(t, err)
	assert.Len(t, alignments, 3)
	for i, al := range in {
		assert.Equal(t, al.ID, alignments[i].ID)
		assert.Equal(t, al.Name, alignments[i].NameEn)
	}

	// ２回目のアップサートで更新されるか確認
	in[0].Name = "Lawful Good"
	err = repo.UpsertBulk(ctx, in)
	assert.NoError(t, err)
	alignments, err = repo.ListAll(ctx)
	assert.NoError(t, err)
	assert.Len(t, alignments, 3)
	assert.Equal(t, "Lawful Good", alignments[0].NameEn)

	// 更新したデータ以外が変更されていないか確認
	for i, al := range in {
		if i == 0 {
			continue
		}
		assert.Equal(t, al.ID, alignments[i].ID)
		assert.Equal(t, al.Name, alignments[i].NameEn)
	}

	// 存在しないIDを含むデータをアップサート
	new_alignment := model.OrderAlignment{ID: 4, Name: "Lawful Evil"}
	in = append(in, new_alignment)

	err = repo.UpsertBulk(ctx, []model.OrderAlignment{new_alignment})
	require.NoError(t, err)

	// 新しいデータが追加されたか確認
	alignments, err = repo.ListAll(ctx)
	assert.NoError(t, err)
	assert.Len(t, alignments, 4)
	for i, oa := range in {
		assert.Equal(t, oa.ID, alignments[i].ID)
		assert.Equal(t, oa.Name, alignments[i].NameEn)
	}

	// 手動で日本語のデータを入力した時の挙動
	_, err = client.OrderAlignment.Update().
		SetNameJa("秩序").
		Where(orderalignment.ID(1)).
		Save(ctx)
	require.NoError(t, err)
	
	// 再度アップサートして日本語のデータが上書きされないことを確認
	err = repo.UpsertBulk(ctx, in)
	require.NoError(t, err)
	alignments, err = repo.ListAll(ctx)
	assert.NoError(t, err)
	assert.Len(t, alignments, 4)
	assert.Equal(t, "秩序", alignments[0].NameJa) // 日本語のデータが保持されていることを確認
	for i, oa := range in {
		assert.Equal(t, oa.ID, alignments[i].ID)
		assert.Equal(t, oa.Name, alignments[i].NameEn)
	}

	// 名前を変更しても日本語のデータが保持されることを確認
	in[0].Name = "Lawful"
	err = repo.UpsertBulk(ctx, in)
	require.NoError(t, err)
	alignments, err = repo.ListAll(ctx)
	assert.NoError(t, err)
	assert.Len(t, alignments, 4)
	assert.Equal(t, "秩序", alignments[0].NameJa) // 日本語のデータが保持されていることを確認
	for i, oa := range in {
		assert.Equal(t, oa.ID, alignments[i].ID)
		assert.Equal(t, oa.Name, alignments[i].NameEn)
	}
}