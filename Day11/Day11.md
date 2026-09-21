String package:
Go has a built-in package called strings that provides functions for working with strings.
we have import it first.
* import "strings"

example:
package main

import (
	"fmt"
	"strings"
)

func main() {
	name := "hello"

	fmt.Println(strings.ToUpper(name))
}

strings functions:

| Function            | Purpose                        |
| ------------------- | ------------------------------ |
| strings.ToUpper()   | Convert to uppercase           |
| strings.ToLower()   | Convert to lowercase           |
| strings.Contains()  | Check whether text exists      |
| strings.HasPrefix() | Check beginning of string      |
| strings.HasSuffix() | Check ending of string         |
| strings.TrimSpace() | Remove spaces at beginning/end |
| strings.Replace()   | Replace text                   |
| strings.Count()     | Count occurrences              |
| strings.Index()     | Find position of text          |
| strings.Split()     | Split a string                 |
| strings.Join()      | Join strings                   |


lets see some of them with examples:

*strings.Contains():
Check whether text exists.
Syntax:
strings.Contains(mainString, searchString)

example:
package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "I am learning Go"

	result := strings.Contains(text, "Go")

	fmt.Println(result)
}

*strings.Split():

It splits one string into multiple pieces and returns a slice of strings.

Example:
package main
import (
	"strings"
	"fmt"
	)
func main(){
text := "apple,banana,mango"

fruits := strings.Split(text, ",")

fmt.Println(fruits)		//[apple banana mango]
}


*strings.Join():

Join() does almost the opposite of Split().

example:
package main
import (
	"strings"
	"fmt"
	)
func main(){
 fruits := []string{"apple", "banana", "mango"}
 result := strings.Join(fruits, ", ")
 fmt.Println(result)
}

*strings.ToUpper():

Converts a string to uppercase.

package main

import (
	"fmt"
	"strings"
)

func main() {
	name := "hello"

	result := strings.ToUpper(name)

	fmt.Println(result)
}


*strings.TrimSpace():

Removes spaces from the beginning and end of a string.

package main
import ( 
	"strings"
	"fmt"
	)
text := "   hello   "

result := strings.TrimSpace(text)

fmt.Println(result)
