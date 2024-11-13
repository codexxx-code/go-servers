package service

import (
	"context"

	"templater/internal/services/templater/model"
	templaterRepository "templater/internal/services/templater/repository"
)

type TemplaterService struct {
	templaterRepository TemplaterRepository
}

var _ TemplaterRepository = new(templaterRepository.TemplaterRepository)

type TemplaterRepository interface {
	GetTemplate(context.Context, model.GetTemplateReq) (model.GetTemplateRes, error)
}

func NewTemplaterService(
	templaterRepository TemplaterRepository,
) *TemplaterService {
	return &TemplaterService{
		templaterRepository: templaterRepository,
	}
}
