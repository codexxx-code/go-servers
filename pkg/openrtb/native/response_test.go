package native_test

import (
	"testing"

	"pkg/openrtb/native"
)

func TestNativeResponse_Validate(t *testing.T) {
	testcases := []Testcase{
		{
			Name: "valid title",
			Validater: &native.NativeResponse{
				Assets: []native.AssetResponse{
					{Title: &native.TitleResponse{}},
				},
			},
		},
		{
			Name: "valid image",
			Validater: &native.NativeResponse{
				Assets: []native.AssetResponse{
					{Image: &native.ImageResponse{}},
				},
			},
		},
		{
			Name: "valid video",
			Validater: &native.NativeResponse{
				Assets: []native.AssetResponse{
					{Video: &native.VideoResponse{}},
				},
			},
		},
		{
			Name: "valid data",
			Validater: &native.NativeResponse{
				Assets: []native.AssetResponse{
					{Data: &native.DataResponse{}},
				},
			},
		},
		{
			Name:      "no assets",
			Validater: &native.NativeResponse{},
			Err:       native.ErrInvalidRequestNoAssets,
		},
		{
			Name: "multiple assets",
			Validater: &native.NativeResponse{
				Assets: []native.AssetResponse{
					{Title: &native.TitleResponse{}, Image: &native.ImageResponse{}},
				},
			},
			Err: native.ErrInvalidMultiAssets,
		},
	}

	assertValidate(t, testcases)
}

func TestNativeResponse_UnmarshalJSON(t *testing.T) {
	testcases := []struct {
		file string
		exp  native.NativeResponse
	}{
		{
			file: "nativeresponse_v1.1",
			exp: native.NativeResponse{
				Ver: "1.1",
				Link: native.Link{
					URL: "http://i.am.a/URL",
				},
				Assets: []native.AssetResponse{
					{
						ID:       123,
						Required: 1,
						Title: &native.TitleResponse{
							Text: "Learn about this awesome thing",
						},
					},
					{
						ID:       124,
						Required: 1,
						Image: &native.ImageResponse{
							URL: "http://www.myads.com/thumbnail1.png",
						},
					},
					{
						ID:       128,
						Required: 1,
						Image: &native.ImageResponse{
							URL: "http://www.myads.com/largethumb1.png",
						},
					},
					{
						ID:       126,
						Required: 1,
						Data: &native.DataResponse{
							Value: "My Brand",
						},
					},
					{
						ID:       127,
						Required: 1,
						Data: &native.DataResponse{
							Value: "Learn all about this awesome story of someone using my product.",
						},
					},
				},
			},
		},
		{
			file: "nativeresponse_video_v1.1",
			exp: native.NativeResponse{
				Ver: "1.1",
				Link: native.Link{
					URL: "http://i.am.a/URL",
				},
				Assets: []native.AssetResponse{
					{
						ID: 4,
						Video: &native.VideoResponse{
							VAST: "<VAST version='2.0'></VAST>",
						},
					},
					{
						ID:       123,
						Required: 1,
						Title: &native.TitleResponse{
							Text: "Watch this awesome thing",
						},
					},
					{
						ID:       124,
						Required: 1,
						Image: &native.ImageResponse{
							URL: "http://www.myads.com/thumbnail1.png",
						},
					},
					{
						ID:       128,
						Required: 1,
						Image: &native.ImageResponse{
							URL: "http://www.myads.com/largethumb1.png",
						},
					},
					{
						ID:       126,
						Required: 1,
						Data: &native.DataResponse{
							Value: "My Brand",
						},
					},
					{
						ID:       127,
						Required: 1,
						Data: &native.DataResponse{
							Value: "Watch all about this awesome story of someone using my product.",
						},
					},
				},
			},
		},
		{
			file: "nativeresponse_v1.2",
			exp: native.NativeResponse{
				Ver: "1.2",
				Link: native.Link{
					URL: "http://i.am.a/URL",
				},
				Assets: []native.AssetResponse{
					{
						ID:       123,
						Required: 1,
						Title: &native.TitleResponse{
							Text: "Learn about this awesome thing",
						},
					},
					{
						ID:       124,
						Required: 1,
						Image: &native.ImageResponse{
							URL: "http://www.myads.com/thumbnail1.png",
						},
					},
					{
						ID:       128,
						Required: 1,
						Image: &native.ImageResponse{
							URL: "http://www.myads.com/largethumb1.png",
						},
					},
					{
						ID:       126,
						Required: 1,
						Data: &native.DataResponse{
							Value: "My Brand",
						},
					},
					{
						ID:       127,
						Required: 1,
						Data: &native.DataResponse{
							Value: "Learn all about this awesome story of someone using my product.",
						},
					},
				},
				EventTrackers: []native.EventTrackerResponse{
					{
						Type:   native.EventTypeImpression,
						Method: native.EventTrackingMethodJS,
						URL:    "http://www.mytracker.com/tracker.js",
					},
					{
						Type:   native.EventTypeViewableMRC50,
						Method: native.EventTrackingMethodImage,
						URL:    "http://www.mytracker.com/tracker.php",
					},
				},
				Privacy: "http://www.myprivacyurl.com",
			},
		},
		{
			file: "nativeresponse_video_v1.2",
			exp: native.NativeResponse{
				Ver: "1.2",
				Link: native.Link{
					URL: "http://i.am.a/URL",
				},
				Assets: []native.AssetResponse{
					{
						ID: 4,
						Video: &native.VideoResponse{
							VAST: "<VAST version='2.0'></VAST>",
						},
					},
					{
						ID:       123,
						Required: 1,
						Title: &native.TitleResponse{
							Text: "Watch this awesome thing",
						},
					},
					{
						ID:       124,
						Required: 1,
						Image: &native.ImageResponse{
							URL: "http://www.myads.com/thumbnail1.png",
						},
					},
					{
						ID:       128,
						Required: 1,
						Image: &native.ImageResponse{
							URL: "http://www.myads.com/largethumb1.png",
						},
					},
					{
						ID:       126,
						Required: 1,
						Data: &native.DataResponse{
							Value: "My Brand",
						},
					},
					{
						ID:       127,
						Required: 1,
						Data: &native.DataResponse{
							Value: "Watch all about this awesome story of someone using my product.",
						},
					},
				},
			},
		},
		{
			file: "nativeresponse_third_party_v1.2",
			exp: native.NativeResponse{
				Ver:       "1.2",
				AssetsURL: "http://www.myadserver.com/getad123nativejson.php",
				Link: native.Link{
					URL: "http://i.am.a/URL",
				},
				Assets: []native.AssetResponse{
					{
						ID:       123,
						Required: 1,
						Title: &native.TitleResponse{
							Text: "Learn about this awesome thing",
						},
					},
					{
						ID:       124,
						Required: 1,
						Image: &native.ImageResponse{
							URL: "http://www.myads.com/thumbnail1.png",
						},
					},
					{
						ID:       128,
						Required: 1,
						Image: &native.ImageResponse{
							URL: "http://www.myads.com/largethumb1.png",
						},
					},
					{
						ID:       126,
						Required: 1,
						Data: &native.DataResponse{
							Value: "My Brand",
						},
					},
					{
						ID:       127,
						Required: 1,
						Data: &native.DataResponse{
							Value: "Learn all about this awesome story of someone using my product.",
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
