Arrays:
Arrays are used to store multiple values of the same type in a single variable, instead of declaring separate variables for each value.

Array declaring:
there are two ways to declare an array:

1. With the var keyword:
Syntax
var array_name = [length]datatype{values} // here length is defined

or

var array_name = [...]datatype{values} // here length is inferred

2. With the := sign:
Syntax
array_name := [length]datatype{values} // here length is defined

or

array_name := [...]datatype{values} // here length is inferred


* The length specifies the number of elements to store in the array. In Go, arrays have a fixed length. The length of the array is either defined by a number or is inferred (means that the compiler decides the length of the array, based on the number of values).

example:

package main
import "fmt"

func main() {
  var arr1 = [3]int{1,2,3}
  arr2 := [5]int{4,5,6,7,8}

  fmt.Println(arr1)
  fmt.Println(arr2)
}

*we can access a specific array element by referring to the index number.
 array indexes start at 0. That means that [0] is the first element, [1] is the second element, etc.
 also change the value of a specific array element by referring to the index number.

Array Initialization:
If an array or one of its elements has not been initialized in the code, it is assigned the default value of its type.

Tip: The default value for int is 0, and the default value for string is "".

Example:

package main
import "fmt"

func main() {
  arr1 := [5]int{} 		//not initialized
  arr2 := [5]int{1,2} 		//partially initialized
  arr3 := [5]int{1,2,3,4,5} 	//fully initialized

  fmt.Println(arr1)
  fmt.Println(arr2)
  fmt.Println(arr3)
}

*The len() function is used to find the length of an array:

Slices:
Slices are similar to arrays, but are more powerful and flexible.
Like arrays, slices are also used to store multiple values of the same type in a single variable.

In Go, there are several ways to create a slice:
*Using the []datatype{values} format
*Create a slice from an array
*Using the make() function

Create a Slice With []datatype{values}:

Syntax
slice_name := []datatype{values}
A common way of declaring a slice is like this:

myslice := []int{}

The code above declares an empty slice of 0 length and 0 capacity.
To initialize the slice during declaration, use this:
myslice := []int{1,2,3}

The code above declares a slice of integers of length 3 and also the capacity of 3.

In Go, there are two functions that can be used to return the length and capacity of a slice:

len() function - returns the length of the slice (the number of elements in the slice)
cap() function - returns the capacity of the slice (the number of elements the slice can grow or shrink to)
Example:
This example shows how to create slices using the []datatype{values} format:

package main
import "fmt"

func main() {
  myslice1 := []int{}
  fmt.Println(len(myslice1))
  fmt.Println(cap(myslice1))
  fmt.Println(myslice1)

  myslice2 := []string{"Go", "Slices", "Are", "Powerful"}
  fmt.Println(len(myslice2))
  fmt.Println(cap(myslice2))
  fmt.Println(myslice2)
}

Create a Slice From an Array:
we can create a slice by slicing an array:

Syntax:

var myarray = [length]datatype{values} 		// An array
myslice := myarray[start:end] 			// A slice made from the array

Example:

package main
import "fmt"

func main() {
  arr1 := [6]int{10, 11, 12, 13, 14,15}
  myslice := arr1[2:4]

  fmt.Printf("myslice = %v\n", myslice)			//myslice = [12 13]
  fmt.Printf("length = %d\n", len(myslice))		//length = 2
  fmt.Printf("capacity = %d\n", cap(myslice))		//capacity = 4
}

In the example above myslice is a slice with length 2. It is made from arr1 which is an array with length 6.

The slice starts from the third element of the array which has value 12 (remember that array indexes start at 0. That means that [0] is the first element, [1] is the second element, etc.). The slice can grow to the end of the array. This means that the capacity of the slice is 4.
If myslice started from element 0, the slice capacity would be 6.


Create a Slice With The make() Function:
The make() function can also be used to create a slice.

Syntax:

slice_name := make([]type, length, capacity)

Note: If the capacity parameter is not defined, it will be equal to length.

Example:

package main
import "fmt"

func main() {
  myslice1 := make([]int, 5, 10)
  fmt.Printf("myslice1 = %v\n", myslice1)
  fmt.Printf("length = %d\n", len(myslice1))
  fmt.Printf("capacity = %d\n", cap(myslice1))

  // with omitted capacity
  myslice2 := make([]int, 5)
  fmt.Printf("myslice2 = %v\n", myslice2)
  fmt.Printf("length = %d\n", len(myslice2))
  fmt.Printf("capacity = %d\n", cap(myslice2))
}


*can access a specific slice element by referring to the index number.
 indexes start at 0. That means that [0] is the first element, [1] is the second element, etc.

example:
package main
import "fmt"

func main() {
  prices := []int{10,20,30}

  fmt.Println(prices[0])		//10
  fmt.Println(prices[2])		//30
}

*we can also change a specific slice element by referring to the index number.

Append Elements To a Slice:

You can append elements to the end of a slice using the append()function:

Syntax:
slice_name = append(slice_name, element1, element2, ...)

Example:

package main
import "fmt"

func main() {
  myslice1 := []int{1, 2, 3, 4, 5, 6}
  fmt.Printf("myslice1 = %v\n", myslice1)
  fmt.Printf("length = %d\n", len(myslice1))
  fmt.Printf("capacity = %d\n", cap(myslice1))

  myslice1 = append(myslice1, 20, 21)
  fmt.Printf("myslice1 = %v\n", myslice1)
  fmt.Printf("length = %d\n", len(myslice1))
  fmt.Printf("capacity = %d\n", cap(myslice1))
}


Append One Slice To Another Slice:
To append all the elements of one slice to another slice, use the append()function:

Syntax:
slice3 = append(slice1, slice2...)

Note: The '...' after slice2 is necessary when appending the elements of one slice to another.


Change The Length of a Slice:
Unlike arrays, it is possible to change the length of a slice.

Example"

package main
import "fmt"

func main() {
  arr1 := [6]int{9, 10, 11, 12, 13, 14} // An array
  myslice1 := arr1[1:5] // Slice array
  fmt.Printf("myslice1 = %v\n", myslice1)			//myslice1 = [10 11 12 13]
  fmt.Printf("length = %d\n", len(myslice1))			//length=4			
  fmt.Printf("capacity = %d\n", cap(myslice1))			//capacity=5

  myslice1 = arr1[1:3] // Change length by re-slicing the array
  fmt.Printf("myslice1 = %v\n", myslice1)			//myslice1 = [10 11]
  fmt.Printf("length = %d\n", len(myslice1))			//length=2
  fmt.Printf("capacity = %d\n", cap(myslice1))			//capacity=5

  myslice1 = append(myslice1, 20, 21, 22, 23) // Change length by appending items
  fmt.Printf("myslice1 = %v\n", myslice1)			//myslice1 = [10 11 20 21 22 23]
  fmt.Printf("length = %d\n", len(myslice1))			//length=6
  fmt.Printf("capacity = %d\n", cap(myslice1))			//capacity=10
}


RANGE: 
it is used in both arrays and slices to loop through their elements.
syntax:
for index, value := range arrayOrSlice {
    // use index and value
}

Example with an array

package main

import "fmt"

func main() {

    numbers := [5]int{10, 20, 30, 40, 50}

    for index, value := range numbers {
        fmt.Println(index, value)
    }
}


Example with a slice

package main

import "fmt"

func main() {

    numbers := []int{10, 20, 30, 40, 50}

    for index, value := range numbers {
        fmt.Println(index, value)
    }
}

Difference betweem  arrays and slices in Go is :

Array has fixed size while
Slice has flexible/dynamic size.
