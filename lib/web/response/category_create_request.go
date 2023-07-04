package response

type CategoryCreateRequest struct {
	Name string `validate:"required"`
}
