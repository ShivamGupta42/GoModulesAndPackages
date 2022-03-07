package packc

import (
	"fmt"
	"sammysambomb/app/src/packa"
	packa "sammysambomb/app/src/packa/packt"
)

func PrintNameB() {
	fmt.Println("Package B calling " + packd.PrintNameA() + packa.PrintNameT())
}
