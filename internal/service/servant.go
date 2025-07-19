package service

import (
	"context"

	"github.com/koo-arch/servant-trait-filter-backend/internal/repository"
	"github.com/koo-arch/servant-trait-filter-backend/internal/search"
)

type ServantService interface {
	Search(ctx context.Context, query search.ServantSearchQuery) (SearchResponseDTO, error)
}

// ServantService provides methods to interact with the Servant entity.
type ServantServiceImpl struct {
	svtRepo repository.ServantRepository
}

// NewServantsService creates a new ServantsService.
func NewServantsServiceImpl(svtRepo repository.ServantRepository) ServantService {
	return &ServantServiceImpl{
		svtRepo: svtRepo,
	}
}

func (s *ServantServiceImpl) Search(ctx context.Context, req search.ServantSearchQuery) (SearchResponseDTO, error) {
	// データベースからServantを検索
	servants, err := s.svtRepo.Search(ctx, req)
	if err != nil {
		return SearchResponseDTO{}, err
	}

	// ServantをDTOに変換
	dtos := make([]ServantDTO, 0, len(servants.Servants))
	for _, svt := range servants.Servants {
		dto := ServantDTO{
			ID:              svt.ID,
			Name:            svt.Name,
			CollectionNo:    svt.CollectionNo,
			Face:            svt.Face,
		}
		dtos = append(dtos, dto)}
	

	return SearchResponseDTO{
		Total: servants.Total,
		Offset: req.Offset,
		Limit: req.Limit,
		Items: dtos,
	}, nil
}

