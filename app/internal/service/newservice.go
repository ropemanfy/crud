package service

import (
	"context"
	"crud/internal/errors"
	"crud/internal/models"
)

type svc struct {
	storage Storage
}

func NewSvc(storage Storage) *svc {
	return &svc{storage: storage}
}
func (s *svc) Create(ctx context.Context, req models.UserRequest) (resp models.UserResponse, err error) {
	id, err := s.storage.Create(models.User{
		Name:  req.Name,
		Email: req.Email,
	})
	if err != nil {
		err = errors.ErrInternal.SetStatus(500).SetCause(err.Error())
		return
	}
	resp.ID = id
	return
}
func (s *svc) GetUser(ctx context.Context) (err error) {
	return
}
func (s *svc) GetUsers(ctx context.Context) (err error) {
	return
}
func (s *svc) Update(ctx context.Context) (err error) {
	return
}
func (s *svc) Delete(ctx context.Context) (err error) {
	return
}
