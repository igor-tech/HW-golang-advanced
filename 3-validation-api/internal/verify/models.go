package verify

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
	"verify/email/configs"
)

type VerificationToken struct {
	Email     string    `json:"email"`
	ExpiresAt time.Time `json:"expires_at"`
	Used      bool      `json:"used"`
	CreatedAt time.Time `json:"created_at"`
	Token     string    `json:"token"`
}

type TokenStorage struct {
	Tokens []VerificationToken `json:"tokens"`
}

type EmailService struct {
	config   *configs.Config
	storage  *TokenStorage
	filename string
}

func NewEmailService(config *configs.Config) *EmailService {
	filename := "tokens.json"
	storage := &TokenStorage{
		Tokens: []VerificationToken{},
	}

	if data, err := os.ReadFile(filename); err == nil {
		json.Unmarshal(data, &storage)
	}

	return &EmailService{
		config:   config,
		storage:  storage,
		filename: filename,
	}
}

func (s *EmailService) saveTokens() error {
	data, err := json.MarshalIndent(s.storage, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile(s.filename, data, 0644)
}

func (s *EmailService) addToken(email, token string) error {
	newToken := VerificationToken{
		Email:     email,
		Token:     token,
		ExpiresAt: time.Now().Add(time.Hour * 24),
		Used:      false,
		CreatedAt: time.Now(),
	}
	s.storage.Tokens = append(s.storage.Tokens, newToken)
	return s.saveTokens()
}

func (s *EmailService) findToken(token string) (*VerificationToken, error) {
	for i := range s.storage.Tokens {
		if s.storage.Tokens[i].Token == token {
			return &s.storage.Tokens[i], nil
		}
	}
	return nil, fmt.Errorf("token not found")
}

func (s *EmailService) markTokenUsed(token string) error {
	for i, t := range s.storage.Tokens {
		if t.Token == token {
			s.storage.Tokens[i].Used = true
			return s.saveTokens()
		}
	}
	return fmt.Errorf("token not found")
}

func (s *EmailService) removeToken(token string) error {
	for i, t := range s.storage.Tokens {
		if t.Token == token {
			s.storage.Tokens = append(s.storage.Tokens[:i], s.storage.Tokens[i+1:]...)
			return s.saveTokens()
		}
	}
	return fmt.Errorf("token not found")
}

