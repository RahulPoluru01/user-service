package service

import (
	"errors"
	"time"

	"github.com/RahulPoluru01/user-service/internal/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func calculateAge(dob time.Time) int {
	now := time.Now()
	age := now.Year() - dob.Year()

	if now.YearDay() < dob.YearDay() {
		age--
	}

	return age
}

type UserResponse struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
	Dob  string `json:"dob"`
	Age  int    `json:"age"`
}

func (s *UserService) CreateUser(name string, dob time.Time) (repository.User, error) {
	if name == "" {
		return repository.User{}, errors.New("name cannot be empty")
	}

	return s.repo.Create(name, dob)
}

func (s *UserService) GetUserByID(id int32) (UserResponse, error) {
	u, err := s.repo.GetByID(id)
	if err != nil {
		return UserResponse{}, err
	}

	return UserResponse{
		ID:   u.ID,
		Name: u.Name,
		Dob:  u.Dob.Format("2006-01-02"),
		Age:  calculateAge(u.Dob),
	}, nil
}

func (s *UserService) ListUsers() ([]UserResponse, error) {
	users, err := s.repo.List()
	if err != nil {
		return nil, err
	}

	result := make([]UserResponse, 0, len(users))
	for _, u := range users {
		result = append(result, UserResponse{
			ID:   u.ID,
			Name: u.Name,
			Dob:  u.Dob.Format("2006-01-02"),
			Age:  calculateAge(u.Dob),
		})
	}

	return result, nil
}

func (s *UserService) UpdateUser(id int32, name string, dob time.Time) (repository.User, error) {
	return s.repo.Update(id, name, dob)
}

func (s *UserService) DeleteUser(id int32) error {
	return s.repo.Delete(id)
}
