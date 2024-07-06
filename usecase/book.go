package usecase

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go-boilerplate/client/redis"
	"go-boilerplate/model"
	"go-boilerplate/repo"
	"go-boilerplate/util/logger"
	"time"
)

type BookUc interface {
	CreateBook(ctx context.Context, book *model.Book) (*model.Book, error)
	GetBookByID(ctx context.Context, id uint) (*model.Book, error)
	UpdateBook(ctx context.Context, book *model.Book) (*model.Book, error)
	DeleteBook(ctx context.Context, id uint) error
}

type usecase struct {
	repo repo.BookRepo
}

func NewUsecase(repo repo.BookRepo) BookUc {
	return &usecase{
		repo: repo,
	}
}

func (uc *usecase) CreateBook(ctx context.Context, book *model.Book) (*model.Book, error) {
	if err := uc.repo.Create(ctx, book); err != nil {
		return nil, err
	}

	return book, nil
}

func (uc *usecase) GetBookByID(ctx context.Context, id uint) (*model.Book, error) {
	cacheKey := "book:" + fmt.Sprint(id)
	cachedBook, err := redis.RedisClient.Get(context.Background(), cacheKey).Result()
	if err == nil {
		var book model.Book
		if err := json.Unmarshal([]byte(cachedBook), &book); err == nil {
			return &book, nil
		}
	}

	book, err := uc.repo.FindByID(ctx, id)
	if err != nil {
		logger.Errorf(ctx, "uc.repo.FindByID: %v", err)
		return nil, err
	}

	// set to redis
	bookJson, err := json.Marshal(book)
	if err == nil {
		go redis.RedisClient.Set(context.TODO(), cacheKey, bookJson, 1*time.Hour) // Adjust expiration as needed
	}

	return book, nil
}

func (uc *usecase) UpdateBook(ctx context.Context, book *model.Book) (*model.Book, error) {
	findBook, err := uc.repo.FindByID(ctx, book.ID)
	if err != nil {
		return nil, err
	}

	if findBook == nil {
		return nil, errors.New("book not found")
	}

	if err := uc.repo.Update(ctx, book); err != nil {
		return nil, err
	}

	// delete from redis
	cacheKey := "book:" + fmt.Sprint(book.ID)
	redis.RedisClient.Del(context.TODO(), cacheKey)

	return book, nil
}

func (uc *usecase) DeleteBook(ctx context.Context, id uint) error {
	err := uc.repo.Delete(ctx, id)
	if err != nil {
		return err
	}

	// delete from redis
	cacheKey := "book:" + fmt.Sprint(id)
	redis.RedisClient.Del(context.TODO(), cacheKey)

	return nil
}
