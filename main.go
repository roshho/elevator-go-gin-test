// elevator project and unit testing

/*
this patches?
localhost:8081/?UserCurrentF=6&UserFinalF=3


*/

package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	// elevatorA := Elevator{
	// 	SN:               1,
	// 	ElevatorCurrentF: 1,
	// 	UserCurrentF:     1,
	// 	UserFinalF:       1,
	// 	DoorOpen:         false,
	// 	ElevatorLog:      make([]int, 0),
	// }

	// // functional
	// // CLI input test
	// elevatorStartingTemp := elevatorA.ElevatorCurrentF
	// fmt.Println("Enter your current floor:")
	// fmt.Scanln(&elevatorA.UserCurrentF)
	// fmt.Println("Enter your destination floor:")
	// fmt.Scan(&elevatorA.UserFinalF)

	// fmt.Printf("Starting at %dF, ending at %dF", elevatorStartingTemp, elevatorA.ElevatorOperation())

	router := gin.Default()
	router.GET("/", getElevator)
	router.PUT("/", putElevator)

	router.Run("localhost:8080")

	// Complete by 230pm
	// CLI move to API
	// connect DB for new "elevators"
	// upload to git
}

type Elevator struct {
	SN               int   `json:"id"`
	ElevatorCurrentF int   `json:"elevatorcurrentf"`
	UserCurrentF     int   `json:"usercurrentf"` // Floor user is currently on
	UserFinalF       int   `json:"userfinalf"`   // Floor is trying to get to
	DoorOpen         bool  `json:"dooropen"`     // 1 is open
	ElevatorLog      []int `json:"elevatorlog"`
}

var elevators = Elevator{
	SN: 1, ElevatorCurrentF: 1, UserCurrentF: 3, UserFinalF: 8, DoorOpen: false, ElevatorLog: []int{1, 2, 3},
}

func getElevator(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, elevators)
}

func putElevator(c *gin.Context) {
	userCurrent := c.Query("usercurrentf") // ??? is this the most efficient way
	userFinal := c.Query("userfinalf")
	elevators.UserCurrentF, _ = strconv.Atoi(userCurrent) // ??? not sure if this is best method
	elevators.UserFinalF, _ = strconv.Atoi(userFinal)
	// c.String(http.StatusOK, "You're: %v, going to: %v", elevators.UserCurrentF, elevators.UserFinalF)
}

// Method for Elevator struct
func (e *Elevator) ElevatorOperation() int {
	fmt.Printf("Elevator on: %dF\n", e.ElevatorCurrentF)
	fmt.Printf("You're on %dF and want to move to %dF\n", e.UserCurrentF, e.UserFinalF)

	e.ElevatorLog = append(e.ElevatorLog, e.ElevatorCurrentF) // log starting floor

	tempDoWhile := false
	for tempDoWhile == false { // e.userCurrentF == e.userFinalF doesn't work bcos user floor would never change, e.elevatorCurrentF != e.userFinalF also doesn't work if currentF == finalF, but user on different F
		if e.UserCurrentF == e.UserFinalF {
			fmt.Println("Select a different destination floor")
			break // or to make it run repeatedly use tempDoWhile = true
		}

		e.DoorOpen = false
		fmt.Printf("Door open: (%v)\n", e.DoorOpen)

		// call move up function or move down function
		if e.ElevatorCurrentF < e.UserCurrentF {
			e.moveUp(e.ElevatorCurrentF, e.UserCurrentF)
		} else if e.ElevatorCurrentF > e.UserCurrentF {
			e.moveDown(e.ElevatorCurrentF, e.UserCurrentF)
		} else {
			fmt.Println("Same floor, no function call needed")
		}

		e.DoorOpen = true
		fmt.Printf("Door open: (%v)\n", e.DoorOpen)

		e.DoorOpen = false
		fmt.Printf("Door open: (%v)\n", e.DoorOpen)

		// call move up function or move down function
		if e.ElevatorCurrentF < e.UserFinalF {
			e.moveUp(e.ElevatorCurrentF, e.UserFinalF)
		} else {
			e.moveDown(e.ElevatorCurrentF, e.UserFinalF)
		}

		e.DoorOpen = true
		fmt.Printf("Door open: (%v)\n", e.DoorOpen)
		tempDoWhile = true
	}

	return e.ElevatorCurrentF
}

func (e *Elevator) moveUp(current int, destination int) {
	for i := current; i < destination; i++ {
		e.ElevatorCurrentF++
		fmt.Printf("Moving up: %dF\n", e.ElevatorCurrentF)
		e.ElevatorLog = append(e.ElevatorLog, e.ElevatorCurrentF)
	}
}

func (e *Elevator) moveDown(current int, destination int) {
	for i := current; i > destination; i-- {
		e.ElevatorCurrentF--
		fmt.Printf("Moving down: %dF\n", e.ElevatorCurrentF)
		e.ElevatorLog = append(e.ElevatorLog, e.ElevatorCurrentF)
	}
}
