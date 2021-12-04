package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"path/filepath"
)

func listDir(listPath string) {
	items, err := ioutil.ReadDir(listPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: could not list directory '%s'\n", listPath)
		os.Exit(1)
	}

	for _, item := range items {
		itemPath := item.Name()
		if listPath != "." {
			itemPath = filepath.Join(listPath, itemPath)
		}

		if (item.IsDir()) {
			listDir(itemPath)
		} else {
			fmt.Println(itemPath)
		}
	}
}

func main() {
	listPath := "."
	if len(os.Args) == 2 {
		listPath = os.Args[1]
	}

	listDir(listPath)
}
