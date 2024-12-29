package service

import (
	dspModel "exchange/internal/services/dsp/model"
	"exchange/internal/services/exchange/model"
	"exchange/internal/enum/sourceTrafficType"
	"pkg/errors"
	"pkg/pointer"
)

func (s *ExchangeService) getDSPs(dto *model.AuctionDTO) (dsps []dspModel.DSP, err error) {

	if len(dto.SSP.DSPs) == 0 {
		return dsps, errors.NotFound.New("У SSP не включена ни одна DSP", errors.ParamsOption(
			"ssp", dto.SSP.Slug,
		))
	}

	// Получаем список DSP по фильтрам
	if dsps, err = s.dspService.GetDSPs(dto.Ctx, dspModel.GetDSPsReq{ //nolint:exhaustruct
		Slugs:             dto.SSP.DSPs,
		IsEnable:          pointer.Pointer(true),
		SourceTrafficType: []sourceTrafficType.SourceTrafficType{dto.SourceTrafficType},
	}); err != nil {
		return dsps, err
	}

	// Проверяем количество DSP
	if len(dsps) == 0 {
		return dsps, errors.NotFound.New("Не найдены DSP, подходящие для этого запроса")
	}

	return dsps, nil
}
