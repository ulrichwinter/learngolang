package countwords

import (
	"bufio"
	"io"
	"unicode"
)

// Countwords counts the words in its input
func Countwords(source io.Reader) (int, error) {
	const bufferSize = 8 * 1024
	var inword bool
	var count int

	input := bufio.NewReaderSize(source, bufferSize)
	for {
		r, _, err := input.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			return 0, err
		}
		if !unicode.IsSpace(r) {
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
