package admin

import (
	"github.com/spf13/cobra"
)

func NewCmdAdmin() *cobra.Command {
	cmd := &cobra.Command{
		Use: "admin",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
	return cmd
}
