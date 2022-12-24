package main

import (
	"fmt"
	"strings"
)

func main() {
	// strings package
	// bool: contains as substring or not
	text := "Ishtiaq"
	pattern := "sh"
	fmt.Println(strings.Contains(text, pattern))

	// int: counts the number of times pattern occurs as substring
	text = "aabbaaccdd"
	pattern = "aa"
	fmt.Println(strings.Count(text, pattern))

	// bool: checks if a string has occured as a prefix/suffix of another string
	text = "appscode"
	prefix := "apps"
	suffix := "code"
	fmt.Println(strings.HasPrefix(text, prefix))
	fmt.Println(strings.HasSuffix(text, suffix))

	// first occurrence index if present, -1 otherwise
	fmt.Println(strings.Index(text, suffix))

	sep := "#"
	fmt.Println(
		// Joining with a separator
		strings.Join([]string{"abc", "def", "ghi"}, sep),

		// Reapeat the text n-times
		strings.Repeat(text, 3),

		// Replaces the substring with the value given
		strings.Replace("aaaddaaa", "a", "b", 5),
		// output: "bbbddbba", replaces first 5 occurrences
	)

	// if n < 0, there is no limit on the number of replacements.
	info := "ishtiaq.islam.appscode.dhaka-1230,bangladesh"
	fmt.Println(strings.Replace(info, ".", "-", -1))

	// split by the second param string: output: [x y z], returns a list of string
	charList := strings.Split("xabyabz", "ab")
	fmt.Println(charList)
	fmt.Println(strings.Split("a-b-c-d-e", "-"))

	// to change case of string
	fmt.Println(strings.ToLower("ISHTIAQ"), strings.ToUpper("islam"))

	// Sometimes we need to work with strings as binary data. To convert a string to a slice of bytes (and vice-versa) do this:
	arr := []byte("test")
	str := string([]byte{'t', 'e', 's', 't'})
	fmt.Println(arr, str)

	// see the following:
	// fileIO part for file handling
	// code frequently used part for sorting array, slices or struct
	// mutex example for Mutexes
	// error for error handling
}
