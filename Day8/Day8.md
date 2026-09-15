
FUNCTONS IN GO:
A function is a block of statements that can be used repeatedly in a program.
A function will not execute automatically when a page loads.
A function will be executed by a call to the function.

Function declaration:
Use the func keyword.
Specify a name for the function, followed by parentheses ().
Finally, add code that defines what the function should do, inside curly braces {}.

Syntax:
func FunctionName() {
  // code to be executed
}

Function calling:
Functions are not executed immediately. They are "saved for later use", and will be executed when they are called.
The opening curly brace ( { ) indicates the beginning of the function code, and the closing curly brace ( } ) indicates the end of the function.
To call the function, just write its name followed by two parentheses ().

example:
package main
import "fmt"

func myMessage() {
  fmt.Println("I just got executed!")
}

func main() {
  myMessage() // call the function
}

*A function can be called multiple times.

Naming Rules for Go Functions
-A function name must start with a letter.
-A function name can only contain alpha-numeric characters and underscores (A-z, 0-9, and _ ).
-Function names are case-sensitive.
-A function name cannot contain spaces.

FUNCTION PARAMETERS:

Information can be passed to functions as a parameter. Parameters act as variables inside the function.
Parameters and their types are specified after the function name, inside the parentheses. You can add as many parameters as you want, just separate them with a comma:
-When a parameter is passed to the function, it is called an argument.
Syntax
func FunctionName(param1 type, param2 type, param3 type) {
  // code to be executed
}

Example
package main
import "fmt"

func familyName(fname string) {
  fmt.Println("Hello", fname, "Refsnes")
}

func main() {
  familyName("Liam")			Hello Liam Refsnes
  familyName("Jenny")			Hello Jenny Refsnes			
  familyName("Anja")			Hello Anja Refsnes
}

Multiple Parameters
Inside the function, you can add as many parameters as you want:
- When you are working with multiple parameters, the function call must have the same number of arguments as there are parameters, 
  and the arguments must be passed in the same order.
Example
package main
import "fmt"

func familyName(fname string, age int) {
  fmt.Println("Hello", age, "year old", fname, "Refsnes")
}

func main() {
  familyName("Liam", 3)
  familyName("Jenny", 14)
  familyName("Anja", 30)
}

FUNCTION RETURNS:
If we want the function to return a value, you need to define the data type of the return value (such as int, string, etc), 
and also use the return keyword inside the function:

Syntax:
func FunctionName(param1 type, param2 type) type {
  // code to be executed
  return output
}

example:
package main
import "fmt"

func myFunction(x int, y int) int {
  return x + y
}

func main() {
  fmt.Println(myFunction(1, 2))
}

Named Return Values:
we can name the return values of a function.
Example
we name the return value as result (of type int), and return the value with a naked return 
(means that we use the return statement without specifying the variable name):

package main
import "fmt"

func myFunction(x int, y int) (result int) {
  result = x + y
  return result
}

func main() {
  fmt.Println(myFunction(1, 2))
}

Multiple Return Values:
Go functions can also return multiple values.
Example:
myFunction() returns one integer (result) and one string (txt1):

package main
import "fmt"

func myFunction(x int, y string) (result int, txt1 string) {
  result = x + x
  txt1 = y + " World!"
  return
}

func main() {
  fmt.Println(myFunction(5, "Hello"))
}

