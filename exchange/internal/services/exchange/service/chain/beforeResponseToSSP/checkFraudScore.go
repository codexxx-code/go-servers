package beforeResponseToSSP

import (
	"time"

	fraudScoreModel "exchange/internal/services/fraudScore/model"
	"exchange/internal/services/fraudScore/service"
	"exchange/internal/services/ssp/model/fraudScore"
	"pkg/errors"
)

// Структура для проверки фрауда с использованием Fraudscore
type FraudChecker struct {
	IsEnabled  bool
	Fraudscore *service.FraudScoreService
}

type checkFraudScore struct {
	baseLink
	fraudScoreService FraudScoreService
}

func (r *checkFraudScore) Apply(dto *beforeResponseToSSP) (err error) {
	if dto.OriginalBidRequest.Device == nil {
		return errors.BadRequest.New("Device object not found")
	}

	if dto.SSP.FraudScore == fraudScore.PreBid {
		isFraud, err := r.fraudScoreService.IsFraud(dto.Ctx, fraudScoreModel.IsFraudReq{
			Event:     "impression",
			IP:        dto.OriginalBidRequest.Device.IP,
			UserAgent: dto.OriginalBidRequest.Device.UserAgent,
			AT:        time.Now().Unix(),
		})
		if err != nil {
			return err
		}
		if isFraud {
			return errors.BadRequest.New("traffic is fraud")
		}
	}
	return nil
}
