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

func TestListClients(t *testing.T) {
	conf := loadTestConfig(t)
	api := NewApi(conf.AccountName, conf.AuthToken)

	clients, err := api.ListClients(Request{})
	assert.NoError(t, err)
	assert.True(t, len(*clients) > 0, "Client length should be greater than zero")

	clients, err = api.ListClients(Request{PerPage: 1})
	assert.NoError(t, err)
	assert.Len(t, *clients, 1, "Client length should be one")

}

func TestListInvoices(t *testing.T) {
	conf := loadTestConfig(t)
	api := NewApi(conf.AccountName, conf.AuthToken)

	now := time.Now()
	firstOfMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.UTC)

	invoices, paging, err := api.ListInvoices(Request{DateFrom: &Date{firstOfMonth}})
	assert.NoError(t, err)
	assert.True(t, len(*invoices) > 0, "Invoices length should be greater than zero")
	assert.Equal(t, paging.Page, 1)
}
