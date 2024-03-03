package interfaces

import (
	"context"
	"crud/internal/models"
)

// @tg http-prefix=api/v1
// @tg http-server log trace metrics disableExchange
type Crud interface {
	// @tg summary=`create`
	// @tg http-method=POST
	// @tg http-path=/users
	Create(ctx context.Context, req models.UserRequest) (resp models.UserResponse, err error)
	GetUser(ctx context.Context) (err error)
	GetUsers(ctx context.Context) (err error)
	Update(ctx context.Context) (err error)
	Delete(ctx context.Context) (err error)
}
