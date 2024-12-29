package network

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/valyala/fasthttp"

	"pkg/errors"
	"pkg/openrtb"
)

func (s *ExchangeNetwork) SendBidRequestToDSP(
	ctx context.Context,
	_ int, // Нужно, чтобы моки нормально отрабатывали на барабане
	url string,
	req openrtb.BidRequest,
) (res openrtb.BidResponse, statusCode int, err error) {

	// Для ручного тестирования
	// return getMockBidResponse(req), 200, nil

	// Сериализуем запрос в JSON
	bodyBytes, err := json.Marshal(req)
	if err != nil {
		return res, 0, errors.InternalServer.Wrap(err)
	}

	// Создаем request
	r := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(r)

	// Устанавливаем URL
	r.SetRequestURI(url)

	// Устанавливаем метод
	r.Header.SetMethod(http.MethodPost)

	// Устанавливаем тело запроса
	r.SetBody(bodyBytes)

	// Определяем response
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	// Смотрим, не закончился ли контекст
	if err = ctx.Err(); err != nil {
		return res, 0, errors.Timeout.Wrap(err)
	}

	done := make(chan error, 1)

	// Делаем запрос
	go func() {
		done <- s.client.Do(r, resp)
	}()

	select {
	case <-ctx.Done(): // Если контекст закончился
		return res, 0, errors.Timeout.Wrap(ctx.Err())
	case err = <-done: // Если запрос завершился
		// Смотрим на результат
		if err != nil {
			switch {
			case errors.Is(err, fasthttp.ErrTimeout):
				return res, 0, errors.Timeout.Wrap(err)
			default:
				return res, 0, errors.InternalServer.Wrap(err)
			}
		}
	}

	// Конфигурируем логику согласно HTTP-кодам
	statusCode = resp.StatusCode()
	switch statusCode {

	// Если успешно
	case http.StatusOK:

		// Декодируем JSON от dsp в стандартный bidResponse
		err = json.Unmarshal(resp.Body(), &res)
		if err != nil {
			return res, statusCode, errors.BadRequest.Wrap(err)
		}

		return res, statusCode, nil

	case http.StatusNoContent:
		return res, statusCode, errors.BadGateway.Wrap(ErrStatusCode204)

	default: // Если что-то пошло не так
		return res, statusCode, errors.BadGateway.Wrap(ErrUnexpectedHTTPCode)
	}
}

var (
	ErrStatusCode204      = errors.New("204 http code")
	ErrUnexpectedHTTPCode = errors.New("unexpected http code")
)

func getMockBidResponse(req openrtb.BidRequest) (res openrtb.BidResponse) { //nolint:unused
	_ = json.Unmarshal([]byte(fmt.Sprintf(handleDebuggingResponse, req.Impressions[0].ID)), &res)
	return res
}

var handleDebuggingResponse = //nolint:unused
`
{ 
	"id": "9a92ee6b-2170-4901-965c-53e73920d8ae",
	"seatbid": [
		{
			"bid": [
				{
					"id": "2df3796364f89bb15f00da086a822469",
					"impid": "%s",
					"price": 1.7579832,
					"nurl": "https://dssa-eu.hybrid.ai/openrtb/win2/124/?r=7827877091867696\u0026sid=__33_uiadhybepwkz8texcsrpn5uu7qky565gdfftfdst7a7dq1c1pnrm8r1kr6mjhp4j1z4u1kb8f6ysg\u0026adid=__33_3dq7bs8natidb5put7cfwbqwz6ccq4ugiix85j4g9ub9ged39wduo9we6ky5556ofrnj8ae3je6fq\u0026p=${AUCTION_PRICE}\u0026mw=",
					"burl": "",
					"lurl": "",
					"adm": "\u003ciframe src=\"https://dsa-eu.hybrid.ai/OpenRtbDelivery/Markup?r=cyheopsmdt\u0026id=124\u0026p=\u0026mw=\u0026ope=__33_m6ctto5758msy48aaihhkeraif4bcn1oogh4f4ongb8rwrkebgmin1uy4b16i4nixb4dxzs6nogdq9htegoqfajjjm3acgczn9e8snfokkcwda4775qcsykfz36h65skubnfbo5qz9bdbjnqt78bzytwugwzxines5xkaeguxcghnr89gafsjurrhmy9qf9sj9w65ptfcpchphnektea6ti4soi9j3u4q8zhzf1irhwo33m3u3t5kdknne9zsbxtg8gwknybodyhgi7mx947sr8js6eyffsdzhg8qjsf1rr3h1m67fdt4frux8zax7tc5mdj4ejnbcr7cga9k5qs8jumnqwszrbbao4f5zx7uttuhetfgxysjc4mb1edrkk6sjnohncurbq5bhfxim6zjkmk37fszfb4obe9diy9t17jnecu5id5ym16jd3bp6u8bb7zqqr6mnm87geb3epg5gft8j67i14wjbbhjg1nxynx6tsf4se84hcjgr4sqwoo\u0026erid=2SDnjcAL8tK\" frameborder=\"0\" scrolling=\"no\" marginwidth=\"0\" marginheight=\"0\" width=\"300\" height=\"250\" allowtransparency=\"true\" style=\"width:300px;height:250px;\"\u003e\u003c/iframe\u003e",
					"adid": "673af72a7bc72f03c066ea90",
					"adomain": [
						"vkusnoitochka.ru"
					],
					"bundle": "",
					"iurl": "https://content.hybrid.ai/Hybrid/A2/D7/A2D72E219EE23DCF24903734B87E2F2D.jpeg",
					"cid": "673af2ad7bc72fac048dae1d",
					"crid": "673af72a7bc72f03c066ea90",
					"tactic": "",
					"cat": [
						"IAB8"
					],
					"attr": null,
					"api": 0,
					"protocol": 0,
					"qagmediarating": 0,
					"language": "ru",
					"dealid": "",
					"w": 300,
					"h": 250,
					"wratio": 0,
					"hratio": 0,
					"exp": 900,
					"ext": {
						"nroa": {
							"erid": "2SDnjcAL8tK",
							"contractor": {
								"inn": "7743068844",
								"name": "Лайон Коммьюникейшнз",
								"legal_form": "ul"
							},
							"client": {
								"inn": "7710044140",
								"name": "ООО \"Система ПБО\"",
								"legal_form": "ul"
							},
							"initial_contract": {
								"id": "CTgZuo8zhxrESjoEYTNpsClw",
								"ado_id": "MTS",
								"type": "intermediary-contract",
								"number": "2301012",
								"subject_type": "mediation",
								"action_type": "distribution",
								"sign_date": "2023-06-30"
							}
						},
						"erid": "2SDnjcAL8tK"
					}
				}
			],
			"seat": "",
			"group": 0
		}
	],
	"bidid": "b5630aa6da2a5e963b5f79ca58ffbd8a",
	"cur": "RUB",
	"customdata": "",
	"nbr": 0
}
`
