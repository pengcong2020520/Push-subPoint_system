package main

import (
	"fmt"
	"subsys/eths"
)

func main() {

	acc, _ := eths.NewKeystore("123")
	fmt.Println(acc)
}
