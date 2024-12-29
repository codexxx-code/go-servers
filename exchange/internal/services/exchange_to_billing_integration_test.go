package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"exchange/internal/config"
	billingEndpointHTTP "exchange/internal/services/billing/endpoint/http"
	billingModel "exchange/internal/services/billing/model"
	exchangeEndpointHTTP "exchange/internal/services/exchange/endpoint/http"
	exchangeModel "exchange/internal/services/exchange/model"
	"pkg/http/fiber"
	"pkg/openrtb"
	"pkg/testUtils"
)

const serviceName = "exchange"
const sspRequestFileTemplate = "%d.SSPRequestImpression.json"
const sspResponseFileTemplate = "%d.SSPResponse.json"
const testcaseDir = "testcase"

type mocks struct {
	exchangeService *exchangeEndpointHTTP.MockExchangeService
	billingService  *billingEndpointHTTP.MockBillingService
}

func getMocks(t *testing.T) mocks {
	return mocks{
		exchangeService: exchangeEndpointHTTP.NewMockExchangeService(t),
		billingService:  billingEndpointHTTP.NewMockBillingService(t),
	}
}

type mockDTO struct {
	sspBidRequest  openrtb.BidRequest  // Запрос от SSP, который автоматически парсится из файла
	sspBidResponse openrtb.BidResponse // Ответ для SSP, который автоматически парсится из файла
}

const defaultSSPSlug = "test_ssp"

func Test_integration_exchange_to_billing(t *testing.T) {

	testUtils.SetDefaultEnvs(t)

	type args struct {
		wantHTTPCodes                          [2]int // Ожидаемый код ответа от сервера на запрос
		isSendNoContentInsteadErrorsInResponse bool   // Параметр, который проставляется в env и отвечает за отправку 204 вместо ошибок
		isSendNoContentForSuccessResponse      bool   // Параметр, который проставляется в env и отвечает за отправку 204 вместо успешных ответов
		needReplaceMacros                      bool   // Параметр, который отвечает за необходимость замены макроса на значение
	}
	t.Setenv("SERVICE_NAME", serviceName)

	// Инициализируем все синглтон зависимости
	config.Reset()
	testUtils.IgnoreError(config.InitSingletones(config.Load()))

	tests := []struct {
		fileNumber  int                  // Номер теста, используется для автоматического парсинга файлов
		name        string               // Название теста, чисто информационная строка
		args        args                 // Входные параметры теста
		mockActions func(mocks, mockDTO) // Перечень вызовов зависимых функций эндпоинтов
	}{
		{
			name:       "1. Отработка крутилки и вызов биллинга с заменой макроса цены",
			fileNumber: 1,
			args: args{
				wantHTTPCodes:                          [2]int{http.StatusOK, http.StatusOK},
				isSendNoContentInsteadErrorsInResponse: false,
				isSendNoContentForSuccessResponse:      false,
				needReplaceMacros:                      true,
			},
			mockActions: func(mocks mocks, dto mockDTO) {

				// Мокаем крутилку
				response := dto.sspBidResponse
				mocks.exchangeService.MockBidSSP(exchangeModel.SSPBidReq{
					BidRequest: dto.sspBidRequest,
					SSPSlug:    defaultSSPSlug,
				},
					response, nil)

				for _, seatBid := range response.SeatBids {
					for range seatBid.Bids {
						// Мокаем биллинг
						mocks.billingService.MockBillURL(billingModel.BillURLReq{
							ID:             "requestID",
							HardcodedPrice: "2.032520325203252",
							MacrosPrice:    "2.032520325203252",
							BillingType:    "nurl",
							IDType:         "bid",
						}, nil)
					}
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d. %s", tt.fileNumber, tt.name), func(t *testing.T) {

			// Создаем моковые зависимости для этого тесткейса
			mocks := getMocks(t)

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
			app := testUtils.IgnoreErrorWithArgument(fiber.GetDefaultServer(serviceName, cfg.FiberServer, nil))
			exchangeEndpointHTTP.MountExchangeEndpoint(app, mocks.exchangeService)
			billingEndpointHTTP.MountBillingEndpoint(app, mocks.billingService)

			// Получаем и парсим запрос от SSP из файла для этого тесткейса
			var sspBidRequest openrtb.BidRequest
			sspBidRequestJSON := testUtils.IgnoreErrorWithArgument(os.ReadFile(fmt.Sprintf("%s/%s", testcaseDir, fmt.Sprintf(sspRequestFileTemplate, tt.fileNumber))))
			testUtils.IgnoreError(json.Unmarshal(sspBidRequestJSON, &sspBidRequest))

			// Получаем и парсим наш ответ к SSP из файла для этого тесткейса
			var sspBidResponse openrtb.BidResponse
			sspBidResponseJSON := testUtils.IgnoreErrorWithArgument(os.ReadFile(fmt.Sprintf("%s/%s", testcaseDir, fmt.Sprintf(sspResponseFileTemplate, tt.fileNumber))))
			testUtils.IgnoreError(json.Unmarshal(sspBidResponseJSON, &sspBidResponse))

			// Мокаем вызовы зависимых функций сервиса
			tt.mockActions(mocks, mockDTO{
				sspBidRequest:  sspBidRequest,
				sspBidResponse: sspBidResponse,
			})

			// Отправляем запрос на аукцион нашему HTTP-серверу
			httpRes, err := app.Test(httptest.NewRequest(http.MethodPost, fmt.Sprintf("/rtb/%s", defaultSSPSlug), bytes.NewReader(sspBidRequestJSON)), 100000000)
			if err != nil {
				t.Fatal(err)
			}
			defer httpRes.Body.Close()

			// Проверяем код ответа на тот, который ожидаем в тесткейсе
			if httpRes.StatusCode != tt.args.wantHTTPCodes[0] {
				t.Fatalf("Expected status code %d, got %d for exchange HTTP call", tt.args.wantHTTPCodes[0], httpRes.StatusCode)
			}

			// Если код ответа 200, то проверяем структуру ответа
			var res openrtb.BidResponse
			if httpRes.StatusCode == http.StatusOK {
				// Парсим ответ
				testUtils.IgnoreError(json.NewDecoder(httpRes.Body).Decode(&res))
				testUtils.CheckStruct(t, res, sspBidResponse)
			}

			for _, seatBid := range res.SeatBids {
				for _, bid := range seatBid.Bids {

					nurlWithPrice := strings.Replace(bid.NoticeURL, "${AUCTION_PRICE}", bid.Price.String(), 1)

					// Отправляем запрос на биллинг нашему HTTP-серверу
					httpRes = testUtils.IgnoreErrorWithArgument(app.Test(httptest.NewRequest(http.MethodGet, nurlWithPrice, nil), 100000000))
					defer httpRes.Body.Close()

					// Проверяем код ответа на тот, который ожидаем в тесткейсе
					if httpRes.StatusCode != tt.args.wantHTTPCodes[1] {
						t.Fatalf("Expected status code %d, got %d for billing HTTP call", tt.args.wantHTTPCodes[1], httpRes.StatusCode)
					}
				}
			}
		})
	}
}
