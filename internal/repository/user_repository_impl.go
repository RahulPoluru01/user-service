package repository

import (
	"context"
	"time"

	dbsqlc "github.com/RahulPoluru01/user-service/db/sqlc"
)

type PostgresUserRepository struct {
	q *dbsqlc.Queries
}

func NewPostgresUserRepository(q *dbsqlc.Queries) UserRepository {
	return &PostgresUserRepository{q: q}
}

func (r *PostgresUserRepository) Create(name string, dob time.Time) (User, error) {
	u, err := r.q.CreateUser(context.Background(), dbsqlc.CreateUserParams{
		Name: name,
		Dob:  dob,
	})
	if err != nil {
		return User{}, err
	}

	return User{
		ID:   u.ID,
		Name: u.Name,
		Dob:  u.Dob,
	}, nil
}

func (r *PostgresUserRepository) GetByID(id int32) (User, error) {
	u, err := r.q.GetUserByID(context.Background(), id)
	if err != nil {
		return User{}, err
	}

	return User{
		ID:   u.ID,
		Name: u.Name,
		Dob:  u.Dob,
	}, nil
}

func (r *PostgresUserRepository) List() ([]User, error) {
	users, err := r.q.ListUsers(context.Background())
	if err != nil {
		return nil, err
	}

	result := make([]User, 0, len(users))
	for _, u := range users {
		result = append(result, User{
			ID:   u.ID,
			Name: u.Name,
			Dob:  u.Dob,
		})
	}

	return result, nil
}

func (r *PostgresUserRepository) Update(id int32, name string, dob time.Time) (User, error) {
	u, err := r.q.UpdateUser(context.Background(), dbsqlc.UpdateUserParams{
		ID:   id,
		Name: name,
		Dob:  dob,
	})
	if err != nil {
		return User{}, err
	}

	return User{
		ID:   u.ID,
		Name: u.Name,
		Dob:  u.Dob,
	}, nil
}

func (r *PostgresUserRepository) Delete(id int32) error {
	return r.q.DeleteUser(context.Background(), id)
}
