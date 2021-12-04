package repository

import (
	"encoding/json"
	"fmt"
	"github.com/konu96/Nolack/internal/domain/entity"
	notiondto "github.com/konu96/Nolack/internal/repository/dto/notion"
	"io/ioutil"
	"net/http"
)

type NotionInterface interface {
	POST(data []byte) (*http.Response, error)
}

type NotionRepository struct {
	Client NotionInterface
}

func NewNotionRepository(client NotionInterface) *NotionRepository {
	return &NotionRepository{
		Client: client,
	}
}

func (r *NotionRepository) CreatePage(page entity.Page) (*notiondto.CreatePageResponse, *notiondto.CreatePageErrorResponse, error) {
	marshaledJSON, err := json.Marshal(page)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal json: %w", err)
	}

	resp, err := r.Client.POST(marshaledJSON)
	if err != nil {
		return nil, nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}

	if resp.StatusCode != 200 {
		var errorResponse notiondto.CreatePageErrorResponse
		if err := json.Unmarshal(body, &errorResponse); err != nil {
			return nil, nil, err
		}

		return nil, &errorResponse, err
	}

	var postResponse notiondto.CreatePageResponse
	if err := json.Unmarshal(body, &postResponse); err != nil {
		return nil, nil, err
	}

	return &postResponse, nil, nil
}
