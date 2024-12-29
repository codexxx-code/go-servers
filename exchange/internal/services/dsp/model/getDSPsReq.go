package model

import "exchange/internal/enum/sourceTrafficType"

type GetDSPsReq struct {
	Slugs             []string
	SourceTrafficType []sourceTrafficType.SourceTrafficType
	IsEnable          *bool
}
