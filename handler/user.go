package handler

import (
	"fmt"

	"github.com/CrazyCatViking/quiz-me/db"
	"github.com/CrazyCatViking/quiz-me/model"
	"github.com/CrazyCatViking/quiz-me/template/user"
)

func ShowUser(db *db.DbContext, requestContext *RequestContext) error {
  var usr model.User
  db.Db.First(&usr)

  fmt.Println(usr)

  c := requestContext.HttpRequestContext

  return render(c, user.Show())
}
