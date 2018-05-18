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

## 2. Overview
(1) Concurrency Overview

In computer science, concurrency refers to the ability of different parts or units of a program, algorithm, or problem to be executed out-of-order or in partial order, without affecting the final outcome. This allows for parallel execution of the concurrent units, which can significantly improve overall speed of the execution in multi-processor and multicore system. 
In many other programming languages, concurrent is made difficult by the subtleties required to implement correct access to shared variables. However, Go encouraged a different approach by passing the shared variable on channels and never actively shared by separate thread of execution. Only one goroutine has access to the value at any given time. 

The concurrency features of Go seemed new.
But they are tooted in a long history, reaching back to Hoare's CSP in 1978 and even Dijkstra's guarded commands(1975).

Programming languages with similar features:
- Occam (May, 1983)
- Erlang (Armstrong, 1986)
- Newsqueak (Pike, 1988)
- Concurrent ML (Reppy, 1993)
- Alef (Winterbottom, 1995)
- Limbo (Doewaed, Pike Winterbottom, 1996)

Distinction\
Go is the latest on the Newsqueak-Alef-Limbo branch, distinguished by first-class channels.\
Erlang is closer to original CSP, where we communicate to a process by name rather than over a channel.\
The models are equivalent but express things differently.\
Rough analogy: writing to a file by name(process, Erlang) vs. writing to a file descriptor(channel, Go).


(2) Goroutines
A Goroutine is a lightweight thread managed by the Go runtime.
A goroutine has a simple model: it is a function executing concurrently with other goroutines in the same address space. It only require a little stack space, which are start small and grow by allocating heap storage as required.

Goroutines are multiplexed onto multiple OS threads so if one should block, such as waiting for I/O, others continue to run. Their design hides many of the complexities of thread creation and management.

```
package main

import (
    "fmt"
    "time"
)

func say(s string) {
    for i := 0; i < 5; i++ {
        time.Sleep(100 * time.Millisecond)
        fmt.Println(s)
    }
}

func main() {
    go say("world")
    say("hello")
}

```


Prefix a function or method call with the go keyword to run the call in a new goroutine. When the call completes, the goroutine exits. 
```
go list.Sort()  // run list.Sort concurrently; don't wait for it.
```

A function literal can be handy in a goroutine invocation.

```
func Announce(message string, delay time.Duration) {
    go func() {
        time.Sleep(delay)
        fmt.Println(message)
    }()  // Note the parentheses - must call the function.
}
```
Goroutines run inthe same address space, so access to shared memory must be synchronized. The sync package provides useful primitives.

In Go, function literals are closures: the implementation makes sure the variables referred to by the function survive as long as they are active.
These examples aren't too practical because the functions have no way of signaling completion. For that, we need channels.

(3) Channels
A channel provides a mechanism for concurrently executing functions to communicate by sending and receiving values of a specified element type. The value of an uninitialized channel is nil.\

The optional <- operator specifies the channel direction, send or receive. If no direction is given, the channel is bidirectional. A channel may be constrained only to send or only to receive by conversion or assignment.
```
chan T          // can be used to send and receive values of type T
chan<- float64  // can only be used to send float64s
<-chan int      // can only be used to receive ints
```

Like maps, channels are allocated with keyword *make*, and the resulting value acts as a reference to an underlying data structure. If an optional integer parameter is provided, it sets the buffer size for the channel. The default is zero, for an unbuffered or synchronous channel.

```
ci := make(chan int)            // unbuffered channel of integers
cj := make(chan int, 0)         // unbuffered channel of integers
cs := make(chan *os.File, 100)  // buffered channel of pointers to Files
```

Unbuffered channels combine communication—the exchange of a value—with synchronization—guaranteeing that two calculations (goroutines) are in a known state.

There are lots of nice idioms using channels. Here's one to get us started. In the previous section we launched a sort in the background. A channel can allow the launching goroutine to wait for the sort to complete.

```
c := make(chan int)  // Allocate a channel.
// Start the sort in a goroutine; when it completes, signal on the channel.
go func() {
    list.Sort()
    c <- 1  // Send a signal; value does not matter.
}()
doSomethingForAWhile()
<-c   // Wait for sort to finish; discard sent value.

```


Receivers always block until there is data to receive. If the channel is unbuffered, the sender blocks until the receiver has received the value. If the channel has a buffer, the sender blocks only until the value has been copied to the buffer; if the buffer is full, this means waiting until some receiver has retrieved a value.

