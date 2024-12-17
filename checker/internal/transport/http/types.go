package http

type subscribeCheckerRequest struct {
	Ip   string `json:"ip"`
	Port int    `json:"port"`
}

type sendServiceStatusRequest struct {
	Uuid string `json:"id"`
}
