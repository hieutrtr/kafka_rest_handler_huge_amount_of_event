package handler

type Handler struct {
	// some storage or database service
}

func New() *Handler {
	return &Handler{}
}
