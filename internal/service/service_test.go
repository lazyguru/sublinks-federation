package service

import (
	"log"
	"sublinks/sublinks-federation/internal/db"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type fakeDB struct {
	*gorm.DB
	Mock sqlmock.Sqlmock
}

func NewFakeDB() db.Database {
	return &fakeDB{}
}
func (d *fakeDB) Connect() error {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})

	if err != nil {
		log.Fatalf("An error '%s' was not expected when opening gorm database", err)
	}
	d.DB = gormDB
	d.Mock = mock
	return nil
}
func (db *fakeDB) Ping() bool {
	return true
}
func (db *fakeDB) RunMigrations() {
}
func (db *fakeDB) Find(interface{}, ...interface{}) error {
	return nil
}
func (db *fakeDB) Preload(model string) *gorm.DB {
	return db.DB
}
func (db *fakeDB) Save(interface{}) error {
	return nil
}
