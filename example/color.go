package main

import (
	"fmt"
	"github.com/khchehab/muzayaf/color"
)

func mainColor() {
	fmt.Println("Color Package Examples:")
	fmt.Println("------------------------")

	// CMYK examples
	fmt.Println("\nCMYK Examples:")
	fmt.Printf("Random CMYK color: %s\n", color.CMYK())
	fmt.Printf("Another random CMYK color: %s\n", color.CMYK())

	// HSL examples
	fmt.Println("\nHSL Examples:")
	fmt.Printf("Random HSL color: %s\n", color.HSL())
	fmt.Printf("Random HSLA color: %s\n", color.HSLA())
	fmt.Printf("Another random HSL color: %s\n", color.HSL())

	// RGBA examples
	fmt.Println("\nRGBA Examples:")
	fmt.Printf("Random RGBA color: %s\n", color.RGBA())
	fmt.Printf("Another random RGBA color: %s\n", color.RGBA())
	fmt.Printf("Random RGB color (full alpha): %s\n", color.RGB())

	// Color name examples
	fmt.Println("\nColor Name Examples:")
	fmt.Printf("Random color name: %s\n", color.ColorName())
	fmt.Printf("Random color name with locale: %s\n", color.ColorName(color.WithLocale("en")))
	fmt.Printf("Another random color name: %s\n", color.ColorName())
}
