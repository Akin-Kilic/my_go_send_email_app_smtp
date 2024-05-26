package services

import (
	"context"
	"my_go_send_email_app_smtp/internal/models"
	"my_go_send_email_app_smtp/internal/repositories"
)

// UserService kullanıcı işlemleri için bir servis
type service struct {
	repo repositories.Repository
}

type Service interface {
	SaveUser(ctx context.Context, user *models.User) error
	GetUserByEmail(ctx context.Context, email string) (models.User, error)
}

// NewUserService UserService tipinde bir örnek döndürür
func NewUserService(db repositories.Repository) Service {
	return &service{repo: db}
}

// SaveUser yeni bir kullanıcıyı veritabanına kaydeder
func (s *service) SaveUser(ctx context.Context, user *models.User) error {
	return s.repo.CreateUser(ctx, user)
}

// GetUserByEmail e-posta adresine göre bir kullanıcı getirir
func (s *service) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	user, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return user, err
	}
	return user, nil
}
