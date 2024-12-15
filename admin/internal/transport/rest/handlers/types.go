package handlers

// структуры для незагрязнения основной сущности

// User запросы

type userRegistrationRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type userLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type userLoginResponse struct {
	JwtToken string `json:"jwt_token"`
}

// Team запросы
type teamRegistrationRequest struct {
	TeamName string `json:"team_name"`
}
