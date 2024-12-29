package beforeRequestToDSP

import (
	"encoding/json"
	"net/url"
	"strings"

	"pkg/slices"

	"exchange/internal/enum/sourceTrafficType"

	"pkg/errors"
	"pkg/openrtb"
	"pkg/uuid"
)

type clickunder struct {
	baseLink
}

func (r *clickunder) Apply(dto *beforeRequestToDSP) error {

	// Если запрос не кликандер, то ничего не делаем
	if !dto.IsClickunder {
		return nil
	}

	// Проходимся по каждому Impression
	for i := range dto.BidRequest.Impressions {

		// Ставим индикатор о FullScreen
		dto.BidRequest.Impressions[i].Interstitial = 1

		// Добавляем объект баннер
		// TODO: Реализовать функционал переключения типа трафика
		/*
			video, err := getDefaultVideo()
			if err != nil {
				return err
			}
			dto.BidRequest.Impressions[i].Video = &video
		*/
		video, _ := getDefaultVideo()
		dto.BidRequest.Impressions[i].Video = &video

		dto.BidRequest.Impressions[i].Ext = nil
	}

	dto.BidRequest.BlockedCategories = nil
	dto.BidRequest.BlockedAdvDomains = nil
	dto.BidRequest.App = nil
	dto.BidRequest.Ext = nil
	if dto.BidRequest.User != nil {
		dto.BidRequest.User.Ext = nil
	}
	if dto.BidRequest.Device != nil {
		dto.BidRequest.Device.Ext = nil
	}

	// Для юзерагента VW (что означает, что запрос пришёл из in-app) отдаём запрос типа App
	if slices.Contains(dto.dsp.SourceTrafficTypes, sourceTrafficType.InApp) {
		if err := makeApp(dto); err != nil {
			return err
		}
	} else {
		if err := makeSite(dto); err != nil {
			return err
		}
	}

	return nil
}

func makeApp(dto *beforeRequestToDSP) (err error) {
	dto.BidRequest.App, err = getClickunderApp(
		"1600034464", // TODO: вынести в настройки
		"1551847165", // TODO: вынести в настройки
	)

	dto.BidRequest.Device.IFA = uuid.New()
	dto.BidRequest.Site = nil

	return err
}

func makeSite(dto *beforeRequestToDSP) (err error) {

	// TODO: Вынести в отдельный чейн и добавлять по отдельной галке в настройке DSP
	if dto.BidRequest.Site, err = getClickunderSite(dto.Settings.ShowcaseURL); err != nil {
		return err
	}

	// Добавляем для DSP к сайту также домен паба
	publisherID := strings.ReplaceAll(dto.PublisherID, ".", "_")
	if publisherID != "" {
		dto.BidRequest.Site.Domain = publisherID + "." + dto.BidRequest.Site.Domain
	}

	dto.BidRequest.App = nil

	return err
}

func getClickunderSite(siteURL string) (site *openrtb.Site, err error) {

	// Парсим URL витрины, чтобы потом достать оттуда только хост
	showcaseURL, err := url.Parse(siteURL)
	if err != nil {
		return site, errors.InternalServer.Wrap(err)
	}

	return &openrtb.Site{
		Inventory: openrtb.Inventory{
			ID:                uuid.New(),
			Name:              "",
			Domain:            showcaseURL.Host,
			Categories:        nil,
			SectionCategories: nil,
			PageCategory:      nil,
			PrivacyPolicy:     0,
			Publisher:         nil,
			Content:           nil,
			Keywords:          "",
			Ext:               nil,
		},
		Page:     siteURL,
		Refferer: "",
		Search:   "",
		Mobile:   0,
	}, nil
}

func getClickunderApp(bundle, publisher string) (*openrtb.App, error) { //nolint:unparam
	return &openrtb.App{
		Inventory: openrtb.Inventory{
			ID:                "",
			Name:              "",
			Domain:            "",
			Categories:        nil,
			SectionCategories: nil,
			PageCategory:      nil,
			PrivacyPolicy:     0,
			Publisher: &openrtb.Publisher{
				ID:         publisher,
				Name:       "",
				Categories: nil,
				Domain:     "",
				Ext:        nil,
			},
			Content:  nil,
			Keywords: "",
			Ext:      nil,
		},
		Bundle:   bundle,
		StoreURL: "",
		Version:  "",
		Paid:     0,
	}, nil
}

