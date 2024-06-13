/*
Reference:
https://stackoverflow.com/questions/67363083/running-unit-tests-multiple-times-with-different-values-in-golang
*/
package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestElevatorLevel(t *testing.T) {
	elevatorA := Elevator{
		ElevatorCurrentF: 1,
		UserCurrentF:     1,
		UserFinalF:       1,
		DoorOpen:         false,
	}

	var tests = []struct {
		a, b int
		want []int
	}{
		{0, 4, []int{1, 0, 1, 2, 3, 4}}, // a = user current floor, b = user destination, c = moveups called, d = movedownscalled
		{3, 1, []int{1, 2, 3, 2, 1}},
		{2, -2, []int{1, 2, 1, 0, -1, -2}},
		{-3, 0, []int{1, 0, -1, -2, -3, -2, -1, 0}},
	}

	for _, tt := range tests { // instead of
		// t.Run enables running "subtests", one for each table entry. These are shown separately when executing `go test -v`.
		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) { // from testing lib
			ans := elevatorA.ElevatorLog
			if reflect.DeepEqual(ans, tt.want) {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}

}
