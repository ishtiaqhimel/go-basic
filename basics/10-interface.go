package main

import (
	"fmt"
	"math"
)

// interface definition
type VowelsFinder interface {
	FindVowels() []rune
}

type MyString string

// MyString implements VowelsFinder
func (ms MyString) FindVowels() []rune {
	var vowels []rune
	for _, x := range ms {
		if x == 'a' || x == 'e' || x == 'i' || x == 'o' || x == 'u' {
			vowels = append(vowels, x)
		}
	}
	return vowels
}

// Example
type circle struct {
	radius float64
}

type rectangle struct {
	length float64
	width  float64
}

type shape interface {
	area() float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r rectangle) area() float64 {
	return r.length * r.width
}

// extra can be implemented but can't access by interface
func (r rectangle) perimeter() float64 {
	return (r.length + r.width) * 2.0
}

// Empty Interface
func havingFun(i interface{}) {
	switch i.(type) { // Type switch
	case string:
		fmt.Println("I'm a string.", i.(string)) // Type assertion: s, ok := i.(string) , if we don't get ok then it will panic for wrong type assertion
	case circle:
		fmt.Println("I'm a circle having radius:", i.(circle).radius)
	}
}

// Embedded interface
type personalData interface {
	printPersonalData()
}

type professionalData interface {
	printProfessionalData()
}

type personInfo interface {
	personalData
	professionalData
}

type person struct {
	name        string
	age         int
	companyName string
	salary      int
}

func (p person) printPersonalData() {
	fmt.Println(p.name, p.age)
}

func (p person) printProfessionalData() {
	fmt.Println(p.companyName, p.salary)
}

func main() {
	name := MyString("Ishtiaq Islam")
	var v VowelsFinder
	v = name // possible since MyString implements VowelsFinder
	fmt.Printf("Vowels are %c\n", v.FindVowels())

	// Example
	c := circle{4.9}
	r := rectangle{5.6, 3.4}
	s := []shape{c, r} // pointer receivers and value receivers act as different for interface implementation
	for _, x := range s {
		fmt.Println(x.area())
	}

	// empty interface
	havingFun("ishtiaq")
	havingFun(c)

	// Under the hood, interface values can be thought of as a tuple of a value and a concrete type
	for _, x := range s {
		fmt.Printf("(%v, %T)\n", x, x)
	}

	// embedded interface
	p := person{"abc", 12, "xyz", 9000}
	var pi personInfo
	pi = p
	pi.printPersonalData()
	pi.printProfessionalData()
}

/*
In Go, an interface is a set of method signatures. When a type provides definition for all the methods in the interface, it is said to implement the interface.
For example WashingMachine can be an interface with method signatures Cleaning() and Drying(). Any type which provides definition for Cleaning() and Drying()
methods is said to implement the WashingMachine interface.

Interfaces are implemented implicitly.

Zero value of interface is nil.

Practical Use of Interface:
- if two different instances share common properties
*/
