package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

func main() {
	fileReader()
	stringsReader()
	connReader()
}

func fileReader() {
	f, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	readerToStdout(f, 4)
}

func stringsReader() {
	s := strings.NewReader("My name is Ishtiaq.")
	readerToStdout(s, 5)
}

func connReader() {
	conn, err := net.Dial("tcp", "google.com:80")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	fmt.Fprint(conn, "GET HTTP 1.0\r\n\r\n")
	readerToStdout(conn, 25)
}

func readerToStdout(r io.Reader, bufSize int) {
	buf := make([]byte, bufSize)
	for {
		// It will use only buf to read, that means it will overwrite the data
		n, err := r.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			break
		}
		if n > 0 {
			// to avoid overwrite issue print n bytes
			fmt.Println(string(buf[:n]))
		}
	}
}
