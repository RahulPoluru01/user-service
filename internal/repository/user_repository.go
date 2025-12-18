package repository

import "time"

type User struct {
	ID   int32
	Name string
	Dob  time.Time
}

type UserRepository interface {
	Create(name string, dob time.Time) (User, error)
	GetByID(id int32) (User, error)
	List() ([]User, error)
	Update(id int32, name string, dob time.Time) (User, error)
	Delete(id int32) error
}
