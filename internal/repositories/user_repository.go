package repositories

import (
	"context"
	"my_go_send_email_app_smtp/internal/models"

	"gorm.io/gorm"
)

// UserService kullanıcı işlemleri için bir servis
type repository struct {
	db *gorm.DB
}

type Repository interface {
	GetUserByEmail(ctx context.Context, email string) (models.User, error)
	CreateUser(ctx context.Context, user *models.User) error
}

// NewUserService UserService tipinde bir örnek döndürür
func NewUserRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	var user models.User
	err := r.db.WithContext(ctx).Model(&models.User{}).Where("email = ?", email).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) CreateUser(ctx context.Context, user *models.User) error {
	return r.db.WithContext(ctx).Model(&models.User{}).Create(&user).Error
}
