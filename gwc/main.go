// countwords counts words in the given files or in stdin
// Usage:
// gwc file1 file2 ... fileN
// gwc (no args: count files in stdin)

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ulrichwinter/learngolang/gwc/countwords"
)

func main() {

	files := os.Args[1:]
	if len(files) == 0 {
		count, err := countwords.Countwords(os.Stdin)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("\n%d\n", count)
	} else {
		var overallcount, numfiles int
		for _, filename := range files {
			file, err := os.Open(filename)
			if err != nil {
				fmt.Printf("Cannot open file %s: %v\n", filename, err)
				continue
			}

			count, err := countwords.Countwords(file)
			if err != nil {
				fmt.Printf("Error counting in file %s: %v", filename, err)
				continue
			}
			fmt.Printf("%s: %d\n", filename, count)
			overallcount += count
			numfiles++
		}
		fmt.Printf("Total of %d files: %d\n", numfiles, overallcount)
	}
}
