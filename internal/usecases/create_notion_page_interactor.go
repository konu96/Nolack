package usecases

import (
	"encoding/json"
	"fmt"
	"github.com/konu96/Nolack/internal/external/slack"
	"github.com/konu96/Nolack/internal/usecases/dto"
)

const pageID = "263a6b171e8049acbecb821b492bfad3"

type NotionInterface interface {
	POST(data []byte) (*dto.PostResponse, error)
}

type CreatePageInteractor struct {
	Slack  slack.Slack
	Notion NotionInterface
}

func NewCreatePageInteractor(slack slack.Slack, notion NotionInterface) CreatePageInteractor {
	return CreatePageInteractor{
		slack,
		notion,
	}
}

func (i *CreatePageInteractor) Exec(channel string) error {
	page := dto.PostRequest{
		Parent: dto.Parent{
			PageID: pageID,
		},
		Cover: dto.Cover{
			Type:     "external",
			External: dto.External{URL: "https://d3bhdfps5qyllw.cloudfront.net/org/63/63516e4f15e183b8925052964a58f077_1080x700_w.jpg"},
		},
		Properties: dto.Properties{
			Title: dto.Title{Title: []dto.TitleInTitle{
				{
					Type: "text",
					Text: dto.TitleContent{
						Content: "Sample page created by nolack",
					},
				},
			}},
		},
	}

	marshaledJSON, err := json.Marshal(page)
	if err != nil {
		return fmt.Errorf("failed to marshal json: %w", err)
	}

	if _, err := i.Notion.POST(marshaledJSON); err != nil {
		return fmt.Errorf("failed to missing post: %w", err)
	}

	return nil
}
