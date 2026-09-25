Structures:
A struct (short for structure) is used to create a collection of members of different data types, into a single variable.
While arrays are used to store multiple values of the same data type into a single variable, structs are used to store 
multiple values of different data types into a single variable.
A struct can be useful for grouping data together to create records.

Struct intialization:
To declare/intialize a structure in Go, use the type and struct keywords:

Syntax:
type struct_name struct {
  member1 datatype;
  member2 datatype;
  member3 datatype;
  ...
}
Example:
 we declare a struct type Person with the following members: name, age, job and salary:

type Person struct {
  name string
  age int
  job string
  salary int
}

Access Struct Members:
To access any member of a structure, use the dot operator (.) between the structure variable name and the structure member:

Example:
package main
import "fmt"

type Person struct {
  name string
  age int
  job string
  salary int
}

func main() {
  var pers1 Person
  var pers2 Person

  // Pers1 specification
  pers1.name = "Hege"
  pers1.age = 45
  pers1.job = "Teacher"
  pers1.salary = 6000

  // Pers2 specification
  pers2.name = "Cecilie"
  pers2.age = 24
  pers2.job = "Marketing"
  pers2.salary = 4500

  // Access and print Pers1 info
  fmt.Println("Name: ", pers1.name)
  fmt.Println("Age: ", pers1.age)
  fmt.Println("Job: ", pers1.job)
  fmt.Println("Salary: ", pers1.salary)

  // Access and print Pers2 info
  fmt.Println("Name: ", pers2.name)
  fmt.Println("Age: ", pers2.age)
  fmt.Println("Job: ", pers2.job)
  fmt.Println("Salary: ", pers2.salary)
}

Pass Struct as Function Arguments
We can also pass a structure as a function argument, like this:

Example
package main
import "fmt"

type Person struct {
  name string
  age int
  job string
  salary int
}

func main() {
  var pers1 Person

  // Pers1 specification
  pers1.name = "Hege"
  pers1.age = 45
  pers1.job = "Teacher"
  pers1.salary = 6000

  // Print Pers1 info by calling a function
  printPerson(pers1)
}

func printPerson(pers Person) {
  fmt.Println("Name: ", pers.name)
  fmt.Println("Age: ", pers.age)
  fmt.Println("Job: ", pers.job)
  fmt.Println("Salary: ", pers.salary)
}

Struct with slice:
A struct with a slice is a struct where one of its fields stores multiple values using a slice.
This is very useful when one object needs to store multiple values.
example:
type Student struct {
    name  string
    marks []int
}



