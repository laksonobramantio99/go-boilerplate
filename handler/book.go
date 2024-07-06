package handler

import (
	"go-boilerplate/usecase"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	Usecase usecase.BookUc
}

func NewBookHandler(usecase usecase.BookUc) *BookHandler {
	return &BookHandler{Usecase: usecase}
}

func (h *BookHandler) Mount(c *gin.RouterGroup) {
	g := c.Group("/books")
	{
		g.POST("", h.CreateBook)
		g.GET("/:id", h.GetBook)
		g.PUT("/:id", h.UpdateBook)
		g.DELETE("/:id", h.DeleteBook)
	}
}

func (h *BookHandler) CreateBook(c *gin.Context) {
	var request struct {
		Title         string    `json:"title" binding:"required"`
		Author        string    `json:"author" binding:"required"`
		Genre         string    `json:"genre"`
		PublishedDate time.Time `json:"published_date"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book, err := h.Usecase.CreateBook(request.Title, request.Author, request.Genre, request.PublishedDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"book": book})
}

func (h *BookHandler) GetBook(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	book, err := h.Usecase.GetBook(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if book == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"book": book})
}

func (h *BookHandler) UpdateBook(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	var request struct {
		Title         string    `json:"title"`
		Author        string    `json:"author"`
		Genre         string    `json:"genre"`
		PublishedDate time.Time `json:"published_date"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book, err := h.Usecase.UpdateBook(uint(id), request.Title, request.Author, request.Genre, request.PublishedDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if book == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"book": book})
}

func (h *BookHandler) DeleteBook(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	if err := h.Usecase.DeleteBook(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}
