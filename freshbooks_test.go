package freshbooks

import (
	"encoding/json"
	"io/ioutil"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type authConfig struct {
	AccountName      string
	AuthToken        string // Token-Based authentication (deprecated)
	ConsumerKey      string // OAuth authentication
	ConsumerSecret   string // OAuth authentication
	OAuthToken       string // OAuth authentication
	OAuthTokenSecret string // OAuth authentication
}

func loadTestConfig(t *testing.T) *authConfig {
	file, e := ioutil.ReadFile("test-config.json")
	if e != nil {
		t.Fatal("Unable to load 'test-config.json'")
	}
	var config authConfig
	if err := json.Unmarshal(file, &config); err != nil {
		t.Fatal("Unable to unmarshal 'test-config.json'")
	}
	return &config
}

// func TestGetUsers(t *testing.T) {
// 	conf := loadTestConfig(t)
// 	api := NewApi(conf.AccountName, conf.AuthToken)
// 	users, err := api.Users()
// 	if err != nil {
// 		t.Fatal("Freshbooks retured an error:", err.Error())
// 	}
// 	if len(users) < 1 {
// 		t.Fatal("There should be at least one user")
// 	}
// }

func TestListUsers(t *testing.T) {
	conf := loadTestConfig(t)
	api := NewApi(conf.AccountName, conf.AuthToken)

	clients, err := api.ListClients(Request{})
	assert.NoError(t, err)
	assert.True(t, len(*clients) > 0, "Client length should be greater than zero")

	clients, err = api.ListClients(Request{PerPage: 1})
	assert.NoError(t, err)
	assert.Len(t, *clients, 1, "Client length should be one")

}

func TestListTimeEntries(t *testing.T) {
	conf := loadTestConfig(t)
	api := NewApi(conf.AccountName, conf.AuthToken)

	timeEntries, paging, err := api.ListTimeEntries(Request{DateFrom: &Date{time.Now()}})
	assert.NoError(t, err)
	assert.True(t, len(*timeEntries) > 0, "Time Entries length should be greater than zero")
	assert.Equal(t, paging.Page, 1)
}

// func TestOAuth(t *testing.T) {
// 	conf := loadTestConfig(t)
// 	token := &oauthplain.Token{
// 		ConsumerKey:      conf.ConsumerKey,
// 		ConsumerSecret:   conf.ConsumerSecret,
// 		OAuthToken:       conf.OAuthToken,
// 		OAuthTokenSecret: conf.OAuthTokenSecret,
// 	}
// 	api := NewApi(conf.AccountName, token)
// 	clients, err := api.ListClients(ClientListRequest{})
// 	if err != nil {
// 		t.Fatal("Freshbooks retured an error:", err.Error())
// 	}
// 	if len(*clients) < 1 {
// 		t.Fatal("There should be at least one client")
// 	}
// }
