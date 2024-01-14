package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/mhmdiamd/go-restapi-future-store/lib/web/dto"
	"github.com/mhmdiamd/go-restapi-future-store/model/domain"
)

type ShippingAddressService interface {
  Create(ctx context.Context, body dto.CreateShippingAddressDTO) domain.ShippingAddress
  Update(ctx context.Context, body dto.UpdateShippingAddressDTO) domain.ShippingAddress
  Delete(ctx context.Context, id uuid.UUID) string
  FindById(ctx context.Context, id uuid.UUID) domain.ShippingAddress
  FindAll(ctx context.Context) []domain.ShippingAddress
  FindAllByIdUser(ctx context.Context, id_user uuid.UUID) []domain.ShippingAddress

  Activate(ctx context.Context, id uuid.UUID) domain.ShippingAddress
}
