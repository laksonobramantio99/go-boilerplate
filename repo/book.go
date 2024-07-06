package repo

import (
	"context"
	"go-boilerplate/model"

	"gorm.io/gorm"
)

type BookRepo interface {
	Create(ctx context.Context, book *model.Book) error
	FindByID(ctx context.Context, id uint) (*model.Book, error)
	Update(ctx context.Context, book *model.Book) error
	Delete(ctx context.Context, id uint) error
}

type repo struct {
	dbMaster *gorm.DB
	dbSlave  *gorm.DB
}

func NewBookRepo(dbMaster, dbSlave *gorm.DB) BookRepo {
	return &repo{
		dbMaster: dbMaster,
		dbSlave:  dbSlave,
	}
}

func (r *repo) Create(ctx context.Context, book *model.Book) error {
	return r.dbMaster.WithContext(ctx).Create(book).Error
}

func (r *repo) FindByID(ctx context.Context, id uint) (*model.Book, error) {
	var book model.Book
	if err := r.dbSlave.WithContext(ctx).First(&book, id).Error; err != nil {
		return nil, err
	}
	return &book, nil
}

func (r *repo) Update(ctx context.Context, book *model.Book) error {
	return r.dbMaster.WithContext(ctx).Save(book).Error
}

func (r *repo) Delete(ctx context.Context, id uint) error {
	return r.dbMaster.WithContext(ctx).Delete(&model.Book{}, id).Error
}
