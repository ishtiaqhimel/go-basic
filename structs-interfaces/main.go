package main

import (
	"fmt"
	"math"
)

// Student struct
type Student struct {
	ID        string
	FirstName string
	LastName  string
	CG        float32
	Credits   int
	Courses   []string
}

// Book struct
type Book struct {
	ID     string
	Title  string
	Isbn   string
	Author *Author
}

// Author struct
type Author struct {
	ID        string
	FirstName string
	LastName  string
}

// Rectangle struct
type Rectangle struct {
	length int
	width  int
}

// Person struct
type Person struct {
	Name string
}

// This would work, but we would rather say an Android is a Person, rather than an Android has a Person.
// Go supports relationships like this by using an embedded type. Also known as anonymous fields, embedded types look like this

// type Android struct {
// 	Person Person
// 	Model  string
// }

// Android struct
type Android struct {
	Person
	Model string
}

func (p *Person) talk() {
	fmt.Println("Hi, my name is", p.Name)
}

func areaWithoutMethod(r *Rectangle) int {
	return r.length * r.width
}

func (r *Rectangle) areaWithMethod() int {
	return r.length * r.width
}

type geometry interface {
	area() float64
	perimeter() float64
}

type circle struct {
	radius float64
}

type square struct {
	side float64
}

func (c *circle) area() float64 {
	return c.radius * c.radius * math.Pi
}
func (c *circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (s *square) area() float64 {
	return s.side * s.side
}
func (s *square) perimeter() float64 {
	return 4 * s.side
}

func main() {
	// create a list of courses & then assign or we can assign in place like line 23
	sub := []string{"maths", "cse"}
	fmt.Println(sub)
	//student1 := Student{ID: "123", FirstName: "Ishtiaq", LastName: "Islam", CG: 3.5, Credits: 130, Courses: sub}
	student1 := Student{ID: "123", FirstName: "Ishtiaq", LastName: "Islam", CG: 3.5, Credits: 130, Courses: []string{"maths", "cse"}}

	// %v prints only the values, if we need the name of the fields, we can use %+v to achieve that
	fmt.Printf("%+v\n", student1)
	fmt.Println(student1)

	// declaring an empty slice of Student type, notice the opening & closing curly braces at the end
	studentsList := []Student{}
	studentsList = append(studentsList, Student{ID: "123", FirstName: "Ishtiaq", LastName: "Islam", CG: 3.5, Credits: 130, Courses: []string{"maths", "cse"}})
	studentsList = append(studentsList, Student{ID: "234", FirstName: "Misir", LastName: "Ali", CG: 3.6, Credits: 120, Courses: []string{"english", "cse"}})
	fmt.Println(studentsList)

	// we can access the field values using the dot(.) operator, courses is a list or slice, so print like this.
	for index, student := range studentsList {
		fmt.Println(index, student.FirstName, student.LastName, student.CG)
		for _, course := range student.Courses {
			fmt.Print(course, " ")
		}
		fmt.Println()
	}

	// var studentsArray []Student
	// studentsArray = []Student or the one below. an array of Student Items
	studentsArray := []Student{
		{ID: "123", FirstName: "Ishtiaq", LastName: "Islam", CG: 3.5, Credits: 130, Courses: []string{"maths", "cse"}},
		{ID: "234", FirstName: "Misir", LastName: "Ali", CG: 3.6, Credits: 120, Courses: []string{"english", "cse"}},
	}

	fmt.Println(studentsArray)

	// other ways of declaration
	var st Student
	fmt.Println(st.FirstName, st.LastName, st.Courses, st.CG, st.Credits)

	// using thie new()
	// This allocates memory for all the fields, sets each of them to their zero value and returns a pointer.
	// (*Student) More often we want to give each of the fields a value. We can do this in two ways. Like this: stdnt := Student{fieldName: initialValue}
	stdnt := new(Student)
	fmt.Println(stdnt.FirstName, stdnt.LastName, stdnt.Courses, stdnt.CG, stdnt.Credits)

	// Book - data, author refernce to another struct Author
	book1 := Book{ID: "1", Title: "book-1", Isbn: "1354", Author: &Author{ID: "12", FirstName: "kazi", LastName: "Nazrul"}}
	fmt.Println(book1.ID, book1.Title, book1.Isbn, book1.Author.FirstName, book1.Author.LastName, book1.Author.ID)

	booksArray := []Book{
		{ID: "123", Title: "book-1", Isbn: "12345", Author: &Author{ID: "12", FirstName: "kazi", LastName: "Nazrul"}},
		{ID: "234", Title: "book-2", Isbn: "56789", Author: &Author{ID: "34", FirstName: "rabindranath", LastName: "tagore"}},
	}
	fmt.Println(booksArray)

	for index, book := range booksArray {
		fmt.Println(index, book.ID, book.Title, book.Isbn, book.Author.ID, book.Author.FirstName, book.Author.LastName)
	}

	// Methods..
	rectangle1 := Rectangle{length: 2, width: 3}
	fmt.Printf("%+v\n", rectangle1)
	fmt.Println("using ref.: ", areaWithoutMethod(&rectangle1))
	// the better alternative of doing this is by using Method like this.
	// In between the keyword func and the name of the function we've added a “receiver”. The receiver is like a parameter – it
	// has a name and a type – but by creating the function in this way it allows us to call the function using the dot (.) operator:
	fmt.Println("using method: ", rectangle1.areaWithMethod())
	// This is much easier to read, we no longer need the & operator (Go automatically knows to pass a pointer to the circle for this method)
	// and because this function can only be used with Circles we can rename the function to just area

	// Embedded Types, description is on the top of structs
	a := Android{Model: "1000", Person: Person{Name: "Ishtiaq"}}
	a.talk()
	a.Person.talk()

	// Intefaces : Interfaces are named collections of method signatures.
	c := circle{radius: 5}
	fmt.Println(c.area(), c.perimeter())

	s := square{side: 5}
	fmt.Println(s.area(), s.perimeter())
}
