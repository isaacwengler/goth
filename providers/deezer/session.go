package deezer

import (
	"encoding/json"
	"errors"
	"strings"
	"time"

	"github.com/isaacwengler/goth"
)

// Session stores data during the auth process with Deezer.
type Session struct {
	AuthURL     string
	AccessToken string
	ExpiresAt   time.Time
}

// GetAuthURL will return the URL set by calling the `BeginAuth` function on the Deezer provider.
func (s Session) GetAuthURL() (string, error) {
	if s.AuthURL == "" {
		return "", errors.New(goth.NoAuthUrlErrorMessage)
	}
	return s.AuthURL, nil
}

// Authorize the session with Deezer and return the access token to be stored for future use.
func (s *Session) Authorize(provider goth.Provider, params goth.Params) (string, error) {
	p := provider.(*Provider)

	ctx := goth.ContextForClient(p.Client())
	token, err := p.config.Exchange(ctx, params.Get("code"))
	if err != nil {
		return "", err
	}

	if !token.Valid() {
		return "", errors.New("Invalid token received from provider")
	}

	expires, ok := token.Extra("expires").(float64)
	if ok != true {
		return "", errors.New("Invalid token received from provider")
	}

	s.AccessToken = token.AccessToken
	s.ExpiresAt = time.Now().Add(time.Second * time.Duration(expires))
	return token.AccessToken, err
}

// Marshal the session into a string
func (s Session) Marshal() string {
	b, _ := json.Marshal(s)
	return string(b)
}

func (s Session) String() string {
	return s.Marshal()
}

// UnmarshalSession will unmarshal a JSON string into a session.
func (p *Provider) UnmarshalSession(data string) (goth.Session, error) {
	sess := &Session{}
	err := json.NewDecoder(strings.NewReader(data)).Decode(sess)
	return sess, err
}
