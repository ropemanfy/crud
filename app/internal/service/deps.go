package service

import "crud/internal/models"

type Storage interface {
	Create(u models.User) (string, error)
	List() ([]models.User, error)
	GetUser(ID string) ([]models.User, error)
	Update(ID, Email, Name string) error
	Delete(ID string) error
}
