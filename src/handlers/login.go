package handlers

import (
	"github.com/CrazyCatViking/quiz-me/src/db"
)

type LoginHandler struct {
  db *db.DbContext
}

func NewLoginHandler(db *db.DbContext) *LoginHandler {
  return &LoginHandler {
    db: db,
  }
} 
