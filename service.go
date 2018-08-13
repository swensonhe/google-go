package google

type Service interface {
	GetTokenInfo(token string) (*User, error)
}
