## IO Package
Go is a programming language built for working with bytes. Whether you have lists of bytes, streams of bytes, or individual bytes, Go makes it easy to process. From these simple primitives we build our abstractions and services.

The io package is one of the most fundamental packages within the standard library. It provides a set of interfaces and helpers for working with streams of bytes.

There are two fundamental operations when working with bytes: reading & writing. Let’s take a look at reading bytes first.

### Reading Bytes
- **Reader Interface:** The basic construct for reading bytes from a stream is the Reader interface:
```go
type Reader interface {
        Read(p []byte) (n int, err error)
}
```
- **func LimitReader(r Reader, n int64) Reader** LimitReader returns a Reader that reads from r but stops with EOF after n bytes. </br>
- **func MultiReader(readers ...Reader) Reader** MultiReader returns a Reader that's the logical concatenation of the provided input readers. They're read sequentially. Once all inputs have returned EOF, Read will return EOF. If any of the readers return a non-nil, non-EOF error, Read will return that error.</br>
- **func TeeReader(r Reader, w Writer) Reader** TeeReader returns a Reader that writes to w what it reads from r. All reads from r performed through it are matched with corresponding writes to w. There is no internal buffering - the write must complete before the read completes. Any error encountered while writing is reported as a read error.
- **Helper Functions:** ReadFull(), ReadAtLeast(), ReadAll()

### Writing Bytes
- **Writer Interface:** The Writer interface is simply the inverse of the Reader. We provide a buffer of bytes to push out onto a stream. Generally speaking writing bytes is simpler than reading them. Readers complicate data handling because they allow partial reads, however, partial writes will always return an error.
```go
type Writer interface {
        Write(p []byte) (n int, err error)
}
```
- **func MultiWriter(writers ...Writer) Writer** MultiWriter creates a writer that duplicates its writes to all the provided writers, similar to the Unix tee(1) command. Each write is written to each listed writer, one at a time. If a listed writer returns an error, that overall write operation stops and returns the error; it does not continue down the list.

- **func WriteString(w Writer, s string) (n int, err error)** WriteString writes the contents of the string s to w, which accepts a slice of bytes. If w implements StringWriter, its WriteString method is invoked directly. Otherwise, w.Write is called exactly once. WriteString() method which can be used to `improve write performance` by not requiring an allocation when converting a string to a byte slice. 

### Copying Bytes
- **Connecting readers and writers** The most basic way to copy a reader to a writer is the aptly named Copy() function. This function uses a 32KB buffer to read from src and then write to dst. If any error besides io.EOF occurs in the read or write then the copy is stopped and the error is returned. One issue with Copy() is that you cannot guarantee a maximum number of bytes. For example, you may want copy a log file up to its current file size. If the log continues to grow during your copy then you’ll end up with more bytes than expected. In this case you can use the CopyN() function to specify an exact number of bytes to be written.
### Closing Streams
The Closer interface is provided as a generic way to close streams:
```go
type Closer interface {
        Close() error
}
```
### Moving Around Within Streams
The Seeker interface is provided to jump around in a stream:
```go
type Seeker interface {
        Seek(offset int64, whence int) (int64, error)
}
```

- [Doc](https://pkg.go.dev/io)
- [Resource](https://medium.com/go-walkthrough/go-walkthrough-io-package-8ac5e95a9fbd)