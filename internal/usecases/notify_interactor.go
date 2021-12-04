package usecases

import (
	"fmt"
	notifydto "github.com/konu96/Nolack/internal/repository/dto/notify"
	"github.com/konu96/Nolack/internal/usecases/repository"
)

type NotifyInteractor struct {
	NotifyRepository repository.NotifyRepository
}

func NewNotifyInteractor(notifyRepository repository.NotifyRepository) *NotifyInteractor {
	return &NotifyInteractor{
		NotifyRepository: notifyRepository,
	}
}

func (i *NotifyInteractor) Exec(input notifydto.NotifyInput) error {
	if err := i.NotifyRepository.Notify(input); err != nil {
		return fmt.Errorf("error notify: %w", err)
	}

	return nil
}
