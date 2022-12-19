package entities

import "context"

type BooksUsecase interface {
	FindBooks(ctx context.Context, req *BooksReq) []Books
	InsertBooks(ctx context.Context, req []*BooksReq) []Books
}

type BooksRepository interface {
	FindBooks(ctx context.Context, req *BooksReq) []Books
	InsertBooks(ctx context.Context, req []*BooksReq) []Books
}

type Books struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

type BooksReq struct {
	Title  string `json:"title" query:"title"`
	Author string `json:"author" query:"author"`
}
