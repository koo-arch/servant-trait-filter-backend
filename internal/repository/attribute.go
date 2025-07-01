package repository

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/koo-arch/servant-trait-filter-backend/ent"
	"github.com/koo-arch/servant-trait-filter-backend/ent/attribute"
	"github.com/koo-arch/servant-trait-filter-backend/internal/model"
)

type AttributeRepository interface {
	ListAll(ctx context.Context) ([]*ent.Attribute, error)
	UpsertBulk(ctx context.Context, attributes []model.Attribute) error
	WithTx(tx *ent.Tx) AttributeRepository
}

type AttributeRepositoryImpl struct {
	client *ent.Client
}

func NewAttributeRepository(client *ent.Client) AttributeRepository {
	return &AttributeRepositoryImpl{
		client: client,
	}
}

func (r *AttributeRepositoryImpl) WithTx(tx *ent.Tx) AttributeRepository {
	return &AttributeRepositoryImpl{
		client: tx.Client(),
	}
}

func (r *AttributeRepositoryImpl) ListAll(ctx context.Context) ([]*ent.Attribute, error) {
	return r.client.Attribute.Query().
		Order(ent.Asc(attribute.FieldID)).
		All(ctx)
}

func (r *AttributeRepositoryImpl) UpsertBulk(ctx context.Context, attributes []model.Attribute) error {
	if len(attributes) == 0 {
		return nil
	}

	builders := make([]*ent.AttributeCreate, 0, len(attributes))
	for _, attr := range attributes {
		builder := r.client.Attribute.Create().
			SetID(attr.ID).
			SetNameEn(attr.Name)
		builders = append(builders, builder)
	}
	err := r.client.Attribute.CreateBulk(builders...).
		OnConflict(
			sql.ConflictColumns(attribute.FieldID),
			sql.UpdateWhere(
				sql.ExprP("attributes.name_en IS DISTINCT FROM EXCLUDED.name_en"),
			),
		).
		UpdateNewValues().
		Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}