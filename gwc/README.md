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
