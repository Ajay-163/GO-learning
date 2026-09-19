
MAPS:
*Maps are used to store data values in key:value pairs.
*Each element in a map is a key:value pair.
*A map is an unordered and changeable collection that does not allow duplicates.
*The length of a map is the number of its elements.
*The default value of a map is nil.
*Maps hold references to an underlying hash table.
*Go has multiple ways for creating maps.

MAP creation:
Create Maps Using var and :=
Syntax:
var a = map[KeyType]ValueType{key1:value1, key2:value2,...}
b := map[KeyType]ValueType{key1:value1, key2:value2,...}

example:
package main
import "fmt"

func main() {
  var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
  b := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}

  fmt.Printf("a\t%v\n", a)
  fmt.Printf("b\t%v\n", b)
}

*The order of the map elements defined in the code is different from the way that they are stored. The data are                     stored in a way to have efficient data retrieval from the map.

Create Maps Using the make() Function:
Syntax:
var a = make(map[KeyType]ValueType)
b := make(map[KeyType]ValueType)

example:
package main
import "fmt"

func main() {
  var a = make(map[string]string) // The map is empty now
  a["brand"] = "Ford"
  a["model"] = "Mustang"Allowed Key Types
The map key can be of any data type for which the equality operator (==) is defined. These include:

Booleans
Numbers
Strings
Arrays
Pointers
Structs
Interfaces (as long as the dynamic type supports equality)
Invalid key types are:

Slices
Maps
Functions
These types are invalid because the equality operator (==) is not defined for them.

Allowed Value Types
The map values can be any type.

Access Map Elements

  a["year"] = "1964"
                                 // a is no longer empty
  b := make(map[string]int)
  b["Oslo"] = 1
  b["Bergen"] = 2
  b["Trondheim"] = 3
  b["Stavanger"] = 4

  fmt.Printf("a\t%v\n", a)
  fmt.Printf("b\t%v\n", b)
}

*The make()function is the right way to create an empty map.

Allowed Key Types:

The map key can be of any data type for which the equality operator (==) is defined. These include:
Booleans
Numbers
Strings
Arrays
Pointers
Structs
Interfaces (as long as the dynamic type supports equality)

Invalid key types are:
Slices
Maps
Functions
These types are invalid because the equality operator (==) is not defined for them.

Allowed Value Types:
*The map values can be any type.

Access Map Elements:
We can access map elements by:

Syntax:
value = map_name[key]

Example:
package main
import "fmt"

func main() {
  var a = make(map[string]string)
  a["brand"] = "Ford"
  a["model"] = "Mustang"
  a["year"] = "1964"

  fmt.Printf(a["brand"])
}

Update and Add Map Elements:
Updating or adding an elements are done by:

Syntax
map_name[key] = value

Example:
package main
import "fmt"

func main() {
  var a = make(map[string]string)
  a["brand"] = "Ford"
  a["model"] = "Mustang"
  a["year"] = "1964"

  fmt.Println(a)

  a["year"] = "1970" // Updating an element
  a["color"] = "red" // Adding an element

  fmt.Println(a)
}

Deleting Map:
Deleting elements is done using the delete() function.

Syntax:
delete(map_name, key)

Example:
package main
import "fmt"

func main() {
  var a = make(map[string]string)
  a["brand"] = "Ford"
  a["model"] = "Mustang"
  a["year"] = "1964"

  fmt.Println(a)
  delete(a,"year")
  fmt.Println(a)
}

Search of elements:
Syntax:
val, ok :=map_name[key]

Example:
package main
import "fmt"

func main() {
  var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964", "day":""}

  val1, ok1 := a["brand"] // Checking for existing key and its value
  val2, ok2 := a["color"] // Checking for non-existing key and its value
  val3, ok3 := a["day"]   // Checking for existing key and its value
  _, ok4 := a["model"]    // Only checking for existing key and not its value

  fmt.Println(val1, ok1)
  fmt.Println(val2, ok2)
  fmt.Println(val3, ok3)
  fmt.Println(ok4)
}





