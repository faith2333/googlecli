package account

import "golang.org/x/oauth2"

const (
	ClientID      string = ""
	CallbackUrl   string = "/google/callback"
	ListenPort    int    = 9096
	GoogleAuthURL string = "https://accounts.google.com/o/oauth2/auth/oauthchooseaccount"
)

var defaultOauth2Config = &oauth2.Config{
	ClientID: ClientID,
	Scopes: []string{
		"https://www.googleapis.com/auth/userinfo.email",
		"https://www.googleapis.com/auth/cloud-platform",
		"https://www.googleapis.com/auth/appengine.admin",
		"https://www.googleapis.com/auth/sqlservice.login",
		"https://www.googleapis.com/auth/compute",
		"https://www.googleapis.com/auth/accounts.reauth",
	},
	RedirectURL: "http://localhost:9096" + CallbackUrl,
}

type AuthInfo struct {
	AuthUser []string `json:"authuser"`
	Code     []string `json:"code"`
	HD       []string `json:"hd"`
	Scope    []string `json:"scope"`
	State    []string `json:"state"`
}
