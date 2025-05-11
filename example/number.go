package main

import (
	"fmt"
	"github.com/khchehab/muzayaf/number"
)

func mainNumber() {
	fmt.Println("Number Package Examples:")
	fmt.Println("------------------------")

	// Integer examples
	fmt.Println("\nInteger Examples:")
	fmt.Printf("Random integer: %d\n", number.Int())
	fmt.Printf("Random integer (10-20): %d\n", number.Int(number.WithIntMin(10), number.WithIntMax(20)))
	fmt.Printf("Random integer (multiple of 5): %d\n", number.Int(number.WithIntMin(10), number.WithIntMax(50), number.WithIntMultiple(5)))

	// Float examples
	fmt.Println("\nFloat Examples:")
	fmt.Printf("Random float: %f\n", number.Float())
	fmt.Printf("Random float (1.5-3.5): %f\n", number.Float(number.WithFloatMin(1.5), number.WithFloatMax(3.5)))
	fmt.Printf("Random float (3 decimal places): %f\n", number.Float(number.WithFloatFractionDigits(3)))

	// Binary examples
	fmt.Println("\nBinary Examples:")
	fmt.Printf("Random binary: %s\n", number.Binary())
	fmt.Printf("Random binary (5-10): %s\n", number.Binary(number.WithBinaryMin(5), number.WithBinaryMax(10)))
	fmt.Printf("Random binary with prefix: %s\n", number.Binary(number.WithBinaryPrefix(true)))

	// Octal examples
	fmt.Println("\nOctal Examples:")
	fmt.Printf("Random octal: %s\n", number.Octal())
	fmt.Printf("Random octal (10-20): %s\n", number.Octal(number.WithOctalMin(10), number.WithOctalMax(20)))
	fmt.Printf("Random octal with prefix: %s\n", number.Octal(number.WithOctalPrefix(true)))

	// Hex examples
	fmt.Println("\nHex Examples:")
	fmt.Printf("Random hex: %s\n", number.Hex())
	fmt.Printf("Random hex (10-20): %s\n", number.Hex(number.WithHexMin(10), number.WithHexMax(20)))
	fmt.Printf("Random hex with prefix: %s\n", number.Hex(number.WithHexPrefix(true)))

	// Roman numeral examples
	fmt.Println("\nRoman Numeral Examples:")
	fmt.Printf("Random roman numeral: %s\n", number.Roman())
	fmt.Printf("Random roman numeral (1-10): %s\n", number.Roman(number.WithRomanMin(1), number.WithRomanMax(10)))
	fmt.Printf("Random roman numeral (50-100): %s\n", number.Roman(number.WithRomanMin(50), number.WithRomanMax(100)))
}
