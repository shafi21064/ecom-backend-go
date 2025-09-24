package repo

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	IsOwner  bool   `json:"is_owner"`
}

type UserRepo interface {
	Create(User) (*User, error)
	Get(email string, password string) (*User, error)
}

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}

// Create inserts a new user into the database
func (r *userRepo) Create(u User) (*User, error) {
	query := `
		INSERT INTO users (name, email, password, is_owner)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	// Execute the query and get the inserted user
	row := r.db.QueryRow(query, u.Name, u.Email, u.Password, u.IsOwner)
	err := row.Scan(&u.ID)
	if err != nil {
		return nil, err
	}

	// Don't return the password in response
	u.Password = ""
	return &u, nil
}

func (r *userRepo) Get(email string, password string) (*User, error) {
	var u User

	// sqlx allows struct mapping automatically
	query := `SELECT id, name, email, password, is_owner FROM users WHERE email=$1 AND password=$2`
	err := r.db.Get(&u, query, email, password)
	if err != nil {
		if err == sql.ErrNoRows || err.Error() == "sql: no rows in result set" {
			return nil, nil // user not found
		}
		return nil, err
	}

	u.Password = "" // do not return password
	return &u, nil
}
