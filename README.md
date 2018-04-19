# CS263 Runtime System Project
## Topic: Investigate and empirically evaluate Go concurrency mechanism extensively.

#### Go lang is a statical language with garbage collection, concurrent features developed by Google. Concurrency is known as making progress on multiple tasks simultaneously within one program. Go has a rich support for this using goroutines and channels. In this project, as greenhands in Go, we aims at exploring  the cool features of Go, especially in concurrency. Dive deep inside its concurrency mechanism and evaluate its performance.

## Team members:
[Lei Xu](https://github.com/xuleidawang)
![lei](https://github.com/xuleidawang/CS263Project/images/lei.jpg)
[Yifu Luo](https://github.com/443582555)
![yifu](https://github.com/xuleidawang/CS263Project/images/yifu.jpg)
## Install 
[https://golang.org/doc/install?download=go1.10.1.darwin-amd64.pkg#install](https://golang.org/doc/install?download=go1.10.1.darwin-amd64.pkg#install) 


## Code organization 
Overview:
- Go code are keeped in a single *workspace*
- A workspace contains many version control *repositories* (managed by Git, like this repository)
- Each repository contains one or more *packages*
- Each package consist of one or more Go source files in a single directory.  

\
An example workspace hierarchy:
```
bin/
    hello                          # command executable
    outyet                         # command executable
pkg/
    linux_amd64/
        github.com/golang/example/
            stringutil.a           # package object
src/
    github.com/golang/example/
        .git/                      # Git repository metadata
	hello/
	    hello.go               # command source
	outyet/
	    main.go                # command source
	    main_test.go           # test source
	stringutil/
	    reverse.go             # package source
	    reverse_test.go        # test source
    golang.org/x/image/
        .git/                      # Git repository metadata
	bmp/
	    reader.go              # package source
	    writer.go              # package source
    ... (many more repositories and packages omitted) ...
```
\
[Reference](https://golang.org/doc/code.html#Workspaces)