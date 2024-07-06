package repo

import (
	"go-boilerplate/model"

	"gorm.io/gorm"
)

type BookRepo interface {
	Create(book *model.Book) error
	FindByID(id uint) (*model.Book, error)
	Update(book *model.Book) error
	Delete(id uint) error
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

func (r *repo) Create(book *model.Book) error {
	return r.dbMaster.Create(book).Error
}

func (r *repo) FindByID(id uint) (*model.Book, error) {
	var book model.Book
	if err := r.dbSlave.First(&book, id).Error; err != nil {
		return nil, err
	}
	return &book, nil
}

func (r *repo) Update(book *model.Book) error {
	return r.dbMaster.Save(book).Error
}

func (r *repo) Delete(id uint) error {
	return r.dbMaster.Delete(&model.Book{}, id).Error
}
