package model

import (
	"time"

	"partners/internal/services/ebay/model"
)

type GetItemDetailsRes struct {
	ItemId           string           `json:"itemId"`
	Title            string           `json:"title"`
	ShortDescription string           `json:"shortDescription"`
	Price            model.PriceModel `json:"price"`
	CategoryPath     string           `json:"categoryPath"`
	CategoryIdPath   string           `json:"categoryIdPath"`
	Condition        string           `json:"condition"`
	ConditionId      string           `json:"conditionId"`
	ItemLocation     struct {
		City            string `json:"city"`
		StateOrProvince string `json:"stateOrProvince"`
		PostalCode      string `json:"postalCode"`
		Country         string `json:"country"`
	} `json:"itemLocation"`
	Image struct {
		ImageUrl string `json:"imageUrl"`
		Width    int    `json:"width"`
		Height   int    `json:"height"`
	} `json:"image"`
	AdditionalImages []struct {
		ImageUrl string `json:"imageUrl"`
		Width    int    `json:"width"`
		Height   int    `json:"height"`
	} `json:"additionalImages"`
	Brand            string    `json:"brand"`
	ItemCreationDate time.Time `json:"itemCreationDate"`
	Seller           struct {
		Username           string `json:"username"`
		FeedbackPercentage string `json:"feedbackPercentage"`
		FeedbackScore      int    `json:"feedbackScore"`
		SellerLegalInfo    struct {
		} `json:"sellerLegalInfo"`
	} `json:"seller"`
	EstimatedAvailabilities []struct {
		DeliveryOptions             []string `json:"deliveryOptions"`
		EstimatedAvailabilityStatus string   `json:"estimatedAvailabilityStatus"`
		EstimatedAvailableQuantity  int      `json:"estimatedAvailableQuantity"`
		EstimatedSoldQuantity       int      `json:"estimatedSoldQuantity"`
		EstimatedRemainingQuantity  int      `json:"estimatedRemainingQuantity"`
	} `json:"estimatedAvailabilities"`
	ShipToLocations struct {
		RegionIncluded []struct {
			RegionName string `json:"regionName,omitempty"`
			RegionType string `json:"regionType"`
			RegionId   string `json:"regionId"`
		} `json:"regionIncluded"`
		RegionExcluded []struct {
			RegionName string `json:"regionName"`
			RegionType string `json:"regionType"`
			RegionId   string `json:"regionId"`
		} `json:"regionExcluded"`
	} `json:"shipToLocations"`
	ReturnTerms struct {
		ReturnsAccepted bool `json:"returnsAccepted"`
	} `json:"returnTerms"`
	Taxes []struct {
		TaxJurisdiction struct {
			Region struct {
				RegionName string `json:"regionName"`
				RegionType string `json:"regionType"`
			} `json:"region"`
			TaxJurisdictionId string `json:"taxJurisdictionId"`
		} `json:"taxJurisdiction"`
		TaxType                  string `json:"taxType"`
		ShippingAndHandlingTaxed bool   `json:"shippingAndHandlingTaxed"`
		IncludedInPrice          bool   `json:"includedInPrice"`
		EbayCollectAndRemitTax   bool   `json:"ebayCollectAndRemitTax"`
	} `json:"taxes"`
	LocalizedAspects []struct {
		Type  string `json:"type"`
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"localizedAspects"`
	TopRatedBuyingExperience bool     `json:"topRatedBuyingExperience"`
	BuyingOptions            []string `json:"buyingOptions"`
	ItemWebUrl               string   `json:"itemWebUrl"`
	ItemAffiliateWebUrl      string   `json:"itemAffiliateWebUrl"`
	Description              string   `json:"description"`
	PaymentMethods           []struct {
		PaymentMethodType   string `json:"paymentMethodType"`
		PaymentMethodBrands []struct {
			PaymentMethodBrandType string `json:"paymentMethodBrandType"`
		} `json:"paymentMethodBrands,omitempty"`
	} `json:"paymentMethods"`
	EnabledForGuestCheckout   bool   `json:"enabledForGuestCheckout"`
	EligibleForInlineCheckout bool   `json:"eligibleForInlineCheckout"`
	LotSize                   int    `json:"lotSize"`
	LegacyItemId              string `json:"legacyItemId"`
	PriorityListing           bool   `json:"priorityListing"`
	AdultOnly                 bool   `json:"adultOnly"`
	CategoryId                string `json:"categoryId"`
	ListingMarketplaceId      string `json:"listingMarketplaceId"`
}

func (r GetItemDetailsRes) ConvertToBusinessModel() model.ItemDetails {

	images := make([]string, 0, len(r.AdditionalImages)+1)
	images = append(images, r.Image.ImageUrl)
	for _, image := range r.AdditionalImages {
		images = append(images, image.ImageUrl)
	}

	return model.ItemDetails{
		ID:          r.ItemId,
		Title:       r.Title,
		Description: r.Description,
		Price: model.PriceModel{
			Value:    r.Price.Value,
			Currency: r.Price.Currency,
		},
		Images:     images,
		Brand:      r.Brand,
		ItemWebUrl: r.ItemAffiliateWebUrl,
	}
}
