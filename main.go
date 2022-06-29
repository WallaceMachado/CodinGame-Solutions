package main

import "fmt"

func main() {
	ordemcrescentearrayInt()
	ordemcrescenteString()
	ordemcrescentearrayString()
	codiMap()
	ordemdecrescenteString()
	ordemdecrescentearrayInt()
	ordemdecrescentearrayString()
	codiMapDec()
}

func ordemdecrescenteString() {

	a := "abcdefg"
	var b []string

	for _, v := range a {
		b = append(b, string(v))

	}
	var c string
	for j := 1; j <= len(a); j++ {
		for i := 0; i < len(b)-1; i++ {
			if b[i] > b[i+1] {
				c = string(b[i])
				b[i] = b[i+1]
				b[i+1] = c
			}
		}
	}

	fmt.Println(b)
}

func ordemdecrescentearrayInt() {
	a := []int{1, 5, 9, 3, 756145, 566}
	var c int
	for j := 1; j <= len(a); j++ {
		for i := 0; i < len(a)-1; i++ {
			if a[i] < a[i+1] {
				c = a[i]
				a[i] = a[i+1]
				a[i+1] = c
			}
		}
	}
	fmt.Println(a[0])

	a[0] = a[0] * 2
	fmt.Println(a[0])

	fmt.Println(a)
}

func ordemdecrescentearrayString() {
	a := []string{"gdurnat", "n", "a", "a", "g"}
	var c string
	for j := 1; j <= len(a); j++ {
		for i := 0; i < len(a)-1; i++ {
			if a[i] < a[i+1] {
				c = string(a[i])
				a[i] = a[i+1]
				a[i+1] = c
			}
		}
	}

	fmt.Println(a)

}

func ordemcrescenteString() {
	a := "gduranat"
	var b []string
	for _, v := range a {
		b = append(b, string(v))
	}
	var c string
	for j := 1; j <= len(a); j++ {
		for i := 0; i < len(b)-1; i++ {
			if b[i] > b[i+1] {
				c = string(b[i])
				b[i] = b[i+1]
				b[i+1] = c
			}
		}
	}

	fmt.Println(b)
}

func ordemcrescentearrayInt() {
	a := []int{1, 4, 3, 3, 756145, 566}
	var c int
	for j := 1; j <= len(a); j++ {
		for i := 0; i < len(a)-1; i++ {
			if a[i] > a[i+1] {
				c = a[i]
				a[i] = a[i+1]
				a[i+1] = c
			}
		}
	}

	fmt.Println(a)
}

func ordemcrescentearrayString() {
	//	a := []int{1, 5, 9, 3, 756145, 566}
	a := []string{"gdurnat", "n", "a", "a", "g"}
	var c string
	for j := 1; j <= len(a); j++ {
		for i := 0; i < len(a)-1; i++ {
			if a[i] > a[i+1] {
				c = string(a[i])
				a[i] = a[i+1]
				a[i+1] = c
			}
		}
	}

	fmt.Println(a)

}

func codiMap() {
	a := "abcde"
	b := []int{5, 4, 3, 2, 1}
	var c []string
	var d = make(map[int]string)
	for _, v := range a {
		c = append(c, string(v))
	}
	for i := 0; i < len(a); i++ {
		d[b[i]] = c[i]
	}
	s := ""
	var f int
	for j := 1; j <= len(a); j++ {
		for i := 0; i < len(a)-1; i++ {
			if b[i] > b[i+1] {
				f = b[i]
				b[i] = b[i+1]
				b[i+1] = f
			}
		}
	}
	for _, v := range b {
		s = s + d[v]
	}
	fmt.Println(s)
}

func codiMapDec() {
	a := "abcde"
	b := []int{5, 4, 3, 2, 1}
	var c []string
	var d = make(map[int]string)
	for _, v := range a {
		c = append(c, string(v))
	}
	for i := 0; i < len(a); i++ {
		d[b[i]] = c[i]
	}
	s := ""
	var f int
	for j := 1; j <= len(a); j++ {
		for i := 0; i < len(a)-1; i++ {
			if b[i] < b[i+1] {
				f = b[i]
				b[i] = b[i+1]
				b[i+1] = f
			}
		}
	}
	for _, v := range b {
		s = s + d[v]
	}
	fmt.Println(s)
}
