package etl

import (
	"context"
	"fmt"

	"github.com/koo-arch/servant-trait-filter-backend/ent"
	"github.com/koo-arch/servant-trait-filter-backend/internal/atlas"
	"github.com/koo-arch/servant-trait-filter-backend/internal/repository"
	"github.com/koo-arch/servant-trait-filter-backend/internal/transaction"
)

type SyncAtlas struct {
	DB *ent.Client
	Client atlas.Client
	ClassRepo repository.ClassRepository
	AttrRepo repository.AttributeRepository
	MoralRepo repository.MoralAlignmentRepository
	OrderRepo repository.OrderAlignmentRepository
	SvtRepo repository.ServantRepository
	TraitRepo repository.TraitRepository
	AscRepo repository.AscensionRepository
}

func NewSyncAtlas(
	db *ent.Client,
	client atlas.Client,
	classRepo repository.ClassRepository,
	attrRepo repository.AttributeRepository,
	moralRepo repository.MoralAlignmentRepository,
	orderRepo repository.OrderAlignmentRepository,
	svtRepo repository.ServantRepository,
	traitRepo repository.TraitRepository,
	ascRepo repository.AscensionRepository,
) *SyncAtlas {
	return &SyncAtlas{
		DB: db,
		Client: client,
		ClassRepo: classRepo,
		AttrRepo: attrRepo,
		MoralRepo: moralRepo,
		OrderRepo: orderRepo,
		SvtRepo: svtRepo,
		TraitRepo: traitRepo,
		AscRepo: ascRepo,
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
	classRepo := s.ClassRepo.WithTx(tx)
	attrRepo := s.AttrRepo.WithTx(tx)
	moralRepo := s.MoralRepo.WithTx(tx)
	orderRepo := s.OrderRepo.WithTx(tx)
	svtRepo := s.SvtRepo.WithTx(tx)
	traitRepo := s.TraitRepo.WithTx(tx)
	ascRepo := s.AscRepo.WithTx(tx)

	// データをリポジトリにアップサート
	if err := traitRepo.UpsertBulk(ctx, traits); err != nil {
		return fmt.Errorf("failed to upsert traits: %w", err)
	}

	// クラスを抽出
	class := s.extractClass(servants)
	if err := classRepo.UpsertBulk(ctx, class); err != nil {
		return fmt.Errorf("failed to upsert classes: %w", err)
	}

	// 特性から属性とアライメントを抽出
	attributes, orderAlign, moralAlign := s.extractMetaFromTraits(traits)
	if err := attrRepo.UpsertBulk(ctx, attributes); err != nil {
		return fmt.Errorf("failed to upsert attributes: %w", err)
	}
	if err := orderRepo.UpsertBulk(ctx, orderAlign); err != nil {
		return fmt.Errorf("failed to upsert order alignments: %w", err)
	}
	if err := moralRepo.UpsertBulk(ctx, moralAlign); err != nil {
		return fmt.Errorf("failed to upsert moral alignments: %w", err)
	}
	
	// プレイアブルサーヴァントを抽出
	playable, ascs := s.extractPlayable(servants)
	if err := svtRepo.UpsertBulk(ctx, playable); err != nil {
		return fmt.Errorf("failed to upsert servants: %w", err)
	}
	if err := ascRepo.UpsertBulk(ctx, ascs); err != nil {
		return fmt.Errorf("failed to upsert ascensions: %w", err)
	}

	return nil
}
