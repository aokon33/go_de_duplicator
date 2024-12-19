package main

import (
	"fmt"
	"os"

	"duperemover/dr"
)

func main() {
	if len(os.Args[1:]) != 1 {
		fmt.Println("Please pass in one folder to de-duplicate!")
	} else {
		fmt.Println(os.Args[1])
		dr.Find_dupes(os.Args[1])
	}
}
