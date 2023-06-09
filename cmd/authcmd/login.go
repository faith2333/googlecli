/*
Copyright © 2023 Annan Wang <wan199406@applovin.com>
*/
package authcmd

import (
	"context"
	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := auth.Login(context.Background()); err != nil {
			panic(err)
		}
	},
}

func init() {
	AuthCmd.AddCommand(loginCmd)

}
