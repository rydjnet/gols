package main

import (
	"fmt"
	"log"
	"os"
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
	dirRecurse("/tmp/", 0)

}
