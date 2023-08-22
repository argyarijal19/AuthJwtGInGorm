package user

import (
	"context"

	"belajar-restapi/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserTabel struct {
	Gorm *gorm.DB
}

func UserTables(db *gorm.DB) *UserTabel {
	return &UserTabel{Gorm: db.Table("user_simoga").Session(&gorm.Session{}).WithContext(context.Background())}
}

func (db *UserTabel) GetAllUser(dest *[]models.UserSimgoa) error {
	return db.Gorm.Find(dest).Error
}

func (db *UserTabel) CreateUser(user *models.UserSimgoa) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return db.Gorm.Create(user).Error

}

func (db *UserTabel) GetUserByUsername(username string) (*models.UserSimgoa, error) {
	var user models.UserSimgoa
	err := db.Gorm.Model(&user).Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
