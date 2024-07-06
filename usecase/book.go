package usecase

import (
	"errors"
	"go-boilerplate/model"
	"go-boilerplate/repo"

	"time"
)

type BookUc interface {
	CreateBook(title, author, genre string, publishedDate time.Time) (*model.Book, error)
	GetBook(id uint) (*model.Book, error)
	UpdateBook(id uint, title, author, genre string, publishedDate time.Time) (*model.Book, error)
	DeleteBook(id uint) error
}

type usecase struct {
	repo repo.BookRepo
}

func NewUsecase(repo repo.BookRepo) BookUc {
	return &usecase{repo}
}

func (uc *usecase) CreateBook(title, author, genre string, publishedDate time.Time) (*model.Book, error) {
	book := &model.Book{
		Title:         title,
		Author:        author,
		Genre:         genre,
		PublishedDate: publishedDate,
	}
	if err := uc.repo.Create(book); err != nil {
		return nil, err
	}
	return book, nil
}

func (uc *usecase) GetBook(id uint) (*model.Book, error) {
	return uc.repo.FindByID(id)
}

func (uc *usecase) UpdateBook(id uint, title, author, genre string, publishedDate time.Time) (*model.Book, error) {
	book, err := uc.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if book == nil {
		return nil, errors.New("book not found")
	}

	book.Title = title
	book.Author = author
	book.Genre = genre
	book.PublishedDate = publishedDate
	book.UpdatedAt = time.Now()

	if err := uc.repo.Update(book); err != nil {
		return nil, err
	}
	return book, nil
}

func (uc *usecase) DeleteBook(id uint) error {
	return uc.repo.Delete(id)
}
