package service

import (
	"context"

	sspModel "exchange/internal/services/ssp/model"
	"pkg/errors"
	"pkg/openrtb"
	"pkg/slices"
)

func (s *ExchangeService) getSSP(ctx context.Context, sspSlug string, _ openrtb.BidRequest) (ssp sspModel.SSP, err error) {

	// Получаем SSP по slug
	ssp, err = slices.FirstWithError(
		s.sspService.GetSSPs(ctx, sspModel.GetSSPsReq{ //nolint:exhaustruct
			Slugs: []string{sspSlug},
		}),
	)
	if err != nil {
		return ssp, errors.NotFound.Wrap(err,
			errors.ParamsOption("Incoming SSPSlug", sspSlug),
			errors.LogAsOption(errors.LogNone),
		)
	}

	return ssp, nil
}
