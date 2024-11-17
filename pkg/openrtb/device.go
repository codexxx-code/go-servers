package openrtb

import "encoding/json"

// Details of the device on which the content and impressions are displayed.
type Device struct {
	// Browser user agent string.
	//
	// Recommended.
	UserAgent string `json:"ua" bson:"ua"`

	// Location of the device assumed to be the user’s current
	// location defined by a Geo object.
	//
	// Recommended.
	Geo *Geo `json:"geo" bson:"geo"`

	// Standard “Do Not Track” flag as set in the header by the browser, where:
	//    0 = tracking is unrestricted;
	//    1 = do not track.
	//
	// Recommended.
	DNT int `json:"dnt" bson:"dnt"`

	// “Limit Ad Tracking” signal commercially endorsed (e.g., iOS, Android), where:
	//    0 = tracking is unrestricted;
	//    1 = tracking must be limited per commercial guidelines.
	//
	// Recommended.
	LMT int `json:"lmt" bson:"lmt"`

	// IPv4 address closest to device.
	//
	// Recommended.
	IP string `json:"ip" bson:"ip"`

	// IP address closest to device as IPv6.
	IPv6 string `json:"ipv6" bson:"ipv6"`

	// The general type of device.
	DeviceType DeviceType `json:"devicetype" bson:"devicetype"`

	// Device make (e.g., “Apple”).
	Make string `json:"make" bson:"make"`

	// Device model (e.g., “iPhone”).
	Model string `json:"model" bson:"model"`

	// Device operating system (e.g., “iOS”).
	OS string `json:"os" bson:"os"`

	// Device operating system version (e.g., “3.1.2”).
	OSVersion string `json:"osv" bson:"osv"`

	// Hardware version of the device (e.g., “5S” for iPhone 5S).
	HWVersion string `json:"hwv" bson:"hwv"`

	// Physical height of the screen in pixels.
	Height int `json:"h" bson:"h"`

	// Physical width of the screen in pixels.
	Width int `json:"w" bson:"w"`

	// Screen size as pixels per linear inch.
	PPI int `json:"ppi" bson:"ppi"`

	// The ratio of physical pixels to device independent pixels.
	PixelRatio float64 `json:"pxratio" bson:"pxratio"`

	// Support for JavaScript, where:
	//    0 = no;
	//    1 = yes.
	JS int `json:"js" bson:"js"`

	// Indicates if the geolocation API will be available to JavaScript code running
	// in the banner, where:
	//    0 = no;
	//    1 = yes.
	GeoFetch int `json:"geofetch" bson:"geofetch"`

	// Version of Flash supported by the browser.
	FlashVersion string `json:"flashver" bson:"flashver"`

	// Browser language using ISO-639-1-alpha-2.
	Language string `json:"language" bson:"language"`

	// Carrier or ISP (e.g., “VERIZON”) using exchange curated string
	// names which should be published to bidders a priori.
	Carrier string `json:"carrier" bson:"carrier"`

	// Mobile carrier as the concatenated MCC-MNC code (e.g., “310-005” identifies Verizon
	// Wireless CDMA in the USA). Refer to https://en.wikipedia.org/wiki/Mobile_country_code
	// for further examples. Note that the dash between the MCC and MNC parts is required
	// to remove parsing ambiguity.
	MCCMNC string `json:"mccmnc" bson:"mccmnc"`

	// Network connection type.
	ConnectionType ConnectionType `json:"connectiontype" bson:"connectiontype"`

	// ID sanctioned for advertiser use in the clear (i.e., not hashed).
	IFA string `json:"ifa" bson:"ifa"`

	// Hardware device ID (e.g., IMEI); hashed via SHA1.
	IDSHA1 string `json:"didsha1" bson:"didsha1"`

	// Hardware device ID (e.g., IMEI); hashed via MD5.
	IDMD5 string `json:"didmd5" bson:"didmd5"`

	// Platform device ID (e.g., Android ID); hashed via SHA1.
	PIDSHA1 string `json:"dpidsha1" bson:"dpidsha1"`

	// Platform device ID (e.g., Android ID); hashed via MD5.
	PIDMD5 string `json:"dpidmd5" bson:"dpidmd5"`

	// MAC address of the device; hashed via SHA1.
	MacSHA1 string `json:"macsha1" bson:"macsha1"`

	// MAC address of the device; hashed via MD5.
	MacMD5 string `json:"macmd5" bson:"macmd5"`

	// Placeholder for exchange-specific extensions to OpenRTB.
	Ext json.RawMessage `json:"ext,omitempty" bson:"ext"`
}
