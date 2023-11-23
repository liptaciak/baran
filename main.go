package main

import (
	"os"
	"fmt"
	"flag"
	"io/fs"
	"runtime"
	"path/filepath"
)

func walk(s string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}

	if !d.IsDir() {
		fmt.Println(s)
	}

	return nil
}	

func main() {
	help := flag.Bool("h", false, "Shows list of flags and usage.")
	list := flag.Bool("l", false, "List all installed packages.")

	flag.Parse()

	if *list {
		if runtime.GOOS != "windows" {
			filepath.WalkDir(os.Getenv("HOME") + "/.barn/libs", walk)
		} else {
			fmt.Println("This command is not available for Windows.")
		}
	} else if *help {
		flag.PrintDefaults()
	} else {
		flag.PrintDefaults()
	}
}
