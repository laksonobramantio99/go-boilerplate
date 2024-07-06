package usecase

import (
	"errors"
	"go-boilerplate/model"
	"go-boilerplate/repo"
)

type BookUc interface {
	CreateBook(book *model.Book) (*model.Book, error)
	GetBookByID(id uint) (*model.Book, error)
	UpdateBook(book *model.Book) (*model.Book, error)
	DeleteBook(id uint) error
}

type usecase struct {
	repo repo.BookRepo
}

func NewUsecase(repo repo.BookRepo) BookUc {
	return &usecase{repo}
}

func (uc *usecase) CreateBook(book *model.Book) (*model.Book, error) {
	if err := uc.repo.Create(book); err != nil {
		return nil, err
	}

	return book, nil
}

func (uc *usecase) GetBookByID(id uint) (*model.Book, error) {
	return uc.repo.FindByID(id)
}

func (uc *usecase) UpdateBook(book *model.Book) (*model.Book, error) {
	findBook, err := uc.repo.FindByID(book.ID)
	if err != nil {
		return nil, err
	}

	if findBook == nil {
		return nil, errors.New("book not found")
	}

	if err := uc.repo.Update(book); err != nil {
		return nil, err
	}

	return book, nil
}

func (uc *usecase) DeleteBook(id uint) error {
	return uc.repo.Delete(id)
}
