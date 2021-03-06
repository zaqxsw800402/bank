package service

import (
	"red/domain"
	"red/dto"
	"red/errs"
	"time"
)

type UserService interface {
	GetAllUsers() ([]dto.UserResponse, *errs.AppError)
	GetUserByID(string) (*dto.UserResponse, *errs.AppError)
	SaveUser(dto.UserRequest, string) (*dto.UserResponse, *errs.AppError)
	GetUserByEmail(string) (*dto.UserResponse, *errs.AppError)
	GetUserWithToken(string) (*domain.Token, *errs.AppError)
	UpdatePassword(email, password string) *errs.AppError

	SaveToken(dto.UserRequest) (*dto.TokenResponse, *errs.AppError)
	//GetUser(string) (*dto.TokenResponse, *errs.AppError)
}

type DefaultUserService struct {
	repo domain.UserRepository
}

func NewUserService(repo domain.UserRepository) DefaultUserService {
	return DefaultUserService{repo: repo}
}

func (s DefaultUserService) GetAllUsers() ([]dto.UserResponse, *errs.AppError) {
	users, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	response := make([]dto.UserResponse, 0)
	for _, user := range users {
		response = append(response, user.ToDto())
	}

	return response, nil
}

func (s DefaultUserService) SaveUser(req dto.UserRequest, hash string) (*dto.UserResponse, *errs.AppError) {
	u := domain.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  hash,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	user, err := s.repo.SaveUser(u)
	if err != nil {
		return nil, err
	}

	response := user.ToDto()
	return &response, nil
}

func (s DefaultUserService) GetUserByID(id string) (*dto.UserResponse, *errs.AppError) {
	user, err := s.repo.ByID(id)
	if err != nil {
		return nil, err
	}

	response := user.ToDto()
	return &response, nil
}

func (s DefaultUserService) GetUserByEmail(email string) (*dto.UserResponse, *errs.AppError) {
	user, err := s.repo.ByEmail(email)
	if err != nil {
		return nil, err
	}

	response := user.ToDto()
	return &response, nil
}

func (s DefaultUserService) UpdatePassword(email, password string) *errs.AppError {
	u := domain.User{Email: email, Password: password}
	_, err := s.repo.UpdatePassword(u)
	if err != nil {
		return err
	}
	return nil
}

func (s DefaultUserService) SaveToken(userRequest dto.UserRequest) (*dto.TokenResponse, *errs.AppError) {
	user, err := s.repo.ByEmail(userRequest.Email)
	if err != nil {
		return nil, err
	}

	validPassword, err := passwordMatches(user.Password, userRequest.Password)
	if err != nil {
		return nil, err

	}

	if !validPassword {
		return nil, err
	}

	tokenRequest, err := dto.GenerateToken(user.ID, 24*time.Hour, dto.ScopeAuthentication)
	if err != nil {
		return nil, err
	}

	token := domain.Token{
		UserID:    int64(user.ID),
		Name:      user.LastName,
		Email:     user.Email,
		Hash:      tokenRequest.Hash,
		Expiry:    tokenRequest.Expiry,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	t, err := s.repo.SaveToken(token)
	t.PlanText = tokenRequest.PlanText
	if err != nil {
		return nil, err
	}

	response := t.ToDto()
	return &response, nil
}

func (s DefaultUserService) GetUserWithToken(token string) (*domain.Token, *errs.AppError) {
	t, err := s.repo.GetUserWithToken(token)
	if err != nil {
		return nil, err
	}
	//response := user.ToDto()
	return t, nil
}
