package main

import (
	ecc "elliptic_curve"
	"fmt"
)

func main() {
	f44 := ecc.NewFieldElement(57, 44)
	f33 := ecc.NewFieldElement(57, 33)
	addResult := f44.Add(f33)
	negativeResult := addResult.Negative()
	subResult := f44.Subtract(f33)

	fmt.Print("add result: ", addResult)
	fmt.Print("neg result: ", negativeResult)
	fmt.Print("sub result: ", subResult)

}
