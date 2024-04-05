package repository

import (
	"backend/entity"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AuthRepository interface {
	Register(req *entity.User) error
	UsernameExist(username string) bool
	Login(username string) (*entity.User, error)
}

type authRepository struct {
	db *pgxpool.Pool
}

func NewAuthRepository(db *pgxpool.Pool) *authRepository {
	return &authRepository{
		db: db,
	}
}

func (r *authRepository) Register(users *entity.User) error {
	// entries := [][]any{}
	columns := []string{"name", "username", "email", "password"}
	values := []interface{}{users.Name, users.Username, users.Email, users.Password}
	tableName := "users"

	// fields := reflect.VisibleFields(reflect.TypeOf(users))
	// entries = append(entries, []any{users.Name})
	// entries = append(entries, []any{users.Username})
	// entries = append(entries, []any{users.Email})
	// entries = append(entries, []any{users.Password})

	_, err := r.db.CopyFrom(context.Background(),
		pgx.Identifier{"users"},
		columns,
		pgx.CopyFromRows([][]interface{}{values}))

	if err != nil {
		return fmt.Errorf("error copying into %s table: %w", tableName, err)
	}

	return err
}

func (r *authRepository) UsernameExist(username string) bool {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE username = $1)`
	err := r.db.QueryRow(context.Background(), query, username).Scan(&exists)
	if err != nil {
		return false
	}
	return exists
}

func (r *authRepository) Login(username string) (*entity.User, error) {
	var user entity.User
	query := `SELECT username, name, password FROM users WHERE username = $1;`
	err := r.db.QueryRow(context.Background(), query, username).Scan(&user.Username, &user.Name, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, err
}
