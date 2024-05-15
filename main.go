package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
)

const (
	NameMaxLength  = 20
	EmailMaxLength = 40
)

var (
	ErrEmailIsTooLong = errors.New("email is too long")
	ErrNameIsTooLong  = errors.New("name is too long")

	ErrUserCreated = errors.New("user already created")
)

var usersCreated = map[string]User{}

type User struct {
	ID        string
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserReq struct {
	Name  string
	Email string
}

type UserResp struct {
	ID    string
	Name  string
	Email string
}

func NewUser(name, email string) User {
	return User{
		ID:    uuid.NewString(),
		Name:  name,
		Email: email,
	}
}

func NewName(name string) (string, error) {
	if len(name) > NameMaxLength {
		return "", ErrNameIsTooLong
	}

	return strings.TrimSpace(name), nil
}

func NewEmail(email string) (string, error) {
	if len(email) > EmailMaxLength {
		return "", ErrEmailIsTooLong
	}

	return strings.TrimSpace(email), nil
}

type UserRepo interface {
	Create(ctx context.Context, user User) error
	GetUserByID(ctx context.Context, id string) (*User, error)
}

type PostgresUserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *PostgresUserRepo {
	return &PostgresUserRepo{
		db: db,
	}
}

func (p *PostgresUserRepo) Create(ctx context.Context, user User) error {
	if user.ID == "user-created" {
		return ErrUserCreated
	}

	return nil
}

func (p *PostgresUserRepo) GetUserByID(ctx context.Context, id string) (*User, error) {
	return &User{
		ID:        id,
		Name:      "Paulo",
		Email:     "paulo7trindade@gmail.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}

type UserService struct {
	repo UserRepo
}

func NewUserService(ur UserRepo) *UserService {
	return &UserService{
		repo: ur,
	}
}

func (ur *UserService) GetUserByID(ctx context.Context, id string) (UserResp, error) {
	// Business logic
	var err error

	if id == "" {
		return UserResp{}, errors.New("empty ID")
	}

	user, err := ur.repo.GetUserByID(ctx, id)
	if err != nil {
		return UserResp{}, err
	}

	return UserResp{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (ur *UserService) Create(ctx context.Context, user UserReq) error {
	// Business logic
	var err error

	user.Email, err = NewEmail(user.Email)
	if err != nil {
		return err
	}

	user.Name, err = NewName(user.Name)
	if err != nil {
		return err
	}

	u := NewUser(user.Name, user.Email)

	err = ur.repo.Create(ctx, u)
	if err != nil {
		return err
	}

	usersCreated[u.ID] = u

	fmt.Printf("User %s created!\n", u.ID)

	return nil
}

func main() {
	ctx := context.Background()

	postgresUserRepo := NewUserRepo(&sql.DB{})
	userService := NewUserService(postgresUserRepo)

	userReq := UserReq{
		Name:  "Paulo",
		Email: "paulo7trindade@gmail.com",
	}

	err := userService.Create(ctx, userReq)
	if err != nil {
		log.Fatalf("Error trying to create a user: %v\n", err)
	}

	err = userService.Create(ctx, userReq)
	if err != nil {
		log.Fatalf("Error trying to create a user: %v\n", err)
	}

	for _, user := range usersCreated {
		user, err := userService.GetUserByID(ctx, user.ID)
		if err != nil {
			log.Fatalf("Error to get user: %v\n", err)
		}

		fmt.Printf("Welcome, %s! (%s)\n", user.Name, user.ID)
	}
}
