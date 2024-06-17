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
			if err := os.MkdirAll(".dar", 0755); err != nil {
				panic(err)
			}

			configFile, err := os.Create(".dar/config")
			if err != nil {
				panic(err)
			}
			if err = configFile.Close(); err != nil {
				panic(err)
			}

			file, err := os.Create(".dar/HEAD")
			if err != nil {
				panic(err)
			}
			defer file.Close()

			if _, err = file.WriteString("ref: refs/heads/main"); err != nil {
				panic(err)
			}

			if err = os.Mkdir(".dar/objects", 0755); err != nil {
				if !errors.Is(err, os.ErrExist) {
					panic(err)
				}
			}

			if err = os.Mkdir(".dar/tmp", 0755); err != nil {
				if !errors.Is(err, os.ErrExist) {
					panic(err)
				}
			}

			fmt.Println("Initialized empty Dar repository in .dar")
		},
	}
}
