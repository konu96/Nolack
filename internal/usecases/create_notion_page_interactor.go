package usecases

import (
	"fmt"
	"github.com/konu96/Nolack/internal/domain/entity"
	repositorydto "github.com/konu96/Nolack/internal/repository/dto"
	usecasesdto "github.com/konu96/Nolack/internal/usecases/dto"
	"github.com/konu96/Nolack/internal/usecases/repository"
)

type CreateNotionPageUseCase interface {
	Exec(channel string, input usecasesdto.CreatePageInput) error
}

type CreateNotionPageInteractor struct {
	NotionRepository repository.NotionRepository
	NotifyInteractor *NotifyInteractor
}

func NewCreatePageInteractor(
	notionRepository repository.NotionRepository,
	notifyInteractor *NotifyInteractor,
) *CreateNotionPageInteractor {
	return &CreateNotionPageInteractor{
		NotionRepository: notionRepository,
		NotifyInteractor: notifyInteractor,
	}
}

func (i *CreateNotionPageInteractor) Exec(channel string, input usecasesdto.CreatePageInput) error {
	page := entity.NewPage(input.PageID, input.URL)

	if _, _, err := i.NotionRepository.CreatePage(page); err != nil {
		if err := i.NotifyInteractor.Exec(repositorydto.NotifyInput{
			Channel: channel,
			Text:    err.Error(),
		}); err != nil {
			return fmt.Errorf("missing error notify: %w", err)
		}
		return fmt.Errorf("missing create page: %w", err)
	}

	if err := i.NotifyInteractor.Exec(repositorydto.NotifyInput{
		Channel: channel,
		Text:    "ページを作成しました",
	}); err != nil {
		return fmt.Errorf("missing success notify: %w", err)
	}

	return nil
}
