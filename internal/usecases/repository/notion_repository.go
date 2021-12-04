package repository

import (
	"github.com/konu96/Nolack/internal/domain/entity"
	"github.com/konu96/Nolack/internal/repository/dto"
)

type NotionRepository interface {
	CreatePage(page entity.Page) (*dto.CreatePageInput, *dto.CreatePageErrorResponse, error)
}
