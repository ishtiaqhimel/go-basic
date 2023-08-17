## Modules
Go modules are a group of related packages that are versioned and distributed together. They specify the requirements of our project, list all the required dependencies, and help us keep track of the specific versions of installed dependencies.

Modules are identified by a module path that is declared in the first line of the `go.mod` file in our project.

The `go.mod` file defines the module’s module path, which is also the import path used for the root directory, and its dependency requirements, which are the other modules needed for a successful build. Each dependency requirement is written as a module path and a specific semantic version.

## Creating a new module
- Create a new empty dir (lets say dir name is `hello`) somewhere outside `$GOPATH/src`, `cd` into that dir and then create a new source file `hello.go`
- Lets import a package thats not included in go standard library
- At this point the dir contains a package but not a module because there is no `go.mod` file. If we were working in `/home/xyz/hello`, to make the current dir the root of a module by using `go mod init`.
```
$ go mod init example.com/hello
go: creating new go.mod: module example.com/hello
```

It will create a `go.mod` file.
```
$ cat go.mod
module example.com/hello

go 1.xx
```

The `go.mod` file only appears in the root of the module. Packages in subdirectories have import paths consisting of the module path plus the path to the subdirectory. For example, if we created a subdirectory world, we would not need to (nor want to) run go mod init there. The package would automatically be recognized as part of the `example.com/hello` module, with import path `example.com/hello/world`.

## Adding a dependency
The primary motivation for Go modules was to improve the experience of using (that is, adding a dependency on) code written by other developers. After importing a new pkg if we run `go run .` or any pkg building go command then go will automatically looks up the module containing that package and adds it to `go.mod` using latest version.

Note that while the go command makes adding a new dependency quick and easy, it is not without cost. Your module now literally depends on the new dependency in critical areas such as correctness, security, and proper licensing, just to name a few.

The command `go list -m all` lists the current module and all its dependencies.

- `go mod init` creates a new module, initializing the go.mod file that describes it.
- `go build`, `go test`, and other package-building commands add new dependencies to `go.mod` as needed.
- `go list -m all` prints the current module’s dependencies.
- `go get xyz@version` changes the required version of a dependency (or adds a new dependency).
- `go mod tidy` removes unused dependencies.

## Ref
- [Using Go Modules](https://go.dev/blog/using-go-modules)
- [How to Use Go Modules](https://www.digitalocean.com/community/tutorials/how-to-use-go-modules)