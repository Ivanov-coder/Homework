package main

import (
	"fmt"
	"math"
)

type Complex struct{
	real float64
	imag float64
}

func complexAdd(c, d Complex) Complex {
	return Complex{c.real + d.real, c.imag + d.imag}
}

func complexMinus(c, d Complex) Complex {
	return Complex{c.real - d.real, c.imag - d.imag}
}

func complexMultiply(c, d Complex) Complex {
	return Complex{c.real * d.real - c.imag * d.imag, c.real * d.imag + c.imag * d.real}
}

func complexDivide(c, d Complex) Complex {
	under := d.real * d.real + d.imag * d.imag
	upperReal := c.real * d.real + c.imag * d.imag
	upperImag := c.imag * d.real - c.real * d.imag
	return Complex{upperReal / under, upperImag / under}
}

func complexLength(c Complex) float64 {
	return math.Sqrt(c.real * c.real + c.imag * c.imag)
}

func complexToString(c Complex) string{
	return fmt.Sprintf("%.2f + %.2fi", c.real, c.imag)
}

func main() {
	a := Complex{1, 2}
	b := Complex{3, 4}
	fmt.Println(complexToString(complexAdd(a, b)))
	fmt.Println(complexToString(complexMinus(a, b)))
	fmt.Println(complexToString(complexMultiply(a, b)))
	fmt.Println(complexToString(complexDivide(a, b)))
	fmt.Printf("%.2f", complexLength(a))
}