package main

import "fmt"

// named struct
type employee struct {
	firstName string
	lastName  string
	id        int
}

// can't use same type twice
type person struct {
	string
	int
}

// Exported struct (can access from different pkg)
type Spec struct {
	Price int    // exported field
	Maker string // exported field
	model string // unexported field
}

// nested struct
type address struct {
	city  string
	state string
}

type personInfo struct {
	name string
	age  int
	add  address
}

// embedded struct
type person_Info struct {
	name string
	age  int
	address
}

func main() {
	// creating struct specifying field names
	emp1 := employee{
		firstName: "mno",
		lastName:  "pqr",
		id:        10028,
	}

	// creating struct without specifying field names
	// it is necessary to maintain the order of fields to be the same as specified in the struct declaration
	emp2 := employee{"xyz", "abc", 10029}
	fmt.Println(emp1, emp2)

	// anonymous struct
	emp3 := struct {
		firstName string
		lastName  string
		id        int
		age       int
	}{
		firstName: "abc",
		lastName:  "xyz",
		id:        10030,
		age:       24,
	}
	fmt.Println(emp3)

	// accessing individual fields of a struct
	fmt.Println(emp3.firstName)
	emp3.age = 28
	fmt.Println(emp3)

	// zero value of a struct
	var emp4 employee
	fmt.Println(emp4)

	// it is also possible to specify values for some fields and ignore the rest
	// the ignored fields are assigned zero values
	emp5 := employee{
		firstName: "kkk",
		lastName:  "lll",
	}
	fmt.Println(emp5)

	// pointer struct
	emp6 := &employee{
		firstName: "ppp",
		lastName:  "ooo",
	}
	fmt.Println(emp6.firstName) // equals to (*emp6).firstName

	// if emp6 is nil then this will return a panic
	// var emp7 *employee
	// fmt.Println(emp7.firstName)

	// anonymous fields
	p := person{
		string: "ppp",
		int:    90,
	}
	fmt.Println(p.string)
	fmt.Println(p.int)

	// struct equality
	// Structs are value types and are comparable if each of their fields are comparable.
	// Two struct variables are considered equal if their corresponding fields are equal.
	if emp1 != emp2 {
		fmt.Println("Not Equal")
	} else {
		fmt.Println("Equal")
	}

	// embedded struct
	pInfo := person_Info{
		name: "xyz",
		age:  90,
		address: address{
			city:  "okl",
			state: "spo",
		},
	}
	fmt.Println(pInfo.city) // equals to pInfo.address.city
	// we can also access the methods of address in the same way

	// Methods
	// A method is just a function with a special receiver type between the func keyword and the method name.
	// The receiver can either be a struct type or non-struct type.
	// func (t type) methodName(params list) {}
	t := newType("bcde") // try with "abcde" --> true
	fmt.Println(t.containsA())

	// Value Reciever vs Pointer Reciever
	c := car{
		spec:  "omom",
		name:  "xyz",
		price: 9000,
	}
	fmt.Println(c)
	c.changeName("popo")
	fmt.Println(c)
	c.changePrice(5000) // same as (&c).changePrice(5000)
	fmt.Println(c)
}

type newType string

func (s newType) containsA() bool {
	for i := 0; i < len(s); i++ {
		if s[i] == 'a' {
			return true
		}
	}
	return false
}

type car struct {
	spec  string
	name  string
	price int
}

func (c car) changeName(newName string) {
	c.name = newName
}

func (c *car) changePrice(newPrice int) {
	c.price = newPrice
}
