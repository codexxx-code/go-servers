package model

type GetDSPResponsesReq struct {
	BidIDs     []string
	RequestIDs []string
	Limit      *int64
}
