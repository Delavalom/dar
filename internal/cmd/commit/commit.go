package commit

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/delavalom/dar/internal/hash"
	"github.com/delavalom/dar/internal/hashMap"
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

			// Create a file to store the commit message
			file, err := os.Create(".dar/COMMIT_EDITMSG")
			if err != nil {
				panic(err)
			}
			defer file.Close()

			if _, err = file.Write([]byte(msg)); err != nil {
				panic(err)
			}

			// Read the tree snapshot from the tmp file
			treeBytes, err := os.ReadFile(".dar/tmp")
			if err != nil {
				panic(err)
			}

			tree := hashMap.New()

			if err := json.Unmarshal(treeBytes, &tree); err != nil {
				panic(err)
			}

			key := hash.New(treeBytes)

			// Create a snapshot of the repository
			shot := snapshot.New(string(key), tree, msg)

			// marshal the snapshot to bytes
			snapshotBytes, err := json.Marshal(shot)
			if err != nil {
				panic(err)
			}

			// Create a file to store the snapshot of the tree
			snapshotFile, err := os.Create(fmt.Sprintf(".dar/snapshots/%s", key))
			if err != nil {
				panic(err)
			}

			// Write the snapshot to the snapshot file
			if _, err = snapshotFile.Write(snapshotBytes); err != nil {
				panic(err)
			}

			if err = snapshotFile.Close(); err != nil {
				panic(err)
			}

			// Create a file to store the commit key
			indexFile, err := os.Create(".dar/index")
			if err != nil {
				panic(err)
			}
			defer indexFile.Close()

			// Write the key of the snapshot to the index file
			if _, err = indexFile.Write([]byte(key)); err != nil {
				panic(err)
			}

			// Remove the tmp file
			if err = os.Remove(".dar/tmp"); err != nil {
				panic(err)
			}

			fmt.Println("Done!")
		},
	}

	commitCommand.Flags().StringP("message", "m", "", "Use the given <msg> as the commit message")
	commitCommand.MarkFlagRequired("message")

	return commitCommand
}
