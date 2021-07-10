package main

import (
	"fmt"
)

func main() {
	s1 := "package main%c%cimport (%c	\"fmt\"%c)%c%cfunc main() {%c"
	s2 := "	s%c := %c%s%c%c"
	s3 := "	%c := %c%c%c%c%c"
	s4 := "	fmt.Printf("
	s5 := "s1, n, n, n, n, n, n, n);%ss2, '1', q, s1, q, n)%c"
	s6 := "s2, '2', q, s2, q, n);%ss2, '3', q, s3, q, n)%c"
	s7 := "s2, '4', q, s4, q, n);%ss2, '5', q, s5, q, n)%c"
	s8 := "s2, '6', q, s6, q, n);%ss2, '7', q, s7, q, n)%c"
	s9 := "s2, '8', q, s8, q, n);%ss2, '9', q, s9, q, n)%c"
	sx := "s2, '0', q, s0, q, n);%ss2, 'x', q, sx, q, n)%c"
	s0 := "s2, 'a', q, sa, q, n);%ss2, 'b', q, sb, q, n)%c"
	sa := "s3, 'n', r, c, 'n', r, n);%ss3, 'q', r, q, r, ' ', n)"
	sb := "s3, 'r', r, c, r, r, n);%ss3, 'c', r, c, c, r, n)"
	n := '\n'
	q := '"'
	r := '\''
	c := '\\'

	fmt.Printf(s1, n, n, n, n, n, n, n); fmt.Printf(s2, '1', q, s1, q, n)
	fmt.Printf(s2, '2', q, s2, q, n); fmt.Printf(s2, '3', q, s3, q, n)
	fmt.Printf(s2, '4', q, s4, q, n); fmt.Printf(s2, '5', q, s5, q, n)
	fmt.Printf(s2, '6', q, s6, q, n); fmt.Printf(s2, '7', q, s7, q, n)
	fmt.Printf(s2, '8', q, s8, q, n); fmt.Printf(s2, '9', q, s9, q, n)
	fmt.Printf(s2, '0', q, s0, q, n); fmt.Printf(s2, 'x', q, sx, q, n)
	fmt.Printf(s2, 'a', q, sa, q, n); fmt.Printf(s2, 'b', q, sb, q, n)
	fmt.Printf(s3, 'n', r, c, 'n', r, n); fmt.Printf(s3, 'q', r, q, r, ' ', n)
	fmt.Printf(s3, 'r', r, c, r, r, n); fmt.Printf(s3, 'c', r, c, c, r, n)
}
