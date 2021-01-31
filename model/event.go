package model

type Event struct {
	UserName string `json:"username"`
	Age int32 `json:"age"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
}
