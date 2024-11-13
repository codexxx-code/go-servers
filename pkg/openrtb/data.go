package openrtb

import "encoding/json"

// Collection of additional user targeting data from a specific data source.
type Data struct {
	// Exchange-specific ID for the data provider.
	ID string `json:"id" db:"id"`

	// Exchange-specific name for the data provider.
	Name string `json:"name" db:"name"`

	// Array of Segment objects that contain the actual data values.
	Segment []Segment `json:"segment" db:"segment"`

	// Placeholder for exchange-specific extensions to OpenRTB
	Ext json.RawMessage `json:"ext,omitempty" db:"ext"`
}

// Specific data point about a user from a specific data source.
type Segment struct {
	// ID of the data segment specific to the data provider.
	ID string `json:"id" db:"id"`

	// Name of the data segment specific to the data provider.
	Name string `json:"name" db:"name"`

	// String representation of the data segment value.
	Value string `json:"value" db:"value"`

	// Placeholder for exchange-specific extensions to OpenRTB.
	Ext json.RawMessage `json:"ext,omitempty" db:"ext"`
}
