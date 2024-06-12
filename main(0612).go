// elevator project and unit testing
// atomic counter: https://stackoverflow.com/questions/47422284/count-how-many-times-function-was-called
package main

import (
	"fmt"
)

func main() {
	elevatorA := Elevator{
		elevatorCurrentF: 1,
		userCurrentF:     1,
		userFinalF:       1,
		doorOpen:         false,
		elevatorLog:      make([]int, 0),
	}

	elevatorStartingTemp := elevatorA.elevatorCurrentF

	// CLI input test
	fmt.Println("Enter your current floor:")
	fmt.Scanln(&elevatorA.userCurrentF)
	fmt.Println("Enter your destination floor:")
	fmt.Scan(&elevatorA.userFinalF)

	fmt.Printf("Starting at %dF, ending at %dF", elevatorStartingTemp, elevatorA.ElevatorOperation())
}

type Elevator struct {
	elevatorCurrentF int  // current floor, 1 is going up, 0 is coming down
	userCurrentF     int  // Floor user is currently on
	userFinalF       int  // Floor is trying to get to
	doorOpen         bool // 1 is open
	elevatorLog      []int
}

// Method for Elevator struct
// Method for Elevator struct
func (e *Elevator) ElevatorOperation() int {
	fmt.Printf("Elevator on: %dF\n", e.elevatorCurrentF)
	fmt.Printf("You're on %dF and want to move to %dF\n", e.userCurrentF, e.userFinalF)

	e.elevatorLog = append(e.elevatorLog, e.elevatorCurrentF) // log starting floor

	tempDoWhile := false
	for tempDoWhile == false { // e.userCurrentF == e.userFinalF doesn't work bcos user floor would never change, e.elevatorCurrentF != e.userFinalF also doesn't work if currentF == finalF, but user on different F
		if e.userCurrentF == e.userFinalF {
			fmt.Println("Select a different destination floor")
			break
			tempDoWhile = true
		}

		e.doorOpen = false
		fmt.Printf("Door open: (%v)\n", e.doorOpen)

		// call move up function or move down function
		if e.elevatorCurrentF < e.userCurrentF {
			e.moveUp(e.elevatorCurrentF, e.userCurrentF)
		} else if e.elevatorCurrentF > e.userCurrentF {
			e.moveDown(e.elevatorCurrentF, e.userCurrentF)
		} else {
			fmt.Println("Same floor, no function call needed")
		}

		e.doorOpen = true
		fmt.Printf("Door open: (%v)\n", e.doorOpen)

		e.doorOpen = false
		fmt.Printf("Door open: (%v)\n", e.doorOpen)

		// call move up function or move down function
		if e.elevatorCurrentF < e.userFinalF {
			e.moveUp(e.elevatorCurrentF, e.userFinalF)
		} else {
			e.moveDown(e.elevatorCurrentF, e.userFinalF)
		}

		e.doorOpen = true
		fmt.Printf("Door open: (%v)\n", e.doorOpen)
		tempDoWhile = true
	}

	return e.elevatorCurrentF
}

func (e *Elevator) moveUp(current int, destination int) {
	for i := current; i < destination; i++ {
		e.elevatorCurrentF++
		fmt.Printf("Moving up: %dF\n", e.elevatorCurrentF)
		e.elevatorLog = append(e.elevatorLog, e.elevatorCurrentF)
	}
}

func (e *Elevator) moveDown(current int, destination int) {
	for i := current; i > destination; i-- {
		e.elevatorCurrentF--
		fmt.Printf("Moving down: %dF\n", e.elevatorCurrentF)
		e.elevatorLog = append(e.elevatorLog, e.elevatorCurrentF)
	}
}
