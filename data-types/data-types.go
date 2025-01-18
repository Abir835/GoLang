package main

import "fmt"

// int int32,int64,int8, int16 also unit same there
// float 32, 64
// complex
// string (Sequence of char)
// bool
// byte
// rune

// In Go, every data type has a unique format specifier.
//
// Data Type	Format Specifier
// integer	%d
// float	%g
// string	%s
// bool	%t

func main() {
	var annualSalary float64 = 65000.5
	var str string = "Hello World"

	fmt.Printf("Annual Salary: %g ,. %s,   ,  %t, . ., %d \n", annualSalary, str, true, 1)
	println("hello world")
}
