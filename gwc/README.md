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

