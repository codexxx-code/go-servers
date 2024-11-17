package native

import "encoding/json"

// NativeResponse is a Native Markup response object.
type NativeResponse struct {
	// Version of the Native Markup version in use.
	//
	// Default: 1.2.
	Ver string `json:"ver"`

	// List of Native Ad’s assets.
	//
	// Required: if no assetsurl.
	//
	// Recommended: as fallback even if assetsurl is provided.
	Assets []AssetResponse `json:"assets"`

	// URL of an alternate source for the assets object.
	//
	// The expected response is a JSON object mirroring the assets
	// object in the bid response, subject to certain requirements
	// as specified in the individual objects.
	//
	// Where present, overrides the asset object in the response.
	//
	// The provided “assetsurl” or “dcourl” should be expected to
	// provide a valid response when called in any context, including
	// importantly the brand name and example thumbnails and headlines
	// (to use for reporting, blacklisting, etc), but it is understood
	// the final actual call should come from the client device from
	// which the ad will be rendered to give the buyer the benefit of
	// the cookies and header data available in that context.
	AssetsURL string `json:"assetsurl"`

	// URL where a dynamic creative specification may be found for
	// populating this ad, per the Dynamic Content Ads Specification.
	//
	// Note this is a beta option as the interpretation of the Dynamic
	// Content Ads Specification and how to assign those elements into
	// a Native Ad is outside the scope of this spec and must be agreed
	// offline between the parties or as may be specified in a future
	// revision of the Dynamic Content Ads spec.
	//
	// Where present, overrides the asset object in the response.
	DynamicContentURL string `json:"dcourl"`

	// This is default link object for the Native Ad.
	//
	// Individual assets can also have a link object which applies if
	// the asset is activated(clicked).
	//
	// If the asset doesn’t have a link object, the parent link object
	// applies.
	Link Link `json:"link"`

	// Array of impression tracking URLs, expected to return a 1x1
	// image or 204 response - typically only passed when using 3rd
	// party trackers.
	//
	// Deprecated: since version 1.2.
	ImpTrackers []string `json:"imptrackers"`

	// Optional JavaScript impression tracker.
	//
	// This is a valid HTML, Javascript is already wrapped in <script>
	// tags.
	//
	// It should be executed at impression time where it can be
	// supported.
	//
	// Deprecated: since version 1.2.
	JSTracker string `json:"jstracker"`

	// Array of tracking objects to run with the ad, in response to
	// the declared supported methods in the request.
	EventTrackers []EventTrackerResponse `json:"eventtrackers"`

	// If support was indicated in the request, URL of a page informing
	// the user about the buyer’s targeting activity.
	Privacy string `json:"privacy"`

	// Placeholder for exchange-specific extensions to OpenRTB.
	Ext json.RawMessage `json:"ext,omitempty"`
}

type jsonNativeResponse NativeResponse

// UnmarshalJSON custom unmarshaling.
func (r *NativeResponse) UnmarshalJSON(data []byte) error {
	j := jsonNativeResponse{Ver: "1.2"} //nolint:exhaustruct
	if err := json.Unmarshal(data, &j); err != nil {
		return err
	}
	*r = (NativeResponse)(j)
	return nil
}

// Valdidate the Native Response object.
func (r *NativeResponse) Validate() error {
	if len(r.Assets) == 0 {
		return ErrInvalidRequestNoAssets
	}
	for i := range r.Assets {
		asset := r.Assets[i]
		if err := asset.Validate(); err != nil {
			return err
		}
	}
	return nil
}
