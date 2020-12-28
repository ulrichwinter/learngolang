package wordcount_test

import (
	"io"
	"os"
	"strings"
	"testing"
	"umw/wordcount"
)

var tests = []struct {
	words   int
	content string
}{
	{0, ""},
	{0, " "},
	{0, "\t"},
	{0, " \t\n\t   \t"},
	{1, " eins "},
	{2, "eins zwei"},
	{3, "\tone\ttwo\t\n\t\tthree\n\n"},
}

func TestCountWords(t *testing.T) {
	for _, test := range tests {
		var content io.Reader
		content = strings.NewReader(test.content)
		got, _ := wordcount.Countwords(content)
		if got != test.words {
			t.Errorf("want %d words, got %d: %q", test.words, got, test.content)
		}
	}
}

// benchmark using a file containing "moby dick" (see https://gist.githubusercontent.com/StevenClontz/4445774/raw/1722a289b665d940495645a5eaaad4da8e3ad4c7/mobydick.txt)
func BenchmarkCountwords(b *testing.B) {
	for i := 0; i < b.N; i++ {
		file, err := os.Open("mobydick.txt")
		if err != nil {
			b.Fatalf("cannot open file \"mobidick.txt\": %v", err)
		}
		got, err := wordcount.Countwords(file)
		file.Close()

		if err != nil {
			b.Errorf("got unexpected error %v", err)
		}
		if got != 115314 {
			b.Errorf("got wrong word count %d - want %d", got, 115314)
		}
	}
}
