package openrtb

// Header represents the header of the response/request.
type Header string

// Header values.
const (
	OpenRTBVersionHeader Header = "x-openrtb-version"
)

// String implements Stringer interface.
func (h Header) String() string { return string(h) }

// Verson represents the version of the OpenRTB protocol.
type Version string

// Version values.
const (
	OpenRTBVersion30  Version = "3.0"
	OpenRTBVersion26  Version = "2.6"
	OpenRTBVersion25  Version = "2.5"
	OpenRTBVersion24  Version = "2.4"
	OpenRTBVersion231 Version = "2.3.1"
	OpenRTBVersion23  Version = "2.3"
	OpenRTBVersion22  Version = "2.2"
	OpenRTBVersion21  Version = "2.1"
	OpenRTBVersion20  Version = "2.0"
)

// String implements Stringer interface.
func (v Version) String() string { return string(v) }
