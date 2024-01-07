package exception

type ErrorForbidden struct {
  Error string
}

func NewForbiddenError(err string) ErrorForbidden {
  return ErrorForbidden{
    Error: err,
  }
}

