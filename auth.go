package mercadolibre

import (
	"errors"
	"fmt"
	"net/url"
	"time"
)

const (
	AuthURLMLA = "https://auth.mercadolibre.com.ar" // Argentina
	AuthURLMLB = "https://auth.mercadolivre.com.br" // Brasil
	AuthURLMCO = "https://auth.mercadolibre.com.co" // Colombia
	AuthURLMCR = "https://auth.mercadolibre.com.cr" // Costa Rica
	AuthURLMEC = "https://auth.mercadolibre.com.ec" // Ecuador
	AuthURLMLC = "https://auth.mercadolibre.cl"     // Chile
	AuthURLMLM = "https://auth.mercadolibre.com.mx" // Mexico
	AuthURLMLU = "https://auth.mercadolibre.com.uy" // Uruguay
	AuthURLMLV = "https://auth.mercadolibre.com.ve" // Venezuela
	AuthURLMPA = "https://auth.mercadolibre.com.pa" // Panama
	AuthURLMPE = "https://auth.mercadolibre.com.pe" // Peru
	AuthURLMPT = "https://auth.mercadolivre.pt"     // Portugal
	AuthURLMRD = "https://auth.mercadolibre.com.do" // Dominicana

	authorizationCode = "authorization_code"
	refreshToken      = "refresh_token"
)

func (client *Client) GetAuthURL(site, redirectUri string) (string, error) {
	if site != "" {
		u := url.URL{}
		query := u.Query()

		query.Add("response_type", "code")
		query.Add("client_id", client.clientId)
		query.Add("redirect_uri", redirectUri)

		authUrl := fmt.Sprintf("%s/authorization?%s", site, query.Encode())
		return authUrl, nil
	} else {
		return "", errors.New("site URL is required")
	}
}

func (client *Client) Authorize(code, redirectUri string) (*AuthResponse, error) {
	authReq := authRequest{
		GrantType:    authorizationCode,
		ClientID:     client.clientId,
		ClientSecret: client.clientSecret,
		Code:         code,
		RedirectURI:  redirectUri,
	}

	authRes := new(AuthResponse)
	authErr := new(authError)

	_, err := client.sling.New().Post("/oauth/token").QueryStruct(authReq).Receive(authRes, authErr)

	if err != nil {
		return nil, err
	}

	if authErr.Status != 0 {
		return nil, errors.New(fmt.Sprintf("%s - %s", authErr.Error, authErr.Message))
	}

	authRes.ReceivedAt = time.Now().Unix()
	client.Auth = *authRes

	return authRes, nil
}

func (client *Client) RefreshAccessToken() (*AuthResponse, error) {
	authReq := authRequest{
		GrantType:    refreshToken,
		ClientID:     client.clientId,
		ClientSecret: client.clientSecret,
		RefreshToken: client.Auth.RefreshToken,
	}

	authRes := new(AuthResponse)
	authErr := new(authError)

	_, err := client.sling.New().Post("/oauth/token").QueryStruct(authReq).Receive(authRes, authErr)

	if err != nil {
		return nil, err
	}

	if authErr.Status != 0 {
		return nil, errors.New(fmt.Sprintf("%s - %s", authErr.Error, authErr.Message))
	}

	authRes.ReceivedAt = time.Now().Unix()
	client.Auth = *authRes

	return authRes, nil
}

func (client *Client) IsExpired() bool {
	return client.Auth.IsExpired()
}

func (auth AuthResponse) IsExpired() bool {
	return (auth.ReceivedAt + int64(auth.ExpiresIn)) <= (time.Now().Unix() + 60)
}

func (auth AuthResponse) getParams() requestParams {
	params := requestParams{
		AccessToken: auth.AccessToken,
	}

	return params
}

func (client *Client) SetAuth(auth *AuthResponse) {
	client.Auth = *auth
}
