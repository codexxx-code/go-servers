package http

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"exchange/internal/config"
	"exchange/internal/services/exchange/service"
	"pkg/http/fiber"
	"pkg/testUtils"
)

func Test_end_to_end_GetADM(t *testing.T) {

	// Инициализируем все синглтон зависимости
	testUtils.SetDefaultEnvs(t)
	testUtils.IgnoreError(config.InitSingletones(config.Load()))

	type mockDTO struct {
		id  string
		adm string
	}
	type args struct {
		id                                     string // ID ADM
		wantHTTPCode                           int    // Ожидаемый код ответа от сервера на запрос
		isSendNoContentInsteadErrorsInResponse bool
		adm                                    string
	}
	// Тесткейсы
	tests := []struct {
		testNumber  int                          // Номер теста, используется для автоматического парсинга файлов
		name        string                       // Название теста, чисто информационная строка
		args        args                         // Входные параметры теста
		mockActions func(service.Mocks, mockDTO) // Перечень вызовов зависимых функций сервиса
	}{
		{
			testNumber: 1,
			name:       "Отработка успешного запроса получение adm",
			args: args{
				wantHTTPCode:                           http.StatusOK,
				id:                                     "id",
				isSendNoContentInsteadErrorsInResponse: false,
				adm:                                    "<someADM>",
			},
			mockActions: func(mocks service.Mocks, dto mockDTO) {

				mocks.ExchangeRepository.MockGetADM(dto.id,
					dto.adm, nil)
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

			// Получаем конфиг
			config.Reset()
			cfg := config.Load()

			// Создаем сервер и монтируем к нему нашу группу эндпоинтов
			app := testUtils.IgnoreErrorWithArgument(fiber.GetDefaultServer(cfg.ServiceName, cfg.FiberServer, nil))
			MountExchangeEndpoint(app, exchangeService)

			// Мокаем вызовы зависимых функций сервиса
			tt.mockActions(mocks, mockDTO{
				id:  tt.args.id,
				adm: tt.args.adm,
			})

			// Формируем запрос, будто мы SSP
			httpReq := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/rtb/adm/%s", tt.args.id), nil)
			defer httpReq.Body.Close()

			// Отправляем запрос нашему HTTP-серверу
			httpRes := testUtils.IgnoreErrorWithArgument(app.Test(httpReq, 100000000))
			defer httpRes.Body.Close()

			// Проверяем код ответа на тот, который ожидаем в тесткейсе
			if httpRes.StatusCode != tt.args.wantHTTPCode {
				t.Fatalf("Expected status code %d, got %d", tt.args.wantHTTPCode, httpRes.StatusCode)
			}

			// Если код ответа 200, то проверяем структуру ответа
			if httpRes.StatusCode == http.StatusOK {
				bodyBytes := testUtils.IgnoreErrorWithArgument(io.ReadAll(httpRes.Body))
				if string(bodyBytes) != tt.args.adm {
					t.Fatalf("Expected body %s, got %s", tt.args.adm, string(bodyBytes))
				}
			}
		})
	}
}
