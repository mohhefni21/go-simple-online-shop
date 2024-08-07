package auth

import (
	"mohhefni/go-online-shop/infra/response"
	"strings"
	"time"
)

type Role string

const (
	ROLE_ADMIN Role = "admin"
	ROLE_USER  Role = "user"
)

type AuthEntity struct {
	Id         int       `db:"id"`
	Email      string    `db:"email"`
	Password   string    `db:"password"`
	Role       Role      `db:"role"`
	CreatedAt  time.Time `db:"created_at"`
	UpdateddAt time.Time `db:"updated_at"`
}

func NewFromRegisterRequest(req RegisterRequestPayload) AuthEntity {
	return AuthEntity{
		Email:      req.Email,
		Password:   req.Password,
		Role:       ROLE_USER,
		CreatedAt:  time.Now(),
		UpdateddAt: time.Now(),
	}
}

func (a *AuthEntity) Validate() (err error) {
	if err = a.ValidateEmail(); err != nil {
		return
	}
	if err = a.ValidatePassword(); err != nil {
		return
	}
	return
}

func (a *AuthEntity) ValidateEmail() (err error) {
	if a.Email == "" {
		return response.ErrEmailRequired
	}

	splitEmail := strings.Split(a.Email, "@")
	if len(splitEmail) != 2 {
		return response.ErrEmailInvalid
	}

	return
}

func (a *AuthEntity) ValidatePassword() (err error) {
	if a.Password == "" {
		return response.ErrPasswordRequired
	}

	if len(a.Password) < 6 {
		return response.ErrPasswordInvalidLength
	}

	return
}
