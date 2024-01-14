package controller

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
	"github.com/mhmdiamd/go-restapi-future-store/helper"
	"github.com/mhmdiamd/go-restapi-future-store/lib/repository"
	"github.com/mhmdiamd/go-restapi-future-store/lib/service"
	"github.com/mhmdiamd/go-restapi-future-store/lib/web/dto"
	"github.com/mhmdiamd/go-restapi-future-store/model/domain"
	"github.com/mhmdiamd/go-restapi-future-store/model/web"
)

type CartControllerImpl struct {
  cartService service.CartService
  productService service.ProductService
}

func NewCartControllerImpl(cartService service.CartService, productService service.ProductService) CartController {
  return &CartControllerImpl{
    cartService: cartService,
    productService: productService,
  }
}

func (c *CartControllerImpl) FindAll(w http.ResponseWriter, req *http.Request, params httprouter.Params){
  id_user := req.Header.Get("id_user")
  parsedId := uuid.MustParse(id_user)

  result := c.cartService.FindAll(req.Context(), parsedId)

  webResponse := web.WebResponse{
    Status: "Success",
    Code: http.StatusOK,
    Data: result,
  }

  helper.WriteToResponseBody(w, webResponse)
}

func (c *CartControllerImpl) Create(w http.ResponseWriter, req *http.Request, params httprouter.Params){
  id_user := req.Header.Get("id_user")
  parsedId := uuid.MustParse(id_user)

  var body repository.CartCreateBody

  helper.ReadRequestBody(req, &body)
  body.Id_user = parsedId

  result := c.cartService.Create(req.Context(), body)
  product := c.productService.FindById(req.Context(), result.Id_product)

  newProduct := domain.JSONBProduct(product)
  result.Product = newProduct 

  webResponse := web.WebResponse{
    Status: "Success",
    Code: http.StatusOK,
    Data: result,
  }

  helper.WriteToResponseBody(w, webResponse)
}

func (c *CartControllerImpl) Delete(w http.ResponseWriter, req *http.Request, params httprouter.Params){

  var id uuid.UUID
  helper.ReadRequestBody(req, &id)

  result := c.cartService.Delete(req.Context(), id)

  webResponse := web.WebResponse{
    Status: "Success",
    Code: http.StatusOK,
    Data: result,
  }

  helper.WriteToResponseBody(w, webResponse)
}

func (c *CartControllerImpl) Increment(w http.ResponseWriter, req *http.Request, params httprouter.Params){
  var data dto.IncrementCart
  helper.ReadRequestBody(req, &data)

  result := c.cartService.Increment(req.Context(), data.Id_cart)
  product := c.productService.FindById(req.Context(), result.Id_product)

  newProduct := domain.JSONBProduct(product)
  result.Product = newProduct 

  webResponse := web.WebResponse{
    Status: "Success",
    Code: http.StatusOK,
    Data: result,
  }

  helper.WriteToResponseBody(w, webResponse)
}

func (c *CartControllerImpl) Decrement(w http.ResponseWriter, req *http.Request, params httprouter.Params){
  var data dto.IncrementCart
  helper.ReadRequestBody(req, &data)

  result := c.cartService.Decrement(req.Context(), data.Id_cart)
  product := c.productService.FindById(req.Context(), result.Id_product)

  newProduct := domain.JSONBProduct(product)
  result.Product = newProduct 

  webResponse := web.WebResponse{
    Status: "Success",
    Code: http.StatusOK,
    Data: result,
  }

  helper.WriteToResponseBody(w, webResponse)
}
