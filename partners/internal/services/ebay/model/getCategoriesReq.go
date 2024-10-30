package model

import "pkg/errors"

type GetCategoriesReq struct {
	Name         *string `json:"name" schema:"name"`
	ID           *string `json:"id" schema:"id"`
	MaxDeepLevel *uint8  `json:"maxDeepLevel" schema:"maxDeepLevel"`
}

func (s GetCategoriesReq) Validate() error {

	// Если включен фильтр по названию
	if s.Name != nil {

		// Проверяем, что фильтр не пустой
		if *s.Name == "" {
			return errors.BadRequest.New("name filter cannot be empty")
		}

		// И включено обрезание дерева
		if s.MaxDeepLevel != nil {
			return errors.BadRequest.New("name filter and maxDeepLevel cannot be used together")
		}
	}

	return nil
}
