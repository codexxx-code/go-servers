package beforeResponseToSSP

import "pkg/errors"

type checkAllSets struct {
	baseLink
}

func (r *checkAllSets) Apply(dto *beforeResponseToSSP) (err error) {

	chainSettings := dto.chainSettings

	if !chainSettings.nurlAlreadySet {
		return errors.InternalServer.New("NURL не проставлен")
	}

	if !chainSettings.admAlreadySet {
		return errors.InternalServer.New("ADM не проставлен")
	}

	if !chainSettings.priceAlreadySet {
		return errors.InternalServer.New("Цена не проставлена")
	}

	if !chainSettings.bidsAlreadyInit {
		return errors.InternalServer.New("Биды не инициализированы")
	}

	return nil
}
