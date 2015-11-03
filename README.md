go-freshbooks
=====
This project implements a [Go](http://golang.org) client library for the [freshbooks API](http://developers.freshbooks.com/).

Originally forked from [toggl/go-freshbooks](https://github.com/toggl/go-freshbooks) and will hopefully merge together eventually but this has been quite the departure.

Supports token-based and OAuth authentication.

Example usage
---------------

```go
api := freshbooks.NewApi("<<AccountName>>", "<<AuthToken>>")

clients, pageInfo, err := api.ListClients(&freshbooks.Request{})
projects, pageInfo, err := api.ListProjects(&freshbooks.Request{})
invoices, pageInfo, err := api.ListInvoices(&freshbooks.Request{})
timeEntries, pageInfo,err := api.ListTimeEntries(&freshbooks.Request{})
```

OAuth authentication
---------------
The FreshBooks API also supports OAuth to authorize applications. [oauthplain](https://github.com/tambet/oauthplain) package is used to generate 'Authorization' headers.

```go
token := &oauthplain.Token{
  ConsumerKey:      "<<ConsumerKey>>",
  ConsumerSecret:   "<<ConsumerSecret>>",
  OAuthToken:       "<<OAuthToken>>",
  OAuthTokenSecret: "<<OAuthTokenSecret>>",
}

api := freshbooks.NewApi("<<AccountName>>", token)
```
