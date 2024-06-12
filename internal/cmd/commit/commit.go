package commit

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func NewCommitCommand() *cobra.Command {
	commitCommand := &cobra.Command{
		Use:   "commit [-m OPTIONS] <msg>",
		Short: "Record changes to the repository",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Println("error: pathspec required")
				os.Exit(1)
			}
			msg := args[0]

			fmt.Println("Committing changes to the repository...")
			file, err := os.Create(".dar/COMMIT_EDITMSG")
			if err != nil {
				panic(err)
			}
			defer file.Close()

			if _, err = file.WriteString(msg); err != nil {
				panic(err)
			}

			fmt.Println("Done!")
		},
	}

	commitCommand.Flags().StringP("message", "m", "", "Use the given <msg> as the commit message")
	commitCommand.MarkFlagRequired("message")

	return commitCommand
}