A buffered channel can be used like a semaphore, for instance to limit throughput. In this example, incoming requests are passed to handle, which sends a value into the channel, processes the request, and then receives a value from the channel to ready the “semaphore” for the next consumer. The capacity of the channel buffer limits the number of simultaneous calls to process.
```
var sem = make(chan int, MaxOutstanding)

func handle(r *Request) {
    sem <- 1    // Wait for active queue to drain.
    process(r)  // May take a long time.
    <-sem       // Done; enable next request to run.
}

func Serve(queue chan *Request) {
    for {
        req := <-queue
        go handle(req)  // Don't wait for handle to finish.
    }
}
```


See the sample program channels.go 
Send a value into a channel using the channel <- syntax. Here we send "ping" to the messages channel we made above, from a new goroutine.
The <-channel syntax receives a value from the channel. Here we’ll receive the "ping" message we sent above and print it out.
When we run the program the "ping" message is successfully passed from one goroutine to another via our channel.
By default sends and receives block until both the sender and receiver are ready. This property allowed us to wait at the end of our program for the "ping" message without having to use any other synchronization.



## 3. Go statements
A "go" statement starts the execution of a function call as an independent concurrent thread of control, or goroutine, within the same address space.

```
GoStmt = "go" Expression .
```

The expression must be a function or method call; it cannot be parenthesized. Calls of built-in functions are restricted as for expression statements.

The function value and parameters are evaluated as usual in the calling goroutine, but unlike with a regular call, program execution does not wait for the invoked function to complete. Instead, the function begins executing independently in a new goroutine. When the function terminates, its goroutine also terminates. If the function has any return values, they are discarded when the function completes.

```
go Server()
go func(ch chan<- bool) { for { sleep(10); ch <- true }} (c)
```


Select statements\
A "select" statement chooses which of a set of possible send or receive operations will proceed. It looks similar to a "switch" statement but with the cases all referring to communication operations.
```
SelectStmt = "select" "{" { CommClause } "}" .
CommClause = CommCase ":" StatementList .
CommCase   = "case" ( SendStmt | RecvStmt ) | "default" .
RecvStmt   = [ ExpressionList "=" | IdentifierList ":=" ] RecvExpr .
RecvExpr   = Expression .
```
A case with a RecvStmt may assign the result of a RecvExpr to one or two variables, which may be declared using a short variable declaration. The RecvExpr must be a (possibly parenthesized) receive operation. There can be at most one default case and it may appear anywhere in the list of cases.

Execution of a "select" statement proceeds in several steps:

1. For all the cases in the statement, the channel operands of receive operations and the channel and right-hand-side expressions of send statements are evaluated exactly once, in source order, upon entering the "select" statement. The result is a set of channels to receive from or send to, and the corresponding values to send. Any side effects in that evaluation will occur irrespective of which (if any) communication operation is selected to proceed. Expressions on the left-hand side of a RecvStmt with a short variable declaration or assignment are not yet evaluated.
2. If one or more of the communications can proceed, a single one that can proceed is chosen via a uniform pseudo-random selection. Otherwise, if there is a default case, that case is chosen. If there is no default case, the "select" statement blocks until at least one of the communications can proceed.
3. Unless the selected case is the default case, the respective communication operation is executed.
4. If the selected case is a RecvStmt with a short variable declaration or an assignment, the left-hand side expressions are evaluated and the received value (or values) are assigned.
5. The statement list of the selected case is executed.

Since communication on nil channels can never proceed, a select with only nil channels and no default case blocks forever.

```
var a []int
var c, c1, c2, c3, c4 chan int
var i1, i2 int
select {
case i1 = <-c1:
    print("received ", i1, " from c1\n")
case c2 <- i2:
    print("sent ", i2, " to c2\n")
case i3, ok := (<-c3):  // same as: i3, ok := <-c3
    if ok {
        print("received ", i3, " from c3\n")
    } else {
        print("c3 is closed\n")
    }
case a[f()] = <-c4:
    // same as:
    // case t := <-c4
    //  a[f()] = t
default:
    print("no communication\n")
}

for {  // send random sequence of bits to c
    select {
    case c <- 0:  // note: no statement, no fallthrough, no folding of cases
    case c <- 1:
    }
}

select {}  // block forever
```

Return statements\
A "return" statement in a function F terminates the execution of F, and optionally provides one or more result values. Any functions deferred by F are executed before F returns to its caller.
```
ReturnStmt = "return" [ ExpressionList ] .
```
In a function without a result type, a "return" statement must not specify any result values.
```
func noResult() {
    return
}
```

There are three ways to return values from a function with a result type:

