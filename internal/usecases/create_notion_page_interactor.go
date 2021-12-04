package usecases

import (
	"fmt"
	"github.com/konu96/Nolack/internal/domain/entity"
	"github.com/konu96/Nolack/internal/external/slack"
	"github.com/konu96/Nolack/internal/usecases/dto"
	"github.com/konu96/Nolack/internal/usecases/repository"
)

type CreateNotionPageInteractor struct {
	Slack            *slack.Slack
	NotionRepository repository.NotionRepository
}

func NewCreatePageInteractor(slack *slack.Slack, NotionRepository repository.NotionRepository) CreateNotionPageInteractor {
	return CreateNotionPageInteractor{
		slack,
		NotionRepository,
	}
}

func (i *CreateNotionPageInteractor) Exec(channel string, input dto.CreatePageInput) error {
	page := entity.NewPage(input.PageID, input.URL)

	if _, _, err := i.NotionRepository.CreatePage(page); err != nil {
		return fmt.Errorf("create page: %w", err)
	}

	return nil
}
