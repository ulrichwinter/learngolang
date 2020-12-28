# learngolang
learning golang with practical examples

## wordcount
Implement a simple version of a commandline tool to count words in files like the famous "wc -w".

Usage:

```bash
$ go get github.com/ulrichwinter/learngolang/gwc
$ cd $GOPATH/src/github.com/ulrichwinter/learngolang
$ gwc gwc/main.go wordcount/wordcount.go 
gwc/main.go: 102
wordcount/wordcount.go: 84
Total of 2 files: 186
``` 
Assuming $GOPATH/bin is in your PATH.

### Setup GOPATH src package folder for further pushes to origin
In order to have the resulting workspace `$GOROOT/src/github.com/ulrichwinter` set up to be used for furhter commits, the following git config is needed first:
```
git config --global url.git@github.com:.insteadOf https://github.com/
```
This leads to have `go get` create a package src folder with the desired origin url.
