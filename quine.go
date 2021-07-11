package main

// a quine in go by Samyak

import (
	"fmt"
)

func main() {
	begin := "package main\n\n// a quine in go by Samyak\n\nimport (\n\t\"fmt\"\n)\n\nfunc main() {\n"
	newString := "\t%s := %#v\n"
	sPrint := "\tfmt.Printf("
	p1 := "%sbegin)\n%snewString, \"begin\", begin)\n"
	p2 := "%snewString, \"newString\", newString)\n%snewString, \"sPrint\", sPrint)\n"
	p3 := "%snewString, \"p1\", p1)\n%snewString, \"p2\", p2)\n"
	p4 := "%snewString, \"p3\", p3)\n%snewString, \"p4\", p4)\n"
	p5 := "%snewString, \"p5\", p5)\n%snewString, \"p0\", p0)\n%snewString, \"p6\", p6)\n"
	p0 := "%snewString, \"p7\", p7)\n%snewString, \"p8\", p8)\n%snewString, \"p9\", p9)\n"
	p6 := "%sp1, sPrint, sPrint)\n%sp2, sPrint, sPrint)\n"
	p7 := "%sp3, sPrint, sPrint)\n%sp4, sPrint, sPrint)\n"
	p8 := "%sp5, sPrint, sPrint, sPrint)\n%sp0, sPrint, sPrint, sPrint)\n%sp6, sPrint, sPrint)\n"
	p9 := "%sp7, sPrint, sPrint)\n%sp8, sPrint, sPrint, sPrint)\n%sp9, sPrint, sPrint, sPrint)\n}\n"
	fmt.Printf(begin)
	fmt.Printf(newString, "begin", begin)
	fmt.Printf(newString, "newString", newString)
	fmt.Printf(newString, "sPrint", sPrint)
	fmt.Printf(newString, "p1", p1)
	fmt.Printf(newString, "p2", p2)
	fmt.Printf(newString, "p3", p3)
	fmt.Printf(newString, "p4", p4)
	fmt.Printf(newString, "p5", p5)
	fmt.Printf(newString, "p0", p0)
	fmt.Printf(newString, "p6", p6)
	fmt.Printf(newString, "p7", p7)
	fmt.Printf(newString, "p8", p8)
	fmt.Printf(newString, "p9", p9)
	fmt.Printf(p1, sPrint, sPrint)
	fmt.Printf(p2, sPrint, sPrint)
	fmt.Printf(p3, sPrint, sPrint)
	fmt.Printf(p4, sPrint, sPrint)
	fmt.Printf(p5, sPrint, sPrint, sPrint)
	fmt.Printf(p0, sPrint, sPrint, sPrint)
	fmt.Printf(p6, sPrint, sPrint)
	fmt.Printf(p7, sPrint, sPrint)
	fmt.Printf(p8, sPrint, sPrint, sPrint)
	fmt.Printf(p9, sPrint, sPrint, sPrint)
}
