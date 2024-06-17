package add

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/delavalom/dar/internal/storage"
	"github.com/spf13/cobra"
)

func NewCloneCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "clone",
		Short: "Clone file contents to the index",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Cloning file contents to the repository...")

			var tree storage.Tmp

			b, err := os.ReadFile(".dar/tmp")
			if err != nil {
				panic(err)
			}

			if err = json.Unmarshal(b, &tree); err != nil {
				panic(err)
			}

			// Read tmp but actually should read last commit index tree
			//
			for k, v := range tree.Tree {
				fmt.Println(k, v)
			}

			fmt.Println("Done!")
		},
	}
}
