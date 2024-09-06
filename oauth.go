package stayntouch

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

const (
	scope = ""
)

type Oauth2Config struct {
	clientcredentials.Config
}

func NewOauth2Config() *Oauth2Config {
	config := &Oauth2Config{
		Config: clientcredentials.Config{
			ClientID:     "",
			ClientSecret: "",
			Scopes:       []string{scope},
			TokenURL:     "https://auth.eu.stayntouch.com/oauth/token",
			AuthStyle:    oauth2.AuthStyleInParams,
		},
	}

	return config
}
