package service

import (
	"context"

	"github.com/koo-arch/servant-trait-filter-backend/ent"
	"github.com/koo-arch/servant-trait-filter-backend/internal/repository"
	"github.com/koo-arch/servant-trait-filter-backend/internal/util"
)

type OrderAlignmentService interface {
	GetAllOrderAlignments(ctx context.Context) ([]OrderAlignmentDTO, error)
}

type OrderAlignmentServiceImpl struct {
	orderAlignmentRepo repository.OrderAlignmentRepository
}

func NewOrderAlignmentServiceImpl(orderAlignmentRepo repository.OrderAlignmentRepository) OrderAlignmentService {
	return &OrderAlignmentServiceImpl{
		orderAlignmentRepo: orderAlignmentRepo,
	}
}

func (s *OrderAlignmentServiceImpl) GetAllOrderAlignments(ctx context.Context) ([]OrderAlignmentDTO, error) {
	// データベースからOrderAlignmentを取得
	orderAlignments, err := s.orderAlignmentRepo.ListAll(ctx)
	if err != nil {
		return nil, err
	}

	// OrderAlignmentをDTOに変換
	dtos := util.ConvertSlice(orderAlignments, func(oa *ent.OrderAlignment) OrderAlignmentDTO {
		return OrderAlignmentDTO{
			ID:   oa.ID,
			Name: util.FallbackName(oa.NameJa, oa.NameEn),
		}
	})

	return dtos, nil
}