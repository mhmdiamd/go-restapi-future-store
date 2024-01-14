package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/mhmdiamd/go-restapi-future-store/lib/repository"
	"github.com/mhmdiamd/go-restapi-future-store/model/domain"
)

type CartService interface {
  Create(ctx context.Context, body repository.CartCreateBody) domain.Cart
  Delete(ctx context.Context, id uuid.UUID) string
  FindAll(ctx context.Context, id_user uuid.UUID) []domain.Cart

  Increment(ctx context.Context, id uuid.UUID) domain.Cart
  Decrement(ctx context.Context, id uuid.UUID) domain.Cart
}
