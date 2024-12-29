package model

import (
	"fmt"
	"strings"

	"exchange/internal/enum"
	"exchange/internal/enum/billingType"
	"exchange/internal/enum/currency"
	"exchange/internal/enum/formatType"
	"exchange/internal/enum/integrationType"
	"exchange/internal/enum/sourceTrafficType"
	"exchange/internal/services/ssp/model/fraudScore"
	"pkg/errors"
)

type UpdateSSPReq struct {
	Slug                string `validate:"required"`
	Name                string `validate:"required"`
	Timeout             *int32
	IsEnable            bool
	IntegrationType     integrationType.IntegrationType       `validate:"required"`
	SourceTrafficTypes  []sourceTrafficType.SourceTrafficType `validate:"required"`
	BillingType         billingType.BillingType               `validate:"required"`
	AuctionSecondPrice  bool
	Currency            currency.Currency       `validate:"required"`
	FormatTypes         []formatType.FormatType `validate:"required"`
	FraudScore          fraudScore.FraudScore   `validate:"required"`
	ClickunderADMFormat *string
	ClickunderDrumSize  *int32
}

// Валидация входных данных
func (r UpdateSSPReq) Validate() error { //nolint:dupl
	if err := r.IntegrationType.Validate(); err != nil {
		return err
	}
	for _, sourceTrafficType := range r.SourceTrafficTypes {
		if err := sourceTrafficType.Validate(); err != nil {
			return err
		}
	}
	if err := r.BillingType.Validate(); err != nil {
		return err
	}
	if err := r.Currency.Validate(); err != nil {
		return err
	}
	if err := r.FraudScore.Validate(); err != nil {
		return err
	}

	// Флаг для отслеживания формата с поддержкой кликандера
	hasClickunder := false

	// Проверка форматов и выход при первом кликандер-формате
	for _, _formatType := range r.FormatTypes {

		// Валидируем формат
		if err := _formatType.Validate(); err != nil {
			return err
		}

		// Если передан кликандер формат, то устанавливаем флаг
		if _formatType == formatType.Clickunder {
			hasClickunder = true
		}
	}

	// Валидация полей кликандера в зависимости от форматов
	if hasClickunder {
		if r.ClickunderDrumSize == nil || r.ClickunderADMFormat == nil {
			return errors.BadRequest.New("ClickunderDrumSize and ClickunderADMFormat are required for clickunder-enabled formats")
		}

		// Проверка, что в clickunderADMFormat есть макрос
		if !strings.Contains(*r.ClickunderADMFormat, enum.ADMURLMacros) {
			return errors.BadRequest.New(fmt.Sprintf("ClickunderADMFormat must contains macros: %s", enum.ADMURLMacros))
		}
	} else if r.ClickunderDrumSize != nil || r.ClickunderADMFormat != nil {
		return errors.BadRequest.New("ClickunderDrumSize and ClickunderADMFormat must be nil for non-clickunder formats")
	}

	return nil
}
