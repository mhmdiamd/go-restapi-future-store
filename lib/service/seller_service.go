package service

import (
	"context"

	"github.com/mhmdiamd/go-restapi-future-store/lib/web/dto"
	"github.com/mhmdiamd/go-restapi-future-store/model/domain"
)

type SellerService interface {
  Update(ctx context.Context, body dto.SellerUpdateDTO) domain.Seller

  // Upload Photo Seller
  UploadPhoto(ctx context.Context, body dto.UploadSellerPhotoDTO) string 
}
