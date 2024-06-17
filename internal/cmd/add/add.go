package add

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/delavalom/dar/internal/hash"
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

			storage.ReadFiles(files, "", tree)
			fmt.Println(tree)
			fmt.Println(tree["cmd"].SubTree)
			fmt.Println(tree["internal"].SubTree["internal/cmd"].SubTree["internal/cmd/add"].SubTree["internal/cmd/add/add.go"].FileContent)

			b, err := json.Marshal(tree)
			if err != nil {
				panic(err)
			}

			sha1 := hash.New(b)

			tmp := &storage.Tmp{
				Key:  sha1,
				Tree: tree,
			}

			tmpBytes, err := json.Marshal(tmp)
			if err != nil {
				panic(err)
			}

			file, err := os.Create(".dar/tmp")
			if err != nil {
				panic(err)
			}
			defer file.Close()

			_, err = file.Write(tmpBytes)
			if err != nil {
				panic(err)
			}

			fmt.Println("Done!")
		},
	}
}
