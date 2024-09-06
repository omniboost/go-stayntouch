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
)

func TestMain(m *testing.M) {
	baseURLString := os.Getenv("BASE_URL")
	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")
	apiVersion := os.Getenv("API_VERSION")
	hotelID := os.Getenv("HOTEL_ID")
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
	if hotelID != "" {
		i, err := strconv.Atoi(hotelID)
		if err != nil {
			log.Fatal(err)
		}
		client.SetHotelID(i)
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
