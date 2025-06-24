package main

import (
	ecc "elliptic_curve"
	"fmt"
	"math/rand"
)

func SolvField19MultipleSet() {
	// randomly select a num from 1 to 18
	min := 1
	max := 18
	k := rand.Intn(max-min) + min
	fmt.Println("randomly selected k: ", k)
	element := ecc.NewFieldElement(19, uint64(k))

	for i := 0; i < 19; i++ {
		fmt.Printf("element %d multiple with %d is %v", k, i, element.ScalarMul(uint64(i)))
	}

}

func main() {
	// f44 := ecc.NewFieldElement(57, 44)
	// f33 := ecc.NewFieldElement(57, 33)
	// addResult := f44.Add(f33)
	// negativeResult := addResult.Negative()
	// subResult := f44.Subtract(f33)

	// fmt.Print("add result: ", addResult)
	// fmt.Print("neg result: ", negativeResult)
	// fmt.Print("sub result: ", subResult)

	// f46 := ecc.NewFieldElement(57, 46)

	// fmt.Printf("multiple %v", f46.Multiplie(f46))
	// fmt.Printf("element 46 with the power of 2 is %v\n", f46.Power(2))

	SolvField19MultipleSet()
}
