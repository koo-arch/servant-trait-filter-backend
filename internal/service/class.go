package service

import (
	"context"

	"github.com/koo-arch/servant-trait-filter-backend/ent"
	"github.com/koo-arch/servant-trait-filter-backend/internal/repository"
	"github.com/koo-arch/servant-trait-filter-backend/internal/util"
)

type ClassService interface {
	GetAllClasses(ctx context.Context) ([]ClassDTO, error)
}

type ClassServiceImpl struct {
	classRepo repository.ClassRepository
}

func NewClassServiceImpl(classRepo repository.ClassRepository) ClassService {
	return &ClassServiceImpl{
		classRepo: classRepo,
	}
}

func (s *ClassServiceImpl) GetAllClasses(ctx context.Context) ([]ClassDTO, error) {
	// データベースからクラスを取得
	classes, err := s.classRepo.ListAll(ctx)
	if err != nil {
		return nil, err
	}

	// クラスをDTOに変換
	dtos := util.ConvertSlice(classes, func(class *ent.Class) ClassDTO {
		return ClassDTO{
			ID:   class.ID,
			Name: util.FallbackName(class.NameJa, class.NameEn),
		}
	})

	return dtos, nil
}