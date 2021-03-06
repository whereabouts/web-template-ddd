package cmd

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/whereabouts/sdk/cli/command"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// 子命令：重命名项目名
var renameCmd = command.NewCommand(
	command.WithName("rename"),
	command.WithUsage("Rename the mod of the project"),
	command.WithArgsUsage("[new project name]"),
).WithAction(func(v command.Value) error {
	if v.NArg() > 1 {
		return errors.New("Args num must be 1")
	}
	name := v.Args().Get(0)
	if name == "" {
		name = defaultName
	}
	fmt.Println("Start to rename project: ", name)

	if err := initDir(pwd, name); err != nil {
		return err
	}
	fmt.Println("\nRename success!")

	return nil
})

func initDir(dir string, name string) error {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, file := range files {
		path := filepath.Join(dir, file.Name())
		if file.IsDir() {
			if err = initDir(path, name); err != nil {
				return err
			}
		} else {
			data, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			str := strings.ReplaceAll(string(data), templateName, name)
			err = ioutil.WriteFile(path, []byte(str), os.ModePerm)
			if err != nil {
				return err
			}
			rel, _ := filepath.Rel(pwd, path)
			fmt.Println("[Success] ", rel)
		}
	}
	return nil
}
