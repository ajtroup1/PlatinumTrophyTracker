package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/smtp"
	"regexp"

	"github.com/ajtroup1/platinum-trophy-tracker/config"
	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

func init() {
	Validate = validator.New()
	Validate.RegisterValidation("password", validatePassword)
}

// validatePassword checks password complexity requirements
func validatePassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	hasMinLen := len(password) >= 3
	hasMaxLen := len(password) <= 130
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(password)
	hasSpecial := regexp.MustCompile(`[!@#~$%^&*(),.?":{}|<>]`).MatchString(password)

	return hasMinLen && hasMaxLen && hasUpper && hasLower && hasNumber && hasSpecial
}

func ParseJSON(r *http.Request, payload any) error {
	if r.Body == nil {
		return fmt.Errorf("missing request body")
	}

	return json.NewDecoder(r.Body).Decode(payload)
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJSON(w, status, map[string]string{"error": err.Error()})
}

func SendEmail(to string, subject string, body string) error {
	from := config.Envs.Email
	password := config.Envs.EmailPassword

	// SMTP server configuration
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Message
	message := []byte("Subject: " + subject + "\r\n" +
		"\r\n" +
		body + "\r\n")

	// Authentication
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, message)
	if err != nil {
		return err
	}

	return nil
}
