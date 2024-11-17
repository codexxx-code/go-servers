package openrtb

import "encoding/json"

// Location of the device or user’s home base depending on the parent object.
type Geo struct {
	// Latitude from -90.0 to +90.0, where negative is south.
	Latitude float64 `json:"lat" bson:"lat"`

	// Longitude from -180.0 to +180.0, where negative is west.
	Longtitude float64 `json:"lon" bson:"lon"`

	// Source of location data; recommended when passing lat/lon.
	Type LocationType `json:"type" bson:"type"`

	// Estimated location accuracy in meters; recommended when lat/lon are specified
	// and derived from a device’s location services (i.e., type = 1). Note that this is
	// the accuracy as reported from the device. Consult OS specific documentation
	// (e.g., Android, iOS) for exact interpretation.
	Accuracy int `json:"accuracy" bson:"accuracy"`

	// Number of seconds since this geolocation fix was established. Note that devices
	// may cache location data across multiple fetches. Ideally, this value should be
	// from the time the actual fix was taken.
	LastFix int `json:"lastfix" bson:"lastfix"`

	// Service or provider used to determine geolocation from IP address if applicable
	// (i.e., type = 2).
	IPService IPLocation `json:"ipservice" bson:"ipservice"`

	// Country code using ISO-3166-1-alpha-3.
	Country string `json:"country" bson:"country"`

	// Region code using ISO-3166-2; 2-letter state code if USA.
	Region string `json:"region" bson:"region"`

	// Region of a country using FIPS 10-4 notation. While OpenRTB supports this attribute,
	// it has been withdrawn by NIST in 2008.
	RegionFIPS104 string `json:"regionfips104" bson:"regionfips104"`

	// Google metro code; similar to but not exactly Nielsen DMAs. See Appendix A for a link
	// to the codes.
	Metro string `json:"metro" bson:"metro"`

	// City using United Nations Code for Trade & Transport Locations. See Appendix A for a
	// link to the codes.
	City string `json:"city" bson:"city"`

	// ZIP or postal code.
	ZIP string `json:"zip" bson:"zip"`

	// Local time as the number +/- of minutes from UTC.
	UTCOffset int `json:"utcoffset" bson:"utcoffset"`

	// Placeholder for exchange-specific extensions to OpenRTB.
	Ext json.RawMessage `json:"ext,omitempty" bson:"ext"`
}
