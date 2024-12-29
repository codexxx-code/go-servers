package model

type UpdateUserReq struct {
	ID        string
	LastName  string
	FirstName string
	Email     string
	AuthorID  string
}
