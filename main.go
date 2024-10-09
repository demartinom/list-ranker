package main

import (
	"flag"
	"fmt"
)

func main() {
	list := flag.String("list", "cities.csv", "list to use for ranking")
	flag.Parse()

	fmt.Println(*list)
}
