package stayntouch

import (
	"net/url"
	"strings"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

const (
	scope                = ""
	oauthStateString     = ""
	authorizationTimeout = 60 * time.Second
	tokenTimeout         = 5 * time.Second
)

type Oauth2Config struct {
	oauth2.Config
}

func NewOauth2Config() *Oauth2Config {
	config := &Oauth2Config{
		Config: oauth2.Config{
			RedirectURL:  "",
			ClientID:     "",
			ClientSecret: "",
			Scopes:       []string{scope},
			Endpoint: oauth2.Endpoint{
				AuthURL:   "https://auth.eu.stayntouch.com/oauth/authorize",
				TokenURL:  "https://auth.eu.stayntouch.com/oauth/token",
				AuthStyle: oauth2.AuthStyleAutoDetect,
			},
		},
	}

	config.SetBaseURL(&BaseURL)
	return config
}

func (c *Oauth2Config) SetBaseURL(baseURL *url.URL) {
	// Strip trailing slash
	baseURL.Path = strings.TrimSuffix(baseURL.Path, "/")

	c.Config.Endpoint = oauth2.Endpoint{
		AuthURL:  baseURL.String() + "/oauth/authorize",
		TokenURL: baseURL.String() + "/oauth/token",
	}
}

type Oauth2ClientCredentialsConfig struct {
	clientcredentials.Config
}

func NewOauth2ClientCredentialsConfig() *Oauth2ClientCredentialsConfig {
	config := &Oauth2ClientCredentialsConfig{
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

// const (
// 	scope = ""
// )

// type Oauth2Config struct {
// 	clientcredentials.Config
// }

// func NewOauth2Config() *Oauth2Config {
// 	config := &Oauth2Config{
// 		Config: clientcredentials.Config{
// 			ClientID:     "",
// 			ClientSecret: "",
// 			Scopes:       []string{scope},
// 			TokenURL:     "https://auth.eu.stayntouch.com/oauth/token",
// 			AuthStyle:    oauth2.AuthStyleInParams,
// 		},
// 	}

// 	return config
// }
