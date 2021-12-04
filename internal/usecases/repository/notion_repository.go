package repository

import (
	"github.com/konu96/Nolack/internal/domain/entity"
	notiondto "github.com/konu96/Nolack/internal/repository/dto/notion"
)

type NotionRepository interface {
	CreatePage(page entity.Page) (*notiondto.CreatePageResponse, *notiondto.CreatePageErrorResponse, error)
}