func getDefaultBanner() *openrtb.Banner { //nolint:unused

	formats := []openrtb.Format{
		{
			Width:       730,
			Height:      90,
			WidthRatio:  0,
			HeightRatio: 0,
			WidthMin:    0,
			Ext:         nil,
		},
		{
			Width:       320,
			Height:      50,
			WidthRatio:  0,
			HeightRatio: 0,
			WidthMin:    0,
			Ext:         nil,
		},
		{
			Width:       480,
			Height:      320,
			WidthRatio:  0,
			HeightRatio: 0,
			WidthMin:    0,
			Ext:         nil,
		},
		{
			Width:       300,
			Height:      250,
			WidthRatio:  0,
			HeightRatio: 0,
			WidthMin:    0,
			Ext:         nil,
		},
		{
			Width:       320,
			Height:      90,
			WidthRatio:  0,
			HeightRatio: 0,
			WidthMin:    0,
			Ext:         nil,
		},
		{
			Width:       320,
			Height:      480,
			WidthRatio:  0,
			HeightRatio: 0,
			WidthMin:    0,
			Ext:         nil,
		},
	}

	return &openrtb.Banner{
		Formats:      formats,
		Width:        0,
		Height:       0,
		WidthMax:     0,
		HeightMax:    0,
		WidthMin:     0,
		HeightMin:    0,
		BlockedTypes: nil,
		BlockedAttrs: nil,
		Position:     openrtb.AdPositionFullscreen,
		MIMEs:        nil,
		TopFrame:     0,
		ExpDirs:      nil,
		APIs:         nil,
		ID:           "",
		VCM:          0,
		Ext:          nil,
	}
}

func getDefaultVideo() (video openrtb.Video, err error) { //nolint:unused // TODO: Заюзать

	const (
		fullscreenWidth  = 320
		fullscreenHeight = 480
		minDuration      = 5
		maxDuration      = 60
	)
	videoExtMap := map[string]string{
		"videotype": "skippable",
	}
	videoExtJSONBytes, err := json.Marshal(videoExtMap)
	if err != nil {
		return video, errors.InternalServer.Wrap(err)
	}

	return openrtb.Video{
		MIMEs: []string{
			"video/mp4",
			"video/3gpp",
		},
		MinDuration: minDuration,
		MaxDuration: maxDuration,
		Protocols: []openrtb.Protocol{
			openrtb.ProtocolVAST1,
			openrtb.ProtocolVAST2,
			openrtb.ProtocolVAST3,
			openrtb.ProtocolVAST1Wrapper,
			openrtb.ProtocolVAST2Wrapper,
			openrtb.ProtocolVAST3Wrapper,
			openrtb.ProtocolVAST4,
			openrtb.ProtocolVAST4Wrapper,
			openrtb.ProtocolDAAST1,
			openrtb.ProtocolDAAST1Wrapper,
		},
		Protocol:      0,
		Width:         0,
		Height:        0,
		StartDelay:    0,
		Placement:     0,
		Linearity:     openrtb.VideoLinearityNonLinear,
		Skip:          0,
		SkipMin:       0,
		SkipAfter:     0,
		Sequence:      1,
		BlockedAttrs:  nil,
		MaxExtended:   0,
		MinBitrate:    0,
		MaxBitrate:    0,
		BoxingAllowed: 1,
		PlaybackMethods: []openrtb.VideoPlayback{
			openrtb.VideoPlaybackPageLoadSoundOn,
			openrtb.VideoPlaybackPageLoadSoundOff,
			openrtb.VideoPlaybackClickToPlay,
			openrtb.VideoPlaybackMouseOver,
			openrtb.VideoPlaybackEnterSoundOn,
			openrtb.VideoPlaybackEnterSoundOff,
		},
		PlaybackEnd:  0,
		Delivery:     nil,
		Position:     openrtb.AdPositionFullscreen,
		CompanionAds: nil,
		APIs: []openrtb.APIFramework{
			openrtb.APIFrameworkVPAID2,
			openrtb.APIFrameworkVPAID1,
		},

		CompanionTypes: nil,
		Ext:            videoExtJSONBytes,
	}, nil
}
