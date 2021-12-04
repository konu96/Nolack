package repository

import "github.com/konu96/Nolack/internal/usecases/dto"

type NotionRepository interface {
	POST(page dto.PostRequest) (*dto.PostResponse, *dto.PostErrorResponse, error)
}
