package mercadolibre

import (
	"net/http"

	"github.com/dghubble/sling"
)

const (
	version = "0.0.1"
	apiUrl  = "https://api.mercadolibre.com"
)

type Client struct {
	sling        *sling.Sling
	Auth         AuthResponse
	clientId     string
	clientSecret string
}

func MercadoLibreClient(clientId string, clientSecret string) *Client {
	httpClient := http.DefaultClient

	base := sling.New().Client(httpClient).Base(apiUrl)
	return &Client{
		sling:        base,
		clientId:     clientId,
		clientSecret: clientSecret,
	}
}
