package main

import "fmt"

func main() {

	arr := []string{"a", "a", "b", "c"}
	ret := []string{}

	for i, s := range arr {
		if (i == 0) || (ret[len(ret)-1] != arr[i]) {
			ret = append(ret, s)
		}
	}
	fmt.Println(ret)
}
