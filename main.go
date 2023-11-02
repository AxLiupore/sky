package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sky/repl"
)

func main() {
	pwd, _ := os.Getwd()
	content, err := os.ReadFile(filepath.Join(pwd, "banner"))
	if err != nil {
		panic(err)
	}
	sky := string(content)
	fmt.Printf(sky)
	repl.Start(os.Stdin, os.Stdout)
}
