package db

import (
	"errors"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Project represents a project in the database.
type Project struct {
	gorm.Model
	Name        string
	Description string
	Metadata    string
}

// ProjectDB provides an interface for interacting with the project database.
type ProjectDB interface {
	GetProjectByID(id string) (*Project, error)
	CreateProject(project *Project) error
	DeleteProject(id string) error
	GetProjects() ([]*Project, error) // New function for getting all projects
}

// GormProjectDB is a Gorm implementation of the ProjectDB interface.
type GormProjectDB struct {
	DB *gorm.DB
}

// NewGormProjectDB creates a new GormProjectDB instance.
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

// GetProjectByID retrieves a project from the database based on its ID.
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

// CreateProject creates a new project in the database.
func (db *GormProjectDB) CreateProject(project *Project) error {
	return db.DB.Create(project).Error
}

// DeleteProject deletes a project from the database based on its ID.
func (db *GormProjectDB) DeleteProject(id string) error {
	return db.DB.Delete(&Project{}, "id = ?", id).Error
}

// GetProjects retrieves all projects from the database.
func (db *GormProjectDB) GetProjects() ([]*Project, error) {
	var projects []*Project
	err := db.DB.Find(&projects).Error
	if err != nil {
		return nil, err
	}
	return projects, nil
}
