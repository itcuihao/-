package main

import (
	"fmt"
	//	"math/rand"
)

func main() {
	//	pt := Pascals_Triangle(rand.Intn(5))
	pt := Pascals_Triangle(5)
	fmt.Println(pt)
}

func Pascals_Triangle(rows int) [][]int {
	S := [][]int{}
	s1 := []int{1}
	S = append(S, s1)
	for i := 1; i < rows; i++ {

		//		s2:=[]
		//		S[i][0] = 1
		//		fmt.Println(S[i])
		//		S = append(S, S[i][0])
		//		S[i][len(S[i])-1] == 1
		//		S = append(S, S[i][len(S[i])-1])
		//		for j := 1; j < len(S[i])-1; j++ {
		//			S[i][j] = S[i-1][j-1] + S[i-1][j]
		//		}
	}
	return S
}
