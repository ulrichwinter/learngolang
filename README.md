# learngolang
learning golang with practical examples

# Examples (so far)
* [gwc](gwc) count words in files

## Download and build using `go get``

```bash
$ go get github.com/ulrichwinter/learngolang/...
``` 
Add $GOPATH/bin to $PATH to ececute the built commands directly.

## Setup GOPATH src package folder for further pushes to origin
In order to have the resulting workspace `$GOROOT/src/github.com/ulrichwinter` set up to be used for furhter commits, the following git config is needed first:
```bash
git config --global url.git@github.com:.insteadOf https://github.com/
```
This leads to have `go get` create a package src folder with the desired origin url.
