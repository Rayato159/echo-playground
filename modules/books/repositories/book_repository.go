package repositories

import (
	"context"

	"github.com/Rayato159/go-echo/entities"
)

type booksRep struct{}

func NewBooksRepository() *booksRep {
	return &booksRep{}
}

func (r *booksRep) FindBooks(ctx context.Context, req *entities.BooksReq) []entities.Books {
	defer ctx.Deadline()

	books := []entities.Books{
		{
			Id:     1,
			Title:  "86",
			Author: "Asato Asato",
		},
		{
			Id:     2,
			Title:  "Classroom of the Elite",
			Author: "Syougo Kinugasa",
		},
	}
	return books
}

func (r *booksRep) InsertBooks(ctx context.Context, req []*entities.BooksReq) []entities.Books {
	defer ctx.Deadline()

	result := make([]entities.Books, 0)
	for i := range req {
		result = append(result, entities.Books{
			Id:     i + 1,
			Title:  req[i].Title,
			Author: req[i].Author,
		})
	}
	return result
}
