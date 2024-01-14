package dto

import "github.com/google/uuid"

type CreateShippingAddressDTO struct {
  Id_user uuid.UUID 
  Place string 
  Recipient_name string 
  Recipient_phone string 
  Address string 
  Postal_code string 
  City_or_subdistrict string 
  Is_active bool 
}

type UpdateShippingAddressDTO struct {
  Id uuid.UUID 
  Place string 
  Recipient_name string 
  Recipient_phone string 
  Address string 
  Postal_code string 
  City_or_subdistrict string 
  Is_active bool 
}

type ActivateShippingAddressDTO struct {
  Id uuid.UUID 
}
