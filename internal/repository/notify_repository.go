package repository

import (
	"fmt"
	notifydto "github.com/konu96/Nolack/internal/repository/dto"
)

type NotifyClient interface {
	Notify(input notifydto.NotifyInput) error
}

type NotifyRepository struct {
	Client NotifyClient
}

func NewNotifyRepository(client NotifyClient) *NotifyRepository {
	return &NotifyRepository{
		Client: client,
	}
}

func (r *NotifyRepository) Notify(input notifydto.NotifyInput) error {
	if err := r.Client.Notify(input); err != nil {
		return fmt.Errorf("notify: %w", err)
	}

	return nil
}
