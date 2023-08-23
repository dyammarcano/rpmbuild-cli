package database

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	db *gorm.DB
}

// NewDatabase opens a connection to the database file
func NewDatabase(filePath string) (*Database, error) {
	db, err := gorm.Open(sqlite.Open(filePath), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	return &Database{db: db}, nil
}

// Close closes the connection to the database file
func (d *Database) Close() {
	db, err := d.db.DB()
	if err != nil {
		panic("Failed to close database")
	}

	if err := db.Close(); err != nil {
		return
	}
}

// Migrate creates the tables if they don't exist
func (d *Database) Migrate(object ...any) error {
	return d.db.AutoMigrate(object...)
}

// Create adds a new object to the database
func (d *Database) Create(object any) error {
	return d.db.Create(object).Error
}

// Update updates an object in the database
func (d *Database) Update(object any) error {
	return d.db.Save(object).Error
}

// Delete deletes an object from the database
func (d *Database) Delete(object any) error {
	return d.db.Delete(object).Error
}

// Find finds an object in the database
func (d *Database) Find(object any, id string) (any, error) {
	return object, d.db.First(object, id).Error
}

// FindAll finds all objects in the database
func (d *Database) FindAll(object any) (any, error) {
	return object, d.db.Find(object).Error
}

// FindBy finds an object in the database by a field
func (d *Database) FindBy(object any, field string, value string) error {
	return d.db.Where(fmt.Sprintf("%s = ?", field), value).Find(object).Error
}

// FindByAll finds all objects in the database by a field
func (d *Database) FindByAll(object any, field string, value string) error {
	return d.db.Where(fmt.Sprintf("%s = ?", field), value).Find(object).Error
}

// BulkInsert inserts an object in the database
func (d *Database) BulkInsert(object []any) error {
	return d.db.Create(object).Error
}

// BulkUpdate updates an object in the database
func (d *Database) BulkUpdate(object []any) error {
	return d.db.Save(object).Error
}

// BulkDelete deletes an object from the database
func (d *Database) BulkDelete(object []any) error {
	return d.db.Delete(object).Error
}
