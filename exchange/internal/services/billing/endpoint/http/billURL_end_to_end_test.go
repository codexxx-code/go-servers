package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"exchange/internal/config"
	"exchange/internal/enum/billingType"
	billingRepository "exchange/internal/services/billing/repository"
	"exchange/internal/services/billing/service"
	dspModel "exchange/internal/services/dsp/model"
	"exchange/internal/services/exchange/model"
	"pkg/decimal"
	"pkg/http/fiber"
	"pkg/openrtb"
	"pkg/testUtils"
)

const testcaseDir = "testcase/billURL_end_to_end"

func Test_end_to_end_BidSSP(t *testing.T) {

	// Инициализируем все синглтон зависимости
	testUtils.SetDefaultEnvs(t)
	config.Reset()
	testUtils.IgnoreError(config.InitSingletones(config.Load()))

	type args struct {
		url                                    string
		isSendNoContentInsteadErrorsInResponse bool // Параметр, который проставляется в env и отвечает за отправку 204 вместо ошибок
		isSendNoContentForSuccessResponse      bool // Параметр, который проставляется в env и отвечает за отправку 204 вместо успешных ответов
		wantHTTPCode                           int
	}
	// Тесткейсы
	tests := []struct {
		fileNumber  int         // Номер файла, используется для автоматического парсинга файлов
		name        string      // Название теста, чисто информационная строка
		args        args        // Входные параметры теста
		mockActions func(mocks) // Перечень вызовов зависимых функций сервиса
	}{
		{
			name:       "1. Биллинг по раскрытому макросу в NURL'е, обычный запрос",
			fileNumber: 1,
			args: args{
				url:                                    "https://test.com/billing/bidID?id_type=bid&price=0.1&url_type=nurl",
				isSendNoContentInsteadErrorsInResponse: false,
				isSendNoContentForSuccessResponse:      false,
				wantHTTPCode:                           http.StatusOK,
			},
			mockActions: func(mocks mocks) {

				bidResponse := getDSPResponse(1)

				// Получаем ответ от DSP по идентификатору переданной записи
				mongoDSPResponse := model.DSPResponse{
					ExchangeID:                "requestID",
					ExchangeBidID:             "bidID",
					BidResponse:               bidResponse,
					SeatBidIndex:              0,
					BidIndex:                  0,
					BillingPriceInDSPCurrency: decimal.NewFromFloat(0.2),
					SlugDSP:                   "testDSP",
					RecordDateTime:            time.Now(),
					OriginRequestID:           "originRequestID",
					CurrencySSPCoefficient:    decimal.NewFromFloat(100),
					SSPCurrency:               "USD",
					SlugSSP:                   "testSSP",
					CurrencyDSPCoefficient:    decimal.NewFromFloat(110),
					DSPCurrency:               "EUR",
				}
				mocks.exchangeRepository.MockGetDSPResponse(model.GetDSPResponsesReq{
					BidIDs: []string{"bidID"},
				},
					mongoDSPResponse, nil)

				// Получаем DSP по слагу из записи в монге
				mocks.dspService.MockGetDSPs(dspModel.GetDSPsReq{
					Slugs: []string{"testDSP"},
				},
					[]dspModel.DSP{{
						Name:               "testDSP",
						Slug:               "testDSP",
						URL:                "https://test.com/testDSP",
						Currency:           "EUR",
						AuctionSecondPrice: true,
						BillingURLType:     billingType.NURL,
					}}, nil)

				// Создаем ивент победы SSP
				mocks.billingRepository.MockCreateSSPWinEvent(billingRepository.CreateSSPEventWinReq{
					RequestID: "requestID",
					SSP:       "testSSP",
					Price:     "0.01", // 0.1 * 100 / 1000 = 0.01
					Timestamp: time.Time{},
				},
					nil)

				mocks.billingNetwork.MockBillDSP("https://events.adx.com.ru/event?price=0.2&tracking=66cf23ee7011cb000191116b&type=win",
					200, nil)

				// Создаем ивено победы DSP
				mocks.billingRepository.MockCreateDSPWinEvent(billingRepository.CreateDSPEventWinReq{
					RequestID: "bidID",
					DSP:       "testDSP",
					Price:     decimal.NewFromFloat(0.022).String(), // Руками считаем ту цену в системной валюте (RUB), которая должна отправиться в аналитику + делим число на 1000, так как приходит нам цена за 1000, а надо записать за один показ // 0.2 * 110 / 1000 = 0.022
					Timestamp: time.Time{},
				},
					nil)
			},
		},
		{
			name:       "2.Биллинг по запасной цене в NURL'е, обычный запрос",
			fileNumber: 1,
			args: args{
				url:                                    "https://test.com/billing/bidID?id_type=bid&price=${AUCTION_PRICE}&url_type=nurl&bid_price=0.1",
				isSendNoContentInsteadErrorsInResponse: false,
				isSendNoContentForSuccessResponse:      false,
				wantHTTPCode:                           http.StatusOK,
			},
			mockActions: func(mocks mocks) {

				bidResponse := getDSPResponse(1)

				// Получаем ответ от DSP по идентификатору переданной записи
				mongoDSPResponse := model.DSPResponse{
					ExchangeID:                "requestID",
					ExchangeBidID:             "bidID",
					BidResponse:               bidResponse,
					SeatBidIndex:              0,
					BidIndex:                  0,
					BillingPriceInDSPCurrency: decimal.NewFromFloat(0.2),
					SlugDSP:                   "testDSP",
					RecordDateTime:            time.Now(),
					OriginRequestID:           "originRequestID",
					CurrencySSPCoefficient:    decimal.NewFromFloat(100),
					SSPCurrency:               "USD",
					SlugSSP:                   "testSSP",
					CurrencyDSPCoefficient:    decimal.NewFromFloat(110),
					DSPCurrency:               "EUR",
				}
				mocks.exchangeRepository.MockGetDSPResponse(model.GetDSPResponsesReq{
					BidIDs: []string{"bidID"},
				},
					mongoDSPResponse, nil)

				// Получаем DSP по слагу из записи в монге
				mocks.dspService.MockGetDSPs(dspModel.GetDSPsReq{
					Slugs: []string{"testDSP"},
				},
					[]dspModel.DSP{{
						Name:               "testDSP",
						Slug:               "testDSP",
						URL:                "https://test.com/testDSP",
						Currency:           "EUR",
						AuctionSecondPrice: true,
						BillingURLType:     billingType.NURL,
					}}, nil)

				// Создаем ивент победы SSP
				mocks.billingRepository.MockCreateSSPWinEvent(billingRepository.CreateSSPEventWinReq{
					RequestID: "requestID",
					SSP:       "testSSP",
					Price:     "0.01", // 0.1 * 100 / 1000 = 0.01
					Timestamp: time.Time{},
				},
					nil)

				mocks.billingNetwork.MockBillDSP("https://events.adx.com.ru/event?price=0.2&tracking=66cf23ee7011cb000191116b&type=win",
					200, nil)

				// Создаем ивент победы DSP
				mocks.billingRepository.MockCreateDSPWinEvent(billingRepository.CreateDSPEventWinReq{
					RequestID: "bidID",
					DSP:       "testDSP",
					Price:     decimal.NewFromFloat(0.022).String(), // Руками считаем ту цену в системной валюте (RUB), которая должна отправиться в аналитику + делим число на 1000, так как приходит нам цена за 1000, а надо записать за один показ // 0.2 * 110 / 1000 = 0.022
					Timestamp: time.Time{},
				},
					nil)
			},
		},
		{
			name:       "3. Биллинг по раскрытому макросу, кликандер",
			fileNumber: 1,
			args: args{
				url:                                    "https://test.com/billing/requestID?id_type=request&price=0.1&url_type=burl",
				isSendNoContentInsteadErrorsInResponse: false,
				isSendNoContentForSuccessResponse:      false,
				wantHTTPCode:                           http.StatusOK,
			},
			mockActions: func(mocks mocks) {

				bidResponse1 := getDSPResponse(1)
				bidResponse2 := getDSPResponse(2)

				// Получаем ответ от DSP по идентификатору переданной записи
				mongoDSPResponse := []model.DSPResponse{
					{
						ExchangeID:                "requestID",
						ExchangeBidID:             "bidID1",
						BidResponse:               bidResponse1,
						SeatBidIndex:              0,
						BidIndex:                  0,
						BillingPriceInDSPCurrency: decimal.NewFromFloat(0.2),
						SlugDSP:                   "testDSP1",
						RecordDateTime:            time.Now(),
						OriginRequestID:           "",
						CurrencySSPCoefficient:    decimal.NewFromFloat(100),
						SSPCurrency:               "USD",
						SlugSSP:                   "testSSP",
						CurrencyDSPCoefficient:    decimal.NewFromFloat(110),
						DSPCurrency:               "EUR",
					},
					{
						ExchangeID:                "requestID",
						ExchangeBidID:             "bidID2",
						BidResponse:               bidResponse1,
						SeatBidIndex:              0,
						BidIndex:                  0,
						BillingPriceInDSPCurrency: decimal.NewFromFloat(0.06),
						SlugDSP:                   "testDSP1",
						RecordDateTime:            time.Now(),
						OriginRequestID:           "",
						CurrencySSPCoefficient:    decimal.NewFromFloat(100),
						SSPCurrency:               "USD",
						SlugSSP:                   "testSSP",
						CurrencyDSPCoefficient:    decimal.NewFromFloat(1),
						DSPCurrency:               "RUB",
					},
					{
						ExchangeID:                "requestID",
						ExchangeBidID:             "bidID3",
						BidResponse:               bidResponse2,
						SeatBidIndex:              0,
						BidIndex:                  0,
						BillingPriceInDSPCurrency: decimal.NewFromFloat(90),
						SlugDSP:                   "testDSP2",
						RecordDateTime:            time.Now(),
						OriginRequestID:           "",
						CurrencySSPCoefficient:    decimal.NewFromFloat(100),
						SSPCurrency:               "USD",
						SlugSSP:                   "testSSP",
						CurrencyDSPCoefficient:    decimal.NewFromFloat(100),
						DSPCurrency:               "USD",
					},
				}

				mocks.exchangeRepository.MockGetDSPResponses(model.GetDSPResponsesReq{
					RequestIDs: []string{"requestID"},
				},
					mongoDSPResponse, nil)

				// Создаем ивент победы SSP
				mocks.billingRepository.MockCreateSSPWinEvent(billingRepository.CreateSSPEventWinReq{
					RequestID: "requestID",
					SSP:       "testSSP",
					Price:     "0.01", // 0.1 * 100 / 1000
					Timestamp: time.Time{},
				},
					nil)

				// Получаем DSP по слагам
				mocks.dspService.MockGetDSPs(dspModel.GetDSPsReq{
					Slugs: []string{"testDSP1", "testDSP2"},
				},
					[]dspModel.DSP{
						{
							Name:               "testDSP1",
							Slug:               "testDSP1",
							URL:                "https://test.com/testDSP",
							Currency:           "EUR",
							AuctionSecondPrice: true,
							BillingURLType:     billingType.NURL,
						},
						{
							Name:               "testDSP2",
							Slug:               "testDSP2",
							URL:                "https://test.com/testDSP",
							Currency:           "EUR",
							AuctionSecondPrice: true,
							BillingURLType:     billingType.BURL,
						},
					}, nil)

				// Биллим DSP
				mocks.billingNetwork.MockBillDSP("https://events.adx.com.ru/event?price=0.2&tracking=66cf23ee7011cb000191116b&type=win",
					200, nil)

				// Создаем ивент победы DSP
				mocks.billingRepository.MockCreateDSPWinEvent(billingRepository.CreateDSPEventWinReq{
					RequestID: "bidID1",
					DSP:       "testDSP1",
					Price:     "0.022", // 0.2 * 110 / 1000
					Timestamp: time.Time{},
				},
					nil)

				//// 2 итерация

				// Биллим DSP
				mocks.billingNetwork.MockBillDSP("https://events.adx.com.ru/event?price=0.06&tracking=66cf23ee7011cb000191116b&type=win",
					200, nil)

				// Создаем ивент победы DSP
				mocks.billingRepository.MockCreateDSPWinEvent(billingRepository.CreateDSPEventWinReq{
					RequestID: "bidID2",
					DSP:       "testDSP1",
					Price:     "0.00006", // 0.06 / 1000
					Timestamp: time.Time{},
				},
					nil)

				//// 3 итерация

				// Биллим DSP
				mocks.billingNetwork.MockBillDSP("https://adriver.com/event/win?price=90&tracking=GllbnQiOnsibmFtZSI6IiIs",
					200, nil)

				// Создаем ивент победы DSP
				mocks.billingRepository.MockCreateDSPWinEvent(billingRepository.CreateDSPEventWinReq{
					RequestID: "bidID3",
					DSP:       "testDSP2",
					Price:     "9", // 90 * 100 / 1000
					Timestamp: time.Time{},
				},
					nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// Создаем моковые зависимости для этого тесткейса
			mocks := getMocks(t)

			// Создаем настоящий сервис с мокированными зависимостями
			billingService := service.NewBillingService(
				mocks.dspService,
				mocks.billingRepository,
				mocks.exchangeRepository,
				mocks.billingNetwork,
				mocks.analyticWriterSerivce,
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

			// Получаем конфиг
			config.Reset()
			cfg := config.Load()

			// Создаем сервер и монтируем к нему нашу группу эндпоинтов
			app := testUtils.IgnoreErrorWithArgument(fiber.GetDefaultServer("", cfg.FiberServer, nil))
			MountBillingEndpoint(app, billingService)

			// Получаем и парсим запрос от SSP из файла для этого тесткейса
			var dspBidResponse openrtb.BidResponse
			dspResponseFileName := fmt.Sprintf(dspResponseFileTemplate, tt.fileNumber)
			dspResponseJSON := testUtils.IgnoreErrorWithArgument(os.ReadFile(fmt.Sprintf("%s/%s", testcaseDir, dspResponseFileName)))
			testUtils.IgnoreError(json.Unmarshal(dspResponseJSON, &dspBidResponse))

			// Мокаем вызовы зависимых функций сервиса
			tt.mockActions(mocks)

			// Отправляем запрос нашему HTTP-серверу
			httpRes, err := app.Test(httptest.NewRequest(http.MethodGet, tt.args.url, nil), 100000000)
			if err != nil {
				t.Fatal(err)
			}
			defer httpRes.Body.Close()

			// Проверяем код ответа на тот, который ожидаем в тесткейсе
			if httpRes.StatusCode != tt.args.wantHTTPCode {
				t.Fatalf("Expected status code %d, got %d", tt.args.wantHTTPCode, httpRes.StatusCode)
			}
		})
	}
}

// Номер теста - номер файла
const dspResponseFileTemplate = "%d.DSPResponse.json"

// getDSPResponse возвращает спаршенный JSON для конкретного теста и конкретной DSP
func getDSPResponse(fileNumber int) (dspBidResponse openrtb.BidResponse) {

	fileName := fmt.Sprintf(dspResponseFileTemplate, fileNumber)

	// Парсим ответ DSP из файла
	dspBidResponseJSON := testUtils.IgnoreErrorWithArgument(os.ReadFile(fmt.Sprintf("%s/%s", testcaseDir, fileName)))
	testUtils.IgnoreError(json.Unmarshal(dspBidResponseJSON, &dspBidResponse))
	return dspBidResponse
}
