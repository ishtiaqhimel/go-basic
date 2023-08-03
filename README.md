# Go Basic
## What is Go?
Go is expressive, concise, clean, and efficient. Go has gained popularity in recent years for the following reasons. Go was created at Google in 2007, and since then, engineering teams across the globe have adopted Go to build products and services at massive scale.
- **Fast and Lightweight language** - Go is a language that manages memory with its built-in garbage collector. Hence it is fast and lightweight. It has a quick compilation time.
- **Statically Typed Language** - The type of go variables are known at compile time.
- **Clean and Concise Language** - The syntax of go is short and easier to understand.
- **C like Language** - Go has adopted lots of features from the C programming language and is similar to C.

## Why learn Go Programming?
- **Cloud & Network Services** - With a strong ecosystem of tools and APIs on major cloud providers, it is easier than ever to build services with Go.
</br>Go has a strong ecosystem supporting service development. The standard library includes packages for common needs like HTTP servers and clients, JSON/XML parsing, SQL databases, and a range of security/encryption functionality, while the Go runtime includes tools for race detection, benchmarking/profiling, code generation, and static code analysis.
</br>The major Cloud providers (GCP, AWS, Azure) have Go APIs for their services, and popular open source libraries provide support for API tooling (Swagger), transport (protocol buffers, gRPC), monitoring (OpenCensus), Object-Relational Mapping (gORM), and authentication (JWT). The open source community has also provided several service frameworks, including Go Kit, Go Micro, and Gizmo, which can be a great way to get started quickly.
- **Command-line Interfaces (CLIs)** - With popular open source packages and a robust standard library, use Go to create fast and elegant CLIs.
</br>When developing CLIs in Go, two tools are widely used: Cobra & Viper.
</br>Cobra is both a library for creating powerful modern CLI applications and a program to generate applications and CLI applications in Go. Cobra powers most of the popular Go applications including CoreOS, Delve, Docker, Dropbox, Git Lfs, Hugo, Kubernetes, and many more. With integrated command help, autocomplete and documentation, it makes documenting each command really simple.
</br>Viper is a complete configuration solution for Go applications, designed to work within an app to handle configuration needs and formats. Cobra and Viper are designed to work together.
</br>Viper supports nested structures in the configuration, allowing CLI developers to manage the configuration for multiple parts of a large application. Viper also provides all of the tooling need to easily build twelve factor apps.
- **Web Development** - With enhanced memory performance and support for several IDEs, Go powers fast and scalable web applications.
</br>Go is designed to enable developers to rapidly develop scalable and secure web applications. Go ships with an easy to use, secure and performant web server and includes it own web templating library. Go has excellent support for all of the latest technologies from HTTP/2, to databases like MySQL, MongoDB and Elasticsearch, to the latest encryption standards including TLS 1.3. Go web applications run natively on Google App Engine and Google Cloud Run (for easy scaling) or on any environment, cloud, or operating system thanks to Go’s extreme portability.
- **Development Operations & Site Reliability Engineering** - With fast build times, lean syntax, an automatic formatter and doc generator, Go is built to support both DevOps and SRE.
</br>Go serves both siblings, DevOps and SRE, from its fast build times and lean syntax to its security and reliability support. Go’s concurrency and networking features also make it ideal for tools that manage cloud deployment—readily supporting automation while scaling for speed and code maintainability as development infrastructure grows over time.
</br>DevOps/SRE teams write software ranging from small scripts, to command-line interfaces (CLI), to complex automation and services, and Go’s feature set has benefits for every situation.

## How to run Go?
The simplest way of running Go code is by using an online Go compiler like [The Go Playground](https://go.dev/play/).

You can also easily install and run Go programming locally on your computer. For that, visit [Install Go on Windows, Linux, and macOS](https://go.dev/doc/install).

From CLI we can run go program with `go run <file-name>.go`. The `file-name` file must contain `main()` function from `main` package. If we have many file in same package, we must run `go run .`.

## Hello World
Basic Structure of a Go Program
Here are the things to take away from this tutorial.

1. All Go programs start from the main() function.
2. The bare minimum code structure of a Go program is:
```go
package main

fun main() {
    // code starts here
}
```

## Learning Path for Go programming language (Golang)
- Basics
    - Basic Syntax
    - Data Types
    - Variables and declaration
    - Conditionals (if, switch)
    - For loop
    - Array, slice, maps
    - Functions (multiple/named returns, variadic, closure)
    - Structs
    - Type interface
    - Type casting
    - Package, imports and exports
        - IO 
        - Template (text)
    - Errors, Panic, Recover
    - Testing
- Concurrency
    - Goroutines
    - Types, Type Assertions, switches
    - Interfaces
    - Context
    - Channel
    - Buffer
    - Select
    - Mutex


### *First Read*
- [Language specification](https://go.dev/ref/spec)
- [How To Write Go Code](https://go.dev/doc/code)

### *Book*
- [An Introduction to Programming in Go](https://www.golang-book.com/books/intro)

### *Additional Resources/Blog links*
- [A Tour of Go](https://go.dev/tour/welcome/1)
- [GolangBot](https://golangbot.com/learn-golang-series/)
- [GoByExample](https://gobyexample.com/)
- [Closure Usage - calhoun](https://www.calhoun.io/5-useful-ways-to-use-closures-in-go/)