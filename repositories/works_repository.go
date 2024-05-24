package repositories

import (
	"go-scrapper/domain/dao"
	"gorm.io/gorm"
	"time"
)

type WorksRepository interface {
	GetAllWorks() ([]dao.Works, error)
	InsertWork(title string, workUpdatedAt time.Time, author dao.Users) (dao.Works, error)
	GetOrInsertWork(title string, workUpdatedAt time.Time, author dao.Users) (dao.Works, error)
}

type WorksRepositoryImpl struct {
	db *gorm.DB
}

func (r WorksRepositoryImpl) GetAllWorks() ([]dao.Works, error) {
	var works []dao.Works
	err := r.db.Find(&works).Error
	if err != nil {
		return nil, err
	}
	return works, nil
}

func (r WorksRepositoryImpl) GetWorkByTitle(title string) (dao.Works, error) {
	var work dao.Works
	err := r.db.Where("title = ?", title).First(&work).Error
	if err != nil {
		return dao.Works{}, err
	}
	return work, nil
}

func (r WorksRepositoryImpl) InsertWork(title string, workUpdatedAt time.Time, author dao.Users) (dao.Works, error) {

	work := dao.Works{
		Title:           title,
		WorkUpdatedDate: workUpdatedAt,
		AuthorUsername:  author.Username,
	}

	res := r.db.Create(&work)
	if res.Error != nil {
		return dao.Works{}, res.Error
	}

	var insertedWork dao.Works
	r.db.Preload("Users").First(&insertedWork, work.ID)

	return insertedWork, nil

}

func (r WorksRepositoryImpl) GetOrInsertWork(title string, workUpdatedAt time.Time, author dao.Users) (dao.Works, error) {

	work, err := r.GetWorkByTitle(title)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			insertedWork, errInsert := r.InsertWork(title, workUpdatedAt, author)
			if errInsert != nil {
				return dao.Works{}, errInsert
			}
			return insertedWork, nil
		}
	}
	return work, nil

}

func WorksRepositoryInit(db *gorm.DB) *WorksRepositoryImpl {
	return &WorksRepositoryImpl{
		db: db,
	}
}
