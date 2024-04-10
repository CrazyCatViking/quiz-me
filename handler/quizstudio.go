package handler

import (
	"github.com/CrazyCatViking/quiz-me/db"
	"github.com/CrazyCatViking/quiz-me/model"
	quizstudio "github.com/CrazyCatViking/quiz-me/template/quizstudio"
)

func RenderQuizStudio(requestContext *RequestContext, db *db.DbContext) error {
  c := requestContext.HttpRequestContext

  quizzes := []model.Quiz{}
  db.Db.Find(&quizzes);

  return render(c, quizstudio.MainPage(quizzes))
}
