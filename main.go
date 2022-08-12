package main

import (
	"fmt"
	"os"

	"github.com/KNurlan2009/wordcount/countutil"
)

func main() {
	str := os.Args[1]
	fmt.Println(countutil.WordCounter(str))
}
