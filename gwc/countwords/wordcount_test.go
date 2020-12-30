package countwords_test

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/ulrichwinter/learngolang/gwc/countwords"
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
	{3, "one\ttwo\nthree"},
	{2, "one\u00A0two"},   // NO-BREAK SPACE
	{2, "one\u1680two"},   // OGHAM SPACE
	{2, "two\u2028three"}, // Line separator
}

func TestCountWords(t *testing.T) {
	for _, test := range tests {
		var content io.Reader
		content = strings.NewReader(test.content)
		got, _ := countwords.Countwords(content)
		if got != test.words {
			t.Errorf("want %d words, got %d: %q", test.words, got, test.content)
		}
	}
}

// benchmark using a file containing "moby dick" (see https://gist.githubusercontent.com/StevenClontz/4445774/raw/1722a289b665d940495645a5eaaad4da8e3ad4c7/mobydick.txt)
func BenchmarkCountwordsMobyDick(b *testing.B) {
	for i := 0; i < b.N; i++ {
		file, err := os.Open("mobydick.txt")
		if err != nil {
			b.Fatalf("cannot open file \"mobidick.txt\": %v", err)
		}
		got, err := countwords.Countwords(file)
		file.Close()

		if err != nil {
			b.Errorf("got unexpected error %v", err)
		}
		if got != 115314 {
			b.Errorf("got wrong word count %d - want %d", got, 115314)
		}
	}
}

// benchmark using large files with random generated content

func BenchmarkCountwordsLargefile_1K(b *testing.B) {
	benchmarkCountwordsLargefile(b, 1_000)
}

func BenchmarkCountwordsLargefile_10K(b *testing.B) {
	benchmarkCountwordsLargefile(b, 10_000)
}

func BenchmarkCountwordsLargefile_100K(b *testing.B) {
	benchmarkCountwordsLargefile(b, 100_000)
}

func BenchmarkCountwordsLargefile_1M(b *testing.B) {
	benchmarkCountwordsLargefile(b, 1_000_000)
}

func BenchmarkCountwordsLargefile_10M(b *testing.B) {
	benchmarkCountwordsLargefile(b, 10_000_000)
}

func BenchmarkCountwordsLargefile_100M(b *testing.B) {
	benchmarkCountwordsLargefile(b, 100_000_000)
}

func benchmarkCountwordsLargefile(b *testing.B, size int64) {
	filename := generateTestFileWithSize(size)
	defer os.Remove(filename)

	b.SetBytes(size)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		file, err := os.Open(filename)
		if err != nil {
			b.Fatalf("cannot open file %q: %v", filename, err)
		}
		_, err = countwords.Countwords(file)
		file.Close()

		if err != nil {
			b.Errorf("got unexpected error %v", err)
		}
	}
	b.StopTimer()
}

func generateTestFileWithSize(size int64) string {
	// generate txt file with the given size:
	// base64 /dev/urandom | head -c 10000000 > file.txt
	filename := fmt.Sprintf("/tmp/benchfile-%d.txt", rand.Intn(1000))
	err := exec.Command("/bin/bash", "-c", fmt.Sprintf("base64 /dev/urandom | head -c %d > %s", size, filename)).Run()
	if err != nil {
		os.Stderr.WriteString(err.Error())
	}
	return filename
}
