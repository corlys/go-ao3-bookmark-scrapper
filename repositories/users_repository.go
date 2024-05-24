package repositories

import (
	"go-scrapper/domain/dao"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetOrInsert(username string) (dao.Users, error)
	InsertUser(username string) (dao.Users, error)
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func (r UserRepositoryImpl) GetOrInsert(username string) (dao.Users, error) {
	var user dao.Users
	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			newUser, insertErr := r.InsertUser(username)
			if insertErr != nil {
				return dao.Users{}, insertErr
			}
			return newUser, nil
		}
		return dao.Users{}, nil
	}
	return user, nil
}

func (r UserRepositoryImpl) InsertUser(username string) (dao.Users, error) {
	user := dao.Users{
		Username: username,
	}

	err := r.db.Create(&user).Error
	if err != nil {
		return dao.Users{}, err
	}

	return user, err
}

func UserRepositoryInit(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		db: db,
	}
}
