package usecases

import (
	"fmt"
	"github.com/konu96/Nolack/internal/repository"
	"github.com/konu96/Nolack/internal/repository/dto"
)

type NotifyUseCase interface {
	Exec(input dto.NotifyInput) error
}

type NotifyInteractor struct {
	NotifyRepository *repository.NotifyRepository
}

func NewNotifyInteractor(client repository.NotifyClient) *NotifyInteractor {
	return &NotifyInteractor{
		NotifyRepository: repository.NewNotifyRepository(client),
	}
}

func (i *NotifyInteractor) Exec(input dto.NotifyInput) error {
	if err := i.NotifyRepository.Notify(input); err != nil {
		return fmt.Errorf("notify: %w", err)
	}

	return nil
}
