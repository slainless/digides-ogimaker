package main

// // #include <stdlib.h>
// import "C"

import (
	"fmt"

	"github.com/slainless/digides-ogimaker/pkg/wasm"
)

func main() {
	param, err := wasm.ReadParameters()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(param)
	fmt.Println("Hello, World!")
}
