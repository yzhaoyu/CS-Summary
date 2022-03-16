package cgo

/*
int sum(nt a, int b);
*/
import "C"

// export
func sum(a, b C.int) C.int {
	return a + b
}

func Randome() int {
	return int(C.random())
	var a uintptr 
}

func Seed(i int) {
	C.srandom(C.uint(i))
}
