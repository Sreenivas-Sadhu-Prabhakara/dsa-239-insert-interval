package main

import (
	"reflect"
	"testing"
)

func TestInsert(t *testing.T) {
	cases := []struct {
		intervals [][2]int
		new       [2]int
		want      [][2]int
	}{
		{[][2]int{{1, 3}, {6, 9}}, [2]int{2, 5}, [][2]int{{1, 5}, {6, 9}}},
		{[][2]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, [2]int{4, 8}, [][2]int{{1, 2}, {3, 10}, {12, 16}}},
		{[][2]int{}, [2]int{5, 7}, [][2]int{{5, 7}}},
		{[][2]int{{1, 5}}, [2]int{2, 3}, [][2]int{{1, 5}}},         // swallowed whole
		{[][2]int{{3, 5}}, [2]int{1, 2}, [][2]int{{1, 2}, {3, 5}}}, // before everything
		{[][2]int{{1, 2}}, [2]int{4, 5}, [][2]int{{1, 2}, {4, 5}}}, // after everything
		{[][2]int{{1, 3}}, [2]int{3, 4}, [][2]int{{1, 4}}},         // touching merges
	}
	for _, c := range cases {
		if got := Insert(c.intervals, c.new); !reflect.DeepEqual(got, c.want) {
			t.Errorf("Insert(%v,%v) = %v, want %v", c.intervals, c.new, got, c.want)
		}
	}
}
