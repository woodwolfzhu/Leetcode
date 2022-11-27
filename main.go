package main

import (
	"fmt"
	_99 "helloWord/LeetCode/399"
)

func main() {
	fmt.Println("asdfafasfd")
	//equations := [][]string{
	//	{"a", "b"}, {"b", "c"},
	//}
	//values := []float64{2.0, 3.0}
	//queries := [][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}}

	equations := [][]string{
		{"a", "b"}, {"c", "d"},
	}
	values := []float64{1.0, 1.0}
	queries := [][]string{{"a", "c"}, {"b", "d"}, {"b", "a"}, {"d", "c"}, {"x", "x"}}

	//equations := [][]string{
	//	{"a", "b"}, {"b", "c"}, {"a", "c"}, {"d", "e"},
	//}
	//values := []float64{2, 3, 6, 1}
	//queries := [][]string{{"a", "c"}, {"b", "c"}, {"a", "e"}, {"a", "a"}, {"x", "x"}, {"a", "d"}}

	res := _99.CalcEquation(equations, values, queries)
	fmt.Println("真牛逼")
	fmt.Println(res)
}
