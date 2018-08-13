package google

type User struct {
	Sub           string `json:"sub"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
	Email         string `json:"email"`
	EmailVerified string `json:"email_verified"`
	Locale        string `json:"locale"`
	Hd            string `json:"hd"`
}
