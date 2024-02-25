package repository

import (
	"context"
	"errors"
	"github.com/gbascur/internal/adapter/entity"
	"github.com/gbascur/internal/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

func (u *userRepository) Get(ctx context.Context, id uuid.UUID) (*domain.User, error) {

	userEntity := &entity.User{
		ID: id,
	}

	result := u.DB.WithContext(ctx).First(userEntity)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, domain.ErrUserNotFound
		}
		return nil, result.Error
	}

	// se puede usar esto o tambien crear una carpeta mapper
	// y tener todas las func para mapear de entity a domain o al reves
	user := &domain.User{
		ID:   userEntity.ID,
		Name: userEntity.Name,
	}
	return user, nil
}
