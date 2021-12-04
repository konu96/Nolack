package repository

import (
	"encoding/json"
	"fmt"
	"github.com/konu96/Nolack/internal/domain/entity"
	"github.com/konu96/Nolack/internal/repository/dto"
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

func (r *NotionRepository) CreatePage(page entity.Page) (*dto.CreatePageResponse, *dto.CreatePageErrorResponse, error) {
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
		var errorResponse dto.CreatePageErrorResponse
		if err := json.Unmarshal(body, &errorResponse); err != nil {
			return nil, nil, err
		}

		return nil, &errorResponse, err
	}

	var postResponse dto.CreatePageResponse
	if err := json.Unmarshal(body, &postResponse); err != nil {
		return nil, nil, err
	}

	return &postResponse, nil, nil
}
