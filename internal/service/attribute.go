package service

import (
	"context"

	"github.com/koo-arch/servant-trait-filter-backend/ent"
	"github.com/koo-arch/servant-trait-filter-backend/internal/repository"
	"github.com/koo-arch/servant-trait-filter-backend/internal/util"
)

type AttributeService interface {
	GetAllAttributes(ctx context.Context) ([]AttributeDTO, error)
}

type AttributeServiceImpl struct {
	attributeRepo repository.AttributeRepository
}

func NewAttributeServiceImpl(attributeRepo repository.AttributeRepository) AttributeService {
	return &AttributeServiceImpl{
		attributeRepo: attributeRepo,
	}
}

func (s *AttributeServiceImpl) GetAllAttributes(ctx context.Context) ([]AttributeDTO, error) {
	// データベースから属性を取得
	attributes, err := s.attributeRepo.ListAll(ctx)
	if err != nil {
		return nil, err
	}

	// 属性をDTOに変換
	dtos := util.ConvertSlice(attributes, func(attr *ent.Attribute) AttributeDTO {
		return AttributeDTO{
			ID:   attr.ID,
			Name: util.FallbackName(attr.NameJa, attr.NameEn),
		}
	})

	return dtos, nil
}