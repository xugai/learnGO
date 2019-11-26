package main

import "fmt"

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}

	fmt.Println("arr[2:6] =", arr[2:6])  // 2 3 4 5
	fmt.Println("arr[2:] =", arr[2:]) // 2 3 4 5 6 7
	fmt.Println("arr[:6] =", arr[:6]) // 0 1 2 3 4 5
	fmt.Println("arr[:] =", arr[:]) // 0 1 2 3 4 5 6 7

	s1 := arr[2:6] // 2 3 4 5
	fmt.Printf("s1 = %v, len(s1) = %d, cap(s1) = %d\n", s1, len(s1), cap(s1))

	fmt.Println("Reslice")
	s2 := s1[1:] // 3 4 5
	fmt.Printf("s2 = %v, len(s2) = %d, cap(s2) = %d\n", s2, len(s2), cap(s2))

	s3 := s2[2:4] // 5 6
	fmt.Printf("s3 = %v, len(s3) = %d, cap(s3) = %d\n", s3, len(s3), cap(s3))

	fmt.Println("slice append operation:")
	s3 = append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)
	fmt.Printf("s3 = %v, s4 = %v, s5 = %v\n", s3, s4, s5)
	fmt.Printf("arr = %v", arr)
}
