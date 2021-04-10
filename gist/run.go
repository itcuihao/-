package gist

import (
	"fmt"
	"strings"
	"unicode"
)

// 驼峰转下划线
func hump2Underscore(name string) string {
	builder := strings.Builder{}
	for i, r := range name {
		if unicode.IsUpper(r) {
			if i != 0 {
				builder.WriteByte('_')
			}
			builder.WriteRune(unicode.ToLower(r))
		} else {
			builder.WriteRune(r)
		}
	}
	return builder.String()
}

func binarySearch(target int, haystack []int) bool {

	low := 0
	high := len(haystack) - 1

	for low <= high {
		median := (low + high) / 2

		if haystack[median] < target {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	if low == len(haystack) || haystack[low] != target {
		return false
	}

	return true
}

type List struct {
	Value int
	Next  *List
}

func plist() {
	l := &List{Value: 1, Next: &List{Value: 2, Next: &List{Value: 3, Next: &List{Value: 4, Next: &List{Value: 5, Next: &List{Value: 6}}}}}}
	printList(l)
	//deleteList(l, 2)
	removeNthFromEnd(l, 2)
	printList(l)
}

func printList(l *List) {
	ll := make([]int, 0, 3)
	for l != nil {
		ll = append(ll, l.Value)
		l = l.Next
	}
	fmt.Println(ll)
}

func deleteList(l *List, n int) {
	i := 0
	tmp := l
	for tmp != nil {
		tmp = tmp.Next
		i++
	}
	target := i - n
	fmt.Println(target)
	i = 1
	for l != nil {
		fmt.Println(i)
		if i == target {
			l.Next = l.Next.Next
		}
		i++
		l = l.Next
	}
}

// 利用快慢指针
func removeNthFromEnd(head *List, n int) *List {
	p := &List{}
	s := p
	e := p
	p.Next = head
	for i := 0; i <= n; i++ {
		s = s.Next
	}
	for s != nil {
		s = s.Next
		e = e.Next
	}
	e.Next = e.Next.Next
	fmt.Println(e.Value)
	return p
}
