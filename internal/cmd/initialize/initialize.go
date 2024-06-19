package initialize

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func NewInitCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "Create an empty Dar repository or reinitialize an existing one",
		Run: func(cmd *cobra.Command, args []string) {
			// Create a .dar directory to store the repository information
			if err := os.MkdirAll(".dar", 0755); err != nil {
				panic(err)
			}

			// Create a config file to store the repository configuration
			configFile, err := os.Create(".dar/config")
			if err != nil {
				panic(err)
			}
			if err = configFile.Close(); err != nil {
				panic(err)
			}

			// Create a HEAD file to store the reference to the current branch
			headFile, err := os.Create(".dar/HEAD")
			if err != nil {
				panic(err)
			}
			defer headFile.Close()

			// Set the default branch to main
			if _, err = headFile.WriteString("ref: refs/heads/main"); err != nil {
				panic(err)
			}

			// Objects to store the content of files of the repository
			if err = os.Mkdir(".dar/objects", 0755); err != nil {
				if !errors.Is(err, os.ErrExist) {
					panic(err)
				}
			}

			// Trees to store snapshots of the repository
			if err = os.Mkdir(".dar/trees", 0755); err != nil {
				if !errors.Is(err, os.ErrExist) {
					panic(err)
				}
			}

			fmt.Println("Initialized empty Dar repository in .dar")
		},
	}
}
