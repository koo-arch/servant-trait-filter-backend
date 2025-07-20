package service

import (
	"context"

	"github.com/koo-arch/servant-trait-filter-backend/ent"
	"github.com/koo-arch/servant-trait-filter-backend/internal/repository"
	"github.com/koo-arch/servant-trait-filter-backend/internal/search"
	"github.com/koo-arch/servant-trait-filter-backend/internal/util"
)

type ServantService interface {
	GetAllServants(ctx context.Context) ([]ServantDTO, error)
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

func (s *ServantServiceImpl) GetAllServants(ctx context.Context) ([]ServantDTO, error) {
	// データベースからServantを取得
	servants, err := s.svtRepo.ListAll(ctx)
	if err != nil {
		return nil, err
	}

	// ServantをDTOに変換
	dtos := util.ConvertSlice(servants, convertServantDTO)

	return dtos, nil
}

func (s *ServantServiceImpl) Search(ctx context.Context, req search.ServantSearchQuery) (SearchResponseDTO, error) {
	// データベースからServantを検索
	servants, err := s.svtRepo.Search(ctx, req)
	if err != nil {
		return SearchResponseDTO{}, err
	}

	// ServantをDTOに変換
	dtos := util.ConvertSlice(servants.Servants, convertServantDTO)

	return SearchResponseDTO{
		Total: servants.Total,
		Offset: req.Offset,
		Limit: req.Limit,
		Items: dtos,
	}, nil
}

func convertServantDTO(svt *ent.Servant) ServantDTO {
	return ServantDTO{
		ID:           svt.ID,
		CollectionNo: svt.CollectionNo,
		Name:         svt.Name,
		Face:         svt.Face,
	}
}
