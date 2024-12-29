package beforeResponseToSSP

import (
	"testing"

	"pkg/errors"
	"pkg/testUtils"
)

func Test_checkAllSets_Apply(t *testing.T) {

	tests := []struct {
		name    string
		req     beforeResponseToSSP
		wantErr error
	}{
		{
			name: "1. Проверка всех установленных значений",
			req: beforeResponseToSSP{
				chainSettings: chainSettings{
					bidsAlreadyInit: true,
					admAlreadySet:   true,
					priceAlreadySet: true,
					nurlAlreadySet:  true,
				},
			},
			wantErr: nil,
		},
		{
			name: "2. Проверка отсутствия установки bidsAlreadyInit",
			req: beforeResponseToSSP{
				chainSettings: chainSettings{
					bidsAlreadyInit: false,
					admAlreadySet:   true,
					priceAlreadySet: true,
					nurlAlreadySet:  true,
				},
			},
			wantErr: errors.InternalServer.New(""),
		},
		{
			name: "3. Проверка отсутствия установки admAlreadySet",
			req: beforeResponseToSSP{
				chainSettings: chainSettings{
					bidsAlreadyInit: true,
					admAlreadySet:   false,
					priceAlreadySet: true,
					nurlAlreadySet:  true,
				},
			},
			wantErr: errors.InternalServer.New(""),
		},
		{
			name: "4. Проверка отсутствия установки priceAlreadySet",
			req: beforeResponseToSSP{
				chainSettings: chainSettings{
					bidsAlreadyInit: true,
					admAlreadySet:   true,
					priceAlreadySet: false,
					nurlAlreadySet:  true,
				},
			},
			wantErr: errors.InternalServer.New(""),
		},
		{
			name: "5. Проверка отсутствия установки nurlAlreadySet",
			req: beforeResponseToSSP{
				chainSettings: chainSettings{
					bidsAlreadyInit: true,
					admAlreadySet:   true,
					priceAlreadySet: true,
					nurlAlreadySet:  false,
				},
			},
			wantErr: errors.InternalServer.New(""),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			r := &checkAllSets{}

			// Применяем метод
			gotErr := r.Apply(&tt.req)
			testUtils.CheckError(t, gotErr, tt.wantErr, false)
		})
	}
}
