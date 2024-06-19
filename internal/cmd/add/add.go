package add

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/delavalom/dar/internal/hashMap"
	"github.com/delavalom/dar/internal/storage"
	"github.com/spf13/cobra"
)

func NewAddCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "add",
		Short: "Add file contents to the index",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Adding file contents to the index...")
			cwd, err := os.Getwd()
			if err != nil {
				panic(err)
			}

			files, err := os.ReadDir(cwd)
			if err != nil {
				panic(err)
			}

			tree := hashMap.New()

			// Read files and store them in a tree structure
			storage.WriteTree(files, "", tree)

			// Marshal tree to bytes to store value into a file
			b, err := json.Marshal(tree)
			if err != nil {
				panic(err)
			}

			// tmp to store staged files tree structure before commit
			tmpFile, err := os.Create(".dar/tmp")
			if err != nil {
				panic(err)
			}
			defer tmpFile.Close()

			// Write the tmpBytes to the tmp file
			if _, err = tmpFile.Write(b); err != nil {
				panic(err)
			}

			fmt.Println("Done!")
		},
	}
}
