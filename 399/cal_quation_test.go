package _99

import "testing"

func TestCalcEquation(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		equations := [][]string{
			{"a", "b"}, {"b", "c"},
		}
		values := []float64{2.0, 3.0}
		queries := [][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}}
		expect := []float64{6.0, 0.5, -1.0, 1.0, -1.0}
		res1 := calcEquation1(equations, values, queries)
		res2 := calcEquation1(equations, values, queries)
		if !checkResult(res1, expect) {
			t.Errorf("calcEquation1 fail \nexpect=%v\nresult=%v", expect, res1)
		}
		if !checkResult(res2, expect) {
			t.Errorf("calcEquation2 fail \nexpect=%v\nresult=%v", expect, res2)
		}
	})

	t.Run("case2", func(t *testing.T) {
		equations := [][]string{
			{"a", "b"}, {"c", "d"},
		}
		values := []float64{1.0, 1.0}
		queries := [][]string{{"a", "c"}, {"b", "d"}, {"b", "a"}, {"d", "c"}, {"x", "x"}}
		expect := []float64{-1, -1, 1.0, 1.0, -1.0}
		res1 := calcEquation1(equations, values, queries)
		res2 := calcEquation1(equations, values, queries)
		if !checkResult(res1, expect) {
			t.Errorf("calcEquation1 fail \nexpect=%v\nresult=%v", expect, res1)
		}
		if !checkResult(res2, expect) {
			t.Errorf("calcEquation2 fail \nexpect=%v\nresult=%v", expect, res2)
		}
	})

	t.Run("case3", func(t *testing.T) {
		equations := [][]string{
			{"a", "b"}, {"b", "c"}, {"a", "c"}, {"d", "e"},
		}
		values := []float64{2, 3, 6, 1}
		queries := [][]string{{"a", "c"}, {"b", "c"}, {"a", "e"}, {"a", "a"}, {"x", "x"}, {"a", "d"}}
		expect := []float64{6, 3, -1, 1, -1, -1}
		res1 := calcEquation1(equations, values, queries)
		res2 := calcEquation1(equations, values, queries)
		if !checkResult(res1, expect) {
			t.Errorf("calcEquation1 fail \nexpect=%v\nresult=%v", expect, res1)
		}
		if !checkResult(res2, expect) {
			t.Errorf("calcEquation2 fail \nexpect=%v\nresult=%v", expect, res2)
		}
	})

}

func checkResult(res, expect []float64) bool {
	for index, v := range expect {
		if index > len(res) {
			return false
		}
		if v != res[index] {
			return false
		}
	}
	return true
}
