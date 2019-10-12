package mercadolibre

type authRequest struct {
	GrantType    string `json:"grant_type"  url:"grant_type"`
	ClientID     string `json:"client_id"   url:"client_id"`
	ClientSecret string `json:"client_secret"   url:"client_secret"`
	Code         string `json:"code,omitempty"  url:"code,omitempty"`
	RedirectURI  string `json:"redirect_uri,omitempty"  url:"redirect_uri,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty" url:"refresh_token,omitempty"`
}

type AuthResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	Scope        string `json:"scope"`
	UserID       int    `json:"user_id"`
	RefreshToken string `json:"refresh_token"`
	ReceivedAt   int64  `json:"received_at"`
}

type authError struct {
	Message string        `json:"message"`
	Error   string        `json:"error"`
	Status  int           `json:"status"`
	Cause   []interface{} `json:"cause"`
}

type requestParams struct {
	AccessToken string `json:"access_token,omitempty" url:"access_token,omitempty"`
}
