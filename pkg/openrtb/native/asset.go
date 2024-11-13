package native

import (
	"encoding/json"
	"errors"
)

// Validation errors.
var ErrInvalidMultiAssets = errors.New("asset has multiple assets")

// AssetRequest is the main container object for each asset requested or
// supported by Exchange on behalf of the rendering client.
//
// Any object that is required is to be flagged as such.
//
// Only one of the {title,img,video,data} objects should be present in each
// object. All others should be null/absent.
//
// The id is to be unique within the AssetRequest Object array so that the response
// can be aligned.
type AssetRequest struct {
	// Unique asset ID, assigned by exchange.
	//
	// Required.
	ID int `json:"id"`

	// Set to 1 if asset is required (exchange will not accept a bid
	// without it).
	//
	// Default 0.
	Required int `json:"required"`

	// Title object for title assets.
	//
	// Recommended.
	Title *TitleRequest `json:"title"`

	// Image object for image assets.
	//
	// Recommended.
	Image *ImageRequest `json:"img"`

	// Video object for video assets.
	Video *VideoRequest `json:"video"`

	// Data object for brand name, description, ratings, prices etc.
	//
	// Recommended.
	Data *DataRequest `json:"data"`

	// Placeholder for exchange-specific extensions to OpenRTB.
	Ext json.RawMessage `json:"ext,omitempty"`
}

type jsonAsset AssetRequest

func (a *AssetRequest) UnmarshalJSON(data []byte) error {
	var j jsonAsset
	if err := json.Unmarshal(data, &j); err != nil {
		return err
	}
	*a = (AssetRequest)(j)
	return nil
}

// Validate the Asset object.
func (a *AssetRequest) Validate() error {
	if a.assetCount() != 1 {
		return ErrInvalidMultiAssets
	}
	if a.Video != nil {
		return a.Video.Validate()
	}
	return nil
}

func (a *AssetRequest) assetCount() int {
	var n int
	if a.Title != nil {
		n++
	}
	if a.Image != nil {
		n++
	}
	if a.Video != nil {
		n++
	}
	if a.Data != nil {
		n++
	}
	return n
}

// AssetResponse is the main container object for each asset requested or
// supported by Exchange on behalf of the rendering client.
//
// Any object that is required is to be flagged as such.
//
// Only one of the {title,img,video,data} objects should be present in each
// object. All others should be null/absent.
//
// The id is to be unique within the AssetRequest Object array so that the response
// can be aligned.
type AssetResponse struct {
	// Unique asset ID, assigned by exchange.
	//
	// Optional: if assetsurl/dcourl is being used.
	//
	// Required: if embedded asset is being used.
	ID int `json:"id"`

	// Set to 1 if asset is required (bidder requires it to be
	// displayed).
	//
	// Default: 0.
	Required int `json:"required"`

	// Title object for title assets.
	Title *TitleResponse `json:"title"`

	// Image object for image assets.
	Image *ImageResponse `json:"img"`

	// Video object for video assets.
	Video *VideoResponse `json:"video"`

	// Data object for ratings, prices etc.
	Data *DataResponse `json:"data"`

	// Link object for call to actions. The link object applies if
	// the asset item is activated (clicked).
	//
	// If there is no link object on the asset, the parent link object
	// on the bid response applies.
	Link *Link `json:"link"`

	// Placeholder for exchange-specific extensions to OpenRTB.
	Ext json.RawMessage `json:"ext,omitempty"`
}

// Validate the Asset object.
func (a *AssetResponse) Validate() error {
	if a.assetCount() != 1 {
		return ErrInvalidMultiAssets
	}
	return nil
}

func (a *AssetResponse) assetCount() int {
	var n int
	if a.Title != nil {
		n++
	}
	if a.Image != nil {
		n++
	}
	if a.Video != nil {
		n++
	}
	if a.Data != nil {
		n++
	}
	return n
}
