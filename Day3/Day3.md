Operators are used to perform operations on variables and values.
Types of operators:
1.Arithematic operators
2.Assingment operators
3.Comparison operators
4.Logical operators
5.Bitwise operators

1.Arithematic operator:
Arithmetic operators are used to perform common mathematical operations.

Operator	Name			Description					Example	
+		Addition		Adds together two values			x + y	
-		Subtraction		Subtracts one value from another		x - y	
*		Multiplication		Multiplies two values				x * y	
/		Division		Divides one value by another			x / y	
%		Modulus			Returns the division remainder			x % y	
++		Increment		Increases the value of a variable by 1		x++	
--		Decrement		Decreases the value of a variable by 1		x--

	
2.Assignment Operators:
Assignment operators are used to assign values to variables.
ex:we use the assignment operator (=) to assign the value 10 to a variable called x:
The addition assignment operator (+=) adds a value to a variable:

package main
import "fmt"

func main() {
  var x = 10
  x +=5
  fmt.Println(x)
}


A list of all assignment operators:

Operator	Example		Same As		
=		x = 5		x = 5	
+=		x += 3		x = x + 3	
-=		x -= 3		x = x - 3	
*=		x *= 3		x = x * 3	
/=		x /= 3		x = x / 3	
%=		x %= 3		x = x % 3	
&=		x &= 3		x = x & 3	
|=		x |= 3		x = x | 3	
^=		x ^= 3		x = x ^ 3	
>>=		x >>= 3		x = x >> 3	
<<=		x <<= 3		x = x << 3


3.Comparison operators:
Comparison operators are used to compare two values.
-The return value of a comparison is either true (1) or false (0).

A list of all comparison operators:

Operator	Name				Example	
==		Equal to			x == y	
!=		Not equal			x != y	
>		Greater than			x > y	
<		Less than			x < y	
>=		Greater than or equal to	x >= y	
<=		Less than or equal to		x <= y

4.Logical operators:
Logical operators are used to determine the logic between variables or values:

Operator	Name		Description							Example	
&& 		Logical and	Returns true if both statements are true			x < 5 &&  x < 10	
|| 		Logical or	Returns true if one of the statements is true			x < 5 || x < 4	
!		Logical not	Reverse the result, returns false if the result is true		!(x < 5 && x < 10)

5.Bitwise operators:
Bitwise operators are used on (binary) numbers:

Operator	Name			Description												Example	
& 		AND			Sets each bit to 1 if both bits are 1									x & y	
|		OR			Sets each bit to 1 if one of two bits is 1								x | y	
 ^		XOR			Sets each bit to 1 if only one of two bits is 1								x ^ b	
<<		Zero fill left shift	Shift left by pushing zeros in from the right								x << 2	
>>		Signed right shift	Shift right by pushing copies of the leftmost bit in from the left, and let the rightmost bits falloff	x >> 2
