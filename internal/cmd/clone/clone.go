package clone

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/delavalom/dar/internal/snapshot"
	"github.com/delavalom/dar/internal/storage"
	"github.com/spf13/cobra"
)

func NewCloneCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "clone",
		Short: "Clone file contents to the index",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Cloning file contents to the repository...")

			// Read the index file to get the key
			indexBytes, err := os.ReadFile(".dar/index")
			if err != nil {
				panic(err)
			}

			key := string(indexBytes)

			// Read the index snapshot from the snapshots directory
			snapshotBytes, err := os.ReadFile(fmt.Sprintf(".dar/snapshots/%s", key))
			if err != nil {
				panic(err)
			}
			var commit snapshot.Commit
			if err = json.Unmarshal(snapshotBytes, &commit); err != nil {
				panic(err)
			}

			if commit.Key != key {
				panic("There's nothing to commit")
			}

			// Read tmp but actually should read last commit index tree
			// then read the files in the object directory and write on the file system
			storage.WriteFiles(commit.Tree)

			fmt.Println("Done!")
		},
	}
}
