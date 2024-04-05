package service

import (
	"backend/entity"
	"backend/errorHandler"
	"backend/helper"
	"backend/repository"
)

type AuthService interface {
	Register(req *entity.User) (*entity.RegisterResponse, error)
	Login(req *entity.LoginRequest) (*entity.LoginResponse, error)
}

type authService struct {
	repository repository.AuthRepository
}

func NewAuthService(r repository.AuthRepository) *authService {
	return &authService{
		repository: r,
	}
}

func (s *authService) Register(req *entity.User) (*entity.RegisterResponse, error) {
	var data entity.RegisterResponse
	if usernameExist := s.repository.UsernameExist(req.Username); usernameExist {

		return nil, &errorHandler.Error409{Message: "Username already exists!"}
	}

	passwordHash, err := helper.HashPassword(req.Password)
	if err != nil {
		return nil, &errorHandler.Error500{Message: err.Error()}
	}
	user := entity.User{
		Name:     req.Name,
		Username: req.Username,
		Email:    req.Email,
		Password: passwordHash,
	}

	if err := s.repository.Register(&user); err != nil {
		return nil, &errorHandler.Error500{Message: "Internal server error"}
	}
	token, err := helper.GenerateToken(&user)
	if err != nil {
		return nil, &errorHandler.Error500{Message: "Internal server error"}
	}
	data = entity.RegisterResponse{
		Username:    user.Username,
		Name:        user.Name,
		AccessToken: token,
	}
	return &data, nil
}

func (s *authService) Login(req *entity.LoginRequest) (*entity.LoginResponse, error) {
	var data entity.LoginResponse
	user, err := s.repository.Login(req.Username)

	if err != nil {
		return nil, &errorHandler.Error500{Message: "Internal server error"}
	}
	if err := helper.VerifyPassword(user.Password, req.Password); err != nil {
		return nil, &errorHandler.Error400{Message: "Username or password wrong"}
	}
	token, err := helper.GenerateToken(user)
	if err != nil {
		return nil, &errorHandler.Error500{Message: err.Error()}
	}

	data = entity.LoginResponse{
		Username:    user.Username,
		Name:        user.Name,
		AccessToken: token,
	}

	return &data, nil

}
