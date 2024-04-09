package handlers

import (
	"fmt"

	"github.com/CrazyCatViking/quiz-me/db"
	"github.com/CrazyCatViking/quiz-me/models"
	"github.com/CrazyCatViking/quiz-me/templates/user"
)

type UserHandler struct {
  db *db.DbContext
  requestContext *RequestContext
}

func NewUserHandler(db *db.DbContext, requestContext *RequestContext) *UserHandler {
  return &UserHandler {
    db: db,
    requestContext: requestContext,
  }
}

func (h *UserHandler) HandleRoute() error {
  var usr models.User
  h.db.Db.First(&usr)

  fmt.Println(usr)

  c := h.requestContext.HttpRequestContext

  return render(c, user.Show())
}
