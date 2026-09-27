Pointers:
A pointer is a special kind of variable that is not only used to store the memory addresses of other variables but also points 
where the memory is located and provides ways to find out the value stored at that memory location.
It is generally termed as a Special kind of Variable because it is almost declared as a variable but with *(dereferencing operator).

* Operator also termed as the dereferencing operator used to declare pointer variable and access the value stored in the address.
& operator termed as address operator used to returns the address of a variable or to access the address of a variable to a pointer.

Declaring a pointer:

Syntax:
var pointer_name *Data_Type
ex: var s *string

Initialization of Pointer:
// normal variable declaration
var a = 45

// Initialization of pointer s with 
// memory address of variable a
var s *int = &a

The default value or zero-value of a pointer is always nil. Or you can say that an uninitialized pointer will always have a nil value.
Declaration and initialization of the pointers can be done into a single line.

Example: 
var s *int = &a

& — Address operator:
The & operator gives us the memory address of a variable.

* — Dereference operator:
The * operator can be used to get the value stored at an address.

Value					             Pointer
Passes a copy					Passes an address
Original usually stays unchanged		Original can be changed
int						*int
changeAge(age)					changeAge(&age)
Simple to understand				Useful for modifying original data
age gives value					*age gives value through pointer and &age gives address
Value = copy the data				Pointer = point to the original data.
