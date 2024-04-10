package handlers

import "github.com/CrazyCatViking/quiz-me/model"

type RequestContext struct {
  HttpRequestContext *models.CustomContext
}

func NewRequestContext(httpRequestContext *models.CustomContext) *RequestContext {
  return &RequestContext {
    HttpRequestContext: httpRequestContext,
  }
}
