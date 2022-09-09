package main

import (
	"fmt"

	"github.com/jhony/zoo-golang/getspeciesbyids"
)

func main() {
	fmt.Printf(
		"%+v\n",
		getspeciesbyids.GSBI("0938aa23-f153-4937-9f88-4858b24d6bce"),
	)
}
