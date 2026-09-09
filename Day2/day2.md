Variables and Data Types:

In Go, there are different types of variables, for example:

int- stores integers (whole numbers), such as 123 or -123
float32- stores floating point numbers, with decimals, such as 19.99 or -19.99
string - stores text, such as "Hello World". String values are surrounded by double quotes
bool- stores values with two states: true or false

declaring variables:
 with var keyword-use var keyword followed by variable name and type
Syntax:
var variablename type=value

with := sign -use the ":= "sign followed by value

Syntax:
variable name:=value
- In this case, the type of the variable is inferred from the value (means that the compiler decides the type of the variable, based on the value).
- It is not possible to declare a variable using :=, without assigning a value to it.

declaring multiple variables on same line

If you use the type keyword, it is only possible to declare one type of variable per line
ex: 
package main
import "fmt"

func main() {
  var a, b, c, d int = 1, 3, 5, 7

  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
  fmt.Println(d)
}

If the type keyword is not specified, you can declare different types of variables on the same line:
package main
import "fmt"

func main() {
  var a, b = 6, "Hello"
  c, d := 7, "World!"

  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
  fmt.Println(d)
}
Go Variable Declaration in a Block
Multiple variable declarations can also be grouped together into a block for greater readability:
package main
import "fmt"

func main() {
   var (
     a int
     b int = 1
     c string = "hello"
   )

  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(c)
}

Go Constants
If a variable should have a fixed value that cannot be changed, you can use the const keyword.
The const keyword declares the variable as "constant", which means that it is unchangeable and read-only.

Syntax
const CONSTNAME type = value
Note: The value of a constant must be assigned when you declare it.
ex:
package main
import "fmt"

const PI = 3.14

func main() {
  fmt.Println(PI)
}

Constant Rules:
Constant names follow the same naming rules as variables
Constant names are usually written in uppercase letters (for easy identification and differentiation from variables)
Constants can be declared both inside and outside of a function

Constant Types:
There are two types of constants:

Typed constants
Untyped constants

Typed Constants:
Typed constants are declared with a defined type:

Example
package main
import "fmt"

const A int = 1

func main() {
  fmt.Println(A)
}

Untyped Constants:
Untyped constants are declared without a type:

Example:
package main
import "fmt"

const A = 1

func main() {
  fmt.Println(A)
}
In this case, the type of the constant is inferred from the value (means the compiler decides the type of the constant, based on the value).



Go Data Types:
Data type is an important concept in programming. Data type specifies the size and type of variable values.
Go is statically typed, meaning that once a variable type is defined, it can only store data of that type.



bool: represents a boolean value and is either true or false
      The default value of a boolean data type is false.Boolean values are mostly used for conditional testing .
Numeric: represents integer types, floating point values, and complex types
Integer data types are used to store a whole number without decimals, like 35, -50, or 1345000.
The integer data type has two categories:

Signed integers - can store both positive and negative values
Unsigned integers - can only store non-negative values
The default type for integer is int. If you do not specify a type, the type will be int.

Float:The float data types are used to store positive and negative numbers with a decimal point, like 35.3, -2.34, or 3597.34987.

The float data type has two keywords:
float32 and float64
depends on its range.The default type for float is float64. If you do not specify a type, the type will be float64. 

string: represents a string value
The string data type is used to store a sequence of characters (text). String values must be surrounded by double quotes:

Example:

package main
import "fmt"

func main() {
  var a bool = true     // Boolean
  var b int = 5         // Integer
  var c float32 = 3.14  // Floating point number
  var d string = "Hi!"  // String

  fmt.Println("Boolean: ", a)
  fmt.Println("Integer: ", b)
  fmt.Println("Float:   ", c)
  fmt.Println("String:  ", d)
}


TYPE CONVERSION:
In Go, type conversion is explicit and required because the language is statically typed;
it uses the syntax T(value) to convert a value value to type T.

Numeric Types: You can convert between compatible numeric types (e.g., int, float64, uint). Widening conversions (e.g., int to int64) are safe, but narrowing conversions (e.g., float64 to int) truncate decimals and may lose data.
var i int = 42
var f float64 = float64(i) // int to float64
var u uint = uint(f)       // float64 to uint

Strings and Bytes: Converting a string to []byte copies the UTF-8 bytes, while converting []byte to string creates a new string. Converting to []rune decodes UTF-8 into Unicode code points.
s := "hello"
b := []byte(s)      // string to []byte
s2 := string(b)     // []byte to string
r := []rune(s)      // string to []rune

Custom/Named Types: You can convert between a named type and its underlying type (e.g., MyInt and int), or between two named types if they share the same underlying type.
type MyInt int
var i int = 10
var mi MyInt = MyInt(i) // int to MyInt

String/Number Conversion: To convert numbers to strings or vice versa, use the strconv package, as direct conversion like string(42) produces a Unicode character, not a numeric string.
import "strconv"
numStr := strconv.Itoa(42)   // int to string "42"
num, err := strconv.Atoi("42") // string to int 42


