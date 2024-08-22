package request

type RegisterRequestPayload struct {
	Email    string `json:"email" example:"user123@gmail.com"`
	Password string `json:"password" example:"supersecret123"`
}

type LoginRequestPayload struct {
	Email    string `json:"email" example:"user123@gmail.com"`
	Password string `json:"password" example:"supersecret123"`
}
