package native_test

import (
	"testing"

	"pkg/openrtb"
	"pkg/openrtb/native"
)

func TestNativeRequest_Validate(t *testing.T) {
	testcases := []Testcase{
		{
			Name: "valid title",
			Validater: &native.NativeRequest{
				Assets: []native.AssetRequest{
					{Title: &native.TitleRequest{}},
				},
			},
		},
		{
			Name: "valid image",
			Validater: &native.NativeRequest{
				Assets: []native.AssetRequest{
					{Image: &native.ImageRequest{}},
				},
			},
		},
		{
			Name: "valid video",
			Validater: &native.NativeRequest{
				Assets: []native.AssetRequest{
					{Video: &native.VideoRequest{
						MIMEs:    []string{"mime"},
						Protocol: openrtb.ProtocolVAST3Wrapper,
					}},
				},
			},
		},
		{
			Name: "valid data",
			Validater: &native.NativeRequest{
				Assets: []native.AssetRequest{
					{Data: &native.DataRequest{}},
				},
			},
		},
		{
			Name:      "no assets",
			Validater: &native.NativeRequest{},
			Err:       native.ErrInvalidRequestNoAssets,
		},
		{
			Name: "multiple assets",
			Validater: &native.NativeRequest{
				Assets: []native.AssetRequest{
					{Title: &native.TitleRequest{}, Image: &native.ImageRequest{}},
				},
			},
			Err: native.ErrInvalidMultiAssets,
		},
	}

	assertValidate(t, testcases)
}

func TestNativeRequest_UnmarshalJSON(t *testing.T) {
	testcases := []struct {
		file string
		exp  native.NativeRequest
	}{
		{
			file: "nativerequest_v1.1",
			exp: native.NativeRequest{
				Ver:            "1.1",
				Layout:         native.LayoutContentWall,
				AdUnit:         native.AdUnitPaidSearch,
				ContextType:    native.ContextTypeContent,
				ContextSubType: native.ContextSubTypeGeneral,
				PlacementType:  11,
				PlacementCount: 1,
				AURLSupport:    1,
				DURLSupport:    1,
				Assets: []native.AssetRequest{
					{
						ID:       123,
						Required: 1,
						Title: &native.TitleRequest{
							Length: 140,
						},
					},
					{
						ID: 128,
						Image: &native.ImageRequest{
							Type:      native.ImageTypeMain,
							WidthMin:  836,
							HeightMin: 627,
						},
					},
					{
						ID:       126,
						Required: 1,
						Data: &native.DataRequest{
							Type:   native.DataTypeSponsored,
							Length: 25,
						},
					},
					{
						ID:       129,
						Required: 1,
						Video: &native.VideoRequest{
							Width:  640,
							Height: 480,
							APIs: []openrtb.APIFramework{
								openrtb.APIFrameworkVPAID1,
								openrtb.APIFrameworkVPAID2,
							},
							Protocols: []openrtb.Protocol{
								openrtb.ProtocolVAST2,
								openrtb.ProtocolVAST3,
							},
							MIMEs:         []string{"video/mp4"},
							Linearity:     openrtb.VideoLinearityLinear,
							Sequence:      1,
							BoxingAllowed: 1,
						},
					},
				},
			},
		},
		{
			file: "nativerequest_v1.2",
			exp: native.NativeRequest{
				Ver:            "1.2",
				ContextType:    native.ContextTypeSocial,
				ContextSubType: native.ContextSubTypeSocial,
				PlacementType:  11,
				PlacementCount: 1,
				AURLSupport:    1,
				DURLSupport:    1,
				Assets: []native.AssetRequest{
					{
						ID:       123,
						Required: 1,
						Title: &native.TitleRequest{
							Length: 140,
						},
					},
					{
						ID: 128,
						Image: &native.ImageRequest{
							Type:      native.ImageTypeMain,
							WidthMin:  836,
							HeightMin: 627,
						},
					},
					{
						ID:       126,
						Required: 1,
						Data: &native.DataRequest{
							Type:   native.DataTypeSponsored,
							Length: 25,
						},
					},
					{
						ID:       129,
						Required: 1,
						Video: &native.VideoRequest{
							Width:  640,
							Height: 480,
							APIs: []openrtb.APIFramework{
								openrtb.APIFrameworkVPAID1,
								openrtb.APIFrameworkVPAID2,
							},
							Protocols: []openrtb.Protocol{
								openrtb.ProtocolVAST2,
								openrtb.ProtocolVAST3,
							},
							MIMEs:         []string{"video/mp4"},
							Linearity:     openrtb.VideoLinearityLinear,
							Sequence:      1,
							BoxingAllowed: 1,
						},
					},
				},
			},
		},
	}

	for _, tc := range testcases {
		tc := tc
		t.Run(tc.file, func(t *testing.T) {
			t.Parallel()
			assertEqualJSON(t, tc.file, &tc.exp)
		})
	}
}
