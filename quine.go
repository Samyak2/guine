package main

// this is a quine in go by Samyak

import (
	"fmt"
)

func main() {
	begin := "package main\n\n// this is a quine in go by Samyak\n\nimport (\n\t\"fmt\"\n)\n\nfunc main() {\n"
	declareString := "\t%s := %#v\n"
	startPrint := "\tfmt.Printf("
	print1 := "%sbegin)\n%sdeclareString, \"begin\", begin)\n"
	print2 := "%sdeclareString, \"declareString\", declareString)\n%sdeclareString, \"startPrint\", startPrint)\n"
	print3 := "%sdeclareString, \"print1\", print1)\n%sdeclareString, \"print2\", print2)\n"
	print4 := "%sdeclareString, \"print3\", print3)\n%sdeclareString, \"print4\", print4)\n"
	print5 := "%sdeclareString, \"print5\", print5)\n%sdeclareString, \"print6\", print6)\n"
	print6 := "%sprint1, startPrint, startPrint)\n%sprint2, startPrint, startPrint)\n%sprint3, startPrint, startPrint)\n%sprint4, startPrint, startPrint)\n%sprint5, startPrint, startPrint)\n%sprint6, startPrint, startPrint, startPrint, startPrint, startPrint, startPrint)\n}\n"
	fmt.Printf(begin)
	fmt.Printf(declareString, "begin", begin)
	fmt.Printf(declareString, "declareString", declareString)
	fmt.Printf(declareString, "startPrint", startPrint)
	fmt.Printf(declareString, "print1", print1)
	fmt.Printf(declareString, "print2", print2)
	fmt.Printf(declareString, "print3", print3)
	fmt.Printf(declareString, "print4", print4)
	fmt.Printf(declareString, "print5", print5)
	fmt.Printf(declareString, "print6", print6)
	fmt.Printf(print1, startPrint, startPrint)
	fmt.Printf(print2, startPrint, startPrint)
	fmt.Printf(print3, startPrint, startPrint)
	fmt.Printf(print4, startPrint, startPrint)
	fmt.Printf(print5, startPrint, startPrint)
	fmt.Printf(print6, startPrint, startPrint, startPrint, startPrint, startPrint, startPrint)
}
