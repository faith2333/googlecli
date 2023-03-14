package account

import (
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"log"
	"net/http"
	"os/exec"
)

type Interface interface {
	// Login Call Google Account Oauth and get token code which returned.
	Login(ctx context.Context) (*AuthInfo, error)
}

type OptionFunc func(config *oauth2.Config)

func WithClientID(clientID string) OptionFunc {
	return func(config *oauth2.Config) {
		config.ClientID = clientID
	}
}

func WithRedirectURL(redirectURL string) OptionFunc {
	return func(config *oauth2.Config) {
		config.RedirectURL = redirectURL
	}
}

// NewGoogleAccount Construction function for Google Account Interface, use factory method pattern.
func NewGoogleAccount(opts ...OptionFunc) Interface {
	config := defaultOauth2Config
	for _, opt := range opts {
		opt(config)
	}

	return &DefaultGoogleAccount{
		config: config,
	}
}

type DefaultGoogleAccount struct {
	config *oauth2.Config
}

func (gc *DefaultGoogleAccount) Login(ctx context.Context) (*AuthInfo, error) {
	// use default config if config is nil
	if gc.config == nil {
		gc.config = defaultOauth2Config
	}

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	authInfo := make(chan []byte)
	go gc.listen(ctx, authInfo)
	defer close(authInfo)

	err := gc.open(ctx)
	if err != nil {
		return nil, err
	}

	select {
	case resp := <-authInfo:
		aInfo := &AuthInfo{}
		if err = json.Unmarshal(resp, &aInfo); err != nil {
			return nil, err
		}
		return aInfo, nil
	}
}

// listen google account callback
// call it with goroutine and communicate use channel.
func (gc *DefaultGoogleAccount) listen(ctx context.Context, authInfo chan []byte) {
	listenAddr := fmt.Sprintf("localhost:%d", ListenPort)
	mux := http.NewServeMux()
	single := make(chan []byte)

	mux.HandleFunc(CallbackUrl, func(writer http.ResponseWriter, request *http.Request) {
		_, err := writer.Write([]byte("Your have logged in, please close this window and return to console "))
		if err != nil {
			log.Printf("response ok to google callback failed: %v", err)
		}

		reqParams, err := json.Marshal(request.URL.Query())
		if err != nil {
			log.Printf("get google callback params failed:%v", err)
			return
		}
		single <- reqParams
	})

	srv := http.Server{
		Addr:    listenAddr,
		Handler: mux,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("listen %s failed: %v \n", srv.Addr, err)
			return
		}
	}()

	defer func() {
		if err := srv.Shutdown(context.TODO()); err != nil {
			log.Printf("shutdown http server failed,err: %v", err)
		}
	}()

	select {
	case <-ctx.Done():
	case body := <-single:
		authInfo <- body
	}

}

// Open google account login web in browser.
func (gc *DefaultGoogleAccount) open(ctx context.Context) error {
	authURL := gc.config.AuthCodeURL("state", oauth2.AccessTypeOffline)
	if authURL == "" {
		return fmt.Errorf("get AuthCodeURL with config failed")
	}

	fullURL := GoogleAuthURL + authURL

	if err := exec.Command("open", fullURL).Run(); err != nil {
		return fmt.Errorf("open browser failed: %v \n", err)
	}
	fmt.Printf("Your browser has been opened to visit: \n %s \n Please login with Google Account \n", fullURL)

	return nil
}
