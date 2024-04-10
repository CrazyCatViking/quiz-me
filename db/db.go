package db

import (
	"github.com/CrazyCatViking/quiz-me/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DbContext struct {
  Db *gorm.DB
}

func Init() *DbContext {
  dsn := "host=localhost user=postgres password=password dbname=quizme-test port=5432 sslmode=disable TimeZone=Europe/Oslo"
  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

  if err != nil {
    panic("failed to connect database")
  }
  
  Migrate(db)
 
  return &DbContext{ db }
}

func Migrate(db *gorm.DB) {
  db.AutoMigrate(&model.User{})
  db.AutoMigrate(&model.Quiz{})
}
