package native

// ImageType defines an image asset type.
type ImageType int

const (
	ImageTypeUnknown ImageType = 0

	// Icon; Icon image.
	//
	// Optional: max height: at least 50; aspect ratio: 1:1.
	ImageTypeIcon ImageType = 1

	// Logo; Logo image for the brand/app.
	//
	// Deprecated: use ImageTypeIcon in version 1.2.
	ImageTypeLogo ImageType = 2

	// Main; Large image preview for the ad.
	//
	// At least one of 2 size variants required:
	//   Small Variant:
	//     max height: at least 200
	//     max width: at least 200, 267, or 382
	//     aspect ratio: 1:1, 4:3, or 1.91:1
	//   Large Variant:
	//     max height: at least 627
	//     max width: at least 627, 836, or 1198
	//     aspect ratio: 1:1, 4:3, or 1.91:1
	ImageTypeMain ImageType = 3

	// 500+ reserved for Exchange specific image asset types.
)

// DataType defines a data asset type.
type DataType int

const (
	DataTypeUnknown DataType = 0

	// Sponsored By message where response should contain the brand name of
	// the sponsor.
	//
	// Format: text.
	//
	// Required: max 25 or longer.
	DataTypeSponsored DataType = 1

	// Descriptive text associated with the product or service being advertised.
	// Longer length of text in response may be truncated or ellipsed by
	// the exchange.
	//
	// Format: text.
	//
	// Recommended: max 140 or longer.
	DataTypeDesc DataType = 2

	// Rating of the product being offered to the user. For example an app’s
	// rating in an app store from 0-5.
	//
	// Format: number formatted as string.
	//
	// Optional: 0-5 integer formatted as string.
	DataTypeRating DataType = 3

	// Number of social ratings or “likes” of the product being offered to
	// the user.
	//
	// Format: number formatted as string.
	DataTypeLikes DataType = 4

	// Number downloads/installs of this product.
	//
	// Format: number formatted as string.
	DataTypeDownloads DataType = 5

	// Price for product / app / in-app purchase. Value should include currency
	// symbol in localised format.
	//
	// Format: number formatted as string.
	DataTypePrice DataType = 6

	// Sale price that can be used together with price to indicate a discounted
	// price compared to a regular price. Value should include currency symbol
	// in localised format.
	//
	// Format: number formatted as string.
	DataTypeSalePrice DataType = 7

	// Phone number formatted.
	//
	// Format: string.
	DataTypePhone DataType = 8

	// Address.
	//
	// Format: text.
	DataTypeAddress DataType = 9

	// Additional descriptive text associated with the product or service being
	// advertised.
	//
	// Format: text.
	DataTypeDesc2 DataType = 10

	// Display URL for the text ad. To be used when sponsoring entity doesn’t
	// own the content. IE sponsored by BRAND on SITE (where SITE is
	// transmitted in this field).
	//
	// Format: text.
	DataTypeDispayURL DataType = 11

	// CTA description - descriptive text describing a ‘call to action’ button
	// for the destination URL.
	//
	// Format: text
	//
	// Optional: max 15 or longer.
	DataTypeCTAText DataType = 12

	// 500+ reserved for Exchange specific data asset types.
)

// AdUnit defines a Native Ad unit ID.
//
// Deprecated: since version 1.2.
type AdUnit int

const (
	AdUnitUnknown AdUnit = 0

	// Paid Search Units.
	AdUnitPaidSearch AdUnit = 1

	// Recommendation Widgets.
	AdUnitRecommendationWidget AdUnit = 2

	// Promoted Listings.
	AdUnitPromotedListing AdUnit = 3

	// In-Ad (IAB Standard) with Native Element Units.
	AdUnitInAd AdUnit = 4

	// Custom ”Can’t Be Contained”.
	AdUnitCustom AdUnit = 5

	// 500+ Reserved for Exchange specific formats.
)

// Layout defines a Native Ad unit.
//
// Deprecated: since version 1.2.
type Layout int

