// elevator project and unit testing

/*
this patches?
localhost:8081/?UserCurrentF=6&UserFinalF=3

mySQL Source: https://medium.com/@amberkakkar01/gin-and-database-integration-bridging-the-gap-between-sql-and-nosql-9c251a1c9fa9

*/

package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func main() {

	router := gin.Default()
	router.GET("/", getElevator)
	router.PUT("/", updateElevator)
	router.POST("/", createElevator)
	router.DELETE("/", deleteElevator)

	router.Run("localhost:8080")

	// db, err = sqlx.Connect("mysql", "test-user:1234@tcp(127.0.0.1:3306)/test-db")
}

type Elevator struct {
	SN               int   `gorm:"primaryKey;column:sn"`
	ElevatorCurrentF int   `gorm:"column:elevator_current_f"`
	UserCurrentF     int   `gorm:"column:user_current_f"`
	UserFinalF       int   `gorm:"column:user_final_f"`
	DoorOpen         bool  `gorm:"column:door_open"`
	ElevatorLog      []int `gorm:"column:elevator_log;type:integer[]"`
}

var elevators = Elevator{
	SN: 1, ElevatorCurrentF: 1, UserCurrentF: 3, UserFinalF: 8, DoorOpen: false, ElevatorLog: []int{},
}

func getElevator(c *gin.Context) {
	// // API, non-DB
	// c.IndentedJSON(http.StatusOK, elevators)

	// API w/ DB
	var elevatorsTemp []Elevator
	db.Select(&elevatorsTemp, "SELECT * FROM elevators")
	c.JSON(200, elevatorsTemp)
}

func createElevator(c *gin.Context) {
	var elevatorTemp Elevator
	c.BindJSON(&elevatorTemp)
	// db.Exec("INSERT INTO elevators (elevator_current_f, user_current_f, user_final_f, door_open, elevator_log) VALUES (?, ?, ?, ?, ?)", 1, elevatorTemp.ElevatorCurrentF, elevatorTemp.UserCurrentF, elevatorTemp.UserFinalF, elevatorTemp.DoorOpen, []int{})
	db.Create(&elevatorTemp)
	c.JSON(200, elevatorTemp)
}

func updateElevator(c *gin.Context) {
	// // API, non-DB
	// userCurrent := c.Query("usercurrentf") // ??? is this the most efficient way
	// userFinal := c.Query("userfinalf")
	// elevators.UserCurrentF, _ = strconv.Atoi(userCurrent) // ??? not sure if this is best method - considered public scope despite under class from struct
	// elevators.UserFinalF, _ = strconv.Atoi(userFinal)
	// c.String(http.StatusOK, "You're: %v, going to: %v", elevators.UserCurrentF, elevators.ElevatorOperation())

	var elevatorTemp Elevator
	c.BindJSON(&elevatorTemp)
	// db.Exec("INSERT INTO elevators (elevator_current_f, user_current_f, user_final_f, door_open, elevator_log) VALUES (?, ?, ?, ?, ?)", id, elevatorTemp.ElevatorCurrentF, elevatorTemp.UserCurrentF, elevatorTemp.UserFinalF, elevatorTemp.DoorOpen, []int{})
	db.Model(&elevatorTemp).Update(Elevator{SN: elevatorTemp.SN, UserCurrentF: elevatorTemp.UserCurrentF, UserFinalF: elevatorTemp.UserFinalF})
	c.JSON(200, elevatorTemp)

}

func deleteElevator(c *gin.Context) {
	var elevatorTemp Elevator
	c.BindJSON(&elevatorTemp)
	db.Delete(&elevatorTemp, 1)
}

// Method for Elevator struct
func (e *Elevator) ElevatorOperation() int {
	fmt.Printf("Elevator on: %dF\n", e.ElevatorCurrentF)
	fmt.Printf("You're on %dF and want to move to %dF\n", e.UserCurrentF, e.UserFinalF)

	e.ElevatorLog = append(e.ElevatorLog, e.ElevatorCurrentF) // log starting floor

	tempDoWhile := false
	for tempDoWhile == false {
		if e.UserCurrentF == e.UserFinalF {
			fmt.Println("Select a different destination floor")
			break // or to make it run repeatedly use tempDoWhile = true
		}

		e.DoorOpen = false
		fmt.Printf("Door open: (%v)\n", e.DoorOpen)

		e.moveElevator(e.UserCurrentF)

		e.DoorOpen = true
		fmt.Printf("Door open: (%v)\n", e.DoorOpen)
		e.DoorOpen = false
		fmt.Printf("Door open: (%v)\n", e.DoorOpen)

		e.moveElevator(e.UserFinalF)

		e.DoorOpen = true
		fmt.Printf("Door open: (%v)\n", e.DoorOpen)
		tempDoWhile = true
	}

	return e.ElevatorCurrentF
}

func (e *Elevator) moveElevator(destination int) {
	if e.ElevatorCurrentF < destination {
		for i := e.ElevatorCurrentF; i < destination; i++ {
			e.ElevatorCurrentF++
			fmt.Printf("Moving up: %dF\n", e.ElevatorCurrentF)
			e.ElevatorLog = append(e.ElevatorLog, e.ElevatorCurrentF)
		}
	} else {
		for i := e.ElevatorCurrentF; i > destination; i-- {
			e.ElevatorCurrentF--
			fmt.Printf("Moving down: %dF\n", e.ElevatorCurrentF)
			e.ElevatorLog = append(e.ElevatorLog, e.ElevatorCurrentF)
		}
	}
}
