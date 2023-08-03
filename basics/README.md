## Intro
Go is a general-purpose language designed with systems programming in mind. It is strongly typed and garbage-collected and has explicit support for concurrent programming. Programs are constructed from packages, whose properties allow efficient management of dependencies.

## Keywords
```
break        default      func         interface    select
case         defer        go           map          struct
chan         else         goto         package      switch
const        fallthrough  if           range        type
continue     for          import       return       var
```
## Operators and punctuation
```
+    &     +=    &=     &&    ==    !=    (    )
-    |     -=    |=     ||    <     <=    [    ]
*    ^     *=    ^=     <-    >     >=    {    }
/    <<    /=    <<=    ++    =     :=    ,    ;
%    >>    %=    >>=    --    !     ...   .    :
     &^          &^=          ~
```

## Identifiers
Identifiers name program entities such as variables and types. An identifier is a sequence of one or more letters (underscore character _ (U+005F) is considered a lowercase letter), unicode letters and digits. The first character in an identifier must be a letter.

**Predeclared Identifiers:**
```
Types:
	any bool byte comparable
	complex64 complex128 error float32 float64
	int int8 int16 int32 int64 rune string
	uint uint8 uint16 uint32 uint64 uintptr

Constants:
	true false iota

Zero value:
	nil

Functions:
	append cap close complex copy delete imag len
	make new panic print println real recover
```

**Blank Identifier:**

The blank identifier is represented by the underscore character _. It serves as an anonymous placeholder instead of a regular (non-blank) identifier and has special meaning in declarations, as an operand, and in assignment statements.

**Exported Identifier:**

An identifier may be exported to permit access to it from another package. An identifier is exported if both:

1. the first character of the identifier's name is a Unicode uppercase letter (Unicode character category Lu); and
2. the identifier is declared in the package block or it is a field name or method name.

All other identifiers are not exported.

**Uniqueness of identifiers:**

Given a set of identifiers, an identifier is called unique if it is different from every other in the set. Two identifiers are different if they are spelled differently, or if they appear in different packages and are not exported. Otherwise, they are the same.

## Variables
A variable is a storage location for holding a value. The set of permissible values is determined by the variable's type.
- Static Type Variable - the type declared or inferred at compile time
- Dynamic Type Variable - the actual type of the value assigned to the variable at runtime (unless the value is the predeclared identifier `nil`, which has no type).

Zero Value (Default Value): `false` for `booleans`, `0` for `numeric` types, `""` for `strings`, and `nil` for `pointers`, `functions`, `interfaces`, `slices`, `channels`, and `maps`. This initialization is done recursively, so for instance each element of an array of structs will have its fields zeroed if no value is specified.

There are 3 ways to assign values to a variable.
- Method 1: `var x int = 10`
- Method 2: `var x = 10`
</br>(compiler automatically detects the type)
- Method 3: `x := 10`
</br>(shorthand notation to assign values to variables)

## Types
A type determines a set of values together with operations and methods specific to those values. A type may be denoted by a type name, if it has one, which must be followed by type arguments if the type is generic. A type may also be specified using a type literal, which composes a type from existing types.

- **Type Name and Type Arguments:** a type name refers to the name given to a specific type. If a type is generic, it may be parameterized with type arguments. This allows the type to be reused with different underlying data types. Let's consider an example:</br> 
`type Stack[T any] []T`</br>
In this example, `Stack` is the type name, and `[T any]` represents the type argument. The `Stack` type is generic and can be instantiated with various data types. For instance, you can create a stack of integers (`Stack[int]`), a stack of strings (`Stack[string]`), or even a stack of custom types.

- **Type Literal:** a type literal is used to compose a new type by combining existing types. It allows you to define a type that has specific properties or characteristics based on other types. Here's an example:
```go
type Celsius float64

type Temperature struct {
    Value Celsius
    Unit  string
}
```

The language predeclares certain type names. Others are introduced with type declarations or type parameter lists. Composite types—array, struct, pointer, function, interface, slice, map, and channel types—may be constructed using type literals.

Predeclared types, defined types, and type parameters are called named types. An alias denotes a named type if the type given in the alias declaration is a named type.

### Type declarations
A type declaration binds an identifier, the type name, to a type. Type declarations come in two forms: alias declarations and type definitions.

- **Alias declarations:** An alias declaration binds an identifier to the given type.
- **Type definitions:** A type definition creates a new, distinct type with the same underlying type and operations as the given type and binds an identifier, the type name, to it. The new type is called a defined type. It is different from any other type, including the type it is created from.

### Type parameter declarations
A type parameter list declares the type parameters of a generic function or type declaration. The type parameter list looks like an ordinary function parameter list except that the type parameter names must all be present and the list is enclosed in square brackets rather than parentheses.

### Types (need to complete)
- **Boolean Type:** A boolean type represents the set of Boolean truth values denoted by the predeclared constants `true` and `false`. The predeclared boolean type is `bool`; it is a defined type.
- **Numeric Types:** An integer, floating-point, or complex type represents the set of integer, floating-point, or complex values, respectively. They are collectively called numeric types.
```go
uint8       the set of all unsigned  8-bit integers (0 to 255)
uint16      the set of all unsigned 16-bit integers (0 to 65535)
uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)

int8        the set of all signed  8-bit integers (-128 to 127)
int16       the set of all signed 16-bit integers (-32768 to 32767)
int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

float32     the set of all IEEE-754 32-bit floating-point numbers
float64     the set of all IEEE-754 64-bit floating-point numbers

complex64   the set of all complex numbers with float32 real and imaginary parts
complex128  the set of all complex numbers with float64 real and imaginary parts

byte        alias for uint8
rune        alias for int32

// there is also a set of predeclared integer types with implementation-specific sizes:

uint     either 32 or 64 bits
int      same size as uint
uintptr  an unsigned integer large enough to store the uninterpreted bits of a pointer value

// Explicit conversions are required when different numeric types are mixed in an expression or assignment. For instance, int32 and int are not the same type even though they may have the same size on a particular architecture.
```
- **String Types:** A string type represents the set of string values. A string value is a (possibly empty) sequence of bytes. The number of bytes is called the length of the string and is never negative. Strings are immutable: once created, it is impossible to change the contents of a string. The predeclared string type is string; it is a defined type. If s[i] is the i'th byte of a string, &s[i] is invalid.

