package commit

import (
	"fmt"
	"os"

	"github.com/delavalom/dar/internal/snapshot"
	"github.com/spf13/cobra"
)

func NewCommitCommand() *cobra.Command {
	commitCommand := &cobra.Command{
		Use:   "commit [-m | --message] <msg>",
		Short: "Record changes to the repository",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			msg, err := cmd.Flags().GetString("message")
			if err != nil {
				fmt.Println("error: commit message must be wrapped in double quotes")
				os.Exit(1)
			}

			fmt.Println("Committing changes to the repository...")
			file, err := os.Create(".dar/COMMIT_EDITMSG")
			if err != nil {
				panic(err)
			}
			defer file.Close()

			if _, err = file.Write([]byte(msg)); err != nil {
				panic(err)
			}

			indexFile, err := os.Create(".dar/index")
			if err != nil {
				panic(err)
			}
			defer indexFile.Close()

			shot := snapshot.New("key", nil, "Adding new file")

			if _, err = indexFile.Write([]byte(shot.Key)); err != nil {
				panic(err)
			}

			fmt.Println("Done!")
		},
	}

	commitCommand.Flags().StringP("message", "m", "", "Use the given <msg> as the commit message")
	commitCommand.MarkFlagRequired("message")

	return commitCommand
}
