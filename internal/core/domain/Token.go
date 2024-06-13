package domain

type Token struct {
	AccessToken string
	TokenType   string
	ExpiresAt   int64
}
