package add

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewAddCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "add",
		Short: "Add file contents to the index",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Adding file contents to the index...")
			fmt.Println("Done!")
		},
	}
}
