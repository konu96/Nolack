package repository

import notifydto "github.com/konu96/Nolack/internal/repository/dto/notify"

type NotifyRepository interface {
	Notify(input notifydto.NotifyInput) error
}
