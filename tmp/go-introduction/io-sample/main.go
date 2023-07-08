package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// -------------------------- Print Input ------------------------------------------
	// there are 3 functions to print msg in Go
	fmt.Print("Hello, ")
	fmt.Println("World!")

	/*
		Format specifier
		int    -> %d
		float  -> %f
		string -> %s
		bool   -> %t
	*/
	val := 8.009
	fmt.Printf("%0.2f", val)

	// --------------------------- Take Input -----------------------------------------
	// there are 3 functions to take inputs
	// var name string
	// var age int
	// fmt.Scan(&name, &age)
	// fmt.Printf("Name: %s\nAge: %d\n", name, age)

	// to get input values up to the new line
	// fmt.Scanln(&name, &age)
	// fmt.Printf("Name: %s\nAge: %d\n", name, age)

	// similar to Scanln() function and takes inputs using format specifiers
	// fmt.Scanf("%s %d", &name, &age)
	// fmt.Printf("Name: %s\nAge: %d\n", name, age)

	// ---------------------------- File IO --------------------------------------------------
	// Checking if a file exists or not: Stat returns file info., returns an error if there is no file.

	fileInfo, err := os.Stat("test.txt")
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("File does not exist.")
		}
	}
	log.Println(fileInfo.Name())

	// according to the book...
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	// use defer file.Close() right after opening the file to make sure the file is closed as soon as the function completes.
	// see the definitions of defer if needed

	// get the file size
	stat, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	// read the file
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		log.Fatal(err)
	}

	str := string(bs)
	fmt.Println(str)

	// Creating a file
	cfile, err := os.Create("test.txt")
	if err != nil {
		// handle the error here
		return
	}
	defer cfile.Close()
	cfile.WriteString("Ishtiaq\n25")

	dir, err := os.Open(".")
	if err != nil {
		return
	}
	defer dir.Close()

	// To get the contents of a directory we use the same os.
	// Open function but give it a directory path instead of a file name. Then we call the Readdir method
	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		return
	}
	for _, fi := range fileInfos {
		fmt.Println(fi.Name())
	}

	// Often we want to recursively walk a folder (read the folder's contents, all the sub-folders, all the
	// sub-sub-folders, …). To make this easier there's a Walk function provided in the path/filepath package
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	})

	// reading file using ioutil lib.
	fileContent, err := ioutil.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", fileContent) // fileContent -> all the characters from the file
	for _, val := range fileContent {
		if val == '\n' {
			fmt.Println()
		} else {
			fmt.Println(string(val))
		}
	}

	// using the ioutil
	fs, err := ioutil.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	st := string(fs)
	fmt.Println(st)
}

/*
It's recommended to use the fmt package for printing.
We usually use println(), print() only for debugging purposes.
*/

// TODO: bufio package
