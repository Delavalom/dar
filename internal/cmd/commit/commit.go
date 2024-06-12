package commit

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewCommitCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "commit",
		Short: "Record changes to the repository",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Committing changes to the repository...")
			fmt.Println("Done!")
		},
	}
}
