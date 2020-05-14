package repositories

import (
	"context"
	"golang_api/models"
)

type UserRepo interface {
	Add(c context.Context, user models.User) (models.User, error)
	GetList(c context.Context) ([]models.User, error)
	GetUserById(c context.Context, id string) (models.User, error)
	Update(c context.Context, user models.User) (models.User, error)
	Delete(c context.Context, id string) (models.User, error)
}
