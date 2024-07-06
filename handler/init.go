package handler

type MainHandler struct {
	Book *BookHandler
}

func NewHandler(
	Book *BookHandler,
) *MainHandler {
	return &MainHandler{
		Book: Book,
	}
}
