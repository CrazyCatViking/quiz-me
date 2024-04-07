package handlers

import (
	"fmt"

	"github.com/CrazyCatViking/quiz-me/src/db"
	"github.com/CrazyCatViking/quiz-me/src/models"
	"github.com/CrazyCatViking/quiz-me/templates/user"
)

type UserHandler struct {
  db *db.DbContext
}

func NewUserHandler(db *db.DbContext) *UserHandler {
  return &UserHandler {
    db: db,
  }
}

func (h *UserHandler) HandleRoute(c CustomContext) error {
  var usr models.User
  h.db.Db.First(&usr)

  fmt.Println(usr)

  return render(c, user.Show())
}
