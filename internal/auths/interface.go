package auths

import (
	"context"
	"github.com/faith2333/googlecli/pkg/google/account"
)

type Interface interface {
	Login(ctx context.Context) error
}

type Config struct {
	ClientID []string `json:"clientID"`
}

type OptionFunc func(config *Config)

func NewAuth(authType AuthType, opts ...OptionFunc) Interface {
	config := &Config{}

	for _, opt := range opts {
		opt(config)
	}

	switch authType {
	case AuthTypeDefault:
		return &defaultAuth{conf: config}
	default:
		return &defaultAuth{conf: config}
	}
}

type defaultAuth struct {
	conf          *Config
	googleAccount account.Interface
}
