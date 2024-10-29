package service

import (
	ebayNetwork "partners/internal/services/ebay/network"
)

type EbayService struct {
	ebayNetwork EbayNetwork
}

var _ EbayNetwork = new(ebayNetwork.EbayNetwork)

type EbayNetwork interface{}

func NewEbayService(
	ebayNetwork EbayNetwork,
) *EbayService {
	return &EbayService{
		ebayNetwork: ebayNetwork,
	}
}
