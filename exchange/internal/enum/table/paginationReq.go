package table

type PaginationReq struct {
	Page int32 `validate:"required"` // Номер страницы
	Size int32 `validate:"required"` // Количество отображаемых элементов
}
