package add

import (
	"fmt"
	"os"

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
			fmt.Println(cwd)
			files, err := os.ReadDir(cwd)
			fmt.Println(files)
			if err != nil {
				panic(err)
			}
			storage.ReadFiles(files, "")
			// storage.ReadFileMock(storage.Files)
			fmt.Println("Done!")
		},
	}
}
