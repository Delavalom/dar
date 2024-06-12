package push

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewPushCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "push",
		Short: "Push changes to the remote repository",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Pushing changes to the remote repository...")
			fmt.Println("Done!")
		},
	}
}
