package db

import (
	"errors"
	"strconv"

	"github.com/fusionharbor/microservices/api/auth"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UserStore interface {
	GetUser(username string) (*auth.User, error)
	GetUsers() ([]*auth.User, error)
	GetUserByID(id string) (*auth.User, error)
	DeleteUser(id string) error
	CreateUser(user *auth.User) error
}

type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string
	Email    string `gorm:"unique"`
	Role     string
	Team     string
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
	userResponse.Role = user.Role
	userResponse.Password = user.Password
	userResponse.Username = user.Username
	return userResponse, nil
}

func (s *gormUserStore) GetUsers() ([]*auth.User, error) {
	var users []User
	err := s.db.Find(&users).Error
	if err != nil {
		return nil, err
	}

	userResponses := make([]*auth.User, len(users))
	for i, user := range users {
		userResponses[i] = &auth.User{
			Id:       strconv.Itoa(int(user.ID)),
			Username: user.Username,
			Email:    user.Email,
			Role:     user.Role,
			// You may not want to include the password in the response
		}
	}

	return userResponses, nil
}

func (s *gormUserStore) GetUserByID(id string) (*auth.User, error) {
	user := User{}
	err := s.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	userResponse := &auth.User{}
	userResponse.Email = user.Email
	userResponse.Role = user.Role
	userResponse.Username = user.Username
	return userResponse, nil
}

func (s *gormUserStore) CreateUser(user *auth.User) error {
	userDB := User{}
	userDB.Email = user.Email
	userDB.Role = user.Role
	userDB.Password = user.Password
	userDB.Username = user.Username
	err := s.db.Create(&userDB).Error
	return err
}

func (s *gormUserStore) DeleteUser(id string) error {
	user := User{}
	err := s.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("user not found")
		}
		return err
	}
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
