package http

import "gitlab.crja72.ru/gospec/go4/ctfplatform/checker/internal/domain/vulnServiceDomain"

type subscribeCheckerRequest struct {
	CheckerUUID string `json:"checker_uuid"`
	Ip          string `json:"ip"`
	Port        int    `json:"port"`
}

type sendServiceStatusRequest struct {
	RequestUUID string                              `json:"request_uuid"`
	StatusCode  vulnServiceDomain.VulnServiceStatus `json:"status_code"`
	Message     string                              `json:"message,omitempty"`
	WebPort     int                                 `json:"web_port,omitempty"`
	Ip          string                              `json:"ip,omitempty"`
	LastCheck   int64                               `json:"last_check"`
}
