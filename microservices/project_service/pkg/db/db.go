package db

import (
	"errors"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Project struct {
	ID          string `gorm:"primaryKey"`
	Name        string
	Description string
}

type ProjectDB interface {
	GetProjectByID(id string) (*Project, error)
	CreateProject(project *Project) error
}

type GormProjectDB struct {
	DB *gorm.DB
}

func NewGormProjectDB(dsn string) (*GormProjectDB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&Project{})
	if err != nil {
		return nil, fmt.Errorf("failed to auto migrate project: %v", err)
	}

	return &GormProjectDB{DB: db}, nil
}

func (db *GormProjectDB) GetProjectByID(id string) (*Project, error) {
	var project Project
	err := db.DB.First(&project, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("project not found")
		}
		return nil, err
	}
	return &project, nil
}

func (db *GormProjectDB) CreateProject(project *Project) error {
	return db.DB.Create(project).Error
}
