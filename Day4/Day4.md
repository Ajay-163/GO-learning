
IF statement:
if statement used to specify a block of Go code to be executed if a condition is true

Syntax:
if condition {
  // code to be executed if condition is true
}
 example:
package main
import "fmt"

func main() {
  if 20 > 18 {
    fmt.Println("20 is greater than 18")
  }
}

IF-ELSE statement:
it is used to specify a block of code to be executed if the condition is false.

Syntax:
if condition {
  // code to be executed if condition is true
} else {
  // code to be executed if condition is false
}
example:
package main
import "fmt"

func main() {
  temperature := 14
  if (temperature > 15) {
    fmt.Println("It is warm out there")
  } else {
    fmt.Println("It is cold out there")
  }
}

ELSE-IF statement:
it is used to specify a new condition if the first condition is false.
Syntax:
f condition1 {
   // code to be executed if condition1 is true
} else if condition2 {
   // code to be executed if condition1 is false and condition2 is true
} else {
   // code to be executed if condition1 and condition2 are both false
}

example:
package main
import "fmt"

func main() {
  time := 22
  if time < 10 {
    fmt.Println("Good morning.")
  } else if time < 20 {
    fmt.Println("Good day.")
  } else {
    fmt.Println("Good evening.")
  }
}


COMBINING conditions:
 these are like using logical operators in conditional statements to retrive the results.
like AND-&&
     OR-||
     NOT-!


