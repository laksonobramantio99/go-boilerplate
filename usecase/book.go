package usecase

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go-boilerplate/client/redis"
	"go-boilerplate/model"
	"go-boilerplate/repo"
	"time"
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
	return &usecase{
		repo: repo,
	}
}

func (uc *usecase) CreateBook(book *model.Book) (*model.Book, error) {
	if err := uc.repo.Create(book); err != nil {
		return nil, err
	}

	return book, nil
}

func (uc *usecase) GetBookByID(id uint) (*model.Book, error) {
	cacheKey := "book:" + fmt.Sprint(id)
	cachedBook, err := redis.RedisClient.Get(context.Background(), cacheKey).Result()
	if err == nil {
		var book model.Book
		if err := json.Unmarshal([]byte(cachedBook), &book); err == nil {
			return &book, nil
		}
	}

	book, err := uc.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	// set to redis
	bookJson, err := json.Marshal(book)
	if err == nil {
		go redis.RedisClient.Set(context.TODO(), cacheKey, bookJson, 1*time.Hour) // Adjust expiration as needed
	}

	return book, nil
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

	// delete from redis
	cacheKey := "book:" + fmt.Sprint(book.ID)
	redis.RedisClient.Del(context.TODO(), cacheKey)

	return book, nil
}

func (uc *usecase) DeleteBook(id uint) error {
	err := uc.repo.Delete(id)
	if err != nil {
		return err
	}

	// delete from redis
	cacheKey := "book:" + fmt.Sprint(id)
	redis.RedisClient.Del(context.TODO(), cacheKey)

	return nil
}
