package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"sort"

	"github.com/gorilla/pat"
	"github.com/isaacwengler/goth"
	"github.com/isaacwengler/goth/gothic"
	"github.com/isaacwengler/goth/providers/amazon"
	"github.com/isaacwengler/goth/providers/apple"
	"github.com/isaacwengler/goth/providers/auth0"
	"github.com/isaacwengler/goth/providers/azuread"
	"github.com/isaacwengler/goth/providers/battlenet"
	"github.com/isaacwengler/goth/providers/bitbucket"
	"github.com/isaacwengler/goth/providers/box"
	"github.com/isaacwengler/goth/providers/dailymotion"
	"github.com/isaacwengler/goth/providers/deezer"
	"github.com/isaacwengler/goth/providers/digitalocean"
	"github.com/isaacwengler/goth/providers/discord"
	"github.com/isaacwengler/goth/providers/dropbox"
	"github.com/isaacwengler/goth/providers/eveonline"
	"github.com/isaacwengler/goth/providers/facebook"
	"github.com/isaacwengler/goth/providers/fitbit"
	"github.com/isaacwengler/goth/providers/gitea"
	"github.com/isaacwengler/goth/providers/github"
	"github.com/isaacwengler/goth/providers/gitlab"
	"github.com/isaacwengler/goth/providers/google"
	"github.com/isaacwengler/goth/providers/gplus"
	"github.com/isaacwengler/goth/providers/heroku"
	"github.com/isaacwengler/goth/providers/instagram"
	"github.com/isaacwengler/goth/providers/intercom"
	"github.com/isaacwengler/goth/providers/kakao"
	"github.com/isaacwengler/goth/providers/lastfm"
	"github.com/isaacwengler/goth/providers/line"
	"github.com/isaacwengler/goth/providers/linkedin"
	"github.com/isaacwengler/goth/providers/mastodon"
	"github.com/isaacwengler/goth/providers/meetup"
	"github.com/isaacwengler/goth/providers/microsoftonline"
	"github.com/isaacwengler/goth/providers/naver"
	"github.com/isaacwengler/goth/providers/nextcloud"
	"github.com/isaacwengler/goth/providers/okta"
	"github.com/isaacwengler/goth/providers/onedrive"
	"github.com/isaacwengler/goth/providers/openidConnect"
	"github.com/isaacwengler/goth/providers/patreon"
	"github.com/isaacwengler/goth/providers/paypal"
	"github.com/isaacwengler/goth/providers/quickbooks"
	"github.com/isaacwengler/goth/providers/salesforce"
	"github.com/isaacwengler/goth/providers/seatalk"
	"github.com/isaacwengler/goth/providers/shopify"
	"github.com/isaacwengler/goth/providers/slack"
	"github.com/isaacwengler/goth/providers/soundcloud"
	"github.com/isaacwengler/goth/providers/spotify"
	"github.com/isaacwengler/goth/providers/steam"
	"github.com/isaacwengler/goth/providers/strava"
	"github.com/isaacwengler/goth/providers/stripe"
	"github.com/isaacwengler/goth/providers/tiktok"
	"github.com/isaacwengler/goth/providers/twitch"
	"github.com/isaacwengler/goth/providers/twitter"
	"github.com/isaacwengler/goth/providers/twitterv2"
	"github.com/isaacwengler/goth/providers/typetalk"
	"github.com/isaacwengler/goth/providers/uber"
	"github.com/isaacwengler/goth/providers/vk"
	"github.com/isaacwengler/goth/providers/wecom"
	"github.com/isaacwengler/goth/providers/wepay"
	"github.com/isaacwengler/goth/providers/xero"
	"github.com/isaacwengler/goth/providers/yahoo"
	"github.com/isaacwengler/goth/providers/yammer"
	"github.com/isaacwengler/goth/providers/yandex"
	"github.com/isaacwengler/goth/providers/zoom"
)

