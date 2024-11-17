package model

import (
	"time"

	"partners/internal/services/ebay/model"
)

type GetItemsSummaryRes struct {
	Href          string `json:"href"`
	Total         int    `json:"total"`
	Next          string `json:"next"`
	Limit         int    `json:"limit"`
	Offset        int    `json:"offset"`
	ItemSummaries []Item `json:"itemSummaries"`
}

func (r GetItemsSummaryRes) ConvertToBusinessModel() []model.ItemSummary {
	items := make([]model.ItemSummary, 0, len(r.ItemSummaries))
	for _, item := range r.ItemSummaries {
		items = append(items, item.ConvertToBusinessModel())
	}
	return items
}

type Item struct {
	ItemId          string   `json:"itemId"`
	Title           string   `json:"title"`
	LeafCategoryIds []string `json:"leafCategoryIds"`
	Categories      []struct {
		CategoryId   string `json:"categoryId"`
		CategoryName string `json:"categoryName"`
	} `json:"categories"`
	Image struct {
		ImageUrl string `json:"imageUrl"`
	} `json:"image"`
	Price struct {
		Value    string `json:"value"`
		Currency string `json:"currency"`
	} `json:"price"`
	ItemHref string `json:"itemHref"`
	Seller   struct {
		Username           string `json:"username"`
		FeedbackPercentage string `json:"feedbackPercentage"`
		FeedbackScore      int    `json:"feedbackScore"`
	} `json:"seller"`
	MarketingPrice  *MarketingPrice `json:"marketingPrice"`
	ThumbnailImages []struct {
		ImageUrl string `json:"imageUrl"`
	} `json:"thumbnailImages"`
	ShippingOptions []struct {
		ShippingCostType string `json:"shippingCostType"`
		ShippingCost     struct {
			Value    string `json:"value"`
			Currency string `json:"currency"`
		} `json:"shippingCost"`
	} `json:"shippingOptions"`
	BuyingOptions       []string `json:"buyingOptions"`
	ItemWebUrl          string   `json:"itemWebUrl"`
	ItemAffiliateWebUrl string   `json:"itemAffiliateWebUrl"`
	ItemLocation        struct {
		PostalCode string `json:"postalCode"`
		Country    string `json:"country"`
	} `json:"itemLocation"`
	AdditionalImages []struct {
		ImageUrl string `json:"imageUrl"`
	} `json:"additionalImages"`
	AdultOnly                bool      `json:"adultOnly"`
	LegacyItemId             string    `json:"legacyItemId"`
	AvailableCoupons         bool      `json:"availableCoupons"`
	ItemCreationDate         time.Time `json:"itemCreationDate"`
	TopRatedBuyingExperience bool      `json:"topRatedBuyingExperience"`
	PriorityListing          bool      `json:"priorityListing"`
	ListingMarketplaceId     string    `json:"listingMarketplaceId"`
	Condition                string    `json:"condition"`
	ConditionId              string    `json:"conditionId"`
	PickupOptions            []struct {
		PickupLocationType string `json:"pickupLocationType"`
	} `json:"pickupOptions"`
	ItemGroupHref   string `json:"itemGroupHref"`
	ItemGroupType   string `json:"itemGroupType"`
	BidCount        int    `json:"bidCount"`
	CurrentBidPrice struct {
		Value    string `json:"value"`
		Currency string `json:"currency"`
	} `json:"currentBidPrice"`
	ItemEndDate time.Time `json:"itemEndDate"`
	Epid        string    `json:"epid"`
}

func (i *Item) ConvertToBusinessModel() model.ItemSummary {

	images := make([]string, 0, len(i.AdditionalImages)+1)
	images = append(images, i.Image.ImageUrl)
	for _, image := range i.AdditionalImages {
		images = append(images, image.ImageUrl)
	}

	return model.ItemSummary{
		ID:     i.ItemId,
		Title:  i.Title,
		Images: images,
		Price: model.PriceModel{
			Value:    i.Price.Value,
			Currency: i.Price.Currency,
		},
		MarketingPrice: i.MarketingPrice.ConvertToBusinessModel(),
		ItemWebURL:     i.ItemAffiliateWebUrl,
	}
}

type MarketingPrice struct {
	OriginalPrice struct {
		Value    string `json:"value"`
		Currency string `json:"currency"`
	} `json:"originalPrice"`
	DiscountPercentage string `json:"discountPercentage"`
	DiscountAmount     struct {
		Value    string `json:"value"`
		Currency string `json:"currency"`
	} `json:"discountAmount"`
	PriceTreatment string `json:"priceTreatment"`
}

func (m *MarketingPrice) ConvertToBusinessModel() *model.MarketingPrice {

	if m == nil {
		return nil
	}

	return &model.MarketingPrice{
		OriginalPrice: model.PriceModel{
			Value:    m.OriginalPrice.Value,
			Currency: m.OriginalPrice.Currency,
		},
		DiscountPercentage: m.DiscountPercentage,
		DiscountAmount: model.PriceModel{
			Value:    m.DiscountAmount.Value,
			Currency: m.DiscountAmount.Currency,
		},
	}
}
