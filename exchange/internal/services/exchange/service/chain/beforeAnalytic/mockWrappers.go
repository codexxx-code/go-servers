package beforeAnalytic

import (
	"testing"
)

//// Обертки над моками для типизации

type mocks struct {
	exchangeRepository *MockExchangeRepository
}

func getMocks(t *testing.T) mocks {
	return mocks{
		exchangeRepository: NewMockExchangeRepository(t),
	}
}

func (m *MockExchangeRepository) MockGetCountryByIP(ipStr string,
	country string, err error,
) {
	m.On("GetCountryByIP", ipStr).
		Return(country, err)
}
