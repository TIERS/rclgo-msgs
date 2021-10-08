package main

import (
	"flag"
	"fmt"
	"os"

	"golang.org/x/mod/modfile"
)

var (
	modFilePath     string
	rclgoImportPath string
)

func init() {
	flag.StringVar(&modFilePath, "mod-file", "go.mod", "path to the module file")
	flag.StringVar(&rclgoImportPath, "rclgo-import-path", "github.com/TIERS/rclgo", "import path of rclgo library")
	flag.Parse()
}

func run() int {
	modData, err := os.ReadFile(modFilePath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to read module file:", err)
		return 1
	}
	mod, err := modfile.ParseLax(modFilePath, modData, nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to parse module file:", err)
		return 1
	}
	for _, req := range mod.Require {
		if req.Mod.Path == rclgoImportPath {
			fmt.Println(req.Mod.Version)
			return 0
		}
	}
	fmt.Fprintln(os.Stderr, "no rclgo import found")
	return 1
}

func main() {
	os.Exit(run())
}
