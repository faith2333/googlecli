package account

import (
	"context"
	"fmt"
	"os"
	"testing"
)

var defaultGoogleAccount Interface

func TestMain(m *testing.M) {
	defaultGoogleAccount = NewGoogleAccount()
	os.Exit(m.Run())
}

func TestDefaultGoogleAccount_Login(t *testing.T) {
	resp, err := defaultGoogleAccount.Login(context.TODO())
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(resp)
}
