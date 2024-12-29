package beforeRequestToDSP

type reduceTMax struct {
	baseLink
}

func (r *reduceTMax) Apply(dto *beforeRequestToDSP) error {

	// Изменяем максимальное время ответа на коэффициент, заданный в настройках системы // TODO: перенести настройку в настройки SSP/DSP
	dto.BidRequest.TimeMax = int(float64(dto.RequestTimeout.Milliseconds()) * (1 - dto.Settings.ReduceTimeoutCoef))

	return nil
}
