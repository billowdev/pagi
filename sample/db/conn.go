package db

import (
	"fmt"
	"log"
	"sample/configs"
	"sample/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func NewDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(fmt.Sprintf("%s.db?create=true", configs.DB_NAME)), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic("failed to connect database")
	}
	log.Println("connect the db...")
	
	// err = db.AutoMigrate(&models.TodoModel{})
	// if err != nil {
	// 	panic("failed to auto migrate")
	// }

	// if err := SeedTodoRecords(db, 20); err != nil {
	// 	log.Fatalf("Failed to seed todo records: %v", err)
	// }

	return db, nil
}
func SeedTodoRecords(db *gorm.DB, count int) error {
	for i := 0; i < count; i++ {
		todo := &models.TodoModel{
			ID:     uint(i),
			Title:  fmt.Sprintf("Todo Title %d", i),
			Detail: fmt.Sprintf("Todo Detail %d", i),
		}
		if err := db.Save(todo).Error; err != nil {
			return err
		}
	}
	return nil
}
