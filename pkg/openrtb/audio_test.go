package openrtb_test

import (
	"testing"

	"pkg/openrtb"
)

func TestAudio_Validate(t *testing.T) {
	testcases := []Testcase{
		{
			Name:      "valid",
			Validater: &openrtb.Audio{MIMEs: []string{"mime"}}, //nolint:exhaustruct
			Err:       nil,
		},
		{
			Name:      "no valid",
			Validater: &openrtb.Audio{}, //nolint:exhaustruct
			Err:       openrtb.ErrInvalidAudioNoMIMEs,
		},
	}

	assertValidate(t, testcases)
}

func TestAudio_Unmarshal(t *testing.T) {
	expected := openrtb.Audio{
		MIMEs:       []string{"audio/mp4"},
		MinDuration: 5,
		MaxDuration: 30,
		Protocols: []openrtb.Protocol{
			openrtb.ProtocolDAAST1,
			openrtb.ProtocolDAAST1Wrapper,
		},
		BlockedAttrs: []openrtb.CreativeAttribute{
			openrtb.CreativeAttributeUserInitiated,
			openrtb.CreativeAttributeWindowsDialogOrAlert,
		},
		Sequence:    1,
		MaxExtended: 30,
		MinBitrate:  300,
		MaxBitrate:  1500,
		Delivery: []openrtb.ContentDelivery{
			openrtb.ContentDeliveryProgressive,
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
		APIs: []openrtb.APIFramework{
			openrtb.APIFrameworkVPAID1,
			openrtb.APIFrameworkVPAID2,
		},
		CompanionTypes: []openrtb.CompanionType{
			openrtb.CompanionTypeStatic,
			openrtb.CompanionTypeHTML,
		},
	}

	assertEqualJSON(t, "audio", &expected)
}

func TestAudio_Sequence(t *testing.T) {
	testcases := []struct {
		name  string
		audio openrtb.Audio
		exp   int
	}{
		{
			name:  "nil sequence",
			audio: openrtb.Audio{},
			exp:   0,
		},
		{
			name: "not nil sequence",
			audio: openrtb.Audio{
				Sequence: 1,
			},
			exp: 1,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			act := tc.audio.GetSequence()
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
