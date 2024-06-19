package root

import (
	"github.com/delavalom/dar/internal/cmd/add"
	"github.com/delavalom/dar/internal/cmd/checkout"
	"github.com/delavalom/dar/internal/cmd/clone"
	"github.com/delavalom/dar/internal/cmd/commit"
	"github.com/delavalom/dar/internal/cmd/initialize"
	"github.com/delavalom/dar/internal/cmd/push"
	"github.com/spf13/cobra"
)

func NewRootCommand() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "dar",
		Short: "dar is a simple version control system written in Go with git-lfs out of the box",
	}

	cmd.PersistentFlags().Bool("help", false, "Show help for command")

	cmd.AddCommand(
		initialize.NewInitCommand(),
		add.NewAddCommand(),
		commit.NewCommitCommand(),
		push.NewPushCommand(),
		checkout.NewCheckoutCommand(),
		clone.NewCloneCommand(),
	)

	return cmd, nil
}
