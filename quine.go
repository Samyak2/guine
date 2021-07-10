package main

import (
	"fmt"
)

func main() {
	s1 := "package main%c%cimport (%c	%cfmt%c%c)%c%cfunc main() {%c"
	s2 := "	s%c := %c%s%c%c"
	s3 := "	%c := %c%c%c%c%c"
	s4 := "	fmt.Printf("
	s5 := "%ss1, n, n, n, q, q, n, n, n, n);%ss2, '1', q, s1, q, n)%c"
	s6 := "%ss2, '2', q, s2, q, n);%ss2, '3', q, s3, q, n)%c"
	s7 := "%ss2, '4', q, s4, q, n);%ss2, '5', q, s5, q, n)%c"
	s8 := "%ss2, '6', q, s6, q, n);%ss2, '7', q, s7, q, n)%c"
	s9 := "%ss2, '8', q, s8, q, n);%ss2, '9', q, s9, q, n)%c"
	sx := "%ss2, 'x', q, sx, q, n);%ss2, '0', q, s0, q, n);%ss2, 'n', q, sn, q, n);%ss2, 'o', q, so, q, n);%ss2, 'p', q, sp, q, n)%c"
	s0 := "%ss2, 'a', q, sa, q, n);%ss2, 'b', q, sb, q, n)%c"
	sn := "%ss2, 'i', q, si, q, n);%ss2, 'j', q, sj, q, n)%c"
	so := "%ss2, 'k', q, sk, q, n);%ss2, 'l', q, sl, q, n)%c"
	sp := "%ss2, 'm', q, sm, q, n);%ss2, 'r', q, sr, q, n);%ss2, 's', q, ss, q, n);%ss2, 't', q, st, q, n);%ss2, 'u', q, su, q, n)%c"
	sa := "%ss3, 'n', r, c, 'n', r, n);%ss3, 'q', r, q, r, ' ', n)%c"
	sb := "%ss3, 'r', r, c, r, r, n);%ss3, 'c', r, c, c, r, n)%c"
	si := "%ss5, s4, s4, n);%ss6, s4, s4, n)%c"
	sj := "%ss7, s4, s4, n);%ss8, s4, s4, n)%c"
	sk := "%ss9, s4, s4, n);%ssx, s4, s4, s4, s4, s4, n)%c"
	sl := "%ss0, s4, s4, n);%ssn, s4, s4, n)%c"
	sm := "%sso, s4, s4, n);%ssp, s4, s4, s4, s4, s4, n)%c"
	sr := "%ssa, s4, s4, n)%c"
	ss := "%ssb, s4, s4, n);%ssi, s4, s4, n)%c"
	st := "%ssj, s4, s4, n);%ssk, s4, s4, n)%c"
	su := "%ssl, s4, s4, n);%ssm, s4, s4, n);%ssr, s4, n);%sss, s4, s4, n);%sst, s4, s4, n);%ssu, s4, s4, s4, s4, s4, s4, n, n)%c}%c"
	n := '\n'
	q := '"' 
	r := '\''
	c := '\\'
	fmt.Printf(s1, n, n, n, q, q, n, n, n, n);	fmt.Printf(s2, '1', q, s1, q, n)
	fmt.Printf(s2, '2', q, s2, q, n);	fmt.Printf(s2, '3', q, s3, q, n)
	fmt.Printf(s2, '4', q, s4, q, n);	fmt.Printf(s2, '5', q, s5, q, n)
	fmt.Printf(s2, '6', q, s6, q, n);	fmt.Printf(s2, '7', q, s7, q, n)
	fmt.Printf(s2, '8', q, s8, q, n);	fmt.Printf(s2, '9', q, s9, q, n)
	fmt.Printf(s2, 'x', q, sx, q, n);	fmt.Printf(s2, '0', q, s0, q, n);	fmt.Printf(s2, 'n', q, sn, q, n);	fmt.Printf(s2, 'o', q, so, q, n);	fmt.Printf(s2, 'p', q, sp, q, n)
	fmt.Printf(s2, 'a', q, sa, q, n);	fmt.Printf(s2, 'b', q, sb, q, n)
	fmt.Printf(s2, 'i', q, si, q, n);	fmt.Printf(s2, 'j', q, sj, q, n)
	fmt.Printf(s2, 'k', q, sk, q, n);	fmt.Printf(s2, 'l', q, sl, q, n)
	fmt.Printf(s2, 'm', q, sm, q, n);	fmt.Printf(s2, 'r', q, sr, q, n);	fmt.Printf(s2, 's', q, ss, q, n);	fmt.Printf(s2, 't', q, st, q, n);	fmt.Printf(s2, 'u', q, su, q, n)
	fmt.Printf(s3, 'n', r, c, 'n', r, n);	fmt.Printf(s3, 'q', r, q, r, ' ', n)
	fmt.Printf(s3, 'r', r, c, r, r, n);	fmt.Printf(s3, 'c', r, c, c, r, n)
	fmt.Printf(s5, s4, s4, n);	fmt.Printf(s6, s4, s4, n)
	fmt.Printf(s7, s4, s4, n);	fmt.Printf(s8, s4, s4, n)
	fmt.Printf(s9, s4, s4, n);	fmt.Printf(sx, s4, s4, s4, s4, s4, n)
	fmt.Printf(s0, s4, s4, n);	fmt.Printf(sn, s4, s4, n)
	fmt.Printf(so, s4, s4, n);	fmt.Printf(sp, s4, s4, s4, s4, s4, n)
	fmt.Printf(sa, s4, s4, n)
	fmt.Printf(sb, s4, s4, n);	fmt.Printf(si, s4, s4, n)
	fmt.Printf(sj, s4, s4, n);	fmt.Printf(sk, s4, s4, n)
	fmt.Printf(sl, s4, s4, n);	fmt.Printf(sm, s4, s4, n);	fmt.Printf(sr, s4, n);	fmt.Printf(ss, s4, s4, n);	fmt.Printf(st, s4, s4, n);	fmt.Printf(su, s4, s4, s4, s4, s4, s4, n, n)
}
