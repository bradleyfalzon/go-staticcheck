package pkg

func fn1() {
	var s []int
	s[0] = 0 // MATCH /index out of bounds/
}

func fn2() {
	s := make([]int, 2)
	s[2] = 0 // MATCH /index out of bounds/
}

func fn3() {
	var s []int
	s[0] = 0 // MATCH /index out of bounds/

	s = make([]int, 2)
	s[2] = 0 // MATCH /index out of bounds/
}

func fn4() {
	s := make([]int, 2)
	s = append(s, 1)
	s[2] = 0
}

func fn5(s []int) {
	s[2] = 0
}

func fn6(s []int) {
	s = s[:2]
	s[2] = 0 // MATCH /index out of bounds/
}

func fn7() {
	s := make([]int, 2)
	fn(s[2]) // MATCH /index out of bounds/
}

func fn8() {
	s := []int{}
	s[0] = 1 // MATCH /index out of bounds/
}

func fn9() {
	s := []int{}
	ptr(&s)
	s[0] = 1
}

func fn(int)     {}
func ptr(*[]int) {}
