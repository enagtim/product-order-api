package auth

import (
	"errors"
	"order-api/internal/user"
	"order-api/pkg/session"
)

type AuthService struct {
	UserRepository *user.UserRepository
}

func NewAuthService(userRepository *user.UserRepository) *AuthService {
	return &AuthService{
		UserRepository: userRepository,
	}
}

func (s *AuthService) CreateUser(phone string) (*user.User, error) {
	existedUser, _ := s.UserRepository.FindByPhone(phone)
	sessionId, err := session.GenerateSessionId()
	if err != nil {
		return nil, errors.New(ErrorGenerationSessionId)
	}
	if existedUser != nil {
		existedUser.SessionId = sessionId
		_, err := s.UserRepository.Update(existedUser)
		if err != nil {
			return nil, err
		}
		return existedUser, nil
	}
	user := &user.User{
		Phone:     phone,
		SessionId: sessionId,
		Code:      3412,
	}
	_, err = s.UserRepository.Create(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (s *AuthService) VerifyUser(phone, sessionId string) (*user.User, error) {
	user, err := s.UserRepository.FindByPhone(phone)
	if err != nil || user == nil {
		return nil, errors.New(ErrorFoundUser)
	}
	if user.SessionId != sessionId {
		return nil, errors.New(ErrorSession)
	}
	return user, nil
}
