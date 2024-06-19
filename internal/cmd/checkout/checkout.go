package checkout

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func NewCheckoutCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "checkout <branch to switch to>",
		Short: "Switch branches",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Println("error: pathspec required")
				os.Exit(1)
			}
			branch := args[0]

			file, err := os.OpenFile(".dar/HEAD", os.O_RDWR, 0644)
			if err != nil {
				panic(err)
			}
			defer file.Close()

			file.WriteString(fmt.Sprintf("ref: refs/heads/%s", branch))

			fmt.Printf("Switched to branch '%s'\n", branch)
		},
	}
}
