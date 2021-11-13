package usecases

import (
	"encoding/json"
	"github.com/konu96/Nolack/internal/external/slack"
	"github.com/konu96/Nolack/internal/usecases/dto"
	"log"
	"net/http"
)

const pageID = "d364bd662bb0425480c899699de4e4cb"

type NotionInterface interface {
	POST(data []byte) (dto.PostResponse, error)
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

func (i *CreatePageInteractor) Exec(w http.ResponseWriter, channel string) {
	page := dto.PostRequest{
		Parent: dto.Parent{
			Type:   "page_id",
			PageID: pageID,
		},
		Cover: dto.Cover{
			Type:     "external",
			External: dto.External{URL: "https://d3bhdfps5qyllw.cloudfront.net/org/63/63516e4f15e183b8925052964a58f077_1080x700_w.jpg"},
		},
	}

	aaa, err := json.Marshal(page)
	if err != nil {
		log.Printf("can not marshal json: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if _, err := i.Notion.POST(aaa); err != nil {
		log.Printf("can not post: %s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
