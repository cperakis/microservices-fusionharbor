package db

import (
	"errors"

	"github.com/fusionharbor/microservices/api/auth"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UserStore interface {
	GetUser(username string) (*auth.User, error)
	GetUserByID(id string) (*auth.User, error)
	CreateUser(user *auth.User) error
}

type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string
	Email    string `gorm:"unique"`
}

type gormUserStore struct {
	db *gorm.DB
}

func (s *gormUserStore) GetUser(username string) (*auth.User, error) {
	user := User{}
	err := s.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	userResponse := &auth.User{}
	userResponse.Email = user.Email
	userResponse.Password = user.Password
	userResponse.Username = user.Username
	return userResponse, nil
}

func (s *gormUserStore) GetUserByID(id string) (*auth.User, error) {
	user := User{}
	err := s.db.Where("id = ?", id).First(user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	userResponse := &auth.User{}
	userResponse.Email = user.Email
	userResponse.Username = user.Username
	return userResponse, nil
}

func (s *gormUserStore) CreateUser(user *auth.User) error {
	userDB := User{}
	userDB.Email = user.Email
	userDB.Password = user.Password
	userDB.Username = user.Username
	err := s.db.Create(&userDB).Error
	return err
}

func NewGormUserStore(dsn string) (UserStore, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&User{})
	if err != nil {
		return nil, err
	}

	return &gormUserStore{db: db}, nil
}
