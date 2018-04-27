#CS263 Runtime System Project
## Topic: Investigate and empirically evaluate Go concurrency mechanism extensively.

## Project Vision
\
Go lang is a statical language with garbage collection, concurrent features developed by Google. Concurrency is known as making progress on multiple tasks simultaneously within one program. Go has a rich support for this using goroutines and channels. In this project, as greenhands in Go, we aims at exploring  the cool features of Go, especially in concurrency. Dive deep inside its concurrency mechanism and evaluate its performance.

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
\

(3) Concurrency Overview

In computer science, concurrency refers to the ability of different parts or units of a program, algorithm, or problem to be executed out-of-order or in partial order, without affecting the final outcome. This allows for parallel execution of the concurrent units, which can significantly improve overall speed of the execution in multi-processor and multicore system. 
In many other programming languages, concurrent is made difficult by the subtleties required to implement correct access to shared variables. However, Go encouraged a different approach by passing the shared variable on channels and never actively shared by separate thread of execution. Only one goroutine has access to the value at any given time. 

(4) Go Routines




Week one:  
We implemented a simple go program that run two print function concurrently. We built a funtion that print digits and one print alphabets. Using runtime.GOMAXPROCS(2), we are able to see that both function print chars concurrently and each time the print value is not the same.   
Concept learned:  
Any function or method in Go can be created as a goroutine.   
Go runtime schedules goroutines to run within a logical processor that is bound to a single operating system thread. By default, the Go runtime allocates a single logical processor to execute all the goroutines that are created for our program(Need to explore further)
Concurrency is not Parallelism. Parallelism is when two or more threads are executing code simultaneously against different processors.  
If you configure the runtime to use more than one logical processor, the scheduler will distribute goroutines between these logical processors which will result in goroutines running on different operating system threads. However, to have true parallelism you need to run your program on a machine with multiple physical processors. If not, then the goroutines will be running concurrently against a single physical processor, even though the Go runtime is using multiple logical processors.  

Week two:


[Reference]
How to write Go code[https://golang.org/doc/code.html#Workspaces](https://golang.org/doc/code.html#Workspaces)
Learn GO Concurrency [https://github.com/golang/go/wiki/LearnConcurrency](https://github.com/golang/go/wiki/LearnConcurrency)
