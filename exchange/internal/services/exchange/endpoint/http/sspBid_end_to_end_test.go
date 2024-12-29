package http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"exchange/internal/config"
	"exchange/internal/enum/formatType"
	"exchange/internal/enum/sourceTrafficType"
	dspModel "exchange/internal/services/dsp/model"
	"exchange/internal/services/exchange/model"
	"exchange/internal/services/exchange/service"
	settingsModel "exchange/internal/services/setting/model"
	sspModel "exchange/internal/services/ssp/model"
	"pkg/decimal"
	"pkg/errors"
	"pkg/http/fiber"
	"pkg/log"
	"pkg/openrtb"
	"pkg/pointer"
	"pkg/testUtils"
	"pkg/uuid"
)

const testcaseDir = "testcase/sspBid_end_to_end"

type response struct {
	responseFile int
	err          error
}

func Test_end_to_end_BidSSP(t *testing.T) { //nolint:gocyclo

	const sspRequestFileTemplate = "%d.SSPRequestImpression.json"

	const defaultSSPSlug = "test_slug"
	generalCurrencyRates := map[string]decimal.Decimal{
		"RUB": decimal.NewFromFloat(0.01),
		"USD": decimal.NewFromInt(1),
		"EUR": decimal.NewFromFloat(1.1),
	}

	generalSettings := settingsModel.Settings{
		Margin:                     decimal.NewFromFloat(0.1),
		Host:                       "https://test.host.com",
		DefaultTimeout:             100000000,
		EmptySecondPriceReduceCoef: decimal.NewFromFloat(0.1),
		ReduceTimeoutCoef:          0.1,
		ShowcaseURL:                "https://test.showcase.com",
	}

	// Инициализируем все синглтон зависимости
	testUtils.SetDefaultEnvs(t)
	testUtils.IgnoreError(config.InitSingletones(config.Load()))

	type mockDTO struct {
		testNumber      int                // Номер теста
		sspBidRequest   openrtb.BidRequest // Запрос от SSP, который автоматически парсится из файла
		exchangeService *service.ExchangeService
	}
	type args struct {
		isSendNoContentInsteadErrorsInResponse bool // Параметр, который проставляется в env и отвечает за отправку 204 вместо ошибок
		isSendNoContentForSuccessResponse      bool // Параметр, который проставляется в env и отвечает за отправку 204 вместо успешных ответов
		wantHTTPCode                           int  // Ожидаемый код ответа от сервера на запрос
	}
	// Тесткейсы
	tests := []struct {
		testNumber  int                                              // Номер теста, используется для автоматического парсинга файлов
		name        string                                           // Название теста, чисто информационная строка
		args        args                                             // Входные параметры теста
		mockActions func(service.Mocks, mockDTO) openrtb.BidResponse // Перечень вызовов зависимых функций сервиса
	}{
		{
			testNumber: 1,
			name:       "Отработка успешного запроса на кликандер",
			args: args{
				isSendNoContentInsteadErrorsInResponse: false,
				isSendNoContentForSuccessResponse:      false,
				wantHTTPCode:                           http.StatusOK,
			},
			mockActions: func(mocks service.Mocks, dto mockDTO) openrtb.BidResponse {

				// Фиксируем пришедший запрос в кафку
				createSSPBidEventReq := dto.sspBidRequest
				createSSPBidEventReq.Ext = []byte(fmt.Sprintf(`{"ssp-slug":"%s"}`, defaultSSPSlug))
				mocks.EventService.MockCreateSSPBidRequestEvent(createSSPBidEventReq,
					nil)

				// Получаем SSP согласно запросу, в нашем случае тестируется кликандер, поэтому он true
				const sspCurrency = "USD"

				ssp := sspModel.SSP{
					Slug:                defaultSSPSlug,
					ClickunderDrumSize:  pointer.Pointer(int32(4)),
					Timeout:             nil,
					IsEnable:            true,
					ClickunderADMFormat: pointer.Pointer("<?xml version=\"1.0\" encoding=\"ISO-8859-1\"?>\n<ad><popunderAd><url><![CDATA[${ADM_URL}]]></url></popunderAd></ad>"),
					FormatTypes:         []formatType.FormatType{formatType.Clickunder},
				}
				mocks.SSPService.MockGetSSP(sspModel.GetSSPsReq{
					Slugs: []string{defaultSSPSlug},
				},
					ssp, nil)

				// Получаем настройки системы, они общие на все тесткейсы
				settings := generalSettings
				mocks.SettingService.MockGetSettings(
					settings, nil)

				// Получаем курсы валют, они общие на все тесткейсы
				currencyRates := generalCurrencyRates
				mocks.CurrencyService.MockGetRates(
					currencyRates, nil)

				// Общий наш идентификатор запроса, который генерируем перед получением DSP
				uuid.AddMockValues("mongoRequestID")

				mocks.ExchangeRepository.MockGetPublisherVisibility("1511441620319572",
					model.PublisherVisibility{
						PublisherID: "1511441620319572",
						Loads:       100,
						Views:       69,
					}, nil)

				const (
					dsp1Currency = "RUB"
					dsp2Currency = "USD"
					dsp3Currency = "EUR"
				)

				// Генерируем DSP согласно входным данным тесткейса
				dsps := []dspModel.DSP{
					{
						Slug:                     "1",
						URL:                      "https://dsp1.com",
						Currency:                 dsp1Currency,
						AuctionSecondPrice:       false,
						IsSupportMultiimpression: false,
					},
					{
						Slug:                     "2",
						URL:                      "https://dsp2.com",
						Currency:                 dsp2Currency,
						AuctionSecondPrice:       true,
						IsSupportMultiimpression: false,
					},
					{
						Slug:                     "3",
						URL:                      "https://dsp3.com",
						Currency:                 dsp3Currency,
						AuctionSecondPrice:       false,
						IsSupportMultiimpression: false,
					},
				}
				mocks.DSPService.MockGetDSPs(dspModel.GetDSPsReq{
					SourceTrafficType: []sourceTrafficType.SourceTrafficType{sourceTrafficType.Desktop},
					IsEnable:          pointer.Pointer(true),
				},
					dsps, nil)

				// Чекаем в одном из чейнов страну по IP с вызовом в геолайт
				mocks.ExchangeRepository.MockGetCountryByIP(dto.sspBidRequest.Device.IP,
					"RUS", nil)

				changeReqForDSP := func(req openrtb.BidRequest, dspIndex int) openrtb.BidRequest {
					copyReq := req.Copy()
					switch dspIndex {
					case 0:
						copyReq.Impressions[0].BidFloor = decimal.NewFromFloat(2.75) // (0.1 * 100 + 10%) / 4
						copyReq.Impressions[0].BidFloorCurrency = dsp1Currency
						copyReq.Currencies = []string{dsp1Currency}
						copyReq.AuctionType = 1
					case 1:
						copyReq.Impressions[0].BidFloor = decimal.NewFromFloat(0.0275) // (0.1 + 10%) / 4
						copyReq.Impressions[0].BidFloorCurrency = dsp2Currency
						copyReq.Currencies = []string{dsp2Currency}
						copyReq.AuctionType = 2
					case 2:
						copyReq.Impressions[0].BidFloor = decimal.NewFromFloat(0.025) // (0.1 / 1.1 + 10%) / 4
						copyReq.Impressions[0].BidFloorCurrency = dsp3Currency
						copyReq.Currencies = []string{dsp3Currency}
						copyReq.AuctionType = 1
					}
					return copyReq
				}

				// Получаем ожидаемый запрос к DSP согласно входным данным
				const originRequestID = "50c3d53cc39690b5febd02db6b28b206"
				bidRequestTemplate := openrtb.BidRequest{
					ID: "id",
					Impressions: []openrtb.Impression{
						{
							ID: "1",
							Banner: &openrtb.Banner{ // Добавили из шаблона
								Formats: []openrtb.Format{
									{
										Width:  730,
										Height: 90,
									},
									{
										Width:  320,
										Height: 50,
									},
									{
										Width:  480,
										Height: 320,
									},
									{
										Width:  300,
										Height: 250,
									},
									{
										Width:  320,
										Height: 90,
									},
									{
										Width:  320,
										Height: 480,
									},
								},
								Width:        0,
								Height:       0,
								Position:     openrtb.AdPositionFullscreen,
								BlockedAttrs: nil,
								BlockedTypes: nil,
								Ext:          nil,
							},
							Interstitial:     1,
							BidFloor:         decimal.Decimal{},
							BidFloorCurrency: "", // Будет заполняться в фукнции
							Ext:              nil,
						},
					},
					Site: &openrtb.Site{ // Добавили из шаблона
						Inventory: openrtb.Inventory{
							ID:     "id",
							Name:   "",
							Domain: "1511441620319572.test.showcase.com",
						},
						Page: "https://test.showcase.com",
					},
					Device: &openrtb.Device{
						UserAgent: "Mozilla/5.0 (Linux; Android 10; K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Mobile Safari/537.36",
						Geo: &openrtb.Geo{
							Country: "RUS",
						},
						IP:         "106.222.213.16",
						DeviceType: 4,
						Language:   "en",
					},
					User: &openrtb.User{
						ID: "id",
						Geo: &openrtb.Geo{
							Country: "RUS",
						},
					},
					AuctionType: 0,          // Будет заполняться в фукнции
					TimeMax:     132,        // 147 - 10%
					Currencies:  []string{}, // Будет заполняться в фукнции
					Ext:         nil,
				}

				// Номер файла, номер ответа
				responsesMap := make(map[int]openrtb.BidResponse)

				// Индекс DSP - Номер вызова - ответ
				dspResponses := [3][4]response{
					{
						{1, nil}, // 203 рубля // 4 место
						{2, nil}, // 409 рублей // 2 место
						{0, errors.BadRequest.New("test error")},
						{0, errors.BadRequest.New("test error")},
					},
					{
						{3, nil}, // 4 доллара  // 3 место
						{4, nil}, // 4.5 доллара // 1 место
						{0, errors.BadRequest.New("test error")},
						{0, errors.BadRequest.New("test error")},
					},
					{
						{0, errors.BadRequest.New("test error")},
						{0, errors.BadRequest.New("test error")},
						{0, errors.BadRequest.New("test error")},
						{0, errors.BadRequest.New("test error")},
					},
				}

				requestNumber := 0

				// Делаем запросы к DSP
				for dspIndex, responses := range dspResponses {
					dsp := dsps[dspIndex]

					for _, response := range responses {

						// Идентификаторы сайта, юзера и запроса для запроса в DSP
						uuid.AddMockValues("id", "id", "id")

						// Получаем заранее подготовленный ответ DSP
						returnBidResponse := getDSPResponse(dto.testNumber, response.responseFile)
						returnErr := response.err

						// Добавляем ответ из файла в массив ответов
						responsesMap[response.responseFile] = returnBidResponse

						// Формируем запрос к DSP из шаблона
						bidRequestForDSP := changeReqForDSP(bidRequestTemplate, dspIndex)

						// Отправляем запрос в моке
						requestNumber++
						mocks.ExchangeNetwork.MockSendBidRequestToDSP(requestNumber, dsp.URL, bidRequestForDSP,
							returnBidResponse, 200, returnErr)

						bidRequestForDSP.Ext = testUtils.IgnoreErrorWithArgument(json.Marshal(map[string]string{
							"dsp-slug": dsp.Slug,
							"ssp-slug": ssp.Slug,
						}))

						mocks.EventService.MockCreateExchangeBidRequestToDSPEvent(bidRequestForDSP, requestNumber,
							nil)
					}
				}

				// Идентификаторы бидов для сохранения в монгу
				uuid.AddMockValues("bidID", "bidID", "bidID", "bidID")

				// Создаем запись c ответом DSP и доп данными в базе данных
				mongoDSPResponses := []model.DSPResponse{
					{ // 0
						ExchangeID:                "mongoRequestID",
						ExchangeBidID:             "bidID",
						BidResponse:               responsesMap[4], // Ставка 0.45 доллара
						SeatBidIndex:              0,
						BidIndex:                  0,
						BillingPriceInDSPCurrency: decimal.NewFromFloat(0.433), // 43.3 рубля из-за аукциона второй цены
						SlugDSP:                   "2",                         // Аукцион второй цены
						RecordDateTime:            time.Time{},
						OriginRequestID:           originRequestID,
						CurrencySSPCoefficient:    decimal.NewFromInt(100),
						SSPCurrency:               sspCurrency,
						SlugSSP:                   defaultSSPSlug,
						CurrencyDSPCoefficient:    decimal.NewFromInt(100),
						DSPCurrency:               dsp2Currency,
						RequestPublisherID:        "1511441620319572",
					},
					{ // 1
						ExchangeID:                "mongoRequestID",
						ExchangeBidID:             "bidID",
						BidResponse:               responsesMap[2], // Ставка 43.3 рубля
						SeatBidIndex:              0,
						BidIndex:                  0,
						BillingPriceInDSPCurrency: decimal.NewFromFloat(43.3), // Своя ставка
						SlugDSP:                   "1",                        // Аукцион первой цены
						RecordDateTime:            time.Time{},
						OriginRequestID:           originRequestID,
						CurrencySSPCoefficient:    decimal.NewFromInt(100),
						SSPCurrency:               sspCurrency,
						SlugSSP:                   defaultSSPSlug,
						CurrencyDSPCoefficient:    decimal.NewFromFloat(1),
						DSPCurrency:               dsp1Currency,
						RequestPublisherID:        "1511441620319572",
					},
					{ // 2
						ExchangeID:                "mongoRequestID",
						ExchangeBidID:             "bidID",
						BidResponse:               responsesMap[3], // Ставка 0.4 доллара
						SeatBidIndex:              0,
						BidIndex:                  0,
						BillingPriceInDSPCurrency: decimal.NewFromFloat(0.234), // 23.4 рублей из-за аукциона второй цены
						SlugDSP:                   "2",                         // Аукцион второй цены
						RecordDateTime:            time.Time{},
						OriginRequestID:           originRequestID,
						CurrencySSPCoefficient:    decimal.NewFromInt(100),
						SSPCurrency:               sspCurrency,
						SlugSSP:                   defaultSSPSlug,
						CurrencyDSPCoefficient:    decimal.NewFromFloat(100),
						DSPCurrency:               dsp2Currency,
						RequestPublisherID:        "1511441620319572",
					},
					{ // 3
						ExchangeID:                "mongoRequestID",
						ExchangeBidID:             "bidID",
						BidResponse:               responsesMap[1], // Ставка 23.4 рубля
						SeatBidIndex:              0,
						BidIndex:                  0,
						BillingPriceInDSPCurrency: decimal.NewFromFloat(23.4), // Своя ставка
						SlugDSP:                   "1",                        // Аукцион первой цены
						RecordDateTime:            time.Time{},
						OriginRequestID:           originRequestID,
						CurrencySSPCoefficient:    decimal.NewFromInt(100),
						SSPCurrency:               sspCurrency,
						SlugSSP:                   defaultSSPSlug,
						CurrencyDSPCoefficient:    decimal.NewFromFloat(1),
						DSPCurrency:               dsp1Currency,
						RequestPublisherID:        "1511441620319572",
					},
				}

				for i, mongoDSPResponse := range mongoDSPResponses {
					mocks.ExchangeRepository.MockCreateDSPResponse(i, mongoDSPResponse,
						nil)
				}
				for _, mongoDSPResponse := range mongoDSPResponses {
					mocks.ExchangeRepository.MockSaveADM(mongoDSPResponse.ExchangeBidID, mongoDSPResponse.BidResponse.SeatBids[mongoDSPResponse.SeatBidIndex].Bids[mongoDSPResponse.BidIndex].AdMarkup, nil)
				}

				// Идентификатор биддера и бида в ответе к SSP
				uuid.AddMockValues("bidder_id", "seatBid_0_bid_0")

				// Отправляем наш ответ в кафку
				sspBidResponse := openrtb.BidResponse{
					ID: originRequestID, // Идентификатор запроса
					SeatBids: []openrtb.SeatBid{
						{
							Bids: []openrtb.Bid{
								{
									ID:         "seatBid_0_bid_0",            // Новый идентификатор
									ImpID:      "1",                          // Идентификатор первого импрешена
									Price:      decimal.NewFromFloat(1.2006), // Сумма всех ставок
									NoticeURL:  "https://test.host.com/billing/mongoRequestID?id_type=request&price=${AUCTION_PRICE}&ssp_slug=test_slug&url_type=nurl&bid_price=1.2006",
									AdMarkup:   "<?xml version=\"1.0\" encoding=\"ISO-8859-1\"?>\n<ad><popunderAd><url><![CDATA[https://test.showcase.com/game/happy-bucket/play?ad_type=banner&adm_id=bidID,bidID,bidID,bidID&is_adult=true]]></url></popunderAd></ad>",
									AdvDomains: []string{"vidimo.media"},
									ImageURL:   "https://cdn.adx.com.ru/banner/0000000000000000.jpg",
									CampaingID: "66c7182ba897d80001520631",
									CreativeID: "66cd9992a897d800015d8af7",
								},
							},
						},
					},
					BidID:    "bidder_id", // Новый идентификатор
					Currency: sspCurrency, // Валюта SSP
				}

				mocks.EventService.MockCreateExchangeBidResponseEvent(sspBidResponse,
					nil)

				return sspBidResponse
			},
		},
		{
			testNumber: 2,
			name:       "Отработка успешного запроса на баннерный запрос, синглимпрешен",
			args: args{
				isSendNoContentInsteadErrorsInResponse: false,
				isSendNoContentForSuccessResponse:      false,
				wantHTTPCode:                           http.StatusOK,
			},
			mockActions: func(mocks service.Mocks, dto mockDTO) openrtb.BidResponse {

				// Фиксируем пришедший запрос в кафку
				createSSPBidEventReq := dto.sspBidRequest
				createSSPBidEventReq.Ext = []byte(fmt.Sprintf(`{"ssp-slug":"%s"}`, defaultSSPSlug))
				mocks.EventService.MockCreateSSPBidRequestEvent(createSSPBidEventReq,
					nil)

				const sspCurrency = "USD"

				// Получаем SSP согласно запросу, в нашем случае тестируется кликандер, поэтому он true
				ssp := sspModel.SSP{
					Slug:               defaultSSPSlug,
					Timeout:            nil,
					IsEnable:           true,
					ClickunderDrumSize: pointer.Pointer(int32(4)), // Допустим, кто-то забыл убрать
				}
				mocks.SSPService.MockGetSSP(sspModel.GetSSPsReq{
					Slugs: []string{defaultSSPSlug},
				},
					ssp, nil)

				// Получаем настройки системы, они общие на все тесткейсы
				settings := generalSettings
				mocks.SettingService.MockGetSettings(
					settings, nil)

				// Получаем курсы валют, они общие на все тесткейсы
				currencyRates := generalCurrencyRates
				mocks.CurrencyService.MockGetRates(
					currencyRates, nil)

				// Наш общий идентификатор запроса
				uuid.AddMockValues("mongoRequestID")

				mocks.ExchangeRepository.MockGetPublisherVisibility("1511441620319572",
					model.PublisherVisibility{
						PublisherID: "1511441620319572",
						Loads:       100,
						Views:       69,
					}, nil)

				// Генерируем DSP согласно входным данным тесткейса

				const (
					dsp1Currency = "RUB"
					dsp2Currency = "USD"
					dsp3Currency = "EUR"
				)

				dsps := []dspModel.DSP{
					{
						Slug:                     "1",
						URL:                      "https://dsp1.com",
						Currency:                 dsp1Currency,
						AuctionSecondPrice:       true,
						IsSupportMultiimpression: false,
					},
					{
						Slug:                     "2",
						URL:                      "https://dsp2.com",
						Currency:                 dsp2Currency,
						AuctionSecondPrice:       true,
						IsSupportMultiimpression: false,
					},
					{
						Slug:                     "3",
						URL:                      "https://dsp3.com",
						Currency:                 dsp3Currency,
						IsSupportMultiimpression: false,
					},
				}
				mocks.DSPService.MockGetDSPs(dspModel.GetDSPsReq{
					SourceTrafficType: []sourceTrafficType.SourceTrafficType{sourceTrafficType.Desktop},
					IsEnable:          pointer.Pointer(true),
				},
					dsps, nil)

				// Чекаем в одном из чейнов страну по IP с вызовом в геолайт
				mocks.ExchangeRepository.MockGetCountryByIP(dto.sspBidRequest.Device.IP,
					"RUS", nil)

				changeReqForDSP := func(req openrtb.BidRequest, dspIndex int) openrtb.BidRequest {
					copyReq := req.Copy()
					switch dspIndex {
					case 0:
						copyReq.Impressions[0].BidFloor = decimal.NewFromFloat(11) // 0.1 * 100 + 10%
						copyReq.Impressions[0].BidFloorCurrency = dsp1Currency
						copyReq.Currencies = []string{dsp1Currency}
						copyReq.AuctionType = 2
					case 1:
						copyReq.Impressions[0].BidFloor = decimal.NewFromFloat(0.11) // 0.1 + 10%
						copyReq.Impressions[0].BidFloorCurrency = dsp2Currency
						copyReq.Currencies = []string{dsp2Currency}
						copyReq.AuctionType = 2
					case 2:
						copyReq.Impressions[0].BidFloor = decimal.NewFromFloat(0.1) // 0.1 / 1.1 + 10%
						copyReq.Impressions[0].BidFloorCurrency = dsp3Currency
						copyReq.Currencies = []string{dsp3Currency}
						copyReq.AuctionType = 1
					}
					return copyReq
				}

				// Получаем ожидаемый запрос к DSP согласно входным данным
				const originRequestID = "50c3d53cc39690b5febd02db6b28b206"
				bidRequestTemplate := openrtb.BidRequest{
					ID: "id",
					Impressions: []openrtb.Impression{
						{
							ID: "1",
							Banner: &openrtb.Banner{
								Width:        320,
								Height:       480,
								BlockedTypes: []openrtb.BannerType{},
								BlockedAttrs: []openrtb.CreativeAttribute{1, 2, 5, 8, 9, 14, 17},
								Position:     7,
								MIMEs: []string{
									"text/javascript",
									"application/javascript",
									"image/jpeg",
									"image/png",
									"image/gif",
								},
								APIs: []openrtb.APIFramework{3, 5, 6, 7},
							},
							BidFloor:         decimal.Decimal{}, // Будет заполняться в фукнции
							BidFloorCurrency: "",                // Будет заполняться в фукнции
							Ext:              []byte(`{"ad_type":40,"limit":1}`),
						},
					},
					Site: &openrtb.Site{ // Добавили из шаблона
						Inventory: openrtb.Inventory{
							ID:           "1511441620319572",
							Name:         "",
							Domain:       "bco3h97c.xyz",
							Categories:   []openrtb.ContentCategory{"IAB24"},
							PageCategory: []openrtb.ContentCategory{"IAB24"},
							Publisher: &openrtb.Publisher{
								ID:   "51026",
								Name: "bco3h97c.xyz",
							},
						},
						Page: "http://bco3h97c.xyz?",
					},
					Device: &openrtb.Device{
						UserAgent: "Mozilla/5.0 (Linux; Android 10; K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Mobile Safari/537.36",
						Geo: &openrtb.Geo{
							Country: "RUS",
						},
						IP:         "106.222.213.16",
						DeviceType: 4,
						Language:   "en",
					},
					User: &openrtb.User{
						ID: "id",
						Geo: &openrtb.Geo{
							Country: "RUS",
						},
					},
					AuctionType: 0,          // Будет заполняться в фукнции
					TimeMax:     132,        // 147 - 10%
					Currencies:  []string{}, // Будет заполняться в фукнции
					Ext:         []byte(`{"ssp-slug":"kadam"}`),
				}

				// Номер DSP, номер ответа
				responsesMap := make(map[int]openrtb.BidResponse)

				// Индекс DSP - ответ
				dspResponses := [3]response{
					{1, nil}, // 203.2 рубля // 1 место
					{2, nil}, // 0.4 доллара // 2 место
					{0, errors.BadRequest.New("test error")},
				}

				requestNumber := 0

				// Делаем запросы к DSP
				for dspIndex, response := range dspResponses {
					dsp := dsps[dspIndex]

					// Идентификатор юзера и запроса для запроса в DSP
					uuid.AddMockValues("id", "id")

					// Получаем заранее подготовленный ответ DSP
					returnBidResponse := getDSPResponse(dto.testNumber, response.responseFile)
					returnErr := response.err

					// Добавляем ответ из файла в массив ответов
					responsesMap[response.responseFile] = returnBidResponse

					// Формируем запрос к DSP из шаблона
					bidRequestForDSP := changeReqForDSP(bidRequestTemplate, dspIndex)

					// Отправляем запрос в моке
					requestNumber++
					mocks.ExchangeNetwork.MockSendBidRequestToDSP(requestNumber, dsp.URL, bidRequestForDSP,
						returnBidResponse, 200, returnErr)

					bidRequestForDSP.Ext = testUtils.IgnoreErrorWithArgument(json.Marshal(map[string]string{
						"dsp-slug": dsp.Slug,
						"ssp-slug": ssp.Slug,
					}))
					mocks.EventService.MockCreateExchangeBidRequestToDSPEvent(bidRequestForDSP, requestNumber,
						nil)
				}

				// Идентификатор бидов для монги
				uuid.AddMockValues("bidID", "bidID")

				// Создаем запись c ответом DSP и доп данными в базе данных
				mongoDSPResponse := model.DSPResponse{
					ExchangeID:                "mongoRequestID",
					ExchangeBidID:             "bidID",
					BidResponse:               responsesMap[1], // Ставка 203.2 рубля
					SeatBidIndex:              0,
					BidIndex:                  0,
					BillingPriceInDSPCurrency: decimal.NewFromFloat(40), // 40 рублей из-за аукциона второй цены
					SlugDSP:                   "1",                      // Аукцион второй цены
					RecordDateTime:            time.Time{},
					OriginRequestID:           originRequestID,
					CurrencySSPCoefficient:    decimal.NewFromInt(100),
					SSPCurrency:               sspCurrency,
					SlugSSP:                   defaultSSPSlug,
					CurrencyDSPCoefficient:    decimal.NewFromInt(1),
					DSPCurrency:               dsp1Currency,
					RequestPublisherID:        "1511441620319572",
				}

				mocks.ExchangeRepository.MockCreateDSPResponse(0, mongoDSPResponse,
					nil)

				// Идентификатор биддера и бида для ответа в SSP
				uuid.AddMockValues("bidder_id", "seatBid_0_bid_0")

				// Отправляем наш ответ в кафку
				sspBidResponse := openrtb.BidResponse{
					ID: originRequestID, // Идентификатор запроса
					SeatBids: []openrtb.SeatBid{
						{
							Bids: []openrtb.Bid{
								{
									ID:         "seatBid_0_bid_0",          // Новый идентификатор
									ImpID:      "1",                        // Идентификатор первого импрешена
									Price:      decimal.NewFromFloat(0.36), // Ставка, по которой будем биллить DSP
									NoticeURL:  "https://test.host.com/billing/bidID?id_type=bid&price=${AUCTION_PRICE}&ssp_slug=test_slug&url_type=nurl",
									AdMarkup:   "\n<!DOCTYPE html>\n<html>\n<head>\n    <meta name=\"viewport\" content=\"width=device-width, initial-scale=1, user-scalable=no\">\n\n    <style>\n        html, body {\n            border: 0 !important;\n            margin: 0 !important;\n            padding: 0 !important;\n            width: 100vw !important;\n            height: 100vh !important;\n            overflow: hidden !important;\n        }\n\n        #iframe-66cf23ee7011cb000191116b {\n            border: 0;\n            \n            \n                width: 100%;\n                height: 100%;\n            \n        }\n    </style>\n</head>\n<body>\n    \n\n    \n\n    \n\n    \n<script>\n    (() => {\n        const initMraid = (window) => {\n            try {\n                mraid;\n            } catch {\n                mraidConnect(window, false);\n                return;\n            }\n\n            if (mraid.getState() === 'loading') {\n                mraid.addEventListener('ready', () => {\n                    mraidConnect(window, true);\n                });\n            } else {\n                mraidConnect(window, true);\n            }\n        };\n\n        const makeOptions = () => ({\n            placementType: mraid.getPlacementType(),\n            expandProperties: mraid.getExpandProperties(),\n            orientationProperties: mraid.getOrientationProperties(),\n            resizeProperties: mraid.getResizeProperties(),\n            maxSize: mraid.getMaxSize(),\n            features: {\n                sms: mraid.supports('sms'),\n                tel: mraid.supports('tel'),\n                calendar: mraid.supports('calendar'),\n                storePicture: mraid.supports('storePicture'),\n                inlineVideo: mraid.supports('inlineVideo'),\n            },\n        });\n\n        const mraidConnect = (eventSource, proxy) => {\n            eventSource.postMessage({ event: 'init', data: proxy, options: proxy ? makeOptions() : undefined }, '*');\n\n            if (proxy) {\n                mraid.addEventListener('error', (message, action) => {\n                    eventSource.postMessage({ event: 'error', data: { message, action, options: makeOptions() } }, '*');\n                });\n\n                mraid.addEventListener('sizeChange', (width, height) => {\n                    eventSource.postMessage({ event: 'sizeChange', data: { width, height }, options: makeOptions() }, '*');\n                });\n\n                mraid.addEventListener('stateChange', (state) => {\n                    eventSource.postMessage({ event: 'stateChange', data: state, options: makeOptions() }, '*');\n                });\n\n                mraid.addEventListener('viewableChange', (viewable) => {\n                    eventSource.postMessage({ event: 'viewableChange', data: viewable, options: makeOptions() }, '*');\n                });\n            }\n        };\n\n        window.addEventListener('message', (event) => {\n            if (event.source !== window.frames.wrapper.frames.banner) {\n                return;\n            }\n\n            switch (event.data.event) {\n                case 'init':\n                    initMraid(event.source);\n                    break;\n                case 'open':\n                    mraid.open(event.data.data);\n                    break;\n                case 'close':\n                    mraid.close();\n                    break;\n                case 'setExpandProperties':\n                    mraid.setExpandProperties(event.data.data);\n                    break;\n                case 'setOrientationProperties':\n                    mraid.setOrientationProperties(event.data.data);\n                    break;\n                case 'setResizeProperties':\n                    mraid.setResizeProperties(event.data.data);\n                    break;\n                case 'useCustomClose':\n                    mraid.useCustomClose(event.data.data);\n                    break;\n                case 'storePicture':\n                    mraid.storePicture(event.data.data);\n                    break;\n                case 'createCalendarEvent':\n                    mraid.createCalendarEvent(event.data.data);\n                    break;\n                case 'playVideo':\n                    mraid.playVideo(event.data.data);\n                    break;\n            }\n        });\n    })()\n</script>\n\n    <script>\n    document.createElement('img').src = 'https://statistics.adx.com.ru/api/stream/device' +\n        `?dpr=${devicePixelRatio}` +\n        `&width=${screen.width}` +\n        `&height=${screen.height}`;\n</script>\n\n\n    <iframe name=\"wrapper\" id=\"iframe-66cf23ee7011cb000191116b\" scrolling=\"no\" marginwidth=\"0\" marginheight=\"0\" src=\"https://adx.com.ru/banner/66cf23ee7011cb000191116b/?clickMacro=&auctionPrice=${AUCTION_PRICE}\"></iframe>\n\n    <script>\n        const frameWrapper = document.querySelector('iframe[name=\"wrapper\"]');\n        if (frameWrapper.dataset.params !== undefined) {\n            fetch('https://adx.com.ru/banner/66cf23ee7011cb000191116b/').then(resp => resp.text()).then(text => {\n                frameWrapper.srcdoc = text;\n            });\n        }\n    </script>\n</body>\n</html>\n",
									AdvDomains: []string{"vidimo.media"},
									ImageURL:   "https://cdn.adx.com.ru/banner/0000000000000000.jpg",
									CampaingID: "66c7182ba897d80001520631",
									CreativeID: "66cd9992a897d800015d8af7",
								},
							},
						},
					},
					BidID:    "bidder_id", // Новый идентификатор
					Currency: sspCurrency, // Валюта SSP
				}

				mocks.EventService.MockCreateExchangeBidResponseEvent(sspBidResponse,
					nil)

				return sspBidResponse
			},
		},
		{
			testNumber: 3,
			name:       "Успешная отработка на мультиимпрешен запрос. 3 импрешена",
			args: args{
				isSendNoContentInsteadErrorsInResponse: false,
				isSendNoContentForSuccessResponse:      false,
				wantHTTPCode:                           http.StatusOK,
			},
			mockActions: func(mocks service.Mocks, dto mockDTO) openrtb.BidResponse {

				// Фиксируем пришедший запрос в кафку
				createSSPBidEventReq := dto.sspBidRequest
				createSSPBidEventReq.Ext = []byte(fmt.Sprintf(`{"ssp-slug":"%s"}`, defaultSSPSlug))
				mocks.EventService.MockCreateSSPBidRequestEvent(createSSPBidEventReq,
					nil)

				const sspCurrency = "USD"

				ssp := sspModel.SSP{
					Slug:               defaultSSPSlug,
					ClickunderDrumSize: pointer.Pointer(int32(4)),
					Timeout:            nil,
					IsEnable:           true,
				}
				mocks.SSPService.MockGetSSP(sspModel.GetSSPsReq{
					Slugs: []string{defaultSSPSlug},
				},
					ssp, nil)

				// Получаем настройки системы, они общие на все тесткейсы
				settings := generalSettings
				mocks.SettingService.MockGetSettings(
					settings, nil)

				// Получаем курсы валют, они общие на все тесткейсы
				currencyRates := generalCurrencyRates
				mocks.CurrencyService.MockGetRates(
					currencyRates, nil)

				// Наш общий новый идентификатор запроса
				uuid.AddMockValues("mongoRequestID")

				mocks.ExchangeRepository.MockGetPublisherVisibility("1511441620319572",
					model.PublisherVisibility{
						PublisherID: "1511441620319572",
						Loads:       100,
						Views:       69,
					}, nil)

				const (
					dsp1Currency = "RUB"
					dsp2Currency = "USD"
					dsp3Currency = "EUR"
				)

				// Генерируем DSP согласно входным данным тесткейса
				dsps := []dspModel.DSP{
					{
						Slug:                     "1",
						URL:                      "https://dsp1.com",
						Currency:                 dsp1Currency,
						AuctionSecondPrice:       false,
						IsSupportMultiimpression: true,
					},
					{
						Slug:                     "2",
						URL:                      "https://dsp2.com",
						Currency:                 dsp2Currency,
						AuctionSecondPrice:       true,
						IsSupportMultiimpression: false,
					},
					{
						Slug:                     "3",
						URL:                      "https://dsp3.com",
						Currency:                 dsp3Currency,
						AuctionSecondPrice:       false,
						IsSupportMultiimpression: false,
					},
				}
				mocks.DSPService.MockGetDSPs(dspModel.GetDSPsReq{
					SourceTrafficType: []sourceTrafficType.SourceTrafficType{sourceTrafficType.Desktop},
					IsEnable:          pointer.Pointer(true),
				},
					dsps, nil)

				// Чекаем в одном из чейнов страну по IP с вызовом в геолайт
				mocks.ExchangeRepository.MockGetCountryByIP(dto.sspBidRequest.Device.IP,
					"RUS", nil)

				changeReqForDSP := func(req openrtb.BidRequest, dspIndex int, callNumber int) openrtb.BidRequest {
					copyReq := req.Copy()

					switch dspIndex {
					case 0:
						if callNumber == 1 {
							copyReq.Impressions[0].BidFloor = decimal.NewFromFloat(11) // 0.1 * 100 + 10%
							copyReq.Impressions[0].BidFloorCurrency = dsp1Currency

							copyReq1 := req.Copy()
							copyReq1.Impressions[0].BidFloor = decimal.NewFromFloat(22) // 0.2 * 100 + 10%
							copyReq1.Impressions[0].BidFloorCurrency = dsp1Currency
							copyReq1.Impressions[0].ID = "2"
							copyReq.Impressions = append(copyReq.Impressions, copyReq1.Impressions[0])

							copyReq2 := req.Copy()
							copyReq2.Impressions[0].BidFloor = decimal.NewFromFloat(33) // 0.3 * 100 + 10%
							copyReq2.Impressions[0].BidFloorCurrency = dsp1Currency
							copyReq2.Impressions[0].ID = "3"
							copyReq.Impressions = append(copyReq.Impressions, copyReq2.Impressions[0])

							copyReq.Currencies = []string{dsp1Currency}
							copyReq.AuctionType = 1
						}
					case 1:
						switch callNumber {
						case 1:
							copyReq.Impressions[0].BidFloor = decimal.NewFromFloat(0.11) // 0.1 + 10%
							copyReq.Impressions[0].BidFloorCurrency = dsp2Currency
							copyReq.Currencies = []string{dsp2Currency}
							copyReq.AuctionType = 2
						case 2:
							copyReq.Impressions[0].BidFloor = decimal.NewFromFloat(0.22) // 0.2 + 10%
							copyReq.Impressions[0].BidFloorCurrency = dsp2Currency
							copyReq.Impressions[0].ID = "2"
							copyReq.Currencies = []string{dsp2Currency}
							copyReq.AuctionType = 2
						case 3:
							copyReq.Impressions[0].BidFloor = decimal.NewFromFloat(0.33) // 0.3 + 10%
							copyReq.Impressions[0].BidFloorCurrency = dsp2Currency
							copyReq.Impressions[0].ID = "3"
							copyReq.Currencies = []string{dsp2Currency}
							copyReq.AuctionType = 2
						}
					case 2:
						switch callNumber {
						case 1:
							copyReq.Impressions[0].BidFloor = decimal.NewFromFloat(0.1) // 0.1 / 1.1 + 10%
							copyReq.Impressions[0].BidFloorCurrency = dsp3Currency
							copyReq.Currencies = []string{dsp3Currency}
							copyReq.AuctionType = 1
						case 2:
							copyReq.Impressions[0].BidFloor = decimal.NewFromFloat(0.2) // 0.2 / 1.1 + 10%
							copyReq.Impressions[0].BidFloorCurrency = dsp3Currency
							copyReq.Impressions[0].ID = "2"
							copyReq.Currencies = []string{dsp3Currency}
							copyReq.AuctionType = 1
						case 3:
							copyReq.Impressions[0].BidFloor = decimal.NewFromFloat(0.3) // 0.2 / 1.1 + 10%
							copyReq.Impressions[0].BidFloorCurrency = dsp3Currency
							copyReq.Impressions[0].ID = "3"
							copyReq.Currencies = []string{dsp3Currency}
							copyReq.AuctionType = 1
						}
					}
					return copyReq
				}

				// Получаем ожидаемый запрос к DSP согласно входным данным
				const originRequestID = "50c3d53cc39690b5febd02db6b28b206"
				bidRequestTemplate := openrtb.BidRequest{
					ID: "id",
					Impressions: []openrtb.Impression{
						{
							ID: "1",
							Banner: &openrtb.Banner{
								Width:        320,
								Height:       480,
								BlockedTypes: []openrtb.BannerType{},
								BlockedAttrs: []openrtb.CreativeAttribute{1, 2, 5, 8, 9, 14, 17},
								Position:     7,
								MIMEs: []string{
									"text/javascript",
									"application/javascript",
									"image/jpeg",
									"image/png",
									"image/gif",
								},
								APIs: []openrtb.APIFramework{3, 5, 6, 7},
							},
							BidFloor:         decimal.Decimal{}, // Будет заполняться в фукнции
							BidFloorCurrency: "",                // Будет заполняться в фукнции
							Ext:              []byte(`{"ad_type":40,"limit":1}`),
						},
					},
					Site: &openrtb.Site{ // Добавили из шаблона
						Inventory: openrtb.Inventory{
							ID:           "1511441620319572",
							Name:         "",
							Domain:       "bco3h97c.xyz",
							Categories:   []openrtb.ContentCategory{"IAB24"},
							PageCategory: []openrtb.ContentCategory{"IAB24"},
							Publisher: &openrtb.Publisher{
								ID:   "51026",
								Name: "bco3h97c.xyz",
							},
						},
						Page: "http://bco3h97c.xyz?",
					},
					Device: &openrtb.Device{
						UserAgent: "Mozilla/5.0 (Linux; Android 10; K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Mobile Safari/537.36",
						Geo: &openrtb.Geo{
							Country: "RUS",
						},
						IP:         "106.222.213.16",
						DeviceType: 4,
						Language:   "en",
					},
					User: &openrtb.User{
						ID: "id",
						Geo: &openrtb.Geo{
							Country: "RUS",
						},
					},
					AuctionType: 0,          // Будет заполняться в фукнции
					TimeMax:     132,        // 147 - 10%
					Currencies:  []string{}, // Будет заполняться в фукнции
					Ext:         []byte(`{"ssp-slug":"kadam"}`),
				}

				// Номер файла, номер ответа
				responsesMap := make(map[int]openrtb.BidResponse)

				// Индекс DSP - Номер вызова - ответ
				dspResponses := [3][]response{
					{
						// 1 импрешен - 23.4 рубля - 2 место 1 импрешен
						// 2 импрешен - 40.4 рублей - 3 место 2 импрешен
						// 3 импрешен - 65 рублей - 1 место 3 импрешен *** 3
						{1, nil},
					},
					{
						{2, nil}, // 1 импрешен - 0.44 доллара - 1 место 1 импрешен *** 1
						{3, nil}, // 2 импрешен - 0.5 доллара - 2 место 2 импрешен
						{4, nil}, // 3 импрешен - 0.3 доллара - 2 место 3 импрешен
					},
					{
						{0, errors.BadRequest.New("test error")},
						{5, nil}, // 2 импрешен - 0.8 евро - 1 место 2 импрешен *** 2
						{0, errors.BadRequest.New("test error")},
					},
				}

				requestNumber := 0

				// Делаем запросы к DSP
				for dspIndex, responses := range dspResponses {
					dsp := dsps[dspIndex]

					for callNumber, response := range responses {

						// Идентификаторы юзера и запроса для запроса в DSP
						uuid.AddMockValues("id", "id")

						// Получаем заранее подготовленный ответ DSP
						returnBidResponse := getDSPResponse(dto.testNumber, response.responseFile)
						returnErr := response.err

						// Добавляем ответ из файла в массив ответов
						responsesMap[response.responseFile] = returnBidResponse

						// Формируем запрос к DSP из шаблона
						bidRequestForDSP := changeReqForDSP(bidRequestTemplate, dspIndex, callNumber+1)

						// Отправляем запрос в моке
						requestNumber++
						mocks.ExchangeNetwork.MockSendBidRequestToDSP(requestNumber, dsp.URL, bidRequestForDSP,
							returnBidResponse, 200, returnErr)

						bidRequestForDSP.Ext = testUtils.IgnoreErrorWithArgument(json.Marshal(map[string]string{
							"dsp-slug": dsp.Slug,
							"ssp-slug": ssp.Slug,
						}))
						mocks.EventService.MockCreateExchangeBidRequestToDSPEvent(bidRequestForDSP, requestNumber,
							nil)
					}
				}

				// Идентификаторы бидов для монги
				uuid.AddMockValues("bidID", "bidID", "bidID", "bidID", "bidID", "bidID", "bidID")

				// Создаем запись c ответом DSP и доп данными в базе данных
				mongoDSPResponses := []model.DSPResponse{
					{ // 0
						ExchangeID:                "mongoRequestID",
						ExchangeBidID:             "bidID",
						BidResponse:               responsesMap[2], // Ставка 0.44 доллара
						SeatBidIndex:              0,
						BidIndex:                  0,
						BillingPriceInDSPCurrency: decimal.NewFromFloat(0.234), // 24.3 рубля из-за аукциона второй цены
						SlugDSP:                   "2",                         // Аукцион второй цены
						RecordDateTime:            time.Time{},
						OriginRequestID:           originRequestID,
						CurrencySSPCoefficient:    decimal.NewFromInt(100),
						SSPCurrency:               sspCurrency,
						SlugSSP:                   defaultSSPSlug,
						CurrencyDSPCoefficient:    decimal.NewFromInt(100),
						DSPCurrency:               dsp2Currency,
						RequestPublisherID:        "1511441620319572",
					},
					{ // 1
						ExchangeID:                "mongoRequestID",
						ExchangeBidID:             "bidID",
						BidResponse:               responsesMap[5], // Ставка 0.8 евро
						SeatBidIndex:              0,
						BidIndex:                  0,
						BillingPriceInDSPCurrency: decimal.NewFromFloat(0.8), // Своя ставка
						SlugDSP:                   "3",                       // Аукцион первой цены
						RecordDateTime:            time.Time{},
						OriginRequestID:           originRequestID,
						CurrencySSPCoefficient:    decimal.NewFromInt(100),
						SSPCurrency:               sspCurrency,
						SlugSSP:                   defaultSSPSlug,
						CurrencyDSPCoefficient:    decimal.NewFromFloat(110),
						DSPCurrency:               dsp3Currency,
						RequestPublisherID:        "1511441620319572",
					},
					{ // 2
						ExchangeID:                "mongoRequestID",
						ExchangeBidID:             "bidID",
						BidResponse:               responsesMap[1], // Ставка 65 рублей
						SeatBidIndex:              0,
						BidIndex:                  2,
						BillingPriceInDSPCurrency: decimal.NewFromFloat(65), // Своя ставка
						SlugDSP:                   "1",                      // Аукцион первой цены
						RecordDateTime:            time.Time{},
						OriginRequestID:           originRequestID,
						CurrencySSPCoefficient:    decimal.NewFromInt(100),
						SSPCurrency:               sspCurrency,
						SlugSSP:                   defaultSSPSlug,
						CurrencyDSPCoefficient:    decimal.NewFromFloat(1),
						DSPCurrency:               dsp1Currency,
						RequestPublisherID:        "1511441620319572",
					},
				}

				for i, mongoDSPResponse := range mongoDSPResponses {
					mocks.ExchangeRepository.MockCreateDSPResponse(i, mongoDSPResponse,
						nil)
				}

				// Идентификатор биддера и бидов для ответа в SSP
				uuid.AddMockValues("bidder_id", "seatBid_0_bid_0", "seatBid_0_bid_1", "seatBid_0_bid_2")

				// Отправляем наш ответ в кафку
				sspBidResponse := openrtb.BidResponse{
					ID: originRequestID, // Идентификатор запроса
					SeatBids: []openrtb.SeatBid{
						{
							Bids: []openrtb.Bid{
								{
									ID:         "seatBid_0_bid_0",            // Новый идентификатор
									ImpID:      "1",                          // Идентификатор первого импрешена
									Price:      decimal.NewFromFloat(0.2106), // Ставка, по которой будем биллить DSP
									NoticeURL:  "https://test.host.com/billing/bidID?id_type=bid&price=${AUCTION_PRICE}&ssp_slug=test_slug&url_type=nurl",
									AdMarkup:   "\n<!DOCTYPE html>\n<html>\n<head>\n    <meta name=\"viewport\" content=\"width=device-width, initial-scale=1, user-scalable=no\">\n\n    <style>\n        html, body {\n            border: 0 !important;\n            margin: 0 !important;\n            padding: 0 !important;\n            width: 100vw !important;\n            height: 100vh !important;\n            overflow: hidden !important;\n        }\n\n        #iframe-66cf23ee7011cb000191116b {\n            border: 0;\n            \n            \n                width: 100%;\n                height: 100%;\n            \n        }\n    </style>\n</head>\n<body>\n    \n\n    \n\n    \n\n    \n<script>\n    (() => {\n        const initMraid = (window) => {\n            try {\n                mraid;\n            } catch {\n                mraidConnect(window, false);\n                return;\n            }\n\n            if (mraid.getState() === 'loading') {\n                mraid.addEventListener('ready', () => {\n                    mraidConnect(window, true);\n                });\n            } else {\n                mraidConnect(window, true);\n            }\n        };\n\n        const makeOptions = () => ({\n            placementType: mraid.getPlacementType(),\n            expandProperties: mraid.getExpandProperties(),\n            orientationProperties: mraid.getOrientationProperties(),\n            resizeProperties: mraid.getResizeProperties(),\n            maxSize: mraid.getMaxSize(),\n            features: {\n                sms: mraid.supports('sms'),\n                tel: mraid.supports('tel'),\n                calendar: mraid.supports('calendar'),\n                storePicture: mraid.supports('storePicture'),\n                inlineVideo: mraid.supports('inlineVideo'),\n            },\n        });\n\n        const mraidConnect = (eventSource, proxy) => {\n            eventSource.postMessage({ event: 'init', data: proxy, options: proxy ? makeOptions() : undefined }, '*');\n\n            if (proxy) {\n                mraid.addEventListener('error', (message, action) => {\n                    eventSource.postMessage({ event: 'error', data: { message, action, options: makeOptions() } }, '*');\n                });\n\n                mraid.addEventListener('sizeChange', (width, height) => {\n                    eventSource.postMessage({ event: 'sizeChange', data: { width, height }, options: makeOptions() }, '*');\n                });\n\n                mraid.addEventListener('stateChange', (state) => {\n                    eventSource.postMessage({ event: 'stateChange', data: state, options: makeOptions() }, '*');\n                });\n\n                mraid.addEventListener('viewableChange', (viewable) => {\n                    eventSource.postMessage({ event: 'viewableChange', data: viewable, options: makeOptions() }, '*');\n                });\n            }\n        };\n\n        window.addEventListener('message', (event) => {\n            if (event.source !== window.frames.wrapper.frames.banner) {\n                return;\n            }\n\n            switch (event.data.event) {\n                case 'init':\n                    initMraid(event.source);\n                    break;\n                case 'open':\n                    mraid.open(event.data.data);\n                    break;\n                case 'close':\n                    mraid.close();\n                    break;\n                case 'setExpandProperties':\n                    mraid.setExpandProperties(event.data.data);\n                    break;\n                case 'setOrientationProperties':\n                    mraid.setOrientationProperties(event.data.data);\n                    break;\n                case 'setResizeProperties':\n                    mraid.setResizeProperties(event.data.data);\n                    break;\n                case 'useCustomClose':\n                    mraid.useCustomClose(event.data.data);\n                    break;\n                case 'storePicture':\n                    mraid.storePicture(event.data.data);\n                    break;\n                case 'createCalendarEvent':\n                    mraid.createCalendarEvent(event.data.data);\n                    break;\n                case 'playVideo':\n                    mraid.playVideo(event.data.data);\n                    break;\n            }\n        });\n    })()\n</script>\n\n    <script>\n    document.createElement('img').src = 'https://statistics.adx.com.ru/api/stream/device' +\n        `?dpr=${devicePixelRatio}` +\n        `&width=${screen.width}` +\n        `&height=${screen.height}`;\n</script>\n\n\n    <iframe name=\"wrapper\" id=\"iframe-66cf23ee7011cb000191116b\" scrolling=\"no\" marginwidth=\"0\" marginheight=\"0\" src=\"https://adx.com.ru/banner/66cf23ee7011cb000191116b/?clickMacro=&auctionPrice=${AUCTION_PRICE}\"></iframe>\n\n    <script>\n        const frameWrapper = document.querySelector('iframe[name=\"wrapper\"]');\n        if (frameWrapper.dataset.params !== undefined) {\n            fetch('https://adx.com.ru/banner/66cf23ee7011cb000191116b/').then(resp => resp.text()).then(text => {\n                frameWrapper.srcdoc = text;\n            });\n        }\n    </script>\n</body>\n</html>\n",
									AdvDomains: []string{"vidimo.media"},
									ImageURL:   "https://cdn.adx.com.ru/banner/0000000000000000.jpg",
									CampaingID: "66c7182ba897d80001520631",
									CreativeID: "66cd9992a897d800015d8af7",
								},
							},
						},
						{
							Bids: []openrtb.Bid{
								{
									ID:         "seatBid_0_bid_1",           // Новый идентификатор
									ImpID:      "2",                         // Идентификатор первого импрешена
									Price:      decimal.NewFromFloat(0.792), // Ставка, по которой будем биллить DSP
									NoticeURL:  "https://test.host.com/billing/bidID?id_type=bid&price=${AUCTION_PRICE}&ssp_slug=test_slug&url_type=nurl",
									AdMarkup:   "\n<!DOCTYPE html>\n<html>\n<head>\n    <meta name=\"viewport\" content=\"width=device-width, initial-scale=1, user-scalable=no\">\n\n    <style>\n        html, body {\n            border: 0 !important;\n            margin: 0 !important;\n            padding: 0 !important;\n            width: 100vw !important;\n            height: 100vh !important;\n            overflow: hidden !important;\n        }\n\n        #iframe-66cf23ee7011cb000191116b {\n            border: 0;\n            \n            \n                width: 100%;\n                height: 100%;\n            \n        }\n    </style>\n</head>\n<body>\n    \n\n    \n\n    \n\n    \n<script>\n    (() => {\n        const initMraid = (window) => {\n            try {\n                mraid;\n            } catch {\n                mraidConnect(window, false);\n                return;\n            }\n\n            if (mraid.getState() === 'loading') {\n                mraid.addEventListener('ready', () => {\n                    mraidConnect(window, true);\n                });\n            } else {\n                mraidConnect(window, true);\n            }\n        };\n\n        const makeOptions = () => ({\n            placementType: mraid.getPlacementType(),\n            expandProperties: mraid.getExpandProperties(),\n            orientationProperties: mraid.getOrientationProperties(),\n            resizeProperties: mraid.getResizeProperties(),\n            maxSize: mraid.getMaxSize(),\n            features: {\n                sms: mraid.supports('sms'),\n                tel: mraid.supports('tel'),\n                calendar: mraid.supports('calendar'),\n                storePicture: mraid.supports('storePicture'),\n                inlineVideo: mraid.supports('inlineVideo'),\n            },\n        });\n\n        const mraidConnect = (eventSource, proxy) => {\n            eventSource.postMessage({ event: 'init', data: proxy, options: proxy ? makeOptions() : undefined }, '*');\n\n            if (proxy) {\n                mraid.addEventListener('error', (message, action) => {\n                    eventSource.postMessage({ event: 'error', data: { message, action, options: makeOptions() } }, '*');\n                });\n\n                mraid.addEventListener('sizeChange', (width, height) => {\n                    eventSource.postMessage({ event: 'sizeChange', data: { width, height }, options: makeOptions() }, '*');\n                });\n\n                mraid.addEventListener('stateChange', (state) => {\n                    eventSource.postMessage({ event: 'stateChange', data: state, options: makeOptions() }, '*');\n                });\n\n                mraid.addEventListener('viewableChange', (viewable) => {\n                    eventSource.postMessage({ event: 'viewableChange', data: viewable, options: makeOptions() }, '*');\n                });\n            }\n        };\n\n        window.addEventListener('message', (event) => {\n            if (event.source !== window.frames.wrapper.frames.banner) {\n                return;\n            }\n\n            switch (event.data.event) {\n                case 'init':\n                    initMraid(event.source);\n                    break;\n                case 'open':\n                    mraid.open(event.data.data);\n                    break;\n                case 'close':\n                    mraid.close();\n                    break;\n                case 'setExpandProperties':\n                    mraid.setExpandProperties(event.data.data);\n                    break;\n                case 'setOrientationProperties':\n                    mraid.setOrientationProperties(event.data.data);\n                    break;\n                case 'setResizeProperties':\n                    mraid.setResizeProperties(event.data.data);\n                    break;\n                case 'useCustomClose':\n                    mraid.useCustomClose(event.data.data);\n                    break;\n                case 'storePicture':\n                    mraid.storePicture(event.data.data);\n                    break;\n                case 'createCalendarEvent':\n                    mraid.createCalendarEvent(event.data.data);\n                    break;\n                case 'playVideo':\n                    mraid.playVideo(event.data.data);\n                    break;\n            }\n        });\n    })()\n</script>\n\n    <script>\n    document.createElement('img').src = 'https://statistics.adx.com.ru/api/stream/device' +\n        `?dpr=${devicePixelRatio}` +\n        `&width=${screen.width}` +\n        `&height=${screen.height}`;\n</script>\n\n\n    <iframe name=\"wrapper\" id=\"iframe-66cf23ee7011cb000191116b\" scrolling=\"no\" marginwidth=\"0\" marginheight=\"0\" src=\"https://adx.com.ru/banner/66cf23ee7011cb000191116b/?clickMacro=&auctionPrice=${AUCTION_PRICE}\"></iframe>\n\n    <script>\n        const frameWrapper = document.querySelector('iframe[name=\"wrapper\"]');\n        if (frameWrapper.dataset.params !== undefined) {\n            fetch('https://adx.com.ru/banner/66cf23ee7011cb000191116b/').then(resp => resp.text()).then(text => {\n                frameWrapper.srcdoc = text;\n            });\n        }\n    </script>\n</body>\n</html>\n",
									AdvDomains: []string{"vidimo.media"},
									ImageURL:   "https://cdn.adx.com.ru/banner/0000000000000000.jpg",
									CampaingID: "66c7182ba897d80001520631",
									CreativeID: "66cd9992a897d800015d8af7",
								},
							},
						},
						{
							Bids: []openrtb.Bid{
								{
									ID:         "seatBid_0_bid_2",           // Новый идентификатор
									ImpID:      "3",                         // Идентификатор первого импрешена
									Price:      decimal.NewFromFloat(0.585), // Ставка, по которой будем биллить DSP
									NoticeURL:  "https://test.host.com/billing/bidID?id_type=bid&price=${AUCTION_PRICE}&ssp_slug=test_slug&url_type=nurl",
									AdMarkup:   "\n<!DOCTYPE html>\n<html>\n<head>\n    <meta name=\"viewport\" content=\"width=device-width, initial-scale=1, user-scalable=no\">\n\n    <style>\n        html, body {\n            border: 0 !important;\n            margin: 0 !important;\n            padding: 0 !important;\n            width: 100vw !important;\n            height: 100vh !important;\n            overflow: hidden !important;\n        }\n\n        #iframe-66cf23ee7011cb000191116b {\n            border: 0;\n            \n            \n                width: 100%;\n                height: 100%;\n            \n        }\n    </style>\n</head>\n<body>\n    \n\n    \n\n    \n\n    \n<script>\n    (() => {\n        const initMraid = (window) => {\n            try {\n                mraid;\n            } catch {\n                mraidConnect(window, false);\n                return;\n            }\n\n            if (mraid.getState() === 'loading') {\n                mraid.addEventListener('ready', () => {\n                    mraidConnect(window, true);\n                });\n            } else {\n                mraidConnect(window, true);\n            }\n        };\n\n        const makeOptions = () => ({\n            placementType: mraid.getPlacementType(),\n            expandProperties: mraid.getExpandProperties(),\n            orientationProperties: mraid.getOrientationProperties(),\n            resizeProperties: mraid.getResizeProperties(),\n            maxSize: mraid.getMaxSize(),\n            features: {\n                sms: mraid.supports('sms'),\n                tel: mraid.supports('tel'),\n                calendar: mraid.supports('calendar'),\n                storePicture: mraid.supports('storePicture'),\n                inlineVideo: mraid.supports('inlineVideo'),\n            },\n        });\n\n        const mraidConnect = (eventSource, proxy) => {\n            eventSource.postMessage({ event: 'init', data: proxy, options: proxy ? makeOptions() : undefined }, '*');\n\n            if (proxy) {\n                mraid.addEventListener('error', (message, action) => {\n                    eventSource.postMessage({ event: 'error', data: { message, action, options: makeOptions() } }, '*');\n                });\n\n                mraid.addEventListener('sizeChange', (width, height) => {\n                    eventSource.postMessage({ event: 'sizeChange', data: { width, height }, options: makeOptions() }, '*');\n                });\n\n                mraid.addEventListener('stateChange', (state) => {\n                    eventSource.postMessage({ event: 'stateChange', data: state, options: makeOptions() }, '*');\n                });\n\n                mraid.addEventListener('viewableChange', (viewable) => {\n                    eventSource.postMessage({ event: 'viewableChange', data: viewable, options: makeOptions() }, '*');\n                });\n            }\n        };\n\n        window.addEventListener('message', (event) => {\n            if (event.source !== window.frames.wrapper.frames.banner) {\n                return;\n            }\n\n            switch (event.data.event) {\n                case 'init':\n                    initMraid(event.source);\n                    break;\n                case 'open':\n                    mraid.open(event.data.data);\n                    break;\n                case 'close':\n                    mraid.close();\n                    break;\n                case 'setExpandProperties':\n                    mraid.setExpandProperties(event.data.data);\n                    break;\n                case 'setOrientationProperties':\n                    mraid.setOrientationProperties(event.data.data);\n                    break;\n                case 'setResizeProperties':\n                    mraid.setResizeProperties(event.data.data);\n                    break;\n                case 'useCustomClose':\n                    mraid.useCustomClose(event.data.data);\n                    break;\n                case 'storePicture':\n                    mraid.storePicture(event.data.data);\n                    break;\n                case 'createCalendarEvent':\n                    mraid.createCalendarEvent(event.data.data);\n                    break;\n                case 'playVideo':\n                    mraid.playVideo(event.data.data);\n                    break;\n            }\n        });\n    })()\n</script>\n\n    <script>\n    document.createElement('img').src = 'https://statistics.adx.com.ru/api/stream/device' +\n        `?dpr=${devicePixelRatio}` +\n        `&width=${screen.width}` +\n        `&height=${screen.height}`;\n</script>\n\n\n    <iframe name=\"wrapper\" id=\"iframe-66cf23ee7011cb000191116b\" scrolling=\"no\" marginwidth=\"0\" marginheight=\"0\" src=\"https://adx.com.ru/banner/66cf23ee7011cb000191116b/?clickMacro=&auctionPrice=${AUCTION_PRICE}\"></iframe>\n\n    <script>\n        const frameWrapper = document.querySelector('iframe[name=\"wrapper\"]');\n        if (frameWrapper.dataset.params !== undefined) {\n            fetch('https://adx.com.ru/banner/66cf23ee7011cb000191116b/').then(resp => resp.text()).then(text => {\n                frameWrapper.srcdoc = text;\n            });\n        }\n    </script>\n</body>\n</html>\n",
									AdvDomains: []string{"vidimo.media"},
									ImageURL:   "https://cdn.adx.com.ru/banner/0000000000000000.jpg",
									CampaingID: "66c7182ba897d80001520631",
									CreativeID: "66cd9992a897d800015d8af7",
								},
							},
						},
					},
					BidID:    "bidder_id", // Новый идентификатор
					Currency: sspCurrency, // Валюта SSP
				}

				mocks.EventService.MockCreateExchangeBidResponseEvent(sspBidResponse,
					nil)

				return sspBidResponse
			},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d. %s", tt.testNumber, tt.name), func(t *testing.T) {

			// Создаем моковые зависимости для этого тесткейса
			mocks := service.NewMocks(t)

			// Создаем настоящий сервис с мокированными зависимостями
			exchangeService := service.NewExchangeService(
				mocks.ExchangeRepository,
				mocks.SSPService,
				mocks.DSPService,
				mocks.CurrencyService,
				mocks.SettingService,
				mocks.ExchangeNetwork,
				mocks.EventService,
			)

			// Меняем нужные для теста переменные конфига
			if tt.args.isSendNoContentInsteadErrorsInResponse {
				t.Setenv("IS_SEND_NO_CONTENT_INSTEAD_ERRORS_IN_RESPONSE", "true")
			} else {
				t.Setenv("IS_SEND_NO_CONTENT_INSTEAD_ERRORS_IN_RESPONSE", "false")
			}
			if tt.args.isSendNoContentForSuccessResponse {
				t.Setenv("IS_SEND_NO_CONTENT_FOR_SUCCESS_RESPONSE", "true")
			} else {
				t.Setenv("IS_SEND_NO_CONTENT_FOR_SUCCESS_RESPONSE", "false")
			}

			testUtils.IgnoreError(config.InitSingletones(config.Load()))

			// Получаем конфиг
			config.Reset()
			cfg := config.Load()

			// Создаем сервер и монтируем к нему нашу группу эндпоинтов
			app := testUtils.IgnoreErrorWithArgument(fiber.GetDefaultServer(cfg.ServiceName, cfg.FiberServer, nil))
			MountExchangeEndpoint(app, exchangeService)

			// Получаем и парсим запрос от SSP из файла для этого тесткейса
			var sspBidRequest openrtb.BidRequest
			sspRequestFileName := fmt.Sprintf(sspRequestFileTemplate, tt.testNumber)
			sspBidRequestJSON := testUtils.IgnoreErrorWithArgument(os.ReadFile(fmt.Sprintf("%s/%s", testcaseDir, sspRequestFileName)))
			testUtils.IgnoreError(json.Unmarshal(sspBidRequestJSON, &sspBidRequest))

			// Мокаем вызовы зависимых функций сервиса
			sspBidResponse := tt.mockActions(mocks, mockDTO{
				testNumber:      tt.testNumber,
				sspBidRequest:   sspBidRequest,
				exchangeService: exchangeService,
			})

			// Формируем запрос, будто мы SSP
			httpReq := httptest.NewRequest(http.MethodPost, fmt.Sprintf("/rtb/%s", defaultSSPSlug), bytes.NewReader(sspBidRequestJSON))
			httpReq.Header.Set("X-Request-ID", "1")
			defer httpReq.Body.Close()

			// Отправляем запрос нашему HTTP-серверу
			httpRes, err := app.Test(httpReq, 100000000) //nolint:bodyclose - линтер не понимает, что я игнорю ошибку и дальше закрываю тело ответа
			if err != nil {
				t.Fatal(err)
			}
			defer httpRes.Body.Close()

			// Проверяем код ответа на тот, который ожидаем в тесткейсе
			if httpRes.StatusCode != tt.args.wantHTTPCode {
				t.Fatalf("Expected status code %d, got %d", tt.args.wantHTTPCode, httpRes.StatusCode)
			}

			// Если код ответа 200, то проверяем структуру ответа
			if httpRes.StatusCode == http.StatusOK {
				bodyBytes := testUtils.IgnoreErrorWithArgument(io.ReadAll(httpRes.Body))
				log.Debug(context.Background(), string(bodyBytes))
				// Парсим ответ
				var res openrtb.BidResponse
				testUtils.IgnoreError(json.Unmarshal(bodyBytes, &res))
				testUtils.CheckStruct(t, res, sspBidResponse)
			}
		})
	}
}

// Номер теста - номер файла
const dspResponseFileTemplate = "%d.DSPResponse-%d.json"

// getDSPResponse возвращает спаршенный JSON для конкретного теста и конкретной DSP
func getDSPResponse(testNumber, fileNumber int) (dspBidResponse openrtb.BidResponse) {

	if fileNumber == 0 {
		return dspBidResponse
	}

	fileName := fmt.Sprintf(dspResponseFileTemplate, testNumber, fileNumber)

	// Парсим ответ DSP из файла
	dspBidResponseJSON := testUtils.IgnoreErrorWithArgument(os.ReadFile(fmt.Sprintf("%s/%s", testcaseDir, fileName)))
	testUtils.IgnoreError(json.Unmarshal(dspBidResponseJSON, &dspBidResponse))
	return dspBidResponse
}
