package main

import (
	"fmt"
	"strings"
)

type StringUnderlying string
type StringUnderlying2 string

func toUpper(s StringUnderlying) StringUnderlying {
	return StringUnderlying(strings.ToUpper(string(s)))
}

func main() {
	var str string
	var strUnd StringUnderlying
	var strUnd2 StringUnderlying2

	str = "string"
	strUnd = "strUnder"
	strUnd2 = "strUnder"

	fmt.Printf("%t\n", str == string(strUnd))
	fmt.Printf("%s\n", toUpper(StringUnderlying(str)))
	fmt.Printf("%s\n", toUpper(strUnd))
	fmt.Printf("%s\n", toUpper(StringUnderlying(strUnd2)))
}
