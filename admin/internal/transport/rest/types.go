package rest

// структуры для незагрязнения основной сущности

type registrationRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type loginResponse struct {
	JwtToken string `json:"jwt_token"`
}
