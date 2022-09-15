package repositories

import (
	"go-structure-folder-clean/cmd/entity"

	"gorm.io/gorm"
)

type User = entity.User

type RepositoryUser interface {
	Save(u User) (User, error)
	FindByEmail(uInput User) (User, error)
	FindById(id int) (User, error)
	Update(usr User) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(u User) (User, error) {
	err := r.db.Create(&u).Error
	if err != nil {
		return u, err
	}

	return u, nil
}

func (r *repository) FindByEmail(uInput User) (User, error) {
	var user User
	err := r.db.Where("email = ?", uInput.Email).Find(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindById(id int) (User, error) {
	var user User
	err := r.db.Where("id = ?", id).Find(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) Update(usr User) (User, error) {
	err := r.db.Save(&usr).Error
	if err != nil {
		return usr, err
	}

	return usr, nil
}
