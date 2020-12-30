## gwc wordcount using golang
Implement a simple version of a commandline tool to count words in files like the famous "wc -w".

Usage:

```bash
$ go get github.com/ulrichwinter/learngolang/gwc
$ cd $GOPATH/src/github.com/ulrichwinter/learngolang/gwc

$ go build
$ find . -name "*.go" | xargs ./gwc
./countwords/wordcount.go: 80
./countwords/wordcount_test.go: 152
./main.go: 129
Total of 3 files: 361
``` 

## performance compared with `wc`

The performance is still 2-3 times slower than what `wc` needs for the same input on the same machine:

```bash
$ time wc -w countwords/mobydick.txt 
  115314 countwords/mobydick.txt

real	0m0.005s
user	0m0.002s
sys	0m0.002s

$ time ./gwc countwords/mobydick.txt 
countwords/mobydick.txt: 115314
Total of 1 files: 115314

real	0m0.017s
user	0m0.011s
sys	0m0.003s
```

Using larger buffer size does not change much: e.g.: using a buffer of 800*1024 bytes which is larger thant the file content:

```bash
$ time ./gwc countwords/mobydick.txt 
countwords/mobydick.txt: 115314
Total of 1 files: 115314

real	0m0.017s
user	0m0.011s
sys	0m0.004s
```

## build wc from source to analyze its implementation
The `wc` utility is part of Darwin and Apple provides its sourcecode at opensource.apple.com.

I'm not sure about the verion naming schema used. But as the `wc` utility doesnt seem to have changed lately, it does not really matter to use the latest version.

Building `wc` from source needs the following simple steps:
```bash
$ wget https://opensource.apple.com/source/text_cmds/text_cmds-101.40.1/wc/wc.c
$ cc wc.c

$ time ./a.out ~/Downloads/mobydick.txt
   15603  115314  643210 /Users/ulrichwinter/Downloads/mobydick.txt

real	0m0.008s
user	0m0.005s
sys	0m0.002s
```

As one can see at [wc.c](https://opensource.apple.com/source/text_cmds/text_cmds-101.40.1/wc/wc.c), the original `wc` uses the same basic elements as this golang countwords when counting words:
- read the file contents chunkwise into a buffer
- loop over the buffer and check if the current character is whitspace (using `iswspace()`) considering wide characters


The size of the buffer is the "optimal transfer block size" of the underlying file system or 8 kB if that information is unavailable. By adding a `printf()` to wc, we find out, which buffer size is used: 1MB (1048576 u_chars)

But as to be expected, changing the buffer size to 1MB (`bufio.NewReaderSize(in, 1024*1024)`) doesn't speed up the golang version:
```bash
$ time ./gwc ~/Downloads/mobydick.txt 
/Users/ulrichwinter/Downloads/mobydick.txt: 115314
Total of 1 files: 115314

real 0m0.015s
user 0m0.010s
sys  0m0.004s
```

## more benchmarks

The current "mobydick.txt" benchmark file is about 6.5 kB large and fits completely into the currently used buffer of 8k.

Benchmarks are not executed during a normal `go build` and have to be specified explicitly like here:
```bash
$ cd countwords/
$ go test -bench=.
goos: darwin
goarch: amd64
pkg: github.com/ulrichwinter/learngolang/gwc/countwords
BenchmarkCountwordsMobyDick-8         	     163	   7428442 ns/op
BenchmarkCountwordsLargefile_1K-8     	   43468	     28025 ns/op	  35.68 MB/s
BenchmarkCountwordsLargefile_10K-8    	   10000	    110604 ns/op	  90.41 MB/s
BenchmarkCountwordsLargefile_100K-8   	    1255	    941877 ns/op	 106.17 MB/s
BenchmarkCountwordsLargefile_1M-8     	     129	   9223164 ns/op	 108.42 MB/s
BenchmarkCountwordsLargefile_10M-8    	      12	  92148298 ns/op	 108.52 MB/s
BenchmarkCountwordsLargefile_100M-8   	       2	 915957432 ns/op	 109.18 MB/s
PASS
ok  	github.com/ulrichwinter/learngolang/gwc/countwords	14.955s
```

Here a comparative benchmark using different file sizes has been added, which uses a file with random content generated using `base64 /dev/urandom | head -c $SIZE > file.txt`.
The benchmark shows, that the throughput stabilizes around 100 MB/s for file sizes above 10kB.

## Fixing multibyte errors accidentally improves performance

The previous version read the input bytewise using `io.Reader.Read([]byte)`. 
The byte was than converted to a rune and checked for whitespace using `unicode.IsSpace(rune)`.
This obviously misses multi byte white space characters.

After changing the implementation do read a rune at a time from the input reader using `bufio.Reader.ReadRune()` the performance also was improved by around 25% regarding the byte/s throughput for large files:

Before:
```bash
$ go test -bench=.
goos: darwin
goarch: amd64
pkg: github.com/ulrichwinter/learngolang/gwc/countwords
BenchmarkCountwordsMobyDick-8         	     164	   7153355 ns/op
BenchmarkCountwordsLargefile_1K-8     	   45661	     26665 ns/op	  37.50 MB/s
BenchmarkCountwordsLargefile_10K-8    	    9600	    109167 ns/op	  91.60 MB/s
BenchmarkCountwordsLargefile_100K-8   	    1255	    936369 ns/op	 106.80 MB/s
BenchmarkCountwordsLargefile_1M-8     	     128	   9222763 ns/op	 108.43 MB/s
BenchmarkCountwordsLargefile_10M-8    	      13	  92084036 ns/op	 108.60 MB/s
BenchmarkCountwordsLargefile_100M-8   	       2	 922631872 ns/op	 108.39 MB/s
PASS
ok  	github.com/ulrichwinter/learngolang/gwc/countwords	14.662s
```

After:
```bash
$ go test -bench=.
goos: darwin
goarch: amd64
pkg: github.com/ulrichwinter/learngolang/gwc/countwords
BenchmarkCountwordsMobyDick-8         	     214	   5565818 ns/op
BenchmarkCountwordsLargefile_1K-8     	   48391	     24138 ns/op	  41.43 MB/s
BenchmarkCountwordsLargefile_10K-8    	   14089	     85313 ns/op	 117.22 MB/s
BenchmarkCountwordsLargefile_100K-8   	    1712	    691615 ns/op	 144.59 MB/s
BenchmarkCountwordsLargefile_1M-8     	     176	   6728695 ns/op	 148.62 MB/s
BenchmarkCountwordsLargefile_10M-8    	      16	  67035026 ns/op	 149.18 MB/s
BenchmarkCountwordsLargefile_100M-8   	       2	 728145192 ns/op	 137.34 MB/s
PASS
ok  	github.com/ulrichwinter/learngolang/gwc/countwords	14.643s
```