const (
	LayoutUnknown Layout = 0

	// Content Wall.
	LayoutContentWall Layout = 1

	// App Wall.
	LayoutAppWall Layout = 2

	// News Feed.
	LayoutNewsFeed Layout = 3

	// Chat List.
	LayoutChatList Layout = 4

	// Carousel.
	LayoutCarousel Layout = 5

	// Content Stream.
	LayoutContentStream Layout = 6

	// Grid adjoining the content.
	LayoutGrid Layout = 7

	// 500+ Reserved for Exchange specific layouts.
)

// EventType defines an event type.
type EventType int

const (
	EventTypeUnknown EventType = 0

	// Impression.
	EventTypeImpression EventType = 1

	// Visible impression using MRC definition at 50% in view for 1 second.
	EventTypeViewableMRC50 EventType = 2

	// 100% in view for 1 second (ie GroupM standard).
	EventTypeViewableMRC100 EventType = 3

	// Visible impression for video using MRC definition at 50% in view for
	// 2 seconds.
	EventTypeViewableVideo50 EventType = 4

	// 500+ reserved for Exchange specific event types.
)

// EventTrackingMethod defines an event tracking method.
type EventTrackingMethod int

const (
	EventTrackingMethodUnknown EventTrackingMethod = 0

	// Image-pixel tracking - URL provided will be inserted as a 1x1 pixel at
	// the time of the event.
	EventTrackingMethodImage EventTrackingMethod = 1

	// Javascript-based tracking - URL provided will be inserted as a js tag
	// at the time of the event.
	EventTrackingMethodJS EventTrackingMethod = 2

	// 500+ reserved for Exchange specific event tracking methods.
)

// ContextType defines a context type.
type ContextType int

const (
	ContextTypeUnknown ContextType = 0

	// Content-centric context such as newsfeed, article, image gallery, video
	// gallery, or similar.
	ContextTypeContent ContextType = 1

	// Social-centric context such as social network feed, email, chat, or
	// similar.
	ContextTypeSocial ContextType = 2

	// Product context such as product listings, details, recommendations,
	// reviews, or similar.
	ContextTypeProduct ContextType = 3

	// 500+ reserved for Exchange specific context types.
)

// ContextSubType defines a context sub type.
type ContextSubType int

const (
	ContextSubTypeUnknown ContextSubType = 0

	// General or mixed content.
	ContextSubTypeGeneral ContextSubType = 10

	// Primarily article content (which of course could include images, etc as
	// part of the article).
	ContextSubTypeArticle ContextSubType = 11

	// Primarily video content.
	ContextSubTypeVideo ContextSubType = 12

	// Primarily audio content.
	ContextSubTypeAudio ContextSubType = 13

	// Primarily image content.
	ContextSubTypeImage ContextSubType = 14

	// User-generated content - forums, comments, etc.
	ContextSubTypeUserGenerated ContextSubType = 15

	// General social content such as a general social network.
	ContextSubTypeSocial ContextSubType = 20

	// Primarily email content.
	ContextSubTypeEmail ContextSubType = 21

	// Primarily chat/IM content.
	ContextSubTypeChat ContextSubType = 22

	// Content focused on selling products, whether digital or physical.
	ContextSubTypeSelling ContextSubType = 30

	// Application store/marketplace.
	ContextSubTypeAppStore ContextSubType = 31

	// Product reviews site primarily (which may sell product secondarily).
	ContextSubTypeProductReview ContextSubType = 32

	// 500+ reserved for Exchange specific context sub types.
)

// PlacementType defines a placement type.
type PlacementType int

const (
	PlacementTypeUnknown PlacementType = 0

	// In the feed of content - for example as an item inside the organic
	// feed/grid/listing/carousel.
	PlacementTypeFeed PlacementType = 1

	// In the atomic unit of the content - IE in the article page or single
	// image page.
	PlacementTypeAtomicContentUnit PlacementType = 2

	// Outside the core content - for example in the ads section on the right
	// rail, as a banner-style placement near the content, etc.
	PlacementTypeOutsideCoreContent PlacementType = 3

	// Recommendation widget, most commonly presented below the article content.
	PlacementTypeRecommendationWidget PlacementType = 4

	// 500+ reserved for Exchange specific placement types.
)
