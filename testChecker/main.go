package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type Response struct {
	Status string `json:"status"`
}

type checkRequest struct {
	RequestUUID     string `json:"request_uuid"`
	VulnServiceIP   string `json:"vuln_service_ip"`
	VulnServicePort int    `json:"vuln_service_port"`
	Flag            string `json:"flag"`
}

type checkResponse struct {
	IsTaskAccepted bool `json:"is_task_accepted"`
}

type sendServiceStatusRequest struct {
	RequestUUID string `json:"request_uuid"`
	StatusCode  int    `json:"status_code"`
	Message     string `json:"message,omitempty"`
	WebPort     int    `json:"web_port,omitempty"`
	Ip          string `json:"ip,omitempty"`
	LastCheck   int64  `json:"last_check"`
}

func main() {
	uuid := os.Getenv("UUID")
	if uuid == "" {
		log.Fatal("UUID environment variable not set")
	}

	url := "http://localhost:4010/checker/subscribe"

	subscribePayload := fmt.Sprintf(`{"checker_uuid":"%v","ip":"%v","port":%v}`, uuid, "127.0.0.1", 4030)
	resp, err := http.Post(url, "application/json", strings.NewReader(subscribePayload))
	if err != nil {
		log.Fatalf("Failed to send POST request: %v", err)
	}
	defer resp.Body.Close()

	var response Response
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		log.Fatalf("Failed to parse response: %v", err)
	}

	if response.Status != "success" {
		log.Fatalf("Subscription failed, got status: %s", response.Status)
	}

	log.Println("Subscription successful, starting server...")

	http.HandleFunc("/checkVulnService", checkVulnServiceHandler)
	log.Fatal(http.ListenAndServe(":4030", nil))
}

func checkVulnServiceHandler(w http.ResponseWriter, r *http.Request) {
	var req checkRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	log.Printf("Received check request: %+v", req)

	response := checkResponse{
		IsTaskAccepted: true, // Placeholder logic, adjust as needed
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}

	// Execute additional task logic after acceptance
	client := &http.Client{}
	checkReq, err := http.NewRequest("GET", "http://localhost:4040/", nil)
	if err != nil {
		log.Printf("Failed to create GET request: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	checkReq.AddCookie(&http.Cookie{
		Name:  "lang",
		Value: "echo $HOSTNAME",
	})

	resp, err := client.Do(checkReq)
	if err != nil {
		log.Printf("Failed to send GET request: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		statusPayload := sendServiceStatusRequest{
			RequestUUID: req.RequestUUID,
			StatusCode:  http.StatusOK,
			Message:     "Service is operational",
			WebPort:     req.VulnServicePort,
			Ip:          req.VulnServiceIP,
			LastCheck:   time.Now().Unix(),
		}

		statusURL := "http://localhost:4010/checker/sendServiceStatus"
		payloadBytes, err := json.Marshal(statusPayload)
		if err != nil {
			log.Printf("Failed to marshal status payload: %v", err)
			return
		}

		statusResp, err := http.Post(statusURL, "application/json", strings.NewReader(string(payloadBytes)))
		if err != nil {
			log.Printf("Failed to send POST request for service status: %v", err)
			return
		}
		defer statusResp.Body.Close()
		log.Println("Service status sent successfully")
	} else {
		http.Error(w, "Service Check Failed", http.StatusServiceUnavailable)
	}
}
