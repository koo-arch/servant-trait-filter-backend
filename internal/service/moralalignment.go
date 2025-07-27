package service

import (
	"context"

	"github.com/koo-arch/servant-trait-filter-backend/ent"
	"github.com/koo-arch/servant-trait-filter-backend/internal/repository"
	"github.com/koo-arch/servant-trait-filter-backend/internal/model"
	"github.com/koo-arch/servant-trait-filter-backend/internal/util"
)

type MoralAlignmentService interface {
	GetAllMoralAlignments(ctx context.Context) ([]model.MoralAlignmentDTO, error)
}

type MoralAlignmentServiceImpl struct {
	moralAlignmentRepo repository.MoralAlignmentRepository
}

func NewMoralAlignmentServiceImpl(moralAlignmentRepo repository.MoralAlignmentRepository) MoralAlignmentService {
	return &MoralAlignmentServiceImpl{
		moralAlignmentRepo: moralAlignmentRepo,
	}
}

func (s *MoralAlignmentServiceImpl) GetAllMoralAlignments(ctx context.Context) ([]model.MoralAlignmentDTO, error) {
	// データベースからMoralAlignmentを取得
	moralAlignments, err := s.moralAlignmentRepo.ListAll(ctx)
	if err != nil {
		return nil, err
	}

	// MoralAlignmentをDTOに変換
	dtos := util.ConvertSlice(moralAlignments, func(ma *ent.MoralAlignment) model.MoralAlignmentDTO {
		return model.MoralAlignmentDTO{
			ID:   ma.ID,
			Name: util.FallbackName(ma.NameJa, ma.NameEn),
		}
	})

	return dtos, nil
}