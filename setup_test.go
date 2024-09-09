package stayntouch_test

import (
	"context"
	"log"
	"net/url"
	"os"
	"strconv"
	"testing"

	stayntouch "github.com/omniboost/go-stayntouch"
)

var (
	client *stayntouch.Client
	hotelID int
)

func TestMain(m *testing.M) {
	var err error
	baseURLString := os.Getenv("BASE_URL")
	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")
	apiVersion := os.Getenv("API_VERSION")
	tokenURL := os.Getenv("TOKEN_URL")
	debug := os.Getenv("DEBUG")

	oauthConfig := stayntouch.NewOauth2Config()
	oauthConfig.ClientID = clientID
	oauthConfig.ClientSecret = clientSecret

	// set alternative token url
	if tokenURL != "" {
		oauthConfig.TokenURL = tokenURL
	}

	// get http client with automatic oauth logic
	httpClient := oauthConfig.Client(context.Background())

	client = stayntouch.NewClient(httpClient)
	if apiVersion != "" {
		client.SetAPIVersion(apiVersion)
	}
	if t := os.Getenv("HOTEL_ID"); t != "" {
		hotelID, err = strconv.Atoi(t)
		if err != nil {
			log.Fatal(err)
		}
	}
	if debug != "" {
		client.SetDebug(true)
	}

	if baseURLString != "" {
		baseURL, err := url.Parse(baseURLString)
		if err != nil {
			log.Fatal(err)
		}
		client.SetBaseURL(*baseURL)
	}
	m.Run()
}
