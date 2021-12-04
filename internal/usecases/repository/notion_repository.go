package repository

import "github.com/konu96/Nolack/internal/usecases/dto"

type NotionRepository interface {
	CreatePage(page dto.CreatePageRequest) (*dto.CreatePageResponse, *dto.CreatePageErrorResponse, error)
}
