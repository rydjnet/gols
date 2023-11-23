package main

import (
	"fmt"
	"log"
	"os"
	"syscall"
)

func dirRecurse(d string, indent int) {
	files, err := os.ReadDir(d)
	if err != nil {
		log.Fatal(err)
	}
	for file := range files {
		if files[file].IsDir() {
			fmt.Printf("%*s%s\n", indent, "", files[file])
			fmt.Printf("  ")
			dirRecurse(d+"/"+files[file].Name(), indent+4)
		}

		fmt.Printf("%*s%s\n", indent, "", files[file])

	}
}
func main() {
	var path string
	switch {
	case len(os.Args) > 2:
		log.Fatal("Error: too many arguments, use only one")

	case len(os.Args) == 1:
		path, _ = syscall.Getwd()

	case len(os.Args) == 2:
		path = os.Args[1]

	}
	dirRecurse(path, 0)

}
