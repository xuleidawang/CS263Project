# CS263 Runtime System Project
## Topic: Investigate and empirically evaluate Go concurrency mechanism extensively.

## Project Vision
\
Go lang is a statical language with garbage collection, concurrent features developed by Google. Concurrency is known as making progress on multiple tasks simultaneously within one program. Go has a rich support for this using goroutines and channels. In this project, as greenhands in Go, we aims at exploring  the cool features of Go, especially in concurrency. Dive deep inside its concurrency mechanism and evaluate its performance.

## Project Goals
- Understand how Go addresses fundamental problems that make concurrency difficult to do correctly.
- Learn the key difference between concurrency and parrallelism.
- Dig into the syntax of Go's memory synchronization primitives.
- Form patterns with these primitives to write maintainable concurrent code.
- Compose patterns into a series of practices that enable us to write large, distributed system that scale
- Learn the sophistication behind goroutines and how Go's runtime stitches everthing.


## Team members:
[Lei Xu](https://github.com/xuleidawang)
![lei](https://raw.githubusercontent.com/xuleidawang/CS263Project/master/images/lei.jpg)
[Yifu Luo](https://github.com/443582555)
![yifu](https://raw.githubusercontent.com/xuleidawang/CS263Project/master/images/yifu.jpg)

## 1.Get Started
(1) Install Golang
[https://golang.org/doc/install?download=go1.10.1.darwin-amd64.pkg#install](https://golang.org/doc/install?download=go1.10.1.darwin-amd64.pkg#install) 

(2) Code organization 
Golang Code Overview:
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

To build our project: Please git clone our repo under the directory  .../go/src/  on your machine. 


[Reference]
How to write Go code[https://golang.org/doc/code.html#Workspaces](https://golang.org/doc/code.html#Workspaces) \
Learn GO Concurrency [https://github.com/golang/go/wiki/LearnConcurrency](https://github.com/golang/go/wiki/LearnConcurrency)
