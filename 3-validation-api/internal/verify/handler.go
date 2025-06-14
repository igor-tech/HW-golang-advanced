package verify

import (
	"fmt"
	"github.com/jordan-wright/email"
	"net/http"
	"net/smtp"
	"strings"
	"verify/email/configs"
)

type VerifyHandler struct {
	*configs.Config
}

type VerifyHandlerDeps struct {
	*configs.Config
}

func NewVerifyHandler(router *http.ServeMux, deps VerifyHandlerDeps) {
	handler := &VerifyHandler{
		deps.Config,
	}
	router.HandleFunc("POST /send", handler.SendEmail())
	router.HandleFunc("GET /verify/", handler.Verify())
}

func (handler *VerifyHandler) SendEmail() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		e := email.NewEmail()
		e.From = "Jordan Wright <test@gmail.com>"
		e.To = []string{"igor20513@gmail.com"}
		e.Bcc = []string{"test_bcc@example.com"}
		e.Cc = []string{"test_cc@example.com"}
		e.Subject = "Awesome Subject"
		e.Text = []byte("Text Body is, of course, supported!")
		e.HTML = []byte("<h1>Fancy HTML is supported, too!</h1>")
		e.Send(handler.Config.Address, smtp.PlainAuth("", handler.Config.Email, handler.Config.Password, "smtp.gmail.com"))
	}
}

func (handler *VerifyHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prefix := "/verify/"
		if !strings.HasPrefix(r.URL.Path, prefix) {
			http.NotFound(w, r)
			return
		}

		hash := r.URL.Path[len(prefix):]
		if hash == "" {
			http.Error(w, "Missing hash", http.StatusBadRequest)
			return
		}

		fmt.Fprintf(w, "Hash received: %s\n", hash)
	}
}
