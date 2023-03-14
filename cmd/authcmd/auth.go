/*
Copyright © 2023 Annan Wang <wan199406@gmail.com>
*/
package authcmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// AuthCmd represents the authcmd command
var AuthCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authentication and Authorization for google account",
	Long:  `Authentication and Authorization for google account`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cmd.Help())
	},
}

func init() {

}
