package handler

import (
	"go-boilerplate/model"
	"go-boilerplate/usecase"
	"go-boilerplate/util/resp"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	uc usecase.BookUc
}

func NewBookHandler(uc usecase.BookUc) *BookHandler {
	return &BookHandler{uc: uc}
}

func (h *BookHandler) Mount(c *gin.RouterGroup) {
	g := c.Group("/books")
	{
		g.POST("", h.CreateBook)
		g.GET("/:id", h.GetBook)
		g.PUT("", h.UpdateBook)
		g.DELETE("/:id", h.DeleteBook)
	}
}

func (h *BookHandler) CreateBook(c *gin.Context) {
	var b model.Book
	if err := c.ShouldBindJSON(&b); err != nil {
		c.JSON(http.StatusBadRequest, resp.Error("err read req body: "+err.Error()))
		return
	}

	res, err := h.uc.CreateBook(&b)
	if err != nil {
		c.JSON(http.StatusInternalServerError, resp.Error("err create book: "+err.Error()))
		return
	}

	c.JSON(http.StatusCreated, resp.Success("success", res))
}

func (h *BookHandler) GetBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, resp.Error("invalid id: "+err.Error()))
		return
	}

	book, err := h.uc.GetBookByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, resp.Error("err get book: "+err.Error()))
		return
	}

	c.JSON(http.StatusOK, book)
}

func (h *BookHandler) UpdateBook(c *gin.Context) {
	var b model.Book
	if err := c.ShouldBindJSON(&b); err != nil {
		c.JSON(http.StatusBadRequest, resp.Error("err read req body: "+err.Error()))
		return
	}

	res, err := h.uc.UpdateBook(&b)
	if err != nil {
		c.JSON(http.StatusInternalServerError, resp.Error("err update book: "+err.Error()))
		return
	}

	c.JSON(http.StatusOK, resp.Success("success", res))
}

func (h *BookHandler) DeleteBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, resp.Error("invalid id: "+err.Error()))
		return
	}

	if err := h.uc.DeleteBook(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, resp.Error("err delete book: "+err.Error()))
		return
	}

	c.JSON(http.StatusOK, resp.Success("successfully deleted", nil))
}
