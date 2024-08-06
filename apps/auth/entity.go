package auth

type Role string

const (
	ROLE_ADMIN Role = "admin"
	ROLE_USER  Role = "user"
)

type AuthEntity struct {
	Id       int
	Email    string
	Password string
	Role     Role
}
