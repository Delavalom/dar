package initialize

import (
	"fmt"
	"io/fs"
	"os"

	"github.com/spf13/cobra"
)

func NewInitCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "push",
		Short: "Push changes to the remote repository",
		Run: func(cmd *cobra.Command, args []string) {
			if err := os.Mkdir(".dar", fs.ModeDir); err != nil {
				panic(err)
			}

			fmt.Println("Initialized empty Dar repository in .dar")
		},
	}
}
