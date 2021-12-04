package repository

import (
	"github.com/konu96/Nolack/internal/repository/dto"
)

type NotifyRepository interface {
	Notify(input dto.NotifyInput) error
}
