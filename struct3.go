package main

import (
	"./struct_pack/structPack"
	"fmt"
)

func main() {
	struct1 := new(structPack.ExpStruct)
	struct1.Mil = 10
	struct1.Mf1 = 16.0

	fmt.Printf("Mi1 = %d\n", struct1.Mil)
	fmt.Printf("Mf1 = %f\n", struct1.Mf1)

}
