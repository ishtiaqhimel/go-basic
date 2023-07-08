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
- Boolean Type
- Numeric Types
- String Types
- Array Types
- Slice Types
- Struct Types
- Pointer Types
- Interface Types
  - Basic Interfaces
  - Embedded Interfaces
  - General Interfaces
- Map Types
- Channel Types

## Data Types
The basic data types in Golang are $-$
| Data Types | Description | Variations |
|:--- |:--- |:---|
| `int` | Integer numbers | int8, int16, int32, int64, int (signed integers), uint8, uint16, uint32, uint32, uint64, uint, uintptr (unsigned integers) |
| `float` | Numbers with decimal points | float32, float64 |
| `complex` | Complex numbers | complex64, complex128 |
| `string` | Sequence of characters (a slice of bytes in Go) |  |
| `bool` | Either true or false | |
| `byte` | A byte (8 bits) of non-negative integers (alias for uint8) | |
| `rune` | Used for characters. Internally used as 32-bit integers (alias for int32) |  |

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