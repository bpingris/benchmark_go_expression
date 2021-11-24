package main

import (
	"fmt"
	"log"

	"github.com/antonmedv/expr"
)

type Temp struct {
	Min int
	Max int
}

func main() {
	out, err := expr.Compile("Min < 10 && Max < 20")
	if err != nil {
		log.Fatal(err)
	}
	res, err := expr.Run(out, Temp{Min: 5, Max: 10})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("res: %v\n", res)
	res, err = expr.Run(out, Temp{Min: 15, Max: 10})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("res: %v\n", res)
	res, err = expr.Run(out, Temp{Min: 5, Max: 10})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("res: %v\n", res)
}
