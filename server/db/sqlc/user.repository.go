package sqlc

import (
	"context"
	"server/internal/user"
)

type userRepsitory struct {
	q *Queries
}

func (r *userRepsitory) CreateUser(ctx context.Context, u *user.User) (*user.User, error) {
	arg := CreateUserParams{
		Username: u.Username,
		Password: u.Password,
		Email:    u.Email,
	}
	result, err := r.q.CreateUser(ctx, arg)
	if err != nil {
		return nil, err
	}

	return &user.User{
		ID:       result.ID,
		Username: result.Username,
		Email:    result.Email,
		Password: result.Password,
	}, nil
}

func (r *userRepsitory) GetUserByEmail(ctx context.Context, email string) (*user.User, error) {
	result, err := r.q.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return &user.User{
		ID:       result.ID,
		Username: result.Username,
		Email:    result.Email,
		Password: result.Password,
	}, nil
}

func NewUserRepository(q *Queries) user.Repository {
	return &userRepsitory{q: q}
}
func (r *userRepsitory) Getuserlist(ctx context.Context) ([]*user.GetuserlistRes, error) {
	rows, err := r.q.Getuserlist(ctx)
	if err != nil {
		return nil, err
	}
	var users []*user.GetuserlistRes
	for _, row := range rows {
		u := &user.GetuserlistRes{
			Userid:   row.ID,
			Username: row.Username,
			Email:    row.Email,
		}
		users = append(users, u)
	}
	return users, nil
}