func main() {
	goth.UseProviders(
		// Use twitterv2 instead of twitter if you only have access to the Essential API Level
		// the twitter provider uses a v1.1 API that is not available to the Essential Level
		twitterv2.New(os.Getenv("TWITTER_KEY"), os.Getenv("TWITTER_SECRET"), "http://localhost:3000/auth/twitterv2/callback"),
		// If you'd like to use authenticate instead of authorize in TwitterV2 provider, use this instead.
		// twitterv2.NewAuthenticate(os.Getenv("TWITTER_KEY"), os.Getenv("TWITTER_SECRET"), "http://localhost:3000/auth/twitterv2/callback"),

		twitter.New(os.Getenv("TWITTER_KEY"), os.Getenv("TWITTER_SECRET"), "http://localhost:3000/auth/twitter/callback"),
		// If you'd like to use authenticate instead of authorize in Twitter provider, use this instead.
		// twitter.NewAuthenticate(os.Getenv("TWITTER_KEY"), os.Getenv("TWITTER_SECRET"), "http://localhost:3000/auth/twitter/callback"),

		tiktok.New(os.Getenv("TIKTOK_KEY"), os.Getenv("TIKTOK_SECRET"), "http://localhost:3000/auth/tiktok/callback"),
		facebook.New(os.Getenv("FACEBOOK_KEY"), os.Getenv("FACEBOOK_SECRET"), "http://localhost:3000/auth/facebook/callback"),
		fitbit.New(os.Getenv("FITBIT_KEY"), os.Getenv("FITBIT_SECRET"), "http://localhost:3000/auth/fitbit/callback"),
		google.New(os.Getenv("GOOGLE_KEY"), os.Getenv("GOOGLE_SECRET"), "http://localhost:3000/auth/google/callback"),
		gplus.New(os.Getenv("GPLUS_KEY"), os.Getenv("GPLUS_SECRET"), "http://localhost:3000/auth/gplus/callback"),
		github.New(os.Getenv("GITHUB_KEY"), os.Getenv("GITHUB_SECRET"), "http://localhost:3000/auth/github/callback"),
		spotify.New(os.Getenv("SPOTIFY_KEY"), os.Getenv("SPOTIFY_SECRET"), "http://localhost:3000/auth/spotify/callback"),
		linkedin.New(os.Getenv("LINKEDIN_KEY"), os.Getenv("LINKEDIN_SECRET"), "http://localhost:3000/auth/linkedin/callback"),
		line.New(os.Getenv("LINE_KEY"), os.Getenv("LINE_SECRET"), "http://localhost:3000/auth/line/callback", "profile", "openid", "email"),
		lastfm.New(os.Getenv("LASTFM_KEY"), os.Getenv("LASTFM_SECRET"), "http://localhost:3000/auth/lastfm/callback"),
		twitch.New(os.Getenv("TWITCH_KEY"), os.Getenv("TWITCH_SECRET"), "http://localhost:3000/auth/twitch/callback"),
		dropbox.New(os.Getenv("DROPBOX_KEY"), os.Getenv("DROPBOX_SECRET"), "http://localhost:3000/auth/dropbox/callback"),
		digitalocean.New(os.Getenv("DIGITALOCEAN_KEY"), os.Getenv("DIGITALOCEAN_SECRET"), "http://localhost:3000/auth/digitalocean/callback", "read"),
		bitbucket.New(os.Getenv("BITBUCKET_KEY"), os.Getenv("BITBUCKET_SECRET"), "http://localhost:3000/auth/bitbucket/callback"),
		instagram.New(os.Getenv("INSTAGRAM_KEY"), os.Getenv("INSTAGRAM_SECRET"), "http://localhost:3000/auth/instagram/callback"),
		intercom.New(os.Getenv("INTERCOM_KEY"), os.Getenv("INTERCOM_SECRET"), "http://localhost:3000/auth/intercom/callback"),
		box.New(os.Getenv("BOX_KEY"), os.Getenv("BOX_SECRET"), "http://localhost:3000/auth/box/callback"),
		salesforce.New(os.Getenv("SALESFORCE_KEY"), os.Getenv("SALESFORCE_SECRET"), "http://localhost:3000/auth/salesforce/callback"),
		seatalk.New(os.Getenv("SEATALK_KEY"), os.Getenv("SEATALK_SECRET"), "http://localhost:3000/auth/seatalk/callback"),
		amazon.New(os.Getenv("AMAZON_KEY"), os.Getenv("AMAZON_SECRET"), "http://localhost:3000/auth/amazon/callback"),
		yammer.New(os.Getenv("YAMMER_KEY"), os.Getenv("YAMMER_SECRET"), "http://localhost:3000/auth/yammer/callback"),
		onedrive.New(os.Getenv("ONEDRIVE_KEY"), os.Getenv("ONEDRIVE_SECRET"), "http://localhost:3000/auth/onedrive/callback"),
		azuread.New(os.Getenv("AZUREAD_KEY"), os.Getenv("AZUREAD_SECRET"), "http://localhost:3000/auth/azuread/callback", nil),
		microsoftonline.New(os.Getenv("MICROSOFTONLINE_KEY"), os.Getenv("MICROSOFTONLINE_SECRET"), "http://localhost:3000/auth/microsoftonline/callback"),
		battlenet.New(os.Getenv("BATTLENET_KEY"), os.Getenv("BATTLENET_SECRET"), "http://localhost:3000/auth/battlenet/callback"),
		eveonline.New(os.Getenv("EVEONLINE_KEY"), os.Getenv("EVEONLINE_SECRET"), "http://localhost:3000/auth/eveonline/callback"),
		kakao.New(os.Getenv("KAKAO_KEY"), os.Getenv("KAKAO_SECRET"), "http://localhost:3000/auth/kakao/callback"),

		// Pointed https://localhost.com to http://localhost:3000/auth/yahoo/callback
		// Yahoo only accepts urls that starts with https
		yahoo.New(os.Getenv("YAHOO_KEY"), os.Getenv("YAHOO_SECRET"), "https://localhost.com"),
		typetalk.New(os.Getenv("TYPETALK_KEY"), os.Getenv("TYPETALK_SECRET"), "http://localhost:3000/auth/typetalk/callback", "my"),
		slack.New(os.Getenv("SLACK_KEY"), os.Getenv("SLACK_SECRET"), "http://localhost:3000/auth/slack/callback"),
		stripe.New(os.Getenv("STRIPE_KEY"), os.Getenv("STRIPE_SECRET"), "http://localhost:3000/auth/stripe/callback"),
		wepay.New(os.Getenv("WEPAY_KEY"), os.Getenv("WEPAY_SECRET"), "http://localhost:3000/auth/wepay/callback", "view_user"),
		// By default paypal production auth urls will be used, please set PAYPAL_ENV=sandbox as environment variable for testing
		// in sandbox environment
		paypal.New(os.Getenv("PAYPAL_KEY"), os.Getenv("PAYPAL_SECRET"), "http://localhost:3000/auth/paypal/callback"),
		steam.New(os.Getenv("STEAM_KEY"), "http://localhost:3000/auth/steam/callback"),
		heroku.New(os.Getenv("HEROKU_KEY"), os.Getenv("HEROKU_SECRET"), "http://localhost:3000/auth/heroku/callback"),
		uber.New(os.Getenv("UBER_KEY"), os.Getenv("UBER_SECRET"), "http://localhost:3000/auth/uber/callback"),
		soundcloud.New(os.Getenv("SOUNDCLOUD_KEY"), os.Getenv("SOUNDCLOUD_SECRET"), "http://localhost:3000/auth/soundcloud/callback"),
		gitlab.New(os.Getenv("GITLAB_KEY"), os.Getenv("GITLAB_SECRET"), "http://localhost:3000/auth/gitlab/callback"),
		dailymotion.New(os.Getenv("DAILYMOTION_KEY"), os.Getenv("DAILYMOTION_SECRET"), "http://localhost:3000/auth/dailymotion/callback", "email"),
		deezer.New(os.Getenv("DEEZER_KEY"), os.Getenv("DEEZER_SECRET"), "http://localhost:3000/auth/deezer/callback", "email"),
		discord.New(os.Getenv("DISCORD_KEY"), os.Getenv("DISCORD_SECRET"), "http://localhost:3000/auth/discord/callback", discord.ScopeIdentify, discord.ScopeEmail),
		meetup.New(os.Getenv("MEETUP_KEY"), os.Getenv("MEETUP_SECRET"), "http://localhost:3000/auth/meetup/callback"),

		// Auth0 allocates domain per customer, a domain must be provided for auth0 to work
		auth0.New(os.Getenv("AUTH0_KEY"), os.Getenv("AUTH0_SECRET"), "http://localhost:3000/auth/auth0/callback", os.Getenv("AUTH0_DOMAIN")),
		xero.New(os.Getenv("XERO_KEY"), os.Getenv("XERO_SECRET"), "http://localhost:3000/auth/xero/callback"),
		vk.New(os.Getenv("VK_KEY"), os.Getenv("VK_SECRET"), "http://localhost:3000/auth/vk/callback"),
		naver.New(os.Getenv("NAVER_KEY"), os.Getenv("NAVER_SECRET"), "http://localhost:3000/auth/naver/callback"),
		yandex.New(os.Getenv("YANDEX_KEY"), os.Getenv("YANDEX_SECRET"), "http://localhost:3000/auth/yandex/callback"),
		nextcloud.NewCustomisedDNS(os.Getenv("NEXTCLOUD_KEY"), os.Getenv("NEXTCLOUD_SECRET"), "http://localhost:3000/auth/nextcloud/callback", os.Getenv("NEXTCLOUD_URL")),
		gitea.New(os.Getenv("GITEA_KEY"), os.Getenv("GITEA_SECRET"), "http://localhost:3000/auth/gitea/callback"),
		shopify.New(os.Getenv("SHOPIFY_KEY"), os.Getenv("SHOPIFY_SECRET"), "http://localhost:3000/auth/shopify/callback", shopify.ScopeReadCustomers, shopify.ScopeReadOrders),
		apple.New(os.Getenv("APPLE_KEY"), os.Getenv("APPLE_SECRET"), "http://localhost:3000/auth/apple/callback", nil, apple.ScopeName, apple.ScopeEmail),
		strava.New(os.Getenv("STRAVA_KEY"), os.Getenv("STRAVA_SECRET"), "http://localhost:3000/auth/strava/callback"),
		okta.New(os.Getenv("OKTA_ID"), os.Getenv("OKTA_SECRET"), os.Getenv("OKTA_ORG_URL"), "http://localhost:3000/auth/okta/callback", "openid", "profile", "email"),
		mastodon.New(os.Getenv("MASTODON_KEY"), os.Getenv("MASTODON_SECRET"), "http://localhost:3000/auth/mastodon/callback", "read:accounts"),
		wecom.New(os.Getenv("WECOM_CORP_ID"), os.Getenv("WECOM_SECRET"), os.Getenv("WECOM_AGENT_ID"), "http://localhost:3000/auth/wecom/callback"),
		zoom.New(os.Getenv("ZOOM_KEY"), os.Getenv("ZOOM_SECRET"), "http://localhost:3000/auth/zoom/callback", "read:user"),
		patreon.New(os.Getenv("PATREON_KEY"), os.Getenv("PATREON_SECRET"), "http://localhost:3000/auth/patreon/callback"),
		quickbooks.New(os.Getenv("QUICKBOOKS_KEY"), os.Getenv("QUICKBOOKS_SECRET"), "http://localhost:3000/auth/quickbooks/callback", false, quickbooks.ScopeAccounting, quickbooks.ScopePayments),
	)

	// OpenID Connect is based on OpenID Connect Auto Discovery URL (https://openid.net/specs/openid-connect-discovery-1_0-17.html)
	// because the OpenID Connect provider initialize itself in the New(), it can return an error which should be handled or ignored
	// ignore the error for now
	openidConnect, _ := openidConnect.New(os.Getenv("OPENID_CONNECT_KEY"), os.Getenv("OPENID_CONNECT_SECRET"), "http://localhost:3000/auth/openid-connect/callback", os.Getenv("OPENID_CONNECT_DISCOVERY_URL"))
	if openidConnect != nil {
		goth.UseProviders(openidConnect)
	}

	m := map[string]string{
		"amazon":          "Amazon",
		"apple":           "Apple",
		"auth0":           "Auth0",
		"azuread":         "Azure AD",
		"battlenet":       "Battle.net",
		"bitbucket":       "Bitbucket",
		"box":             "Box",
		"dailymotion":     "Dailymotion",
		"deezer":          "Deezer",
		"digitalocean":    "Digital Ocean",
		"discord":         "Discord",
		"dropbox":         "Dropbox",
		"eveonline":       "Eve Online",
		"facebook":        "Facebook",
		"fitbit":          "Fitbit",
		"gitea":           "Gitea",
		"github":          "Github",
		"gitlab":          "Gitlab",
		"google":          "Google",
		"gplus":           "Google Plus",
		"heroku":          "Heroku",
		"instagram":       "Instagram",
		"intercom":        "Intercom",
		"kakao":           "Kakao",
		"lastfm":          "Last FM",
		"line":            "LINE",
		"linkedin":        "LinkedIn",
		"mastodon":        "Mastodon",
		"meetup":          "Meetup.com",
		"microsoftonline": "Microsoft Online",
		"naver":           "Naver",
		"nextcloud":       "NextCloud",
		"okta":            "Okta",
		"onedrive":        "Onedrive",
		"openid-connect":  "OpenID Connect",
		"patreon":         "Patreon",
		"paypal":          "Paypal",
		"quickbooks":      "Quickbooks",
		"salesforce":      "Salesforce",
		"seatalk":         "SeaTalk",
		"shopify":         "Shopify",
		"slack":           "Slack",
		"soundcloud":      "SoundCloud",
		"spotify":         "Spotify",
		"steam":           "Steam",
		"strava":          "Strava",
		"stripe":          "Stripe",
		"tiktok":          "TikTok",
		"twitch":          "Twitch",
		"twitter":         "Twitter",
		"twitterv2":       "Twitter",
		"typetalk":        "Typetalk",
		"uber":            "Uber",
		"vk":              "VK",
		"wecom":           "WeCom",
		"wepay":           "Wepay",
		"xero":            "Xero",
		"yahoo":           "Yahoo",
		"yammer":          "Yammer",
		"yandex":          "Yandex",
		"zoom":            "Zoom",
	}
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	providerIndex := &ProviderIndex{Providers: keys, ProvidersMap: m}

	p := pat.New()
	p.Get("/auth/{provider}/callback", func(res http.ResponseWriter, req *http.Request) {

		user, err := gothic.CompleteUserAuth(res, req)
		if err != nil {
			fmt.Fprintln(res, err)
			return
		}
		t, _ := template.New("foo").Parse(userTemplate)
		t.Execute(res, user)
	})

	p.Get("/logout/{provider}", func(res http.ResponseWriter, req *http.Request) {
		gothic.Logout(res, req)
		res.Header().Set("Location", "/")
		res.WriteHeader(http.StatusTemporaryRedirect)
	})

	p.Get("/auth/{provider}", func(res http.ResponseWriter, req *http.Request) {
		// try to get the user without re-authenticating
		if gothUser, err := gothic.CompleteUserAuth(res, req); err == nil {
			t, _ := template.New("foo").Parse(userTemplate)
			t.Execute(res, gothUser)
		} else {
			gothic.BeginAuthHandler(res, req)
		}
	})

	p.Get("/", func(res http.ResponseWriter, req *http.Request) {
		t, _ := template.New("foo").Parse(indexTemplate)
		t.Execute(res, providerIndex)
	})

	log.Println("listening on localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", p))
}

type ProviderIndex struct {
	Providers    []string
	ProvidersMap map[string]string
}

var indexTemplate = `{{range $key,$value:=.Providers}}
    <p><a href="/auth/{{$value}}">Log in with {{index $.ProvidersMap $value}}</a></p>
{{end}}`

var userTemplate = `
<p><a href="/logout/{{.Provider}}">logout</a></p>
<p>Name: {{.Name}} [{{.LastName}}, {{.FirstName}}]</p>
<p>Email: {{.Email}}</p>
<p>NickName: {{.NickName}}</p>
<p>Location: {{.Location}}</p>
<p>AvatarURL: {{.AvatarURL}} <img src="{{.AvatarURL}}"></p>
<p>Description: {{.Description}}</p>
<p>UserID: {{.UserID}}</p>
<p>AccessToken: {{.AccessToken}}</p>
<p>ExpiresAt: {{.ExpiresAt}}</p>
<p>RefreshToken: {{.RefreshToken}}</p>
`
