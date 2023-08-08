package service

import (
	"strings"

	"github.com/google/uuid"
	"github.com/williaminfante/go_test_starter/config"
	"github.com/williaminfante/go_test_starter/entity"
	"gorm.io/gorm"
)

func CreateUser(input entity.User) (*entity.User, error) {
	input.Email = strings.ToLower(input.Email)

	db := config.GetDb() 
	input.ID = uuid.New().String() 
	
	db = db.Create(&input)
	if err := db.Error; err != nil {
		return nil, err 
	}


	return &input, nil 
}

func GetUserByEmail(email string) (*entity.User, error) {
	db := config.GetDb() 
	var user entity.User 

	if err := db.Model(user).Where("email = ?", email).Take(&user).Error; err != nil {
		return nil, err 
	}

	return &user, nil 
}


func GetAllUsers() ([]*entity.User, error) {
	db := config.GetDb() 
	var users []*entity.User 
	
	db = db.Find(&users)
	if err := db.Error; err != nil {
		return nil, err 
	}

	return users, nil 
}


func ClearAllUsers() {
	db := config.GetDb() 
	db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&entity.User{})
}