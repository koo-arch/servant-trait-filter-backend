package etl

import (
	"context"
	"fmt"

	"github.com/koo-arch/servant-trait-filter-backend/ent"
	"github.com/koo-arch/servant-trait-filter-backend/internal/atlas"
	"github.com/koo-arch/servant-trait-filter-backend/internal/di"
	"github.com/koo-arch/servant-trait-filter-backend/internal/transaction"
)

type SyncAtlas struct {
	DB *ent.Client
	Client atlas.Client
	Repos *di.Repos
}

func NewSyncAtlas(
	db *ent.Client,
	client atlas.Client,
	repos *di.Repos,
) *SyncAtlas {
	return &SyncAtlas{
		DB: db,
		Client: client,
		Repos: repos,
	}
}

func (s *SyncAtlas) Sync(ctx context.Context) error {
	// サーヴァントを取得
	servants, err := s.Client.FetchServants(ctx, "JP")
	if err != nil {
		return fmt.Errorf("failed to fetch servants: %w", err)
	}
	// 特性を取得
	traits, err := s.Client.FetchTraits(ctx, "JP")
	if err != nil {
		return fmt.Errorf("failed to fetch traits: %w", err)
	}

	if len(servants) == 0 && len(traits) == 0 {
		return nil // No data to sync
	}

	// トランザクションを開始
	tx, err := s.DB.Tx(ctx)
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}

	defer transaction.HandleTransaction(tx, &err)

	// リポジトリをトランザクションに紐付け
	repos := s.execWithTx(tx)

	// データをリポジトリにアップサート
	if err := repos.Trait.UpsertBulk(ctx, traits); err != nil {
		return fmt.Errorf("failed to upsert traits: %w", err)
	}

	// クラスを抽出
	class := s.extractClass(servants)
	if err := repos.Class.UpsertBulk(ctx, class); err != nil {
		return fmt.Errorf("failed to upsert classes: %w", err)
	}

	// 特性から属性とアライメントを抽出
	attributes, orderAlign, moralAlign := s.extractMetaFromTraits(traits)
	if err := repos.Attribute.UpsertBulk(ctx, attributes); err != nil {
		return fmt.Errorf("failed to upsert attributes: %w", err)
	}
	if err := repos.OrderAlign.UpsertBulk(ctx, orderAlign); err != nil {
		return fmt.Errorf("failed to upsert order alignments: %w", err)
	}
	if err := repos.MoralAlign.UpsertBulk(ctx, moralAlign); err != nil {
		return fmt.Errorf("failed to upsert moral alignments: %w", err)
	}
	
	// プレイアブルサーヴァントを抽出
	playable, ascs := s.extractPlayable(servants)
	if err := repos.Servant.UpsertBulk(ctx, playable); err != nil {
		return fmt.Errorf("failed to upsert servants: %w", err)
	}
	if err := repos.Ascension.UpsertBulk(ctx, ascs); err != nil {
		return fmt.Errorf("failed to upsert ascensions: %w", err)
	}

	return nil
}

func (s *SyncAtlas) execWithTx(tx *ent.Tx) *di.Repos {
	return &di.Repos{
		Servant: s.Repos.Servant.WithTx(tx),
		Class: s.Repos.Class.WithTx(tx),
		Attribute: s.Repos.Attribute.WithTx(tx),
		Trait: s.Repos.Trait.WithTx(tx),
		MoralAlign: s.Repos.MoralAlign.WithTx(tx),
		OrderAlign: s.Repos.OrderAlign.WithTx(tx),
		Ascension: s.Repos.Ascension.WithTx(tx),
	}
}