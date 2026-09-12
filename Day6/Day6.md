
LOOPS in GO:
The only loop available in GO id FOR loop.

FOR LOOP:
The for loop loops through a block of code a specified number of times.
Each execution of a loop is called an iteration.

The for loop can take up to three statements:
1.Intialization:Initializes the loop counter value.

2.Condition:Evaluated for each loop iteration. If it evaluates to TRUE, the loop continues. 
If it evaluates to FALSE, the loop ends.

3.Increment/Decrement: Increases/Decreases the loop counter value.

Syntax:
for statement1; statement2; statement3 {
   // code to be executed for each iteration
}

The CONTINUE Statement:
The continue statement is used to skip one or more iterations in the loop. It then continues with the next iteration in the loop.
example:

package main
import "fmt"
func main() {
  for i:=0; i < 5; i++ {
    if i == 3 {
      continue		//skips the value of 3
    }
   fmt.Println(i)
  }
}


The BREAK Statement:
The break statement is used to break/terminate the loop execution.

Example:

package main
import "fmt"

func main() {
  for i:=0; i < 5; i++ {
    if i == 3 {
      break		//breaks out of the loop when i is equal to 3
    }
   fmt.Println(i)
  }
}

*continue and break are usually used with conditions.

Nested Loops:
It is possible to place a loop inside another loop.
*the "inner loop" will be executed one time for each iteration of the "outer loop":

Example
package main
import "fmt"

func main() {
  adj := [2]string{"big", "tasty"}
  fruits := [3]string{"apple", "orange", "banana"}
  for i:=0; i < len(adj); i++ {
    for j:=0; j < len(fruits); j++ {
      fmt.Println(adj[i],fruits[j])
    }
  }
}

The RANGE Keyword:
The range keyword is used to more easily iterate through the elements of an array, slice or map. It returns both the index and the value.

Syntax:
for index, value := range array|slice|map {
   // code to be executed for each iteration
}

Example:
package main
import "fmt"

func main() {
  fruits := [3]string{"apple", "orange", "banana"}
  for idx, val := range fruits {
     fmt.Printf("%v\t%v\n", idx, val)	//prints both the indexes and the values at each (idx stores the index, val stores the value):
  }
}

* an underscore (_) can be used when we don't want to print any values of declared variables.

