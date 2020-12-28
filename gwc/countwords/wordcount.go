package countwords

import (
	"bufio"
	"io"
	"unicode"
)

// Countwords counts the words in its input
func Countwords(in io.Reader) (int, error) {
	var inword bool
	var count int

	var current [1]byte

	input := bufio.NewReaderSize(in, 8*1024)
	for {
		_, err := input.Read(current[:])
		if err == io.EOF {
			break
		}
		if err != nil {
			return 0, err
		}

		if !unicode.IsSpace(rune(current[0])) {
			if !inword {
				count++
				inword = true
			}
		} else {
			inword = false
		}
	}
	return count, nil
}
