package service

import (
	"context"

	ebayNetwork "partners/internal/services/ebay/network"
	ebayNetworkModel "partners/internal/services/ebay/network/model"
)

type EbayService struct {
	ebayNetwork EbayNetwork
}

var _ EbayNetwork = new(ebayNetwork.EbayNetwork)

type EbayNetwork interface {
	// Categories
	GetCategoryTreeID(ctx context.Context, req ebayNetworkModel.GetCategoryTreeIDReq) (ebayNetworkModel.GetCategoryTreeIDRes, error)
	GetCategories(ctx context.Context, req ebayNetworkModel.GetCategoriesReq) (ebayNetworkModel.GetCategoriesRes, error)

	// Items
	GetItems(ctx context.Context, req ebayNetworkModel.GetItemsReq) (ebayNetworkModel.GetItemsRes, error)
}

func NewEbayService(
	ebayNetwork EbayNetwork,
) *EbayService {
	return &EbayService{
		ebayNetwork: ebayNetwork,
	}
}
