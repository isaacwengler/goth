package tiktok_test

import (
	"testing"

	"github.com/isaacwengler/goth"
	"github.com/isaacwengler/goth/providers/tiktok"
	"github.com/stretchr/testify/assert"
)

func Test_Implements_Session(t *testing.T) {
	t.Parallel()
	a := assert.New(t)
	s := &tiktok.Session{}

	a.Implements((*goth.Session)(nil), s)
}

func Test_GetAuthURL(t *testing.T) {
	t.Parallel()
	a := assert.New(t)
	s := &tiktok.Session{}

	_, err := s.GetAuthURL()
	a.Error(err)

	s.AuthURL = "/foo"

	url, _ := s.GetAuthURL()
	a.Equal(url, "/foo")
}

func Test_ToJSON(t *testing.T) {
	t.Parallel()
	a := assert.New(t)
	s := &tiktok.Session{}

	data := s.Marshal()
	a.Equal(data, `{"AuthURL":"","AccessToken":"","ExpiresAt":"0001-01-01T00:00:00Z","OpenID":"","RefreshToken":"","RefreshExpiresAt":"0001-01-01T00:00:00Z"}`)
}

func Test_String(t *testing.T) {
	t.Parallel()
	a := assert.New(t)
	s := &tiktok.Session{}

	a.Equal(s.String(), s.Marshal())
}
