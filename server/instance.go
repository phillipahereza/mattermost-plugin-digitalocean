package main

import (
	"context"

	digitalocean "github.com/digitalocean/godo"
	"golang.org/x/oauth2"
)

// DOInstance represents a DO account. This allows connecting to different DO accounts.
type DOInstance interface {
	GetPlugin() *Plugin
	GetClient() *digitalocean.Client
	GetIntanceToken() string
}

// DOAccountInstance is a struct with Plugin and some information about the connected to
// DO Account.
type DOAccountInstance struct {
	*Plugin
	AccURL string
	Token  string
}

type TokenSource struct {
	AccessToken string
}

func (t *TokenSource) Token() (*oauth2.Token, error) {
	token := &oauth2.Token{
		AccessToken: t.AccessToken,
	}
	return token, nil
}

func NewDOAccount(p *Plugin, url string) *DOAccountInstance {
	return &DOAccountInstance{
		Plugin: p,
		AccURL: url,
	}
}

func (doac *DOAccountInstance) GetPlugin() *Plugin {
	return doac.Plugin
}

func (doac *DOAccountInstance) GetIntanceToken() string {
	return doac.Token
}

func (doac *DOAccountInstance) GetClient() *digitalocean.Client {
	tokenSource := &TokenSource{
		AccessToken: doac.GetIntanceToken(),
	}

	oauthClient := oauth2.NewClient(context.Background(), tokenSource)
	client := digitalocean.NewClient(oauthClient)
	return client
}
