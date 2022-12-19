package usecases

import (
	"context"

	"github.com/Rayato159/go-echo/entities"
)

type booksUse struct {
	BooksRep entities.BooksRepository
}

func NewBooksUsecase(rep entities.BooksRepository) *booksUse {
	return &booksUse{
		BooksRep: rep,
	}
}

func (u *booksUse) FindBooks(ctx context.Context, req *entities.BooksReq) []entities.Books {
	defer ctx.Deadline()

	books := u.BooksRep.FindBooks(ctx, req)
	newBooks := make([]entities.Books, 0)
	if req.Author != "" || req.Title != "" {
		for i := range books {
			if books[i].Author == req.Author || books[i].Title == req.Title {
				newBooks = append(newBooks, books[i])
			}
		}
		return newBooks
	} else {
		return books
	}
}

func (u *booksUse) InsertBooks(ctx context.Context, req []*entities.BooksReq) []entities.Books {
	defer ctx.Deadline()
	return u.BooksRep.InsertBooks(ctx, req)
}
