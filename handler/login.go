package handlers

import (
	"github.com/CrazyCatViking/quiz-me/template/login"
)

type LoginHandler struct {
  requestContext *RequestContext
}

func (h *LoginHandler) HandleRoute() error {
  c := h.requestContext.HttpRequestContext

  return render(c, login.LoginPage())
}

func NewLoginHandler(requestContext *RequestContext) *LoginHandler {
  return &LoginHandler {
    requestContext: requestContext,
  }
} 
