// Package gitea implements the OAuth2 protocol for authenticating users through gitea.
// This package can be used as a reference implementation of an OAuth2 provider for Goth.
package gitea

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/isaacwengler/goth"
	"golang.org/x/oauth2"
)

// These vars define the default Authentication, Token, and Profile URLS for Gitea.
//
// Examples:
//
//	gitea.AuthURL = "https://gitea.acme.com/oauth/authorize
//	gitea.TokenURL = "https://gitea.acme.com/oauth/token
//	gitea.ProfileURL = "https://gitea.acme.com/api/v3/user
var (
	AuthURL    = "https://gitea.com/login/oauth/authorize"
	TokenURL   = "https://gitea.com/login/oauth/access_token"
	ProfileURL = "https://gitea.com/api/v1/user"
)

// Provider is the implementation of `goth.Provider` for accessing Gitea.
type Provider struct {
	ClientKey    string
	Secret       string
	CallbackURL  string
	HTTPClient   *http.Client
	config       *oauth2.Config
	providerName string
	authURL      string
	tokenURL     string
	profileURL   string
}

// New creates a new Gitea provider and sets up important connection details.
// You should always call `gitea.New` to get a new provider.  Never try to
// create one manually.
func New(clientKey, secret, callbackURL string, scopes ...string) *Provider {
	return NewCustomisedURL(clientKey, secret, callbackURL, AuthURL, TokenURL, ProfileURL, scopes...)
}

// NewCustomisedURL is similar to New(...) but can be used to set custom URLs to connect to
func NewCustomisedURL(clientKey, secret, callbackURL, authURL, tokenURL, profileURL string, scopes ...string) *Provider {
	p := &Provider{
		ClientKey:    clientKey,
		Secret:       secret,
		CallbackURL:  callbackURL,
		providerName: "gitea",
		profileURL:   profileURL,
	}
	p.config = newConfig(p, authURL, tokenURL, scopes)
	return p
}

// Name is the name used to retrieve this provider later.
func (p *Provider) Name() string {
	return p.providerName
}

// SetName is to update the name of the provider (needed in case of multiple providers of 1 type)
func (p *Provider) SetName(name string) {
	p.providerName = name
}

func (p *Provider) Client() *http.Client {
	return goth.HTTPClientWithFallBack(p.HTTPClient)
}

// Debug is a no-op for the gitea package.
func (p *Provider) Debug(debug bool) {}

// BeginAuth asks Gitea for an authentication end-point.
func (p *Provider) BeginAuth(state string) (goth.Session, error) {
	return &Session{
		AuthURL: p.config.AuthCodeURL(state),
	}, nil
}

// FetchUser will go to Gitea and access basic information about the user.
func (p *Provider) FetchUser(session goth.Session) (goth.User, error) {
	sess := session.(*Session)
	user := goth.User{
		AccessToken:  sess.AccessToken,
		Provider:     p.Name(),
		RefreshToken: sess.RefreshToken,
		ExpiresAt:    sess.ExpiresAt,
	}

	if user.AccessToken == "" {
		// data is not yet retrieved since accessToken is still empty
		return user, fmt.Errorf("%s cannot get user information without accessToken", p.providerName)
	}

	response, err := p.Client().Get(p.profileURL + "?access_token=" + url.QueryEscape(sess.AccessToken))
	if err != nil {
		if response != nil {
			response.Body.Close()
		}
		return user, err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return user, fmt.Errorf("%s responded with a %d trying to fetch user information", p.providerName, response.StatusCode)
	}

	bits, err := io.ReadAll(response.Body)
	if err != nil {
		return user, err
	}

	err = json.NewDecoder(bytes.NewReader(bits)).Decode(&user.RawData)
	if err != nil {
		return user, err
	}

	err = userFromReader(bytes.NewReader(bits), &user)

	return user, err
}

func newConfig(provider *Provider, authURL, tokenURL string, scopes []string) *oauth2.Config {
	c := &oauth2.Config{
		ClientID:     provider.ClientKey,
		ClientSecret: provider.Secret,
		RedirectURL:  provider.CallbackURL,
		Endpoint: oauth2.Endpoint{
			AuthURL:  authURL,
			TokenURL: tokenURL,
		},
		Scopes: []string{},
	}

	if len(scopes) > 0 {
		for _, scope := range scopes {
			c.Scopes = append(c.Scopes, scope)
		}
	}
	return c
}

func userFromReader(r io.Reader, user *goth.User) error {
	u := struct {
		Name      string `json:"full_name"`
		Email     string `json:"email"`
		NickName  string `json:"login"`
		ID        int    `json:"id"`
		AvatarURL string `json:"avatar_url"`
	}{}
	err := json.NewDecoder(r).Decode(&u)
	if err != nil {
		return err
	}
	user.Email = u.Email
	user.Name = u.Name
	user.NickName = u.NickName
	user.UserID = strconv.Itoa(u.ID)
	user.AvatarURL = u.AvatarURL
	return nil
}

// RefreshTokenAvailable refresh token is provided by auth provider or not
func (p *Provider) RefreshTokenAvailable() bool {
	return true
}

// RefreshToken get new access token based on the refresh token
func (p *Provider) RefreshToken(refreshToken string) (*oauth2.Token, error) {
	token := &oauth2.Token{RefreshToken: refreshToken}
	ts := p.config.TokenSource(goth.ContextForClient(p.Client()), token)
	newToken, err := ts.Token()
	if err != nil {
		return nil, err
	}
	return newToken, err
}
