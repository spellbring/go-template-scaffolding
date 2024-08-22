package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"io/ioutil"
	"path/filepath"

	"github.com/spf13/cobra"
)

var (
	templateDir      string
	targetDir        string
	modName          string
	placeholder      string
	authorPlaceHoder string
	author           string
)

func main() {
	placeholder = "{{bootstrap_template}}"
	authorPlaceHoder = "{{author}}"
	templateDir = "template"
	var rootCmd = &cobra.Command{
		Use:   "template-go-cli",
		Short: "CLI to setup project templates in golang",
		Run: func(cmd *cobra.Command, args []string) {
			err := copyTemplate(templateDir, targetDir, modName, placeholder, authorPlaceHoder, author)
			if err != nil {
				fmt.Println("Error:", err)
				os.Exit(1)
			}
			fmt.Println("Template copied and module name replaced successfully.")

			if err := os.Chdir(targetDir); err != nil {
				fmt.Println("Error changing directory:", err)
				os.Exit(1)
			}

			if err := executeGoModInit(); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			if err := executeGoModTidy(); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
	}

	rootCmd.Flags().StringVarP(&targetDir, "target", "d", "", "Target directory path (required)")
	rootCmd.Flags().StringVarP(&modName, "modname", "m", "", "Go module name to replace (required)")
	rootCmd.Flags().StringVarP(&author, "author", "a", "", "Name of team or author (required)")

	rootCmd.MarkFlagRequired("target")
	rootCmd.MarkFlagRequired("modname")
	rootCmd.MarkFlagRequired("author")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

func copyTemplate(templateDir, targetDir, modName, placeholder, authorPlaceHolder, author string) error {
	err := filepath.Walk(templateDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath := strings.TrimPrefix(path, templateDir)
		targetPath := filepath.Join(targetDir, relPath)

		if info.IsDir() {
			return os.MkdirAll(targetPath, info.Mode())
		}

		content, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		contentStr := strings.ReplaceAll(string(content), placeholder, modName)
		contentStr = strings.ReplaceAll(contentStr, authorPlaceHolder, author)

		err = ioutil.WriteFile(targetPath, []byte(contentStr), info.Mode())
		if err != nil {
			return err
		}

		return nil
	})

	return err
}

func executeGoModInit() error {
	cmd := exec.Command("go", "mod", "init", modName)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
func executeGoModTidy() error {
	cmd := exec.Command("go", "mod", "tidy")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
