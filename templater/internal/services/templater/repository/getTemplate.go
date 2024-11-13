package service

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"pkg/ddlHelper"
	"templater/internal/services/templater/model"
	"templater/internal/services/templater/repository/templateDDL"
)

func (r *TemplaterRepository) GetTemplate(ctx context.Context, req model.GetTemplateReq) (res model.GetTemplateRes, err error) {

	filters := make(sq.Eq)

	if len(req.SSPSlug) != 0 {
		filters[templateDDL.ColumnSSPSlug] = req.SSPSlug
	}

	// Выполняем запрос
	return res, r.sql.Get(ctx, &res, sq.
		Select(ddlHelper.SelectAll).
		From(templateDDL.Table).
		Where(filters),
	)
}
