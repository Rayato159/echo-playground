package handlers

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/Rayato159/go-echo/entities"
	"github.com/labstack/echo/v4"
)

type booksHan struct {
	BooksUse entities.BooksUsecase
}

func NewBooksHandler(e *echo.Echo, v string, u entities.BooksUsecase) {
	h := &booksHan{
		BooksUse: u,
	}

	e.GET(fmt.Sprintf("%s/books", v), h.FindBooks)
	e.POST(fmt.Sprintf("%s/books", v), h.InsertBooks)
}

func (h *booksHan) FindBooks(c echo.Context) error {
	ctx := context.Background()

	req := &entities.BooksReq{
		Title:  strings.ToLower(c.QueryParam("title")),
		Author: strings.ToLower(c.QueryParam("author")),
	}

	// Find books
	books := h.BooksUse.FindBooks(ctx, req)
	if len(books) == 0 {
		return echo.NewHTTPError(http.StatusInternalServerError, "error, books not found")
	}
	return c.JSON(http.StatusOK, books)
}

func (h *booksHan) InsertBooks(c echo.Context) error {
	ctx := context.Background()

	req := make([]*entities.BooksReq, 0)
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "error, body parser error")
	}
	if len(req) == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "error, books are empty to insert")
	}

	// Insert books
	return c.JSON(http.StatusCreated, h.BooksUse.InsertBooks(ctx, req))
}