- **Array Types:** An array is a numbered sequence of elements of a single type, called the element type. The number of elements is called the length of the array and is never negative. Array types are always one-dimensional but may be composed to form multi-dimensional types.
- **Slice Types:** A slice is a descriptor for a contiguous segment of an underlying array and provides access to a numbered sequence of elements from that array. A slice type denotes the set of all slices of arrays of its element type. The number of elements is called the length of the slice and is never negative. The value of an uninitialized slice is `nil`. The slice index of a given element may be less than the index of the same element in the underlying array.</br>
A slice, once initialized, is always associated with an underlying array that holds its elements. A slice therefore shares storage with its array and with other slices of the same array; by contrast, distinct arrays always represent distinct storage.</br>
The array underlying a slice may extend past the end of the slice. The capacity is a measure of that extent: it is the sum of the length of the slice and the length of the array beyond the slice; a slice of length up to that capacity can be created by slicing a new one from the original slice. The capacity of a slice a can be discovered using the built-in function `cap(a)`.</br>
A new, initialized slice value for a given element type `T` may be made using the built-in function `make`, which takes a slice type and parameters specifying the length and optionally the capacity. A slice created with make always allocates a new, hidden array to which the returned slice value refers. That is, executing `make([]T, length, capacity)`. Produces the same slice as allocating an array and slicing it, so these two expressions are equivalent: `make([]int, 50, 100)` and `new([100]int)[0:50]`. Like arrays, slices are always one-dimensional but may be composed to construct higher-dimensional objects. With arrays of arrays, the inner arrays are, by construction, always the same length; however with slices of slices (or arrays of slices), the inner lengths may vary dynamically. Moreover, the inner slices must be initialized individually.
- **Struct Types:** A struct is a sequence of named elements, called fields, each of which has a name and a type. Field names may be specified explicitly (IdentifierList) or implicitly (EmbeddedField). Within a struct, non-blank field names must be unique. A field declared with a type but no explicit field name is called an embedded field. An embedded field must be specified as a type name T or as a pointer to a non-interface type name *T, and T itself may not be a pointer type. The unqualified type name acts as the field name.
- **Pointer Types:** A pointer type denotes the set of all pointers to variables of a given type, called the base type of the pointer. The value of an uninitialized pointer is `nil`.
- **Function Type:** A function type denotes the set of all functions with the same parameter and result types. The value of an uninitialized variable of function type is nil.
</br>Within a list of parameters or results, the names (IdentifierList) must either all be present or all be absent. If present, each name stands for one item (parameter or result) of the specified type and all non-blank names in the signature must be unique. If absent, each type stands for one item of that type. Parameter and result lists are always parenthesized except that if there is exactly one unnamed result it may be written as an unparenthesized type.
</br>The final incoming parameter in a function signature may have a type prefixed with `...`. A function with such a parameter is called variadic and may be invoked with zero or more arguments for that parameter.
- **Interface Types:** An interface type defines a type set. A variable of interface type can store a value of any type that is in the type set of the interface. Such a type is said to implement the interface. The value of an uninitialized variable of interface type is `nil`.
</br>An interface type is specified by a list of interface elements. An interface element is either a method or a type element, where a type element is a union of one or more type terms. A type term is either a single type or a single underlying type.
  - **Basic Interfaces:** In its most basic form an interface specifies a (possibly empty) list of methods. The type set defined by such an interface is the set of types which implement all of those methods, and the corresponding method set consists exactly of the methods specified by the interface. Interfaces whose type sets can be defined entirely by a list of methods are called basic interfaces.
  - **Embedded Interfaces:** In a slightly more general form an interface T may use a (possibly qualified) interface type name E as an interface element. This is called embedding interface E in T. 
  - **General Interfaces:** In their most general form, an interface element may also be an arbitrary type term T, or a term of the form ~T specifying the underlying type T, or a union of terms t1|t2|…|tn. 
- **Map Types:** A map is an unordered group of elements of one type, called the element type, indexed by a set of unique keys of another type, called the key type. The value of an uninitialized map is nil.
- **Channel Types:** A channel provides a mechanism for concurrently executing functions to communicate by sending and receiving values of a specified element type. The value of an uninitialized channel is nil.

## I/O in Go
- Using `fmt` package
- Using the bufio.NewReader(os.Stdin)
- File I/O

## Type Casting
Golang provides various predefined functions like `int()`, `float32()`, `string()`, etc to perform explicit type casting.
```go
package main
import "fmt"

func main() {
  
  var floatValue float32 = 5.45

  // type conversion from float to int
  var intValue int = int(floatValue)
 
  fmt.Printf("Float Value is %f\n", floatValue)
  fmt.Printf("Integer Value is %d", intValue)
}
```
 Go doesn't support implicit type casting. 
 ```go
 package main
import "fmt"

func main() {
  
  // initialize integer variable to a floating-point number
  var number int = 4.34 // Error

  fmt.Printf("Number is %f", number)
}
 ```