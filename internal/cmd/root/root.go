package root

import (
	"dar/internal/cmd/commit"
	"dar/internal/cmd/initialize"
	"dar/internal/cmd/push"
	"fmt"

	"github.com/spf13/cobra"
)

func NewRootCommand() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "dar",
		Short: "dar is a simple version control system written in Go with git-lfs out of the box",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("dar is a dominican version control system")
		},
	}

	cmd.PersistentFlags().Bool("help", false, "Show help for command")

	cmd.AddCommand(initialize.NewInitCommand())
	cmd.AddCommand(commit.NewCommitCommand())
	cmd.AddCommand(push.NewPushCommand())

	return cmd, nil
}
