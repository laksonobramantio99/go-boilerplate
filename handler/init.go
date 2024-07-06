package handler

type MainHandler struct {
	Book *BookHandler
}

func NewMainHandler(
	Book *BookHandler,
) *MainHandler {
	return &MainHandler{
		Book: Book,
	}
}