1. The return value or values may be explicitly listed in the "return" statement. Each expression must be single-valued and assignable to the corresponding element of the function's result type.
```
func simpleF() int {
    return 2
}

func complexF1() (re float64, im float64) {
    return -7.0, -4.0
}
```
2. The expression list in the "return" statement may be a single call to a multi-valued function. The effect is as if each value returned from that function were assigned to a temporary variable with the type of the respective value, followed by a "return" statement listing these variables, at which point the rules of the previous case apply.

```
func complexF2() (re float64, im float64) {
    return complexF1()
}
```

3. The expression list may be empty if the function's result type specifies names for its result parameters. The result parameters act as ordinary local variables and the function may assign values to them as necessary. The "return" statement returns the values of these variables.
```
func complexF3() (re float64, im float64) {
    re = 7.0
    im = 4.0
    return
}

func (devnull) Write(p []byte) (n int, _ error) {
    n = len(p)
    return
}
```

- Why is concurrency supported?
interaction environment. 
Writing a program deal with real world.
Concurrency is the composition of independently executing computations.
Concurrency is a way to structure software, particularly as a way to write clean code that interacts well with the real world.

Not parallelism.
One processor can still be concurrent but cannot be parallel.

A model for software construction
Easy to understande.
Easy to use.
Easy to reason about.

CSP paper 


- What is concurrecny, anyway?
- Where does the idea come from?
- What is it good for?
- How do we use it ? 

## 4. Concurrency features in GO and Java
To analyze concurrency features, we implement simple matrix multiplication programs in both Go and Java. Java implementation uses Java Thread, and Go implementation uses Goroutine and Channel. From the experiment, Go derived better performance than Java in both compile time and concurrency. Moreover, Go code shows the ease of concurrent programming. Go is still young, but we are convinced that Go will become the mainstream.

The Java Programming Language was released by Sun Microsystems in 1995. It is a concurrent, object-oriented, and garbage-collected programming language, which is widely used for several areas of application. Java has built-in support for concurrency: the Thread class, Runnable interface, and java.util.concurrent concurrent package. They provide powerful features for multi-thread programming.

In this section, we focus on concurrency feature, which is the one of common features of Go and Java. The purpose of this project is to analyze the performance of Go and compare it with Java on two aspects: compile time and the concurrency feature. To analyze the performance, we prepare simple matrix multiplication programs, implemented in Go and Java.

Tang presented two multi-core parallel programs in Go in order to show the ease of multi-core parallel programming using Go. Implementations of benchmarks are parallel integration and parallel dynamic programming. He also measured performance of Go with shifting the number of cores used. He concludes that it is easy to write multi-core programs in Go, and the highest speed of benchmarks are derived on an octal-core AMD chip.

### Method
This section explains the methods for performance comparison. Matrix multiplication is often used for performance evaluation of programming languages as shown in [5], [6], and [7]. In this experiment, we use simple matrix multiplication that is calculated by C=AB. To simplify the problem, we define two matrices that have same length of row and column. First, we divide matrix A into some processes, and calculate the product of a part of A and matrix B. After completion of all calculation, we combine them. Figure 1 illustrates the part of calculation of matrix C.
All the measurements were performed on the same machine. Details of the hardware and software are given below:


fig

The benchmarks are implementations of simple matrix multiplication in both Go and Java. In order to measure performance, we prepare four benchmark sources: matrix.go, parallel_matrix.go, Matrix.java and ParallelMatrix.java.



Week one:  
We implemented a simple go program that run two print function concurrently. We built a funtion that print digits and one print alphabets. Using runtime.GOMAXPROCS(2), we are able to see that both function print chars concurrently and each time the print value is not the same.   
Concept learned:  
Any function or method in Go can be created as a goroutine.   
Go runtime schedules goroutines to run within a logical processor that is bound to a single operating system thread. By default, the Go runtime allocates a single logical processor to execute all the goroutines that are created for our program(Need to explore further)
Concurrency is not Parallelism. Parallelism is when two or more threads are executing code simultaneously against different processors.  
If you configure the runtime to use more than one logical processor, the scheduler will distribute goroutines between these logical processors which will result in goroutines running on different operating system threads. However, to have true parallelism you need to run your program on a machine with multiple physical processors. If not, then the goroutines will be running concurrently against a single physical processor, even though the Go runtime is using multiple logical processors.  



[Reference]
How to write Go code[https://golang.org/doc/code.html#Workspaces](https://golang.org/doc/code.html#Workspaces) \
Learn GO Concurrency [https://github.com/golang/go/wiki/LearnConcurrency](https://github.com/golang/go/wiki/LearnConcurrency)
