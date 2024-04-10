package handler

import "github.com/CrazyCatViking/quiz-me/model"


type RequestContext struct {
  HttpRequestContext *model.CustomContext
}

func NewRequestContext(httpRequestContext *model.CustomContext) *RequestContext {
  return &RequestContext {
    HttpRequestContext: httpRequestContext,
  }
}
