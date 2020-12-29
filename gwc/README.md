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