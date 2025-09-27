package handlers

import "github/Chidi-creator/go-medic-server/internal/usecases"

type AuthHandler struct {
	uc usecases.UserUseCase
}

func NewAuthHandler(uc usecases.UserUseCase) *AuthHandler {
	return &AuthHandler{
		uc: uc,
	}
}
