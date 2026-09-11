SWITCH case:

statement:SWITCH  is used to select one of many code blocks to be executed.
The difference is that it only runs the matched case so it does not need a break statement.
Switch has two types:
1.Single-case switch
2.Multi-case switch

Single-case switch:
The expression is evaluated once,
The value of the switch expression is compared with the values of each case,
If there is a match, the associated block of code is executed,
The default keyword is optional. It specifies some code to run if there is no case match.

Syntax:
switch expression {
case x:
   // code block
case y:
   // code block
case z:
...
default:
   // code block
}

example:
package main   
import "fmt" 

func main() {
  var day = 2
  switch day{
  
case
(1):
    fmt.Print("Saturday")
  
case
(2):
    fmt.Print("Sunday")    
  }
}


Multi-case switch:
It is possible to have multiple values for each case in the switch statement.

Syntax:
switch expression {
case x,y:
   // code block if expression is evaluated to x or y
case v,w:
   // code block if expression is evaluated to v or w
case z:
...
default:
   // code block if expression is not found in any cases
}

package main
import "fmt"

func main() {
   n := 5

   switch n {
   case 1,3,5:
    fmt.Println("odd numbers")
   case 2,4:
     fmt.Println("Even number")
  default:
    fmt.Println("Invalid  number")
  }
}



