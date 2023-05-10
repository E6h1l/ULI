# Long arithmetic own implementation in Go
Implemented logical and bitwise operations:

Inv (bitwise inversion); Xor (bitwise exclusive or); Or (bitwise exclusive or); AND (bitwise and); 

shiftR (right shift by n bits); shiftL (left shift by n bits);

Implemented arithmetic operations:

Add (addition); other arithmetic operations will be implemented later...

# Usage
It is important to note that the program accepts numbers only in HEX as input.

Initializing a long number:
```golang
newInt := NewUli("Your Number")
```
Examples of using bitwise and arithmetic operations:
```golang
a := NewUli("1")
b := NewUli("2")

c := a.Add(b) //result of addition

xorRes := a.Xor() //result of xor
andRes := a.And(b) //result of and
orRes := a.Or(b) //result of or
orRes := b.Inv() //result of inv

shiftR := a.RightShift(shift) // result of right shift
shiftL := b.LeftShift(shift) // result of right shift

```
