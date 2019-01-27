package topic12

import "fmt"

func f1(n int) {
	number := 1
	i := 0
	for i < n {
		number *= 10
		i++
	}
	for i = 1; i < number; i++ {
		fmt.Println(i)
	}
}

func f2(n int) {
	if n <= 0 {
		return
	}

	s := new(Strn)
	number := make([]rune, n)
	for i := 0; i < n; i++ {
		number[i] = '0'
	}
	s.Number = number

	for !s.increment() {
		printn(s.Number)
	}
}

type Strn struct {
	Number []rune
}

// 字符串+1
func (s *Strn) increment() bool {
	number := s.Number
	// isOverflow 是否溢出
	// nTakeOver 是否进位
	isOverflow, nTakeOver, nlen := false, 0, len(number)
	for i := nlen - 1; i >= 0; i-- {
		n := int(number[i] - '0')
		nsum := n + nTakeOver
		if i == nlen-1 {
			nsum++
		}
		switch {
		case nsum >= 10:
			if i == 0 {
				isOverflow = true
			} else {
				nsum -= 10
				nTakeOver = 1
				number[i] = '0' + rune(nsum)
			}
		default:
			number[i] = '0' + rune(nsum)
			break
		}
	}
	s.Number = number
	return isOverflow
}

func printn(number []rune) {
	isBeginning0 := true
	nlen := len(number)
	for i := 0; i < nlen; i++ {
		if isBeginning0 && number[i] != '0' {
			isBeginning0 = false
		}
		if !isBeginning0 {
			fmt.Printf("%c", number[i])
		}
	}
	fmt.Print("\n")
}
