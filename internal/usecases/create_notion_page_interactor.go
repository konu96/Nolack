package usecases

import (
	"fmt"
	"github.com/konu96/Nolack/internal/domain/entity"
	"github.com/konu96/Nolack/internal/external/slack"
	"github.com/konu96/Nolack/internal/usecases/repository"
)

const pageID = "263a6b171e8049acbecb821b492bfad3"

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

func (i *CreateNotionPageInteractor) Exec(channel string) error {
	page := entity.NewPage(pageID, "https://d3bhdfps5qyllw.cloudfront.net/org/63/63516e4f15e183b8925052964a58f077_1080x700_w.jpg")

	if _, _, err := i.NotionRepository.CreatePage(page); err != nil {
		return fmt.Errorf("create page: %w", err)
	}

	return nil
}
