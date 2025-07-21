package service

import (
	"context"

	"github.com/koo-arch/servant-trait-filter-backend/ent"
	"github.com/koo-arch/servant-trait-filter-backend/internal/repository"
	"github.com/koo-arch/servant-trait-filter-backend/internal/util"
)

type TraitService interface {
	GetAllTraits(ctx context.Context) ([]TraitDTO, error)
}

type TraitServiceImpl struct {
	traitRepo repository.TraitRepository
}

func NewTraitServiceImpl(traitRepo repository.TraitRepository) TraitService {
	return &TraitServiceImpl{
		traitRepo: traitRepo,
	}
}

func (s *TraitServiceImpl) GetAllTraits(ctx context.Context) ([]TraitDTO, error) {
	// データベースからTraitを取得
	traits, err := s.traitRepo.ListAll(ctx)
	if err != nil {
		return nil, err
	}

	// TraitをDTOに変換
	dtos := util.ConvertSlice(traits, func(trait *ent.Trait) TraitDTO {
		return TraitDTO{
			ID:   trait.ID,
			Name: util.FallbackName(trait.NameJa, trait.NameEn),
		}
	})

	return dtos, nil
}