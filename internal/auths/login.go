package auths

import (
	"context"
	"fmt"
	"github.com/faith2333/googlecli/pkg/google/account"
)

// Login implement of Auths Interface login method
func (dAuth *defaultAuth) Login(ctx context.Context) error {
	if dAuth.googleAccount == nil {
		dAuth.googleAccount = account.NewGoogleAccount()
	}

	authInfo, err := dAuth.googleAccount.Login(ctx)
	if err != nil {
		return err
	}

	fmt.Println(authInfo)
	return nil
}
