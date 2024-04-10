package handlers

import quizstudio "github.com/CrazyCatViking/quiz-me/template/quizstudio"

func RenderQuizStudio(requestContext *RequestContext) error {
  c := requestContext.HttpRequestContext

  return render(c, quizstudio.MainPage())
}
