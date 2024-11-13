package openrtb_test

import (
	"testing"

	"pkg/openrtb"
)

func TestVideo_Validate(t *testing.T) {
	testcases := []Testcase{
		{
			Name: "valid",
			Validater: &openrtb.Video{
				MIMEs:     []string{"mime"},
				Linearity: openrtb.VideoLinearityLinear,
				Protocols: []openrtb.Protocol{openrtb.ProtocolDAAST1},
				Sequence:  1,
			},
		},
		{
			Name:      "missing mimes",
			Validater: &openrtb.Video{},
			Err:       openrtb.ErrInvalidVideoNoMIMEs,
		},
		{
			Name: "missing protocols",
			Validater: &openrtb.Video{
				MIMEs:     []string{"mime"},
				Linearity: openrtb.VideoLinearityLinear,
				Sequence:  1,
			},
			Err: openrtb.ErrInvalidVideoNoProtocols,
		},
	}

	assertValidate(t, testcases)
}

func TestVideo_BoxingAllowed(t *testing.T) {
	testcases := []struct {
		name  string
		video openrtb.Video
		exp   int
	}{
		{
			name:  "nil boxingallowed",
			video: openrtb.Video{},
			exp:   0,
		},
		{
			name: "not nil boxingallowed",
			video: openrtb.Video{
				BoxingAllowed: 1,
			},
			exp: 1,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			act := tc.video.GetBoxingAllowed()
			if act != tc.exp {
				t.Errorf(
					"\nExpected a '%d' value but got '%d' value.",
					tc.exp,
					act,
				)
			}
		})
	}
}

func TestVideo_Linearity(t *testing.T) {
	testcases := []struct {
		name  string
		video openrtb.Video
		exp   openrtb.VideoLinearity
	}{
		{
			name:  "nil linearity",
			video: openrtb.Video{},
			exp:   0,
		},
		{
			name: "not nil linearity",
			video: openrtb.Video{
				Linearity: openrtb.VideoLinearityLinear,
			},
			exp: openrtb.VideoLinearityLinear,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			act := tc.video.GetLinearity()
			if act != tc.exp {
				t.Errorf(
					"\nExpected a '%d' value but got '%d' value.",
					tc.exp,
					act,
				)
			}
		})
	}
}

func TestVideo_GetSequence(t *testing.T) {
	testcases := []struct {
		name  string
		video openrtb.Video
		exp   int
	}{
		{
			name:  "nil sequence",
			video: openrtb.Video{},
			exp:   0,
		},
		{
			name: "not nil sequence",
			video: openrtb.Video{
				Sequence: 1,
			},
			exp: 1,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			act := tc.video.GetSequence()
			if act != tc.exp {
				t.Errorf(
					"\nExpected a '%d' value but got '%d' value.",
					tc.exp,
					act,
				)
			}
		})
	}
}

func TestVideo_Unmarshal(t *testing.T) {
	expected := openrtb.Video{
		Width:       640,
		Height:      480,
		Position:    openrtb.AdPositionAboveFold,
		StartDelay:  0,
		MinDuration: 5,
		MaxDuration: 30,
		MaxExtended: 30,
		MinBitrate:  300,
		MaxBitrate:  1500,
		APIs: []openrtb.APIFramework{
			openrtb.APIFrameworkVPAID1,
			openrtb.APIFrameworkVPAID2,
		},
		Protocols: []openrtb.Protocol{
			openrtb.ProtocolVAST2,
			openrtb.ProtocolVAST3,
		},
		MIMEs: []string{
			"video/x-flv",
			"video/mp4",
			"application/x-shockwave-flash",
			"application/javascript",
		},
		Linearity: openrtb.VideoLinearityLinear,
		Sequence:  1,
		PlaybackMethods: []openrtb.VideoPlayback{
			openrtb.VideoPlaybackPageLoadSoundOn,
			openrtb.VideoPlaybackClickToPlay,
		},
		Delivery: []openrtb.ContentDelivery{
			openrtb.ContentDeliveryProgressive,
		},
		BlockedAttrs: []openrtb.CreativeAttribute{
			openrtb.CreativeAttributeUserInitiated,
			openrtb.CreativeAttributeWindowsDialogOrAlert,
		},
		CompanionAds: []openrtb.Banner{
			{
				ID:       "1234567893-1",
				Width:    300,
				Height:   250,
				Position: openrtb.AdPositionAboveFold,
				BlockedAttrs: []openrtb.CreativeAttribute{
					openrtb.CreativeAttributeUserInitiated,
					openrtb.CreativeAttributeWindowsDialogOrAlert,
				},
				ExpDirs: []openrtb.ExpDir{
					openrtb.ExpDirRight,
					openrtb.ExpDirDown,
				},
			},
		},
		CompanionTypes: []int{
			1,
			2,
		},
		Placement:     openrtb.VideoPlacementInStream,
		BoxingAllowed: 1,
	}

	assertEqualJSON(t, "video", &expected)
}
