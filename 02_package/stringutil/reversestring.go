package stringutil

import "fmt"

func reverseString(value string) string {
	r := []rune(value)
	for i:=0 ; i < len(r); i++ {
		j := (len(r)-1)-i
		fmt.Println(i, j)
		if i >= j {
			break;
		}
		temp := r[i]
		r[i] = r[j]
		r[j] = temp
	}

	return string(r)
}
