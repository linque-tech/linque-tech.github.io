package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/smtp"
	"os"
	"strings"
)

type contactRequest struct {
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Subject string `json:"subject"`
	Message string `json:"message"`
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/contact", handleContact)

	addr := env("LISTEN_ADDR", "127.0.0.1:8081")
	log.Printf("contact-handler listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}

func handleContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", env("ALLOWED_ORIGIN", "*"))
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req contactRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	req.Name = strings.TrimSpace(req.Name)
	req.Email = strings.TrimSpace(req.Email)
	req.Subject = strings.TrimSpace(req.Subject)
	req.Message = strings.TrimSpace(req.Message)

	if req.Name == "" || req.Email == "" || req.Subject == "" || req.Message == "" {
		http.Error(w, "missing required fields", http.StatusUnprocessableEntity)
		return
	}

	if err := sendMail(req); err != nil {
		log.Printf("sendMail error: %v", err)
		http.Error(w, "mail delivery failed", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func sendMail(req contactRequest) error {
	host := env("SMTP_HOST", "")
	port := env("SMTP_PORT", "587")
	user := env("SMTP_USER", "")
	pass := env("SMTP_PASS", "")
	from := env("MAIL_FROM", user)
	to := env("MAIL_TO", "contact@linque.eu")

	auth := smtp.PlainAuth("", user, pass, host)

	body := fmt.Sprintf(
		"Name:    %s\nCompany: %s\nEmail:   %s\n\n%s",
		req.Name, req.Company, req.Email, req.Message,
	)
	msg := fmt.Sprintf(
		"From: %s\r\nTo: %s\r\nReply-To: %s\r\nSubject: [Contact Form] %s\r\nMIME-Version: 1.0\r\nContent-Type: text/plain; charset=UTF-8\r\n\r\n%s",
		from, to, req.Email, req.Subject, body,
	)

	return smtp.SendMail(net.JoinHostPort(host, port), auth, from, []string{to}, []byte(msg))
}

func env(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
