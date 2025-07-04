package verify

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/smtp"
	"time"
	"verify/email/configs"

	"github.com/jordan-wright/email"
)

type VerifyHandler struct {
	*configs.Config
	emailService *EmailService
}

type VerifyHandlerDeps struct {
	*configs.Config
}

type SendEmailRequest struct {
	Email string `json:"email"`
}

func NewVerifyHandler(router *http.ServeMux, deps VerifyHandlerDeps) {
	handler := &VerifyHandler{
		Config:       deps.Config,
		emailService: NewEmailService(deps.Config),
	}

	router.HandleFunc("GET /verify/{token}", handler.Verify())
}

func (handler *VerifyHandler) SendEmail() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req SendEmailRequest

		// Парсинг JSON
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		// Валидация email
		if !isValidEmail(req.Email) {
			http.Error(w, "Invalid email address", http.StatusBadRequest)
			return
		}

		// Генерация токена
		token, err := generateToken()
		if err != nil {
			http.Error(w, "Failed to generate token", http.StatusInternalServerError)
			return
		}

		// Сохранение токена
		if err = handler.emailService.addToken(req.Email, token); err != nil {
			http.Error(w, "Ошибка сохранения токена", http.StatusInternalServerError)
			return
		}

		// Формирование ссылки
		verificationURL := fmt.Sprintf("%s/verify/%s", handler.Config.BaseUrl, token)

		// Создание содержимого письма
		emailData := EmailData{
			VerificationURL: verificationURL,
			Email:           req.Email,
		}

		content, err := renderEmailTemplate(emailData)

		if err != nil {
			http.Error(w, "Failed to generate email content", http.StatusInternalServerError)
			return
		}

		// Отправка письма
		e := email.NewEmail()
		e.From = fmt.Sprintf("Igor Shargin <%s>", handler.Config.SMTPEmail)
		e.To = []string{req.Email}
		e.Subject = "Подтверждение email | GO app Verify Email"
		e.HTML = []byte(content)

		// SMTP авторизация
		auth := smtp.PlainAuth("", handler.Config.SMTPEmail, handler.Config.SMTPPassword, handler.Config.SMTPHost)

		smtpAddr := fmt.Sprintf("%s:%s", handler.Config.SMTPHost, handler.Config.SMTPPort)

		if err := e.Send(smtpAddr, auth); err != nil {
			http.Error(w, fmt.Sprintf("Ошибка отправки email: %v", err), http.StatusInternalServerError)
			return
		}

		response := map[string]string{
			"message": "Email sent successfully",
			"email":   req.Email,
		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(response)
	}
}

func (handler *VerifyHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.PathValue("token")

		if token == "" {
			errorPage, _ := renderErrorPage("Missing token")
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(errorPage))
			return
		}

		// Поиск токена
		tokenData, err := handler.emailService.findToken(token)
		if err != nil {
			errorPage, _ := renderErrorPage("Invalid token")
			w.Header().Set("Content-Type", "text/html")
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(errorPage))
			return
		}

		// Проверка срока действия токена
		if time.Now().After(tokenData.ExpiresAt) {
			errorPage, _ := renderErrorPage("Token expired")
			w.Header().Set("Content-Type", "text/html")
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(errorPage))
			return
		}

		// Проверка использования токена
		if tokenData.Used {
			errorPage, _ := renderErrorPage("Token already used")
			w.Header().Set("Content-Type", "text/html")
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(errorPage))
			return
		}

		if err := handler.emailService.markTokenUsed(token); err != nil {
			errorPage, _ := renderErrorPage("Ошибка подтверждения")
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(errorPage))
			return
		}

		// Успешная страница
		successPage, _ := renderSuccessPage(tokenData.Email)
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(successPage))

	}
}
